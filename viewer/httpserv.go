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

func buildScatter(fn string, off int, mul float32, bm int) (*viewerdbpb.ViewerData, error) {
	// mul = 0.0

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
		// v.Destmoney = int32(bm) + int32(float32(v.Destmoney-int32(bm))*mul)

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

func buildScatter2(token string, lstfn []string, off int, mul []float32, bm int) (*viewerdbpb.ViewerData, error) {
	vjd := &viewerdbpb.ViewerData_Scatter2{
		Scatter2: &viewerdbpb.Scatter2Data{},
	}

	for i, v := range lstfn {
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
			v.Destmoney = int32(bm) + int32(float32(v.Destmoney-int32(bm))*mul[i])

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
		Type:  viewerdbpb.ViewerType_SCATTER2,
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
		// if bm < int32(200000-basemoney) {
		// 	bm = 0
		// } else {
		// 	bm = bm - int32(200000-basemoney)
		// }

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

type boxplot2 struct {
	fn []string
}

func buildBoxplot2(arr []boxplot2) (*viewerdbpb.ViewerData, error) {
	pd := &viewerdbpb.Boxplot2Data{}

	for _, v := range arr {
		pd1 := &viewerdbpb.BoxplotData{}

		for _, n := range v.fn {
			bn, _ := buildBoxplotNode(n, 0)
			pd1.Data = append(pd1.Data, bn)
		}

		pd.Lst = append(pd.Lst, pd1)
	}

	vjd := &viewerdbpb.ViewerData_Boxplot2{
		Boxplot2: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_BOXPLOT2,
		Title: "user money boxplot 2",
		Token: "usermoneyboxplot2",
		Data:  vjd,
	}

	return vd, nil
}

type boxplot3 struct {
	name string
	fn   []string
}

func buildDataset(fn string) (*viewerdbpb.Dataset, error) {
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

	pd := &viewerdbpb.Dataset{
		Name:    fn,
		ValType: viewerdbpb.ValueType_FLOAT32,
	}

	for _, v := range up.Arr {
		bm := v.Destmoney

		bmf := float32(bm)

		for i := 0; i < v.Nums; i++ {
			pd.ArrFloat32 = append(pd.ArrFloat32, bmf)
		}
	}

	return pd, nil
}

func buildBoxplot3(arr []boxplot3, xval []string, datasetname string) (*viewerdbpb.ViewerData, error) {
	pd := &viewerdbpb.Boxplot3Data{}

	if datasetname == "" {
		for _, v := range arr {
			b3 := &viewerdbpb.Boxplot3{
				Category: v.name,
			}

			for _, n := range v.fn {
				b3.DatasetName = append(b3.DatasetName, n)
			}

			pd.Boxplots = append(pd.Boxplots, b3)
		}

		pd.XVal = &viewerdbpb.Dataset{
			Name:    "xVal",
			ValType: viewerdbpb.ValueType_STRING,
		}

		for _, v := range xval {
			pd.XVal.ArrString = append(pd.XVal.ArrString, v)
		}

		vjd := &viewerdbpb.ViewerData_Boxplot3{
			Boxplot3: pd,
		}

		vd := &viewerdbpb.ViewerData{
			Type:  viewerdbpb.ViewerType_BOXPLOT3,
			Title: "user money boxplot 3",
			Token: "usermoneyboxplot100",
			Data:  vjd,
		}

		return vd, nil
	}

	bn, _ := buildDataset(datasetname)
	pd.Datasets = make(map[string]*viewerdbpb.Dataset)
	pd.Datasets[datasetname] = bn

	vjd := &viewerdbpb.ViewerData_Boxplot3{
		Boxplot3: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_BOXPLOT3,
		Title: "user money boxplot 3",
		Token: "usermoneyboxplot10",
		Data:  vjd,
	}

	return vd, nil
}

type scatter3 struct {
	name string
	fn   []string
}

func buildDataset2D(fn string, off int) (*viewerdbpb.Dataset2D, error) {
	// vjd := &viewerdbpb.ViewerData_Scatter2{
	// 	Scatter2: &viewerdbpb.Scatter2Data{},
	// }

	// for _, v := range lstfn {
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

	pd := &viewerdbpb.Dataset2D{
		XType: viewerdbpb.ValueType_INT32,
		YType: viewerdbpb.ValueType_INT32,
		Name:  fn,
	}

	mapnums := make(map[int32]*DTUserPieNode)

	for _, v := range up.Arr {
		insMapPieNode(mapnums, v, int32(off))
	}

	for _, v := range mapnums {
		pd.XArrInt32 = append(pd.XArrInt32, v.Destmoney)
		pd.YArrInt32 = append(pd.YArrInt32, int32(v.Nums))
	}

	// vjd.Scatter2.Lst = append(vjd.Scatter2.Lst, pd)
	// }

	// vjd := &viewerdbpb.ViewerData_Scatter{
	// 	Scatter: pd,
	// }

	// vd := &viewerdbpb.ViewerData{
	// 	Type:  viewerdbpb.ViewerType_SCATTER2,
	// 	Title: token,
	// 	Token: token,
	// 	Data:  vjd,
	// }

	return pd, nil
}

const maxint32 = int32(^uint32(0) >> 1)
const minint32 = ^maxint32

func countDataset2DRange(fn string, off int) (int32, int32, int32, int32, error) {
	fi, err := os.Open("./test/" + fn + ".json")
	if err != nil {
		return minint32, maxint32, minint32, maxint32, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return minint32, maxint32, minint32, maxint32, err
	}

	up := &DTUserPie{}
	err = json.Unmarshal(fd, &up)
	if err != nil {
		return minint32, maxint32, minint32, maxint32, err
	}

	xmin := maxint32
	ymin := maxint32
	xmax := minint32
	ymax := minint32

	// pd := &viewerdbpb.Dataset2D{
	// 	XType: viewerdbpb.ValueType_INT32,
	// 	YType: viewerdbpb.ValueType_INT32,
	// 	Name:  fn,
	// }

	// mapnums := make(map[int32]*DTUserPieNode)

	for _, v := range up.Arr {
		if v.Destmoney < xmin {
			xmin = v.Destmoney
		}

		if v.Destmoney > xmax {
			xmax = v.Destmoney
		}

		if int32(v.Nums) < ymin {
			ymin = int32(v.Nums)
		}

		if int32(v.Nums) > ymax {
			ymax = int32(v.Nums)
		}

		// insMapPieNode(mapnums, v, int32(off))
	}

	// for _, v := range mapnums {
	// 	pd.XArrInt32 = append(pd.XArrInt32, v.Destmoney)
	// 	pd.YArrInt32 = append(pd.YArrInt32, int32(v.Nums))
	// }

	// vjd.Scatter2.Lst = append(vjd.Scatter2.Lst, pd)
	// }

	// vjd := &viewerdbpb.ViewerData_Scatter{
	// 	Scatter: pd,
	// }

	// vd := &viewerdbpb.ViewerData{
	// 	Type:  viewerdbpb.ViewerType_SCATTER2,
	// 	Title: token,
	// 	Token: token,
	// 	Data:  vjd,
	// }

	return xmin, xmax, ymin, ymax, nil
}

func buildScatter3(token string, arr []scatter3, zval []string, datasetname string, off int) (*viewerdbpb.ViewerData, error) {
	pd := &viewerdbpb.Scatter3Data{}

	if datasetname == "" {
		xmin := maxint32
		ymin := maxint32
		xmax := minint32
		ymax := minint32

		for _, v := range arr {
			b3 := &viewerdbpb.Scatter3{
				Category: v.name,
			}

			for _, n := range v.fn {
				b3.DatasetName = append(b3.DatasetName, n)

				cxmin, cxmax, cymin, cymax, err := countDataset2DRange(n, off)
				if err != nil {
					if cxmin < xmin {
						xmin = cxmin
					}

					if cxmax > xmax {
						xmax = cxmax
					}

					if cymin < ymin {
						ymin = cymin
					}

					if cymax > ymax {
						ymax = cymax
					}
				}
			}

			pd.Scatter3 = append(pd.Scatter3, b3)
		}

		pd.XAxis = &viewerdbpb.CoordinateConfig{
			ValType:     viewerdbpb.ValueType_INT32,
			CType:       viewerdbpb.CoordinateType_MINMAXRANGE,
			MinValInt32: xmin,
			MaxValInt32: xmax,
		}

		pd.YAxis = &viewerdbpb.CoordinateConfig{
			ValType:     viewerdbpb.ValueType_INT32,
			CType:       viewerdbpb.CoordinateType_MINMAXRANGE,
			MinValInt32: ymin,
			MaxValInt32: ymax,
		}

		pd.ZVal = &viewerdbpb.Dataset{
			Name:    "zVal",
			ValType: viewerdbpb.ValueType_STRING,
		}

		for _, v := range zval {
			pd.ZVal.ArrString = append(pd.ZVal.ArrString, v)
		}

		vjd := &viewerdbpb.ViewerData_Scatter3{
			Scatter3: pd,
		}

		vd := &viewerdbpb.ViewerData{
			Type:  viewerdbpb.ViewerType_SCATTER3,
			Title: "scatter 3",
			Token: token,
			Data:  vjd,
		}

		return vd, nil
	}

	bn, _ := buildDataset2D(datasetname, off)
	pd.Datasets = make(map[string]*viewerdbpb.Dataset2D)
	pd.Datasets[datasetname] = bn

	vjd := &viewerdbpb.ViewerData_Scatter3{
		Scatter3: pd,
	}

	vd := &viewerdbpb.ViewerData{
		Type:  viewerdbpb.ViewerType_SCATTER3,
		Title: "scatter 3",
		Token: token,
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
		vd, err := buildScatter("pie_10whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20whackamole_scatter" {
		vd, err := buildScatter("pie_20whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30whackamole_scatter" {
		vd, err := buildScatter("pie_30whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40whackamole_scatter" {
		vd, err := buildScatter("pie_40whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50whackamole_scatter" {
		vd, err := buildScatter("pie_50whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60whackamole_scatter" {
		vd, err := buildScatter("pie_60whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70whackamole_scatter" {
		vd, err := buildScatter("pie_70whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80whackamole_scatter" {
		vd, err := buildScatter("pie_80whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90whackamole_scatter" {
		vd, err := buildScatter("pie_90whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100whackamole_scatter" {
		vd, err := buildScatter("pie_100whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150whackamole_scatter" {
		vd, err := buildScatter("pie_150whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200whackamole_scatter" {
		vd, err := buildScatter("pie_200whackamole", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot" {
		vd, err := buildBoxplo([]string{"pie_10whackamole", "pie_20whackamole", "pie_30whackamole", "pie_40whackamole", "pie_50whackamole", "pie_60whackamole",
			"pie_70whackamole", "pie_80whackamole", "pie_90whackamole", "pie_100whackamole", "pie_150whackamole", "pie_200whackamole"},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200})
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
		vd, err := buildScatter("pie_10magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20magician_scatter" {
		vd, err := buildScatter("pie_20magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30magician_scatter" {
		vd, err := buildScatter("pie_30magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40magician_scatter" {
		vd, err := buildScatter("pie_40magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50magician_scatter" {
		vd, err := buildScatter("pie_50magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60magician_scatter" {
		vd, err := buildScatter("pie_60magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70magician_scatter" {
		vd, err := buildScatter("pie_70magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80magician_scatter" {
		vd, err := buildScatter("pie_80magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90magician_scatter" {
		vd, err := buildScatter("pie_90magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100magician_scatter" {
		vd, err := buildScatter("pie_100magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150magician_scatter" {
		vd, err := buildScatter("pie_150magician", 100, 1, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200magician_scatter" {
		vd, err := buildScatter("pie_200magician", 100, 1, 200000)
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
		vd, err := buildScatter("pie_10dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20dragonball_scatter" {
		vd, err := buildScatter("pie_20dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30dragonball_scatter" {
		vd, err := buildScatter("pie_30dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40dragonball_scatter" {
		vd, err := buildScatter("pie_40dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50dragonball_scatter" {
		vd, err := buildScatter("pie_50dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60dragonball_scatter" {
		vd, err := buildScatter("pie_60dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70dragonball_scatter" {
		vd, err := buildScatter("pie_70dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80dragonball_scatter" {
		vd, err := buildScatter("pie_80dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90dragonball_scatter" {
		vd, err := buildScatter("pie_90dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100dragonball_scatter" {
		vd, err := buildScatter("pie_100dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150dragonball_scatter" {
		vd, err := buildScatter("pie_150dragonball", 100, 35.0/20.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200dragonball_scatter" {
		vd, err := buildScatter("pie_200dragonball", 100, 35.0/20.0, 200000)
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
		vd, err := buildScatter("pie_10wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20wrathofthor_scatter" {
		vd, err := buildScatter("pie_20wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30wrathofthor_scatter" {
		vd, err := buildScatter("pie_30wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40wrathofthor_scatter" {
		vd, err := buildScatter("pie_40wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50wrathofthor_scatter" {
		vd, err := buildScatter("pie_50wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60wrathofthor_scatter" {
		vd, err := buildScatter("pie_60wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70wrathofthor_scatter" {
		vd, err := buildScatter("pie_70wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80wrathofthor_scatter" {
		vd, err := buildScatter("pie_80wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90wrathofthor_scatter" {
		vd, err := buildScatter("pie_90wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100wrathofthor_scatter" {
		vd, err := buildScatter("pie_100wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150wrathofthor_scatter" {
		vd, err := buildScatter("pie_150wrathofthor", 100, 35.0/30.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200wrathofthor_scatter" {
		vd, err := buildScatter("pie_200wrathofthor", 100, 35.0/30.0, 200000)
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
		vd, err := buildScatter2("10_scatter", []string{"pie_10whackamole", "pie_10magician", "pie_10dragonball", "pie_10wrathofthor", "pie_10magicbean", "pie_10dnp", "pie_10legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "20_scatter" {
		vd, err := buildScatter2("20_scatter", []string{"pie_20whackamole", "pie_20magician", "pie_20dragonball", "pie_20wrathofthor", "pie_20magicbean", "pie_20dnp", "pie_20legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "30_scatter" {
		vd, err := buildScatter2("30_scatter", []string{"pie_30whackamole", "pie_30magician", "pie_30dragonball", "pie_30wrathofthor", "pie_30magicbean", "pie_30dnp", "pie_30legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "40_scatter" {
		vd, err := buildScatter2("40_scatter", []string{"pie_40whackamole", "pie_40magician", "pie_40dragonball", "pie_40wrathofthor", "pie_40magicbean", "pie_40dnp", "pie_40legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "50_scatter" {
		vd, err := buildScatter2("50_scatter", []string{"pie_50whackamole", "pie_50magician", "pie_50dragonball", "pie_50wrathofthor", "pie_50magicbean", "pie_50dnp", "pie_50legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "60_scatter" {
		vd, err := buildScatter2("60_scatter", []string{"pie_60whackamole", "pie_60magician", "pie_60dragonball", "pie_60wrathofthor", "pie_60magicbean", "pie_60dnp", "pie_60legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "70_scatter" {
		vd, err := buildScatter2("70_scatter", []string{"pie_70whackamole", "pie_70magician", "pie_70dragonball", "pie_70wrathofthor", "pie_70magicbean", "pie_70dnp", "pie_70legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "80_scatter" {
		vd, err := buildScatter2("80_scatter", []string{"pie_80whackamole", "pie_80magician", "pie_80dragonball", "pie_80wrathofthor", "pie_80magicbean", "pie_80dnp", "pie_80legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "90_scatter" {
		vd, err := buildScatter2("90_scatter", []string{"pie_90whackamole", "pie_90magician", "pie_90dragonball", "pie_90wrathofthor", "pie_90magicbean", "pie_90dnp", "pie_90legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "100_scatter" {
		vd, err := buildScatter2("100_scatter", []string{"pie_100whackamole", "pie_100magician", "pie_100dragonball", "pie_100wrathofthor", "pie_100magicbean", "pie_100dnp", "pie_100legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "200_scatter" {
		vd, err := buildScatter2("200_scatter", []string{"pie_200whackamole", "pie_200magician", "pie_200dragonball", "pie_200wrathofthor", "pie_200magicbean", "pie_200dnp", "pie_200legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "150_scatter" {
		vd, err := buildScatter2("150_scatter", []string{"pie_150whackamole", "pie_150magician", "pie_150dragonball", "pie_150wrathofthor", "pie_150magicbean", "pie_150dnp", "pie_150legendary"}, 100, []float32{
			1.0, 1.0, 35 / 20.0, 35 / 30.0, 35 / 40.0, 35 / 60.0, 35 / 30.0,
		}, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline100" {
		vd, err := buildMulti("multiline100", []string{"10_scatter", "20_scatter", "30_scatter", "40_scatter", "50_scatter", "60_scatter", "70_scatter", "80_scatter",
			"90_scatter", "100_scatter", "150_scatter", "200_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot2" {
		vd, err := buildBoxplot2([]boxplot2{boxplot2{fn: []string{"pie_10whackamole", "pie_20whackamole", "pie_30whackamole", "pie_40whackamole", "pie_50whackamole", "pie_60whackamole",
			"pie_70whackamole", "pie_80whackamole", "pie_90whackamole", "pie_100whackamole", "pie_150whackamole", "pie_200whackamole"}},
			boxplot2{fn: []string{"pie_10magician", "pie_20magician", "pie_30magician", "pie_40magician", "pie_50magician", "pie_60magician",
				"pie_70magician", "pie_80magician", "pie_90magician", "pie_100magician", "pie_150magician", "pie_200magician"}},
			boxplot2{fn: []string{"pie_10dragonball", "pie_20dragonball", "pie_30dragonball", "pie_40dragonball", "pie_50dragonball", "pie_60dragonball",
				"pie_70dragonball", "pie_80dragonball", "pie_90dragonball", "pie_100dragonball", "pie_150dragonball", "pie_200dragonball"}},
			boxplot2{fn: []string{"pie_10wrathofthor", "pie_20wrathofthor", "pie_30wrathofthor", "pie_40wrathofthor", "pie_50wrathofthor", "pie_60wrathofthor",
				"pie_70wrathofthor", "pie_80wrathofthor", "pie_90wrathofthor", "pie_100wrathofthor", "pie_150wrathofthor", "pie_200wrathofthor"}}})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot3" {
		vd, err := buildBoxplo([]string{"pie_10magician", "pie_20magician", "pie_30magician", "pie_40magician", "pie_50magician", "pie_60magician",
			"pie_70magician", "pie_80magician", "pie_90magician", "pie_100magician", "pie_150magician", "pie_200magician"},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot4" {
		vd, err := buildBoxplo([]string{"pie_10dragonball", "pie_20dragonball", "pie_30dragonball", "pie_40dragonball", "pie_50dragonball", "pie_60dragonball",
			"pie_70dragonball", "pie_80dragonball", "pie_90dragonball", "pie_100dragonball", "pie_150dragonball", "pie_200dragonball"},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot5" {
		vd, err := buildBoxplo([]string{"pie_10wrathofthor", "pie_20wrathofthor", "pie_30wrathofthor", "pie_40wrathofthor", "pie_50wrathofthor", "pie_60wrathofthor",
			"pie_70wrathofthor", "pie_80wrathofthor", "pie_90wrathofthor", "pie_100wrathofthor", "pie_150wrathofthor", "pie_200wrathofthor"},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10magicbean" {
		vd, err := buildPie("pie_10magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20magicbean" {
		vd, err := buildPie("pie_20magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30magicbean" {
		vd, err := buildPie("pie_30magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40magicbean" {
		vd, err := buildPie("pie_40magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50magicbean" {
		vd, err := buildPie("pie_50magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60magicbean" {
		vd, err := buildPie("pie_60magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70magicbean" {
		vd, err := buildPie("pie_70magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80magicbean" {
		vd, err := buildPie("pie_80magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90magicbean" {
		vd, err := buildPie("pie_90magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100magicbean" {
		vd, err := buildPie("pie_100magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150magicbean" {
		vd, err := buildPie("pie_150magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200magicbean" {
		vd, err := buildPie("pie_200magicbean", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10magicbean_scatter" {
		vd, err := buildScatter("pie_10magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20magicbean_scatter" {
		vd, err := buildScatter("pie_20magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30magicbean_scatter" {
		vd, err := buildScatter("pie_30magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40magicbean_scatter" {
		vd, err := buildScatter("pie_40magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50magicbean_scatter" {
		vd, err := buildScatter("pie_50magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60magicbean_scatter" {
		vd, err := buildScatter("pie_60magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70magicbean_scatter" {
		vd, err := buildScatter("pie_70magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80magicbean_scatter" {
		vd, err := buildScatter("pie_80magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90magicbean_scatter" {
		vd, err := buildScatter("pie_90magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100magicbean_scatter" {
		vd, err := buildScatter("pie_100magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150magicbean_scatter" {
		vd, err := buildScatter("pie_150magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200magicbean_scatter" {
		vd, err := buildScatter("pie_200magicbean", 100, 40/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline11" {
		vd, err := buildMulti("multiline11", []string{"pie_10magicbean_scatter", "pie_20magicbean_scatter",
			"pie_30magicbean_scatter", "pie_40magicbean_scatter", "pie_50magicbean_scatter",
			"pie_60magicbean_scatter", "pie_70magicbean_scatter", "pie_80magicbean_scatter",
			"pie_90magicbean_scatter", "pie_100magicbean_scatter", "pie_150magicbean_scatter",
			"pie_200magicbean_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline12" {
		vd, err := buildMulti("multiline12", []string{"pie_10magicbean", "pie_20magicbean", "pie_30magicbean",
			"pie_40magicbean", "pie_50magicbean", "pie_60magicbean", "pie_70magicbean", "pie_80magicbean",
			"pie_90magicbean", "pie_100magicbean", "pie_150magicbean", "pie_200magicbean"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10dnp" {
		vd, err := buildPie("pie_10dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20dnp" {
		vd, err := buildPie("pie_20dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30dnp" {
		vd, err := buildPie("pie_30dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40dnp" {
		vd, err := buildPie("pie_40dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50dnp" {
		vd, err := buildPie("pie_50dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60dnp" {
		vd, err := buildPie("pie_60dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70dnp" {
		vd, err := buildPie("pie_70dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80dnp" {
		vd, err := buildPie("pie_80dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90dnp" {
		vd, err := buildPie("pie_90dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100dnp" {
		vd, err := buildPie("pie_100dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150dnp" {
		vd, err := buildPie("pie_150dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200dnp" {
		vd, err := buildPie("pie_200dnp", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10dnp_scatter" {
		vd, err := buildScatter("pie_10dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20dnp_scatter" {
		vd, err := buildScatter("pie_20dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30dnp_scatter" {
		vd, err := buildScatter("pie_30dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40dnp_scatter" {
		vd, err := buildScatter("pie_40dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50dnp_scatter" {
		vd, err := buildScatter("pie_50dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60dnp_scatter" {
		vd, err := buildScatter("pie_60dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70dnp_scatter" {
		vd, err := buildScatter("pie_70dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80dnp_scatter" {
		vd, err := buildScatter("pie_80dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90dnp_scatter" {
		vd, err := buildScatter("pie_90dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100dnp_scatter" {
		vd, err := buildScatter("pie_100dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150dnp_scatter" {
		vd, err := buildScatter("pie_150dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200dnp_scatter" {
		vd, err := buildScatter("pie_200dnp", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline13" {
		vd, err := buildMulti("multiline13", []string{"pie_10dnp_scatter", "pie_20dnp_scatter",
			"pie_30dnp_scatter", "pie_40dnp_scatter", "pie_50dnp_scatter",
			"pie_60dnp_scatter", "pie_70dnp_scatter", "pie_80dnp_scatter",
			"pie_90dnp_scatter", "pie_100dnp_scatter", "pie_150dnp_scatter",
			"pie_200dnp_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline14" {
		vd, err := buildMulti("multiline14", []string{"pie_10dnp", "pie_20dnp", "pie_30dnp",
			"pie_40dnp", "pie_50dnp", "pie_60dnp", "pie_70dnp", "pie_80dnp",
			"pie_90dnp", "pie_100dnp", "pie_150dnp", "pie_200dnp"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10legendary" {
		vd, err := buildPie("pie_10legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20legendary" {
		vd, err := buildPie("pie_20legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30legendary" {
		vd, err := buildPie("pie_30legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40legendary" {
		vd, err := buildPie("pie_40legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50legendary" {
		vd, err := buildPie("pie_50legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60legendary" {
		vd, err := buildPie("pie_60legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70legendary" {
		vd, err := buildPie("pie_70legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80legendary" {
		vd, err := buildPie("pie_80legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90legendary" {
		vd, err := buildPie("pie_90legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100legendary" {
		vd, err := buildPie("pie_100legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150legendary" {
		vd, err := buildPie("pie_150legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200legendary" {
		vd, err := buildPie("pie_200legendary", 100)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_10legendary_scatter" {
		vd, err := buildScatter("pie_10legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_20legendary_scatter" {
		vd, err := buildScatter("pie_20legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_30legendary_scatter" {
		vd, err := buildScatter("pie_30legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_40legendary_scatter" {
		vd, err := buildScatter("pie_40legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_50legendary_scatter" {
		vd, err := buildScatter("pie_50legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_60legendary_scatter" {
		vd, err := buildScatter("pie_60legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_70legendary_scatter" {
		vd, err := buildScatter("pie_70legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_80legendary_scatter" {
		vd, err := buildScatter("pie_80legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_90legendary_scatter" {
		vd, err := buildScatter("pie_90legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_100legendary_scatter" {
		vd, err := buildScatter("pie_100legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_150legendary_scatter" {
		vd, err := buildScatter("pie_150legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "pie_200legendary_scatter" {
		vd, err := buildScatter("pie_200legendary", 100, 60/35.0, 200000)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline15" {
		vd, err := buildMulti("multiline15", []string{"pie_10legendary_scatter", "pie_20legendary_scatter",
			"pie_30legendary_scatter", "pie_40legendary_scatter", "pie_50legendary_scatter",
			"pie_60legendary_scatter", "pie_70legendary_scatter", "pie_80legendary_scatter",
			"pie_90legendary_scatter", "pie_100legendary_scatter", "pie_150legendary_scatter",
			"pie_200legendary_scatter"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "multiline16" {
		vd, err := buildMulti("multiline16", []string{"pie_10legendary", "pie_20legendary", "pie_30legendary",
			"pie_40legendary", "pie_50legendary", "pie_60legendary", "pie_70legendary", "pie_80legendary",
			"pie_90legendary", "pie_100legendary", "pie_150legendary", "pie_200legendary"})
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "usermoneyboxplot100" {
		dataset := r.URL.Query().Get("dataset")

		vd, err := buildBoxplot3([]boxplot3{boxplot3{name: "whackamole", fn: []string{"pie_10whackamole", "pie_20whackamole", "pie_30whackamole", "pie_40whackamole", "pie_50whackamole", "pie_60whackamole",
			"pie_70whackamole", "pie_80whackamole", "pie_90whackamole", "pie_100whackamole", "pie_150whackamole", "pie_200whackamole"}},
			boxplot3{name: "magician", fn: []string{"pie_10magician", "pie_20magician", "pie_30magician", "pie_40magician", "pie_50magician", "pie_60magician",
				"pie_70magician", "pie_80magician", "pie_90magician", "pie_100magician", "pie_150magician", "pie_200magician"}},
			boxplot3{name: "dragonball", fn: []string{"pie_10dragonball", "pie_20dragonball", "pie_30dragonball", "pie_40dragonball", "pie_50dragonball", "pie_60dragonball",
				"pie_70dragonball", "pie_80dragonball", "pie_90dragonball", "pie_100dragonball", "pie_150dragonball", "pie_200dragonball"}},
			boxplot3{name: "wrathofthor", fn: []string{"pie_10wrathofthor", "pie_20wrathofthor", "pie_30wrathofthor", "pie_40wrathofthor", "pie_50wrathofthor", "pie_60wrathofthor",
				"pie_70wrathofthor", "pie_80wrathofthor", "pie_90wrathofthor", "pie_100wrathofthor", "pie_150wrathofthor", "pie_200wrathofthor"}}},
			[]string{"10", "20", "30", "40", "50", "60", "70", "80", "90", "100", "150", "200"}, dataset)
		if err != nil {
			return
		}

		jsonBytes, err := json.Marshal(vd)
		if err != nil {
			return
		}

		w.Write(jsonBytes)
	} else if token == "allgamescatter3" {
		dataset := r.URL.Query().Get("dataset")

		vd, err := buildScatter3(token, []scatter3{scatter3{name: "whackamole", fn: []string{"pie_10whackamole", "pie_20whackamole", "pie_30whackamole", "pie_40whackamole", "pie_50whackamole", "pie_60whackamole",
			"pie_70whackamole", "pie_80whackamole", "pie_90whackamole", "pie_100whackamole", "pie_150whackamole", "pie_200whackamole"}},
			scatter3{name: "magician", fn: []string{"pie_10magician", "pie_20magician", "pie_30magician", "pie_40magician", "pie_50magician", "pie_60magician",
				"pie_70magician", "pie_80magician", "pie_90magician", "pie_100magician", "pie_150magician", "pie_200magician"}},
			scatter3{name: "dragonball", fn: []string{"pie_10dragonball", "pie_20dragonball", "pie_30dragonball", "pie_40dragonball", "pie_50dragonball", "pie_60dragonball",
				"pie_70dragonball", "pie_80dragonball", "pie_90dragonball", "pie_100dragonball", "pie_150dragonball", "pie_200dragonball"}},
			scatter3{name: "wrathofthor", fn: []string{"pie_10wrathofthor", "pie_20wrathofthor", "pie_30wrathofthor", "pie_40wrathofthor", "pie_50wrathofthor", "pie_60wrathofthor",
				"pie_70wrathofthor", "pie_80wrathofthor", "pie_90wrathofthor", "pie_100wrathofthor", "pie_150wrathofthor", "pie_200wrathofthor"}}},
			[]string{"10", "20", "30", "40", "50", "60", "70", "80", "90", "100", "150", "200"}, dataset, 100)
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
