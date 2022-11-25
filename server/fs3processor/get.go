package fs3processor

import (
	"errors"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"os"
)

func (*Fs3RequestProcessor) Get(req *fs3.GetRequest) (reply *fs3.GetReply, err error) {
	path := req.GetFilePath()
	reply = &fs3.GetReply{
		FilePath: path,
	}
	if !IsPathSafe(path) {
		reply.Status = fs3.Status_ILLEGAL_PATH
		return reply, errors.New("Requested path is not allowed")
	}

	serverPath := MakeServerSidePath(path)

	if FileNotExists(serverPath) {
		reply.Status = fs3.Status_NOT_FOUND
		return reply, errors.New("file not found")
	}

	content, err := os.ReadFile(serverPath)
	if err != nil {
		reply.Status = fs3.Status_INTERNAL_ERROR
		return reply, errors.New("could not read file")
	}
	reply.FileContent = content
	reply.Status = fs3.Status_GREAT_SUCCESS

	return reply, err
}
