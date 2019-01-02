package viewer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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

func buildDTGame6lovefighters() (*viewerdbpb.ViewerData, error) {
	row0 := &viewerdbpb.DTGame6List{
		Arr: []int32{7, 2, 6},
	}

	row1 := &viewerdbpb.DTGame6List{
		Arr: []int32{3, 7, 0, 5},
	}

	row2 := &viewerdbpb.DTGame6List{
		Arr: []int32{3, 6, 2, 2},
	}

	row3 := &viewerdbpb.DTGame6List{
		Arr: []int32{7, 3, 8, 2, 3},
	}

	row4 := &viewerdbpb.DTGame6List{
		Arr: []int32{3, 2, 7, 1, 5},
	}

	dtgame6 := &viewerdbpb.DTGame6Data{
		GameCode:   "lovefighters",
		GameModule: "basegame",
		ArrType:    viewerdbpb.DTGame6ArrayType_COLUMN,
		Arr:        []*viewerdbpb.DTGame6List{row0, row1, row2, row3, row4},
		Bet:        100,
		RealBet:    2500,
		Win:        1000,
		RealWin:    1200,
	}

	rs0 := &viewerdbpb.DTGame6Result{
		Type:       "line",
		Bet:        100,
		RealBet:    2500,
		BaseMul:    10,
		Win:        1000,
		RealWin:    1200,
		BonusPrize: 200,
		OtherMul:   0,
		Pos: []*viewerdbpb.DTGame6Pos{
			&viewerdbpb.DTGame6Pos{
				X: 0,
				Y: 1,
			},
			&viewerdbpb.DTGame6Pos{
				X: 1,
				Y: 2,
			},
			&viewerdbpb.DTGame6Pos{
				X: 2,
				Y: 2,
			},
			&viewerdbpb.DTGame6Pos{
				X: 3,
				Y: 3,
			},
		},
	}

	rs1 := &viewerdbpb.DTGame6Result{
		Type:       "scatter",
		Bet:        100,
		RealBet:    2500,
		BaseMul:    10,
		Win:        1000,
		RealWin:    1200,
		BonusPrize: 200,
		OtherMul:   0,
		Pos: []*viewerdbpb.DTGame6Pos{
			&viewerdbpb.DTGame6Pos{
				X: 0,
				Y: 1,
			},
			&viewerdbpb.DTGame6Pos{
				X: 1,
				Y: 2,
			},
			&viewerdbpb.DTGame6Pos{
				X: 2,
				Y: 3,
			},
			&viewerdbpb.DTGame6Pos{
				X: 3,
				Y: 3,
			},
		},
	}

	dtgame6.Results = append(dtgame6.Results, rs0)
	dtgame6.Results = append(dtgame6.Results, rs1)

	dtgi := &viewerdbpb.ViewerData_Dtgame6{
		Dtgame6: dtgame6,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_DTGAME6,
		Title: "lovefighters game result",
		Token: "lovefighters001",
		Data:  dtgi,
	}

	return vd, nil
}

