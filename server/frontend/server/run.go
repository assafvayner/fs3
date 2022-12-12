package server

import (
	"fmt"
	"net/http"
)

func (server *FrontendServer) Run() error {
	if server.Port == 0 {
		server.Logger.Fatalln("failed to start, port may not be 0")
		return fmt.Errorf("port may not be 0")
	}

	mux := http.NewServeMux()
	mux.Handle("/copy", http.HandlerFunc(server.Copy))
	mux.Handle("/get", http.HandlerFunc(server.Get))
	mux.Handle("/remove", http.HandlerFunc(server.Remove))
	mux.Handle("/describe", http.HandlerFunc(server.Describe))
	mux.Handle("/newuser", http.HandlerFunc(server.NewUser))
	mux.Handle("/login", http.HandlerFunc(server.Login))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Port),
		Handler: mux,
	}
	return srv.ListenAndServe()
}
