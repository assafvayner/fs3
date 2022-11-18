package fs3handler

import (
  "context"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3Handler) Get(ctx context.Context, req *fs3.GetRequest) (reply *fs3.GetReply, err error) {
  // do something with req
  return reply, err
}