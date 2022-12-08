package fs3processor

import (
	"errors"
	"os"

	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/jwtutils"
)

func (handler *Fs3RequestProcessor) Get(req *fs3.GetRequest) (reply *fs3.GetReply, err error) {
	path := req.GetFilePath()
	reply = &fs3.GetReply{
		FilePath: path,
	}
	if !IsPathSafe(path) {
		handler.Logger.Printf("Flagged get request with illegal path: %s\n", req.GetFilePath())
		reply.Status = fs3.Status_ILLEGAL_PATH
		return reply, errors.New("Requested path is not allowed")
	}

	// empty string if global scope
	username := jwtutils.GetUsernameFromTokenNoVerify(req.GetToken())
	serverPath := MakeServerSidePath(path, username)

	if FileNotExists(serverPath) {
		handler.Logger.Printf("Get request path: %s does not exist\n", req.GetFilePath())
		reply.Status = fs3.Status_NOT_FOUND
		return reply, errors.New("file not found")
	}

	content, err := os.ReadFile(serverPath)
	if err != nil {
		handler.Logger.Printf("Get request path: %s read error\n", req.GetFilePath())
		reply.Status = fs3.Status_INTERNAL_ERROR
		return reply, errors.New("could not read file")
	}
	reply.FileContent = content
	reply.Status = fs3.Status_GREAT_SUCCESS

	handler.Logger.Printf("Get request path: %s success\n", req.GetFilePath())
	return reply, err
}
