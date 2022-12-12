package fs3processor

import (
	"errors"
	"os"
	"path"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/jwtutils"
)

func (handler *Fs3RequestProcessor) Describe(req *fs3.DescribeRequest) (reply *fs3.DescribeReply, err error) {
	requestedPath := req.GetPath()
	reply = &fs3.DescribeReply{
		Path: requestedPath,
	}
	if !IsPathSafe(requestedPath) {
		handler.Logger.Printf("flagged describe request with illegal path: %s\n", requestedPath)
		reply.Status = fs3.Status_ILLEGAL_PATH
		err = errors.New("requested path is not allowed")
		return
	}

	username := jwtutils.GetUsernameFromTokenNoVerify(req.GetToken())
	serverPath := MakeServerSidePath(requestedPath, username)

	finfo, err := os.Stat(serverPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		handler.Logger.Printf("requested resource %s does not exist", serverPath)
		reply.Status = fs3.Status_NOT_FOUND
		err = errors.New("requested resource does not exist")
		return
	}
	if err != nil {
		handler.Logger.Printf("stat failed on path: %s err: %s\n", req.GetPath(), err)
		reply.Status = fs3.Status_INTERNAL_ERROR
		err = errors.New("failed to analyze requested resource")
		return
	}

	if finfo.Mode().IsRegular() {
		handler.Logger.Printf("Describe successfully %s is neither regular file or directory", serverPath)
		reply.Resource = &fs3.DescribeReply_File_{
			File: &fs3.DescribeReply_File{
				Filename: finfo.Name(),
			},
		}
		reply.Status = fs3.Status_GREAT_SUCCESS
		err = nil
		return
	}
	if !finfo.Mode().IsDir() {
		handler.Logger.Printf("user requested resource: %s is neither regular file or directory", serverPath)
		reply.Status = fs3.Status_INTERNAL_ERROR
		err = errors.New("requested resource is neither a file not a directory")
		return
	}

	// handle dir
	// assumed capacity of 10, just to not autoreallocate space for each append
	files := make([]string, 0, 10)
	subdirs := make([]string, 0, 10)

	dirents, err := os.ReadDir(serverPath)
	if err != nil {
		handler.Logger.Printf("error from os.Readdir for path: %s, err: %s", serverPath, err)
		reply.Status = fs3.Status_INTERNAL_ERROR
		err = errors.New("failed to look at directory")
		return
	}
	for _, dirent := range dirents {
		if dirent.Type().IsDir() {
			subdirs = append(subdirs, dirent.Name())
			continue
		}
		if dirent.Type().IsRegular() {
			files = append(files, dirent.Name())
		}
	}
	dirResource := &fs3.DescribeReply_Directory{
		Directoryname:  path.Base(requestedPath),
		Files:          files,
		Subdirectories: subdirs,
	}
	reply.Resource = &fs3.DescribeReply_Directory_{
		Directory: dirResource,
	}
	reply.Status = fs3.Status_GREAT_SUCCESS
	err = nil
	return
}
