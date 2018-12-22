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

func buildGraph(fn string) (*viewerdbpb.ViewerData, error) {
	fi, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	ng := &NoteGraph{}
	err = json.Unmarshal(fd, &ng)
	if err != nil {
		return nil, err
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

	return vd, nil
}

func buildJSON(fn string) (*viewerdbpb.ViewerData, error) {
	fi, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
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

	return vd, nil
}

func buildLine(userid string) (*viewerdbpb.ViewerData, error) {
	fi, err := os.Open("./test/" + userid + "_whackamole.json")
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	ubl := &DTUserBetList{}
	err = json.Unmarshal(fd, &ubl)
	if err != nil {
		return nil, err
	}

	ld := &viewerdbpb.LineData{
		XAxisType: "int32",
		YAxisType: "int64",
	}

	ldd := &viewerdbpb.LineSeriesData{
		Name: userid + "'s money",
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
		Title: userid + "'s whackamole",
		Token: userid,
		Data:  vjd,
	}

	return vd, nil
}

func buildMulti(mytoken string, tokens []string) (*viewerdbpb.ViewerData, error) {
	md := &viewerdbpb.MultiData{
		Tokens: tokens,
	}

	vjd := &viewerdbpb.ViewerData_Multi{
		Multi: md,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_MULTI,
		Title: "multi",
		Token: mytoken,
		Data:  vjd,
	}

	return vd, nil
}

func buildLines(mytoken string, tokens []string) (*viewerdbpb.ViewerData, error) {
	ld := &viewerdbpb.LineData{
		XAxisType: "int32",
		YAxisType: "int64",
	}

	for ci, v := range tokens {
		fi, err := os.Open("./test/" + v + "_whackamole.json")
		if err != nil {
			return nil, err
		}

		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		if err != nil {
			return nil, err
		}

		ubl := &DTUserBetList{}
		err = json.Unmarshal(fd, &ubl)
		if err != nil {
			return nil, err
		}

		ldd := &viewerdbpb.LineSeriesData{
			Name: v + "'s money",
		}

		for i, v := range ubl.Arr {
			if i > 200 {
				break
			}

			if ci == 0 {
				ld.XAxisInt32 = append(ld.XAxisInt32, int32(i))
			}

			ldd.ValInt64 = append(ldd.ValInt64, v.Destmoney)
		}

		ld.Series = append(ld.Series, ldd)
	}

	vjd := &viewerdbpb.ViewerData_Line{
		Line: ld,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_JSON,
		Title: "some user's whackamole",
		Token: mytoken,
		Data:  vjd,
	}

	return vd, nil
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
		vd, err := buildGraph("./test/graph.json")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "234567" {
		vd, err := buildJSON("./test/graph.json")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "10655" {
		vd, err := buildLine("10655")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "10656" {
		vd, err := buildLine("10656")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "10657" {
		vd, err := buildLine("10657")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline" {
		vd, err := buildMulti("multiline", []string{"10655", "10656", "10657"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "lines" {
		vd, err := buildLines("lines", []string{"10655", "10656", "10657"})
		if err != nil {
			return
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
