package fs3handler

import (
  "context"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3Handler) Remove(ctx context.Context, req *fs3.RemoveRequest) (reply *fs3.RemoveReply, err error) {
  // do something with req
  return reply, err
}