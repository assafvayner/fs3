package primary

import (
	"context"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (handler *PrimaryHandler) Describe(ctx context.Context, req *fs3.DescribeRequest) (*fs3.DescribeReply, error) {
  return handler.Fs3processor.Describe(req)
}