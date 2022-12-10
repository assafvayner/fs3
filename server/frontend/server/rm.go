package server

import (
	"fmt"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (server *FrontendServer) Remove(w http.ResponseWriter, r *http.Request) {
	var err error
	token := r.Header.Get("token")
	if token == "" {
		server.Logger.Println("get request missing token header")
		http.Error(w, "Missing token", http.StatusBadRequest)
	}
	filepath := r.URL.Query().Get("file-path")
	if filepath == "" {
		server.Logger.Println("request missing file-path query param on remove request")
		http.Error(w, "Missing file-path", http.StatusBadRequest)
	}

	req := &fs3.RemoveRequest{
		FilePath: filepath,
		Token:    token,
	}

	server.VerifyFs3Client()
	res, err := server.Fs3Client.Remove(r.Context(), req)
	if err != nil {
		server.Logger.Printf("Error in sending request to primary: %s\n", err)
		errorCode := http.StatusBadRequest
		if res != nil {
			errorCode = GetFs3StatusHttpCode(res.GetStatus())
		}
		http.Error(w, fmt.Sprintf("Error in execution: %s", err), errorCode)
		return
	}
	server.Logger.Printf("successfully removed %s\n", filepath)
	w.Header().Add("Content-type", "application/json")
}
