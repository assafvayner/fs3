package primary

import (
  "fmt"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
	"gitlab.cs.washington.edu/assafv/fs3/server/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/fs3processor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PrimaryHandler struct {
	Fs3processor *fs3processor.Fs3RequestProcessor
	PBClient     primarybackup.BackupClient
	fs3.UnimplementedFs3Server
}

func NewPrimaryHandler() *PrimaryHandler {
	// need to dial connection to backup

	return &PrimaryHandler{
		Fs3processor: fs3processor.NewFs3RequestProcessor(),
		PBClient:     nil,
	}
}

// before using PBClient, functions should call VerifyPBClient
// which lazily creates connection to backup
func (handler *PrimaryHandler) VerifyPBClient() {
	if handler.PBClient != nil {
		return
	}

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(fmt.Sprint("localhost:", config.BACKUP_PORT), opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not start a connection to backup")
		fmt.Fprintln(os.Stderr, "Guess I'll die")
		os.Exit(1)
	}
	handler.PBClient = primarybackup.NewBackupClient(conn)
}
