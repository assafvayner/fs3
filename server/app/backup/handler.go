package backup

import (
	"log"

	fs3 "github.com/assafvayner/fs3/protos/primarybackup"
	"github.com/assafvayner/fs3/server/app/fs3processor"
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
