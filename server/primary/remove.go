package primary

import (
	"context"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

func (handler *PrimaryHandler) Remove(ctx context.Context, req *fs3.RemoveRequest) (reply *fs3.RemoveReply, err error) {
	handler.VerifyPBClient()

	// do something with req
	forwardReq := &primarybackup.ForwardRequest{
		ClientRequest: &primarybackup.ForwardRequest_RemoveRequest{
			RemoveRequest: req,
		},
	}

	_, err = handler.PBClient.Forward(ctx, forwardReq)
	if err != nil {
		handler.Logger.Printf("error on forward for remove file <%s>, err: %s\n", req.GetFilePath(), err)
		return nil, err
	}

	reply, err = handler.Fs3processor.Remove(req)

	return reply, err
}
