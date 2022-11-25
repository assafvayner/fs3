package primary

import (
	"context"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	// primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

func (handler *PrimaryHandler) Get(ctx context.Context, req *fs3.GetRequest) (*fs3.GetReply, error) {
	// handler.VerifyPBClient()

	// do something with req
	// forwardReq := &primarybackup.ForwardRequest{
	// 	ClientRequest: &primarybackup.ForwardRequest_GetRequest{
	// 		GetRequest: req,
	// 	},
	// }

	// _, err = handler.PBClient.Forward(ctx, forwardReq)
	// if err != nil {
	// 	return nil, err
	// }

	return handler.Fs3processor.Get(req)
}
