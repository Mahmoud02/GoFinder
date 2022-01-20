package textindexerapi_test

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	gc "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// Run all gocheck test-suites
	gc.TestingT(t)
}

func mustEncodeTimestamp(c *gc.C, t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	c.Assert(err, gc.IsNil)
	return ts
}
