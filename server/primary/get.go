package primary

import (
	"context"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

// no need to forward, optimization
func (handler *PrimaryHandler) Get(ctx context.Context, req *fs3.GetRequest) (*fs3.GetReply, error) {
	return handler.Fs3processor.Get(req)
}
