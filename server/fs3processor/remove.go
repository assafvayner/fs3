package fs3processor

import (
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3RequestProcessor) Remove(req *fs3.RemoveRequest) (reply *fs3.RemoveReply, err error) {
  // do something with req
  reply = &fs3.RemoveReply{}
  return reply, err
}