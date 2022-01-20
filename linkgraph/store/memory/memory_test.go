package memory

import (
	"GoFinder/linkgraph/graph/graphtest"
	gc "gopkg.in/check.v1"
	"testing"
)

var _ = gc.Suite(new(InMemoryGraphTestSuite))

type InMemoryGraphTestSuite struct {
	graphtest.SuiteBase
}

func (s *InMemoryGraphTestSuite) SetUpTest(c *gc.C) {
	s.SetGraph(NewInMemoryGraph())
}
func Test(t *testing.T) { gc.TestingT(t) }
