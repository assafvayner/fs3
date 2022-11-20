package fs3processor

import (
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (*Fs3RequestProcessor) Get(req *fs3.GetRequest) (reply *fs3.GetReply, err error) {
  // do something with req
  reply = &fs3.GetReply{}
  return reply, err
}