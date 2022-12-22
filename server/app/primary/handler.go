package primary

import (
	"fmt"
	"log"

	fs3 "github.com/assafvayner/fs3/protos/fs3"
	primarybackup "github.com/assafvayner/fs3/protos/primarybackup"
	"github.com/assafvayner/fs3/server/app/config"
	"github.com/assafvayner/fs3/server/app/fs3processor"
	"github.com/assafvayner/fs3/server/shared/tlsutils"
	"google.golang.org/grpc"
)

type PrimaryHandler struct {
	Fs3processor *fs3processor.Fs3RequestProcessor
	PBClient     primarybackup.BackupClient
	Logger       *log.Logger
	fs3.UnimplementedFs3Server
}

func NewPrimaryHandler(logger *log.Logger) *PrimaryHandler {
	// PBClient to be lazily dialed later
	return &PrimaryHandler{
		Fs3processor: fs3processor.NewFs3RequestProcessor(logger),
		PBClient:     nil,
		Logger:       logger,
	}
}

// before using PBClient, functions should call VerifyPBClient
// which lazily creates connection to backup
func (handler *PrimaryHandler) VerifyPBClient() {
	if handler.PBClient != nil {
		return
	}

	tlsCredentials, err := tlsutils.GetClientTLSCredentials()
	if err != nil {
		handler.Logger.Fatalln("Could not get tls credentials to communicate with backup, big problem")
	}

	// TODO: figure out exactly how to determine backup hostname
	conn, err := grpc.Dial(fmt.Sprint("backup.fs3:", config.GetBackupPort()), grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		handler.Logger.Fatalln("Could not start a connection to backup")
	}
	handler.PBClient = primarybackup.NewBackupClient(conn)
}
