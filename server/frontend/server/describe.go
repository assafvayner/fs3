package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/fs3"
)

func (server *FrontendServer) Describe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token == "" {
		server.Logger.Println("describe request missing token header")
		http.Error(w, "Missing token", http.StatusBadRequest)
    return
	}
  path := r.Header.Get("path")
  if path == "" {
		server.Logger.Println("describe request missing path header")
		http.Error(w, "Missing path", http.StatusBadRequest)
		return
	}

  req := &fs3.DescribeRequest{
    Path: path,
    Token: token,
  }
  server.VerifyFs3Client()
  res, err := server.Fs3Client.Describe(r.Context(), req)
  if err != nil {
    server.Logger.Printf("Error from server on describe for path %s: %s", path, err)
   	errorCode := http.StatusBadRequest
		if res != nil {
			errorCode = GetFs3StatusHttpCode(res.GetStatus())
		}
		http.Error(w, fmt.Sprintf("Error in execution: %s", err), errorCode)
    return
  }

  w.Header().Add("Content-type", "application/json")
  err = json.NewEncoder(w).Encode(res)
  if err != nil {
    server.Logger.Printf("error writing json on describe path %s: %s\n", path, err)
    http.Error(w, "failed to write json response", http.StatusInternalServerError)
    return
  }

  server.Logger.Printf("successfully described path: %s", path)
}