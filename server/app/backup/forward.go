package backup

import (
	"context"
	"errors"

	primarybackup "github.com/assafvayner/fs3/protos/primarybackup"
)

func (handler *BackupHandler) Forward(
	ctx context.Context,
	req *primarybackup.ForwardRequest,
) (reply *primarybackup.ForwardReply, err error) {
	reply = &primarybackup.ForwardReply{}

	fs3CopyRequest := req.GetCopyRequest()
	if fs3CopyRequest != nil {
		_, err = handler.Fs3processor.Copy(fs3CopyRequest)
		return reply, err
	}

	fs3RemoveRequest := req.GetRemoveRequest()
	if fs3RemoveRequest != nil {
		_, err = handler.Fs3processor.Remove(fs3RemoveRequest)
		return reply, err
	}

	fs3GetRequest := req.GetGetRequest()
	if fs3GetRequest != nil {
		_, err = handler.Fs3processor.Get(fs3GetRequest)
		return reply, err
	}

	handler.Logger.Println("Failed to process forward request, missing fs3 request")
	err = errors.New("Bad Forward request content")
	return reply, err
}
