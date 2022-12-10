package server

import (
	"fmt"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (server *FrontendServer) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	token := r.Header.Get("token")
	if token == "" {
		server.Logger.Println("get request missing token header")
		http.Error(w, "Missing token", http.StatusBadRequest)
	}
	filepath := r.URL.Query().Get("file-path")
	if filepath == "" {
		server.Logger.Println("get request missing file-path query param")
		http.Error(w, "Missing file-path", http.StatusBadRequest)
	}

	req := &fs3.GetRequest{
		FilePath: filepath,
		Token:    token,
	}

	server.VerifyFs3Client()
	res, err := server.Fs3Client.Get(r.Context(), req)
	if err != nil {
		server.Logger.Printf("Error in sending request to primary: %s\n", err)
		errorCode := http.StatusBadRequest
		if res != nil {
			errorCode = GetFs3StatusHttpCode(res.GetStatus())
		}
		http.Error(w, fmt.Sprintf("Error in execution: %s", err), errorCode)
	}

	responseObj := hasFileContent{
		FileContent: string(res.GetFileContent()),
	}
	w.Header().Add("Content-type", "text/plain")
	_, err = w.Write([]byte(responseObj.FileContent))
	if err != nil {
		server.Logger.Printf("error writing body: %s for get filepath %s\n", err, filepath)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
	server.Logger.Printf("succesfully processed get for %s\n", filepath)
}
