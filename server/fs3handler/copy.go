package fs3handler

import (
  "context"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3Handler) Copy(ctx context.Context, req *fs3.CopyRequest) (reply *fs3.CopyReply, err error) {
  // do something with req
  return reply, err
}