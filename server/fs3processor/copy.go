package fs3processor

import (
	"errors"
	fs3 "gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"os"
)

func (*Fs3RequestProcessor) Copy(req *fs3.CopyRequest) (reply *fs3.CopyReply, err error) {
	path := req.GetFilePath()
	reply = &fs3.CopyReply{
		FilePath: path,
	}
	if !IsPathSafe(path) {
		reply.Status = fs3.Status_ILLEGAL_PATH
		return reply, errors.New("Requested path is not allowed")
	}

	serverPath := MakeServerSidePath(path)
	pathToFile := GetDirFromFilePath(serverPath)
	if pathToFile != "" {
		err = os.MkdirAll(pathToFile, 0777)
		if err != nil {
			reply.Status = fs3.Status_INTERNAL_ERROR
			return reply, err
		}
	}

	f, err := os.OpenFile(serverPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		reply.Status = fs3.Status_INTERNAL_ERROR
		return reply, err
	}
	_, err = f.Write(req.GetFileContent())
	if err != nil {
		reply.Status = fs3.Status_INTERNAL_ERROR
		return reply, err
	}

	reply.Status = fs3.Status_GREAT_SUCCESS

	return reply, err
}
