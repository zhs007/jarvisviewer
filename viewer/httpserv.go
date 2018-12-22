package viewer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/zhs007/jarvisviewer/viewerdb/proto"
)

// NoteGraphKey -
type NoteGraphKey struct {
	ID    int
	Key   string
	Times int
}

// NoteGraphLink -
type NoteGraphLink struct {
	Src   int
	Dest  int
	Times int
}

// NoteGraph -
type NoteGraph struct {
	Keys  []*NoteGraphKey
	Links []*NoteGraphLink
}

// DTUserBetListNode -
type DTUserBetListNode struct {
	Destmoney int64
}

// DTUserBetList -
type DTUserBetList struct {
	Arr []*DTUserBetListNode
}

// HTTPServer -
type HTTPServer struct {
	addr string
	serv *http.Server
}

// func (s *HTTPServer) procGraphQL(w http.ResponseWriter, r *http.Request) []byte {
// 	// ankadbname := r.Header.Get("Ankadbname")

// 	req, _ := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	return req
// }

func (s *HTTPServer) onViewerData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	token := r.URL.Query().Get("token")
	if token == "123456" {
		fi, err := os.Open("./test/graph.json")
		if err != nil {
			return
		}

		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		if err != nil {
			return
		}

		ng := &NoteGraph{}
		err = json.Unmarshal(fd, &ng)
		if err != nil {
			return
		}

		gd := &viewerdbpb.GraphData{}

		for _, v := range ng.Keys {
			k := &viewerdbpb.GraphNode{
				Id:       int32(v.ID),
				Name:     v.Key,
				Size:     int32(v.Times),
				Category: "Category0",
			}

			gd.Nodes = append(gd.Nodes, k)
		}

		for _, v := range ng.Links {
			k := &viewerdbpb.GraphLink{
				Src:  int32(v.Src),
				Dest: int32(v.Dest),
				Size: int32(v.Times),
			}

			gd.Links = append(gd.Links, k)
		}

		vgd := &viewerdbpb.ViewerData_Graph{
			Graph: gd,
		}

		vd := &viewerdbpb.ViewerData{
			Type:  viewerdbpb.ViewerType_GRAPH,
			Title: "知识库关系图",
			Token: "123456",
			Data:  vgd,
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "234567" {
		fi, err := os.Open("./test/graph.json")
		if err != nil {
			return
		}

		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		if err != nil {
			return
		}

		jd := &viewerdbpb.JSONData{}
		jd.Str = string(fd)

		vjd := &viewerdbpb.ViewerData_Json{
			Json: jd,
		}

		vd := &viewerdbpb.ViewerData{
			Type:  viewerdbpb.ViewerType_JSON,
			Title: "知识库",
			Token: "234567",
			Data:  vjd,
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "10655" {
		fi, err := os.Open("./test/10655_whackamole.json")
		if err != nil {
			return
		}

		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		if err != nil {
			return
		}

		ubl := &DTUserBetList{}
		err = json.Unmarshal(fd, &ubl)
		if err != nil {
			return
		}

		ld := &viewerdbpb.LineData{
			XAxisType: "int32",
			YAxisType: "int64",
		}

		ldd := &viewerdbpb.LineSeriesData{
			Name: "10655's money",
		}

		for i, v := range ubl.Arr {
			ld.XAxisInt32 = append(ld.XAxisInt32, int32(i))
			ldd.ValInt64 = append(ldd.ValInt64, v.Destmoney)
		}

		ld.Series = append(ld.Series, ldd)

		vjd := &viewerdbpb.ViewerData_Line{
			Line: ld,
		}

		vd := &viewerdbpb.ViewerData{
			Type:  viewerdbpb.ViewerType_JSON,
			Title: "10655's whackamole",
			Token: "10655",
			Data:  vjd,
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	}
}

// HTTPServer -
func newHTTPServer(addr string) (*HTTPServer, error) {
	s := &HTTPServer{
		addr: addr,
		serv: nil,
	}

	return s, nil
}

func (s *HTTPServer) start(ctx context.Context) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
		s.onViewerData(w, r)
	})

	// fsh := http.FileServer(http.Dir("./www/graphiql"))
	// mux.Handle("/graphiql/", http.StripPrefix("/graphiql/", fsh))

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
