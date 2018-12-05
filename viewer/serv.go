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

	// var mapreq map[string]interface{}
	// err := json.Unmarshal([]byte(req), &mapreq)
	// if err != nil {
	// 	return MakeGraphQLErrorResult(err)
	// }

	// querystr, ok := mapreq["query"].(string)
	// if !ok {
	// 	return MakeGraphQLErrorResult(ErrHTTPNoQuery)
	// }

	// mapval, ok1 := mapreq["variables"].(map[string]interface{})
	// if !ok1 {
	// 	mapval = nil
	// 	// return MakeGraphQLErrorResult(pb.CODE_HTTP_VARIABLE_ERR)
	// }

	// // curdb := s.anka.MgrDB.GetDB(ankadbname)
	// curctx := context.WithValue(r.Context(), interface{}("ankadb"), s.anka)

	// result1, err := s.anka.logic.OnQuery(curctx, querystr, mapval)
	// if err != nil {
	// 	return MakeGraphQLErrorResult(err)
	// }

	// return nil
}

func (s *HTTPServer) onGraphQL(w http.ResponseWriter, r *http.Request) {
	// if r.RequestURI == "/graphql" {
	result := s.procGraphQL(w, r)

	json.NewEncoder(w).Encode(result)
	// }
}

// HTTPServer -
func newHTTPServer() (*HTTPServer, error) {

	s := &HTTPServer{
		// anka: anka,
		// addr: anka.cfg.AddrHTTP,
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
