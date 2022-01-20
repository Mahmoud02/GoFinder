package memory

import (
	"GoFinder/textindexer/index/indextest"
	gc "gopkg.in/check.v1"
	"testing"
)

var _ = gc.Suite(new(InMemoryBleveTestSuite))

func Test(t *testing.T) { gc.TestingT(t) }

type InMemoryBleveTestSuite struct {
	indextest.SuiteBase
	idx *InMemoryBleveIndexer
}

func (s *InMemoryBleveTestSuite) SetUpTest(c *gc.C) {
	idx, err := NewInMemoryBleveIndexer()
	c.Assert(err, gc.IsNil)
	s.SetIndexer(idx)
	s.idx = idx
}

func (s *InMemoryBleveTestSuite) TearDownTest(c *gc.C) {
	c.Assert(s.idx.Close(), gc.IsNil)
}
