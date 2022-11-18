package primarybackuphandler

import (
	"context"
  primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

func (*PrimaryBackupHandler) Forward(ctx context.Context, req *primarybackup.ForwardRequest) (reply *primarybackup.ForwardReply, err error) {
  return reply, err
}