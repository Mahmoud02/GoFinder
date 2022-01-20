package postgresql

import (
	"GoFinder/linkgraph/graph/graphtest"
	"database/sql"
	gc "gopkg.in/check.v1"
	"os"
	"testing"
)

var _ = gc.Suite(new(PostgreDBTestSuite))

func Test(t *testing.T) { gc.TestingT(t) }

type PostgreDBTestSuite struct {
	graphtest.SuiteBase
	db *sql.DB
}

func (s *PostgreDBTestSuite) SetUpSuite(c *gc.C) {
	dsn := os.Getenv("CDB_DSN")
	if dsn == "" {
		c.Skip("Missing CDB_DSN envvar; skipping cockroachdb-backed graph test suite")
	}

	g, err := NewPostgreDBGraph(dsn)
	c.Assert(err, gc.IsNil)
	s.SetGraph(g)
	s.db = g.db
}

func (s *PostgreDBTestSuite) SetUpTest(c *gc.C) {
	s.flushDB(c)
}

func (s *PostgreDBTestSuite) TearDownSuite(c *gc.C) {
	if s.db != nil {
		s.flushDB(c)
		c.Assert(s.db.Close(), gc.IsNil)
	}
}

func (s *PostgreDBTestSuite) flushDB(c *gc.C) {
	_, err := s.db.Exec("DELETE FROM links")
	c.Assert(err, gc.IsNil)
	_, err = s.db.Exec("DELETE FROM edges")
	c.Assert(err, gc.IsNil)
}