func buildDTGame6(gamecode string) (*viewerdbpb.ViewerData, error) {
	row0 := &viewerdbpb.DTGame6List{
		Arr: []int32{8, 2, 3, 8, 7},
	}

	row1 := &viewerdbpb.DTGame6List{
		Arr: []int32{10, 10, 8, 1, 3},
	}

	row2 := &viewerdbpb.DTGame6List{
		Arr: []int32{10, 5, 6, 1, 6},
	}

	dtgame6 := &viewerdbpb.DTGame6Data{
		GameCode:   gamecode,
		GameModule: "basegame",
		ArrType:    viewerdbpb.DTGame6ArrayType_ROW,
		Arr:        []*viewerdbpb.DTGame6List{row0, row1, row2},
		Bet:        100,
		RealBet:    2500,
		Win:        1000,
		RealWin:    1200,
	}

	rs0 := &viewerdbpb.DTGame6Result{
		Type:       "line",
		Bet:        100,
		RealBet:    2500,
		BaseMul:    10,
		Win:        1000,
		RealWin:    1200,
		BonusPrize: 200,
		OtherMul:   0,
		Pos: []*viewerdbpb.DTGame6Pos{
			&viewerdbpb.DTGame6Pos{
				X: 0,
				Y: 1,
			},
			&viewerdbpb.DTGame6Pos{
				X: 1,
				Y: 2,
			},
			&viewerdbpb.DTGame6Pos{
				X: 2,
				Y: 0,
			},
		},
	}

	rs1 := &viewerdbpb.DTGame6Result{
		Type:       "scatter",
		Bet:        100,
		RealBet:    2500,
		BaseMul:    10,
		Win:        1000,
		RealWin:    1200,
		BonusPrize: 200,
		OtherMul:   0,
		Pos: []*viewerdbpb.DTGame6Pos{
			&viewerdbpb.DTGame6Pos{
				X: 0,
				Y: 2,
			},
			&viewerdbpb.DTGame6Pos{
				X: 2,
				Y: 0,
			},
			&viewerdbpb.DTGame6Pos{
				X: 4,
				Y: 0,
			},
		},
	}

	dtgame6.Results = append(dtgame6.Results, rs0)
	dtgame6.Results = append(dtgame6.Results, rs1)

	dtgi := &viewerdbpb.ViewerData_Dtgame6{
		Dtgame6: dtgame6,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_DTGAME6,
		Title: "clash game result",
		Token: "clash001",
		Data:  dtgi,
	}

	return vd, nil
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

func insMapPieNode(mappn map[int32]*DTUserPieNode, n *DTUserPieNode, off int32) {
	dm := n.Destmoney - n.Destmoney%off
	for k, v := range mappn {
		if n.Destmoney >= k && n.Destmoney < k+off {
			v.Nums = v.Nums + n.Nums

			return
		}
	}

	mappn[dm] = n
}

func buildPie(fn string, off int) (*viewerdbpb.ViewerData, error) {
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

	mapnums := make(map[int32]*DTUserPieNode)

	for _, v := range up.Arr {
		insMapPieNode(mapnums, v, int32(off))
	}

	var arrnums []*DTUserPieNode
	for _, v := range mapnums {
		arrnums = append(arrnums, v)
	}

	sort.Slice(arrnums, func(i, j int) bool {
		if arrnums[i].Destmoney < arrnums[j].Destmoney {
			return true
		}

		return false
	})

	for _, v := range arrnums {
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

	mapnums := make(map[int32]*DTUserPieNode)

	for _, v := range up.Arr {
		insMapPieNode(mapnums, v, int32(off))
	}

	for _, v := range mapnums {
		pd.Data = append(pd.Data, &viewerdbpb.ScatterNode{
			XInt32: v.Destmoney,
			YInt32: int32(v.Nums),
		})
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

func buildScatter2(token string, lstfn []string, off int) (*viewerdbpb.ViewerData, error) {
	vjd := &viewerdbpb.ViewerData_Scatter2{
		Scatter2: &viewerdbpb.Scatter2Data{},
	}

	for _, v := range lstfn {
		fi, err := os.Open("./test/" + v + ".json")
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
			Name:  v,
		}

		mapnums := make(map[int32]*DTUserPieNode)

		for _, v := range up.Arr {
			insMapPieNode(mapnums, v, int32(off))
		}

		for _, v := range mapnums {
			pd.Data = append(pd.Data, &viewerdbpb.ScatterNode{
				XInt32: v.Destmoney,
				YInt32: int32(v.Nums),
			})
		}

		vjd.Scatter2.Lst = append(vjd.Scatter2.Lst, pd)
	}

	// vjd := &viewerdbpb.ViewerData_Scatter{
	// 	Scatter: pd,
	// }

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_SCATTER,
		Title: token,
		Token: token,
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
		vd, err := buildPie("pie_10whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20whackamole" {
		vd, err := buildPie("pie_20whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30whackamole" {
		vd, err := buildPie("pie_30whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40whackamole" {
		vd, err := buildPie("pie_40whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50whackamole" {
		vd, err := buildPie("pie_50whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60whackamole" {
		vd, err := buildPie("pie_60whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70whackamole" {
		vd, err := buildPie("pie_70whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80whackamole" {
		vd, err := buildPie("pie_80whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90whackamole" {
		vd, err := buildPie("pie_90whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100whackamole" {
		vd, err := buildPie("pie_100whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150whackamole" {
		vd, err := buildPie("pie_150whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200whackamole" {
		vd, err := buildPie("pie_200whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10whackamole_scatter" {
		vd, err := buildScatter("pie_10whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20whackamole_scatter" {
		vd, err := buildScatter("pie_20whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30whackamole_scatter" {
		vd, err := buildScatter("pie_30whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40whackamole_scatter" {
		vd, err := buildScatter("pie_40whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50whackamole_scatter" {
		vd, err := buildScatter("pie_50whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60whackamole_scatter" {
		vd, err := buildScatter("pie_60whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70whackamole_scatter" {
		vd, err := buildScatter("pie_70whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80whackamole_scatter" {
		vd, err := buildScatter("pie_80whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90whackamole_scatter" {
		vd, err := buildScatter("pie_90whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100whackamole_scatter" {
		vd, err := buildScatter("pie_100whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150whackamole_scatter" {
		vd, err := buildScatter("pie_150whackamole", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200whackamole_scatter" {
		vd, err := buildScatter("pie_200whackamole", 100)
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
	} else if token == "clash001" {
		vd, err := buildDTGame6("clash")
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "lovefighters001" {
		vd, err := buildDTGame6lovefighters()
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline2" {
		vd, err := buildMulti("multiline2", []string{"pie_10whackamole_scatter", "pie_20whackamole_scatter", "pie_30whackamole_scatter",
			"pie_40whackamole_scatter", "pie_50whackamole_scatter", "pie_60whackamole_scatter", "pie_70whackamole_scatter", "pie_80whackamole_scatter", "pie_90whackamole_scatter",
			"pie_100whackamole_scatter", "pie_150whackamole_scatter", "pie_200whackamole_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline3" {
		vd, err := buildMulti("multiline3", []string{"pie_10whackamole", "pie_20whackamole", "pie_30whackamole",
			"pie_40whackamole", "pie_50whackamole", "pie_60whackamole", "pie_70whackamole", "pie_80whackamole", "pie_90whackamole",
			"pie_100whackamole", "pie_150whackamole", "pie_200whackamole"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10magician" {
		vd, err := buildPie("pie_10magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20magician" {
		vd, err := buildPie("pie_20magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30magician" {
		vd, err := buildPie("pie_30magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40magician" {
		vd, err := buildPie("pie_40magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50magician" {
		vd, err := buildPie("pie_50magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60magician" {
		vd, err := buildPie("pie_60magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70magician" {
		vd, err := buildPie("pie_70magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80magician" {
		vd, err := buildPie("pie_80magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90magician" {
		vd, err := buildPie("pie_90magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100magician" {
		vd, err := buildPie("pie_100magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150magician" {
		vd, err := buildPie("pie_150magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200magician" {
		vd, err := buildPie("pie_200magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10magician_scatter" {
		vd, err := buildScatter("pie_10magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20magician_scatter" {
		vd, err := buildScatter("pie_20magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30magician_scatter" {
		vd, err := buildScatter("pie_30magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40magician_scatter" {
		vd, err := buildScatter("pie_40magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50magician_scatter" {
		vd, err := buildScatter("pie_50magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60magician_scatter" {
		vd, err := buildScatter("pie_60magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70magician_scatter" {
		vd, err := buildScatter("pie_70magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80magician_scatter" {
		vd, err := buildScatter("pie_80magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90magician_scatter" {
		vd, err := buildScatter("pie_90magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100magician_scatter" {
		vd, err := buildScatter("pie_100magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150magician_scatter" {
		vd, err := buildScatter("pie_150magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200magician_scatter" {
		vd, err := buildScatter("pie_200magician", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline5" {
		vd, err := buildMulti("multiline5", []string{"pie_10magician_scatter", "pie_20magician_scatter", "pie_30magician_scatter",
			"pie_40magician_scatter", "pie_50magician_scatter", "pie_60magician_scatter", "pie_70magician_scatter", "pie_80magician_scatter", "pie_90magician_scatter",
			"pie_100magician_scatter", "pie_150magician_scatter", "pie_200magician_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline6" {
		vd, err := buildMulti("multiline6", []string{"pie_10magician", "pie_20magician", "pie_30magician",
			"pie_40magician", "pie_50magician", "pie_60magician", "pie_70magician", "pie_80magician", "pie_90magician",
			"pie_100magician", "pie_150magician", "pie_200magician"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10dragonball" {
		vd, err := buildPie("pie_10dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20dragonball" {
		vd, err := buildPie("pie_20dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30dragonball" {
		vd, err := buildPie("pie_30dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40dragonball" {
		vd, err := buildPie("pie_40dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50dragonball" {
		vd, err := buildPie("pie_50dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60dragonball" {
		vd, err := buildPie("pie_60dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70dragonball" {
		vd, err := buildPie("pie_70dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80dragonball" {
		vd, err := buildPie("pie_80dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90dragonball" {
		vd, err := buildPie("pie_90dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100dragonball" {
		vd, err := buildPie("pie_100dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150dragonball" {
		vd, err := buildPie("pie_150dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200dragonball" {
		vd, err := buildPie("pie_200dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10dragonball_scatter" {
		vd, err := buildScatter("pie_10dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20dragonball_scatter" {
		vd, err := buildScatter("pie_20dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30dragonball_scatter" {
		vd, err := buildScatter("pie_30dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40dragonball_scatter" {
		vd, err := buildScatter("pie_40dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50dragonball_scatter" {
		vd, err := buildScatter("pie_50dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60dragonball_scatter" {
		vd, err := buildScatter("pie_60dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70dragonball_scatter" {
		vd, err := buildScatter("pie_70dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80dragonball_scatter" {
		vd, err := buildScatter("pie_80dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90dragonball_scatter" {
		vd, err := buildScatter("pie_90dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100dragonball_scatter" {
		vd, err := buildScatter("pie_100dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150dragonball_scatter" {
		vd, err := buildScatter("pie_150dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200dragonball_scatter" {
		vd, err := buildScatter("pie_200dragonball", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline7" {
		vd, err := buildMulti("multiline7", []string{"pie_10dragonball_scatter", "pie_20dragonball_scatter", "pie_30dragonball_scatter",
			"pie_40dragonball_scatter", "pie_50dragonball_scatter", "pie_60dragonball_scatter", "pie_70dragonball_scatter", "pie_80dragonball_scatter", "pie_90dragonball_scatter",
			"pie_100dragonball_scatter", "pie_150dragonball_scatter", "pie_200dragonball_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline8" {
		vd, err := buildMulti("multiline8", []string{"pie_10dragonball", "pie_20dragonball", "pie_30dragonball",
			"pie_40dragonball", "pie_50dragonball", "pie_60dragonball", "pie_70dragonball", "pie_80dragonball", "pie_90dragonball",
			"pie_100dragonball", "pie_150dragonball", "pie_200dragonball"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10wrathofthor" {
		vd, err := buildPie("pie_10wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20wrathofthor" {
		vd, err := buildPie("pie_20wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30wrathofthor" {
		vd, err := buildPie("pie_30wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40wrathofthor" {
		vd, err := buildPie("pie_40wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50wrathofthor" {
		vd, err := buildPie("pie_50wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60wrathofthor" {
		vd, err := buildPie("pie_60wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70wrathofthor" {
		vd, err := buildPie("pie_70wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80wrathofthor" {
		vd, err := buildPie("pie_80wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90wrathofthor" {
		vd, err := buildPie("pie_90wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100wrathofthor" {
		vd, err := buildPie("pie_100wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150wrathofthor" {
		vd, err := buildPie("pie_150wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200wrathofthor" {
		vd, err := buildPie("pie_200wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10wrathofthor_scatter" {
		vd, err := buildScatter("pie_10wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20wrathofthor_scatter" {
		vd, err := buildScatter("pie_20wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30wrathofthor_scatter" {
		vd, err := buildScatter("pie_30wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40wrathofthor_scatter" {
		vd, err := buildScatter("pie_40wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50wrathofthor_scatter" {
		vd, err := buildScatter("pie_50wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60wrathofthor_scatter" {
		vd, err := buildScatter("pie_60wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70wrathofthor_scatter" {
		vd, err := buildScatter("pie_70wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80wrathofthor_scatter" {
		vd, err := buildScatter("pie_80wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90wrathofthor_scatter" {
		vd, err := buildScatter("pie_90wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100wrathofthor_scatter" {
		vd, err := buildScatter("pie_100wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150wrathofthor_scatter" {
		vd, err := buildScatter("pie_150wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200wrathofthor_scatter" {
		vd, err := buildScatter("pie_200wrathofthor", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline9" {
		vd, err := buildMulti("multiline9", []string{"pie_10wrathofthor_scatter", "pie_20wrathofthor_scatter",
			"pie_30wrathofthor_scatter", "pie_40wrathofthor_scatter", "pie_50wrathofthor_scatter",
			"pie_60wrathofthor_scatter", "pie_70wrathofthor_scatter", "pie_80wrathofthor_scatter",
			"pie_90wrathofthor_scatter", "pie_100wrathofthor_scatter", "pie_150wrathofthor_scatter",
			"pie_200wrathofthor_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline10" {
		vd, err := buildMulti("multiline10", []string{"pie_10wrathofthor", "pie_20wrathofthor", "pie_30wrathofthor",
			"pie_40wrathofthor", "pie_50wrathofthor", "pie_60wrathofthor", "pie_70wrathofthor", "pie_80wrathofthor",
			"pie_90wrathofthor", "pie_100wrathofthor", "pie_150wrathofthor", "pie_200wrathofthor"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "10_scatter" {
		vd, err := buildScatter2("10_scatter", []string{"pie_10whackamole", "pie_10magician", "pie_10dragonball", "pie_10wrathofthor"}, 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "20_scatter" {
		vd, err := buildScatter2("20_scatter", []string{"pie_20whackamole", "pie_20magician", "pie_20dragonball", "pie_20wrathofthor"}, 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "30_scatter" {
		vd, err := buildScatter2("30_scatter", []string{"pie_30whackamole", "pie_30magician", "pie_30dragonball", "pie_30wrathofthor"}, 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline100" {
		vd, err := buildMulti("multiline100", []string{"10_scatter", "20_scatter", "30_scatter"})
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
