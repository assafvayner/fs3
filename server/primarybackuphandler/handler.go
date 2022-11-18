package primarybackuphandler

import (
  "gitlab.cs.washington.edu/assafv/fs3/server/fs3handler"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

type PrimaryBackupHandler struct {
  fs3handler *fs3handler.Fs3Handler
  fs3.UnimplementedBackupServer
}

func NewPrimaryBackupHandler() (*PrimaryBackupHandler) {
  return &PrimaryBackupHandler{
    fs3handler: fs3handler.NewFs3Handler(),
  }
}