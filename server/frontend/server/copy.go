package server

import (
	"fmt"
	"io"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (server *FrontendServer) Copy(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token == "" {
		server.Logger.Println("get request missing token header")
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}
	filepath := r.URL.Query().Get("file-path")
	if filepath == "" {
		server.Logger.Println("request missing file-path query param on remove request")
		http.Error(w, "Missing file-path", http.StatusBadRequest)
		return
	}

	reqContentType := r.Header.Get("Content-type")
	if reqContentType != "text/plain" && reqContentType != "" {
		http.Error(w, "Expecting plain text body", http.StatusBadRequest)
		return
	}

	filecontent, err := io.ReadAll(r.Body) //json.NewDecoder(r.Body).Decode(&hasFileContentInstance)
	if err != nil {
		server.Logger.Printf("error decoding body: %s\n", err)
		http.Error(w, "Could not parse file content", http.StatusBadRequest)
		return
	}

	req := &fs3.CopyRequest{
		FilePath:    filepath,
		FileContent: filecontent,
		Token:       token,
	}

	server.VerifyFs3Client()
	res, err := server.Fs3Client.Copy(r.Context(), req)
	if err != nil {
		server.Logger.Printf("Error in sending request to primary: %s\n", err)
		errorCode := http.StatusBadRequest
		if res != nil {
			server.Logger.Println("res in not nil in receiving error from client")
			errorCode = GetFs3StatusHttpCode(res.GetStatus())
		}
		http.Error(w, fmt.Sprintf("Error in execution: %s", err), errorCode)
		return
	}
	server.Logger.Printf("successfully copied %s\n", filepath)
	w.Header().Add("Content-type", "application/json")
}
