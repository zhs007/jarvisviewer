package viewer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPServer -
type HTTPServer struct {
	addr string
	serv *http.Server
}

func (s *HTTPServer) procGraphQL(w http.ResponseWriter, r *http.Request) []byte {
	// ankadbname := r.Header.Get("Ankadbname")

	req, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	return req
}

func (s *HTTPServer) onGraphQL(w http.ResponseWriter, r *http.Request) {
	result := s.procGraphQL(w, r)

	json.NewEncoder(w).Encode(result)
}

// HTTPServer -
func newHTTPServer() (*HTTPServer, error) {

	s := &HTTPServer{
		serv: nil,
	}

	return s, nil
}

func (s *HTTPServer) start(ctx context.Context) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		s.onGraphQL(w, r)
	})

	fsh := http.FileServer(http.Dir("./www/graphiql"))
	mux.Handle("/graphiql/", http.StripPrefix("/graphiql/", fsh))

	server := &http.Server{
		Addr:         s.addr,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      mux,
	}

	s.serv = server

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *HTTPServer) stop() {
	if s.serv != nil {
		s.serv.Close()
	}

	return
}
