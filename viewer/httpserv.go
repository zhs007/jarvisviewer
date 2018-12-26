package viewer

import (
	"context"
	"encoding/json"
	"fmt"
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

// DTUserPieNode -
type DTUserPieNode struct {
	Destmoney int32
	Nums      int
}

// DTUserPie -
type DTUserPie struct {
	Arr []*DTUserPieNode
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
		Type:  viewerdbpb.ViewerType_LINE,
		Title: userid + "'s whackamole",
		Token: userid,
		Data:  vjd,
	}

	return vd, nil
}

func buildPie(fn string) (*viewerdbpb.ViewerData, error) {
	fi, err := os.Open("./test/" + fn + ".json")
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	up := &DTUserPie{}
	err = json.Unmarshal(fd, &up)
	if err != nil {
		return nil, err
	}

	pd := &viewerdbpb.PieData{
		Name: "user money range",
	}

	for _, v := range up.Arr {
		pd.Data = append(pd.Data, &viewerdbpb.PieDataInfo{
			Name:     fmt.Sprintf("user money %v", v.Destmoney),
			ValInt32: int32(v.Nums),
		})
	}

	vjd := &viewerdbpb.ViewerData_Pie{
		Pie: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_PIE,
		Title: fn,
		Token: fn,
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
		Type:  viewerdbpb.ViewerType_LINE,
		Title: "some user's whackamole",
		Token: mytoken,
		Data:  vjd,
	}

	return vd, nil
}

func buildScatter(fn string, off int) (*viewerdbpb.ViewerData, error) {
	fi, err := os.Open("./test/" + fn + ".json")
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	up := &DTUserPie{}
	err = json.Unmarshal(fd, &up)
	if err != nil {
		return nil, err
	}

	pd := &viewerdbpb.ScatterData{
		XType: viewerdbpb.ValueType_INT32,
		YType: viewerdbpb.ValueType_INT32,
	}

	for _, v := range up.Arr {
		if len(pd.Data) == 0 {
			pd.Data = append(pd.Data, &viewerdbpb.ScatterNode{
				XInt32: v.Destmoney,
				YInt32: int32(v.Nums),
			})
		} else if v.Destmoney > pd.Data[len(pd.Data)-1].XInt32-int32(off) && v.Destmoney < pd.Data[len(pd.Data)-1].XInt32+int32(off) {
			pd.Data[len(pd.Data)-1].YInt32 = pd.Data[len(pd.Data)-1].YInt32 + int32(v.Nums)
		} else {
			pd.Data = append(pd.Data, &viewerdbpb.ScatterNode{
				XInt32: v.Destmoney,
				YInt32: int32(v.Nums),
			})
		}
	}

	vjd := &viewerdbpb.ViewerData_Scatter{
		Scatter: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_SCATTER,
		Title: fn + " scatter",
		Token: fn,
		Data:  vjd,
	}

	return vd, nil
}

func buildBoxplotNode(fn string, basemoney int) (*viewerdbpb.BoxplotNode, error) {
	fi, err := os.Open("./test/" + fn + ".json")
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	up := &DTUserPie{}
	err = json.Unmarshal(fd, &up)
	if err != nil {
		return nil, err
	}

	pd := &viewerdbpb.BoxplotNode{
		Name:    fn,
		ValType: viewerdbpb.ValueType_FLOAT32,
	}

	for _, v := range up.Arr {
		bm := v.Destmoney
		if bm < int32(200000-basemoney) {
			bm = 0
		} else {
			bm = bm - int32(200000-basemoney)
		}

		bmf := float32(bm)

		for i := 0; i < v.Nums; i++ {
			pd.ArrFloat32 = append(pd.ArrFloat32, bmf)
		}
	}

	return pd, nil

	// vjd := &viewerdbpb.ViewerData_Scatter{
	// 	Scatter: pd,
	// }

	// vd := &viewerdbpb.ViewerData{
	// 	Type:  viewerdbpb.ViewerType_SCATTER,
	// 	Title: fn + " scatter",
	// 	Token: fn,
	// 	Data:  vjd,
	// }

	// return vd, nil
}

func buildBoxplo(arrname []string, arrbm []int) (*viewerdbpb.ViewerData, error) {
	pd := &viewerdbpb.BoxplotData{}

	for i, n := range arrname {
		bn, _ := buildBoxplotNode(n, arrbm[i])
		pd.Data = append(pd.Data, bn)
	}

	vjd := &viewerdbpb.ViewerData_Boxplot{
		Boxplot: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_BOXPLOT,
		Title: "user money boxplot",
		Token: "usermoneyboxplot",
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
	} else if token == "pie_10whackamole" {
		vd, err := buildPie("pie_10whackamole")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50whackamole" {
		vd, err := buildPie("pie_50whackamole")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100whackamole" {
		vd, err := buildPie("pie_100whackamole")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10whackamole_scatter" {
		vd, err := buildScatter("pie_10whackamole", 20000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50whackamole_scatter" {
		vd, err := buildScatter("pie_50whackamole", 20000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100whackamole_scatter" {
		vd, err := buildScatter("pie_100whackamole", 20000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot" {
		vd, err := buildBoxplo([]string{"pie_10whackamole", "pie_50whackamole", "pie_100whackamole"}, []int{350 * 10, 350 * 50, 350 * 100})
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

	fsh := http.FileServer(http.Dir("./www/static"))
	mux.Handle("/", http.StripPrefix("/", fsh))

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
