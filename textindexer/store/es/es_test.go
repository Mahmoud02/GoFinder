package es

import (
	"GoFinder/textindexer/index/indextest"
	gc "gopkg.in/check.v1"
	"os"
	"strings"
	"testing"
)

var _ = gc.Suite(new(ElasticSearchTestSuite))

func Test(t *testing.T) { gc.TestingT(t) }

type ElasticSearchTestSuite struct {
	indextest.SuiteBase
	idx *ElasticSearchIndexer
}

func (s *ElasticSearchTestSuite) SetUpSuite(c *gc.C) {
	nodeList := os.Getenv("ES_NODES")
	println(nodeList)
	if nodeList == "" {
		c.Skip("Missing ES_NODES envvar; skipping elasticsearch-backed index test suite")
	}

	idx, err := NewElasticSearchIndexer(strings.Split(nodeList, ","), true)
	c.Assert(err, gc.IsNil)
	s.SetIndexer(idx)
	s.idx = idx
}

func (s *ElasticSearchTestSuite) SetUpTest(c *gc.C) {
	if s.idx.es != nil {
		_, err := s.idx.es.Indices.Delete([]string{indexName})
		c.Assert(err, gc.IsNil)
		err = ensureIndex(s.idx.es)
		c.Assert(err, gc.IsNil)
	}
}
