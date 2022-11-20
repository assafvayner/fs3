package primary

import (
	"context"
  "fmt"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	primarybackup "gitlab.cs.washington.edu/assafv/fs3/protos/primarybackup"
)

func (handler *PrimaryHandler) Get(ctx context.Context, req *fs3.GetRequest) (reply *fs3.GetReply, err error) {
	handler.VerifyPBClient()

	// do something with req
	forwardReq := &primarybackup.ForwardRequest{
		ClientRequest: &primarybackup.ForwardRequest_GetRequest{
			GetRequest: req,
		},
	}

  fmt.Println(forwardReq.GetGetRequest().GetFilePath())
  fmt.Println(handler.PBClient)
	_, err = handler.PBClient.Forward(ctx, forwardReq)
	if err != nil {
		return nil, err
	}

	reply, err = handler.Fs3processor.Get(req)

	return reply, err
}
