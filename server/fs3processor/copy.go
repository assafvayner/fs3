package fs3processor

import (
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3RequestProcessor) Copy(req *fs3.CopyRequest) (reply *fs3.CopyReply, err error) {
  reply = &fs3.CopyReply{}
  return reply, err
}
