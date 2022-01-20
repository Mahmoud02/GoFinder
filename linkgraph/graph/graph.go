package graph

import (
	"github.com/google/uuid"
	"time"
)

/*
	The LinkIterator and EdgeIterator types, which are returned by  methods, are
	interfaces themselves.their internal implementation details will
    depend on the database technology that we select for the link graph persistence layer.
*/

// Iterator is implemented by graph objects that can be iterated.
// The common logic between the  iterators has been extracted into a separate interface called Iterator
type Iterator interface {
	// Next advances the iterator. If no more items are available or an
	// error occurs, calls to Next() return false.
	Next() bool

	// Error returns the last error encountered by the iterator.
	Error() error

	// Close releases any resources associated with an iterator.
	Close() error
}

// LinkIterator is implemented by objects that can iterate the graph links.
type LinkIterator interface {
	Iterator

	// Link returns the currently fetched link object.
	Link() *Link
}

// EdgeIterator is implemented by objects that can iterate the graph edges.
type EdgeIterator interface {
	Iterator

	// Edge returns the currently fetched edge objects.
	Edge() *Edge
}

// Link
/*
 Link model instances represent the set of web pages that have been processed or discovered
 by the crawler component.
*/
type Link struct {
	// A unique identifier for the link.
	ID uuid.UUID

	// The link target.
	URL string

	// The timestamp when the link was last retrieved.
	RetrievedAt time.Time
}

// Edge
/*
 Each web page in the graph may contain zero or more outgoing links to other web pages. An
 Edge model instance represents a uni-directional connection between two links in the
 graph
*/
type Edge struct {
	// A unique identifier for the edge.
	ID uuid.UUID

	// The origin link.
	Src uuid.UUID

	// The destination link.
	Dst uuid.UUID

	// The timestamp when the link was last updated.
	UpdatedAt time.Time
}

// Graph interface is implemented by objects that can mutate or query a link graph.
/*
The crawler retrieves web page links and discovers connections between websites, it
makes sense for us to use a graph-based representation for our system modeling.
*/
type Graph interface {
	// InsertLink creates a new link or updates an existing link.
	InsertLink(link *Link) error

	// FindLink looks up a link by its ID.
	FindLink(id uuid.UUID) (*Link, error)

	// Links returns an iterator for the set of links whose IDs belong to the
	// [fromID, toID) range and were retrieved before the provided timestamp.
	Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (LinkIterator, error)

	// InsertEdge creates a new edge or updates an existing edge.
	InsertEdge(edge *Edge) error

	// Edges returns an iterator for the set of edges whose source vertex IDs
	// belong to the [fromID, toID) range and were updated before the provided
	// timestamp.
	Edges(fromID, toID uuid.UUID, updatedBefore time.Time) (EdgeIterator, error)

	// RemoveStaleEdges removes any edge that originates from the specified
	// link ID and was updated before the specified timestamp.
	RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error
}
