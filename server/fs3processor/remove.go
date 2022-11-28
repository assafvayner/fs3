package fs3processor

import (
	"errors"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"os"
)

func (handler *Fs3RequestProcessor) Remove(req *fs3.RemoveRequest) (reply *fs3.RemoveReply, err error) {
	path := req.GetFilePath()
	reply = &fs3.RemoveReply{
		FilePath: path,
	}
	if !IsPathSafe(path) {
		handler.Logger.Printf("Flagged remove request with illegal path: %s\n", req.GetFilePath())
		reply.Status = fs3.Status_ILLEGAL_PATH
		return reply, errors.New("Requested path is not allowed")
	}

	serverPath := MakeServerSidePath(path)

	if FileNotExists(serverPath) {
		handler.Logger.Printf("Remove request path: %s does not exist\n", req.GetFilePath())
		reply.Status = fs3.Status_NOT_FOUND
		return reply, errors.New("file not found")
	}

	err = os.Remove(serverPath)
	if err != nil {
		handler.Logger.Printf("Remove request path: %s failed to remove\n", req.GetFilePath())
		reply.Status = fs3.Status_INTERNAL_ERROR
		return reply, errors.New("could not remove file")
	}
	reply.Status = fs3.Status_GREAT_SUCCESS

	handler.Logger.Printf("Remove request path: %s success\n", req.GetFilePath())
	return reply, err
}
