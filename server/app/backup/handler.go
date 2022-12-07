package backup

import (
	"log"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
	"gitlab.cs.washington.edu/assafv/fs3/server/app/fs3processor"
)

type BackupHandler struct {
	Fs3processor *fs3processor.Fs3RequestProcessor
	Logger       *log.Logger
	fs3.UnimplementedBackupServer
}

func NewBackupHandler(logger *log.Logger) *BackupHandler {
	return &BackupHandler{
		Fs3processor: fs3processor.NewFs3RequestProcessor(logger),
		Logger:       logger,
	}
}
