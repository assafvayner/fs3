package server

import (
	"fmt"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
)

func (server *FrontendServer) Login(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	password := r.Header.Get("password")
	if username == "" {
		http.Error(w, "missing username", http.StatusBadRequest)
    return
	}
	if password == "" {
		http.Error(w, "missing password", http.StatusBadRequest)
    return
	}
  prevToken := r.Header.Get("token")

  req := &authservice.GetNewTokenRequest{
    Username: username,
    Password: password,
    PreviousToken: prevToken,
  }

  server.VerifyAuthClient()
  reply, err := server.AuthClient.GetToken(r.Context(), req)
  if err != nil {
    // TODO determine user not authenticated
    server.Logger.Printf("authserver returned error for get token for user %s\n", username)
    http.Error(w, fmt.Sprintf("error from authserver: %s", err), http.StatusInternalServerError)
    return
  }

  if !reply.GetStatus().GetSuccess() {
    server.Logger.Printf("authserver get token not sucessful: %s\n", reply.GetStatus().GetMessage())
    http.Error(w, fmt.Sprintf("authserver get token not successful: %s", reply.GetStatus().GetMessage()), http.StatusInternalServerError)
    return    
  }
  w.Header().Add("token", reply.GetToken())
}
