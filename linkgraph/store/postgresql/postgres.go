package postgresql

import (
	"GoFinder/linkgraph/graph"
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/xerrors"

	"time"
)

var (
	InsertLinkQuery = `INSERT INTO links (url, retrieved_at) VALUES ($1, $2) 
         ON CONFLICT (url) DO UPDATE SET retrieved_at=GREATEST(links.retrieved_at, $2)
         RETURNING id, retrieved_at`

	findLinkQuery         = "SELECT url, retrieved_at FROM links WHERE id=$1"
	linksInPartitionQuery = "SELECT id, url, retrieved_at FROM links WHERE id >= $1 AND id < $2 AND retrieved_at < $3"

	InsertEdgeQuery = `INSERT INTO edges (src, dst, updated_at) VALUES ($1, $2, NOW())
         ON CONFLICT (src,dst) DO UPDATE SET updated_at=NOW()
         RETURNING id, updated_at`

	edgesInPartitionQuery = "SELECT id, src, dst, updated_at FROM edges WHERE src >= $1 AND src < $2 AND updated_at < $3"
	removeStaleEdgesQuery = "DELETE FROM edges WHERE src=$1 AND updated_at < $2"

	// Compile-time check for ensuring CockroachDbGraph implements Graph.
	_ graph.Graph = (*PostgreDBGraph)(nil)
)

// PostgreDBGraph implements a graph that persists its links and edges to a  PostgreSQL instance.
type PostgreDBGraph struct {
	db *sql.DB
}

// NewPostgreDBGraph returns a CockroachDbGraph instance that connects to the cockroachdb
// instance specified by dsn.
func NewPostgreDBGraph(dsn string) (*PostgreDBGraph, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &PostgreDBGraph{db: db}, nil
}

// Close terminates the connection to the backing cockroachdb instance.
func (c *PostgreDBGraph) Close() error {
	return c.db.Close()
}

// InsertLink creates a new link or updates an existing link.
func (c *PostgreDBGraph) InsertLink(link *graph.Link) error {
	row := c.db.QueryRow(InsertLinkQuery, link.URL, link.RetrievedAt.UTC())
	if err := row.Scan(&link.ID, &link.RetrievedAt); err != nil {
		return xerrors.Errorf("upsert link: %w", err)
	}

	link.RetrievedAt = link.RetrievedAt.UTC()
	return nil
}

// FindLink looks up a link by its ID.
func (c *PostgreDBGraph) FindLink(id uuid.UUID) (*graph.Link, error) {
	row := c.db.QueryRow(findLinkQuery, id)
	link := &graph.Link{ID: id}
	if err := row.Scan(&link.URL, &link.RetrievedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, xerrors.Errorf("find link: %w", graph.ErrNotFound)
		}

		return nil, xerrors.Errorf("find link: %w", err)
	}

	link.RetrievedAt = link.RetrievedAt.UTC()
	return link, nil
}

// Links returns an iterator for the set of links whose IDs belong to the
// [fromID, toID) range and were last accessed before the provided value.
func (c *PostgreDBGraph) Links(fromID, toID uuid.UUID, accessedBefore time.Time) (graph.LinkIterator, error) {
	rows, err := c.db.Query(linksInPartitionQuery, fromID, toID, accessedBefore.UTC())
	if err != nil {
		return nil, xerrors.Errorf("links: %w", err)
	}

	return &linkIterator{rows: rows}, nil
}

// InsertEdge creates a new edge or updates an existing edge.
func (c *PostgreDBGraph) InsertEdge(edge *graph.Edge) error {
	row := c.db.QueryRow(InsertEdgeQuery, edge.Src, edge.Dst)
	if err := row.Scan(&edge.ID, &edge.UpdatedAt); err != nil {
		if isForeignKeyViolationError(err) {
			err = graph.ErrUnknownEdgeLinks
		}
		return xerrors.Errorf("upsert edge: %w", err)
	}

	edge.UpdatedAt = edge.UpdatedAt.UTC()
	return nil
}

// Edges returns an iterator for the set of edges whose source vertex IDs
// belong to the [fromID, toID) range and were last updated before the provided
// value.
func (c *PostgreDBGraph) Edges(fromID, toID uuid.UUID, updatedBefore time.Time) (graph.EdgeIterator, error) {
	rows, err := c.db.Query(edgesInPartitionQuery, fromID, toID, updatedBefore.UTC())
	if err != nil {
		return nil, xerrors.Errorf("edges: %w", err)
	}

	return &edgeIterator{rows: rows}, nil
}

// RemoveStaleEdges removes any edge that originates from the specified link ID
// and was updated before the specified timestamp.
func (c *PostgreDBGraph) RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error {
	_, err := c.db.Exec(removeStaleEdgesQuery, fromID, updatedBefore.UTC())
	if err != nil {
		return xerrors.Errorf("remove stale edges: %w", err)
	}

	return nil
}

// isForeignKeyViolationError returns true if err indicates a foreign key constraint violation.
func isForeignKeyViolationError(err error) bool {
	pqErr, valid := err.(*pq.Error)
	if !valid {
		return false
	}

	return pqErr.Code.Name() == "foreign_key_violation"
}
