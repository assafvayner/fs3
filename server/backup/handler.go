package backup

import (
  "gitlab.cs.washington.edu/assafv/fs3/server/fs3processor"
  fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

type BackupHandler struct {
  Fs3processor *fs3processor.Fs3RequestProcessor
  fs3.UnimplementedBackupServer
}

func NewBackupHandler() (*BackupHandler) {
  return &BackupHandler{
    Fs3processor: fs3processor.NewFs3RequestProcessor(),
  }
}