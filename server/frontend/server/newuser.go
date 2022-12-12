package server

import (
	"fmt"
	"net/http"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
)

func (server *FrontendServer) NewUser(w http.ResponseWriter, r *http.Request) {
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

	req := &authservice.NewUserRequest{
		Username: username,
		Password: password,
	}

	server.VerifyAuthClient()
	reply, err := server.AuthClient.NewUser(r.Context(), req)
	if err != nil {
		server.Logger.Printf("Error in request to authserver: %s\n", err)
		// TODO: determine if the error is a login error
		http.Error(w, fmt.Sprintf("Error forwarding request: %s", err), http.StatusInternalServerError)
		return
	}
	if !reply.GetStatus().GetSuccess() {
		server.Logger.Printf("authserver failed for user %s\n", username)
		http.Error(w, fmt.Sprintf("Error from authserver: %s", reply.GetStatus().GetMessage()), http.StatusInternalServerError)
		return
	}
}
