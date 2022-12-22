package primary

import (
	"context"

	fs3 "github.com/assafvayner/fs3/protos/fs3"
	primarybackup "github.com/assafvayner/fs3/protos/primarybackup"
)

func (handler *PrimaryHandler) Copy(
	ctx context.Context,
	req *fs3.CopyRequest,
) (reply *fs3.CopyReply, err error) {
	handler.VerifyPBClient()

	// do something with req
	forwardReq := &primarybackup.ForwardRequest{
		ClientRequest: &primarybackup.ForwardRequest_CopyRequest{
			CopyRequest: req,
		},
	}

	_, err = handler.PBClient.Forward(ctx, forwardReq)
	if err != nil {
		handler.Logger.Printf(
			"error on forward for copy file <%s>, err: %s\n",
			req.GetFilePath(),
			err,
		)
		return nil, err
	}

	reply, err = handler.Fs3processor.Copy(req)

	return reply, err
}
