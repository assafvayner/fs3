package fs3handler

import (

  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

type Fs3Handler struct {
  fs3.UnimplementedFs3Server
}

func NewFs3Handler() (*Fs3Handler) {
  return &Fs3Handler{}
}