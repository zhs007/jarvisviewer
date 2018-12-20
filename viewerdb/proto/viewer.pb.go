// Code generated by protoc-gen-go. DO NOT EDIT.
// source: viewer.proto

package viewerdbpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ViewerType - Viewer type
//            - 数据类型
type ViewerType int32

const (
	ViewerType_EMPTY ViewerType = 0
	ViewerType_GRAPH ViewerType = 1
	ViewerType_JSON  ViewerType = 2
	ViewerType_LINE  ViewerType = 3
)

var ViewerType_name = map[int32]string{
	0: "EMPTY",
	1: "GRAPH",
	2: "JSON",
	3: "LINE",
}
var ViewerType_value = map[string]int32{
	"EMPTY": 0,
	"GRAPH": 1,
	"JSON":  2,
	"LINE":  3,
}

func (x ViewerType) String() string {
	return proto.EnumName(ViewerType_name, int32(x))
}
func (ViewerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{0}
}

// GraphNode - Graph node
//           - 图节点
type GraphNode struct {
	// id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// size
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// category
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphNode) Reset()         { *m = GraphNode{} }
func (m *GraphNode) String() string { return proto.CompactTextString(m) }
func (*GraphNode) ProtoMessage()    {}
func (*GraphNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{0}
}
func (m *GraphNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphNode.Unmarshal(m, b)
}
func (m *GraphNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphNode.Marshal(b, m, deterministic)
}
func (dst *GraphNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphNode.Merge(dst, src)
}
func (m *GraphNode) XXX_Size() int {
	return xxx_messageInfo_GraphNode.Size(m)
}
func (m *GraphNode) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphNode.DiscardUnknown(m)
}

var xxx_messageInfo_GraphNode proto.InternalMessageInfo

func (m *GraphNode) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GraphNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GraphNode) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// GraphLink - Graph link info
//           - 图节点连接关系
type GraphLink struct {
	// src
	Src int32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	// dest
	Dest int32 `protobuf:"varint,2,opt,name=dest,proto3" json:"dest,omitempty"`
	// size
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphLink) Reset()         { *m = GraphLink{} }
func (m *GraphLink) String() string { return proto.CompactTextString(m) }
func (*GraphLink) ProtoMessage()    {}
func (*GraphLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{1}
}
func (m *GraphLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphLink.Unmarshal(m, b)
}
func (m *GraphLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphLink.Marshal(b, m, deterministic)
}
func (dst *GraphLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphLink.Merge(dst, src)
}
func (m *GraphLink) XXX_Size() int {
	return xxx_messageInfo_GraphLink.Size(m)
}
func (m *GraphLink) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphLink.DiscardUnknown(m)
}

var xxx_messageInfo_GraphLink proto.InternalMessageInfo

func (m *GraphLink) GetSrc() int32 {
	if m != nil {
		return m.Src
	}
	return 0
}

func (m *GraphLink) GetDest() int32 {
	if m != nil {
		return m.Dest
	}
	return 0
}

func (m *GraphLink) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

// GraphData - Graph data
//           - 图节点完整数据
type GraphData struct {
	// nodes
	Nodes []*GraphNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// links
	Links                []*GraphLink `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GraphData) Reset()         { *m = GraphData{} }
func (m *GraphData) String() string { return proto.CompactTextString(m) }
func (*GraphData) ProtoMessage()    {}
func (*GraphData) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{2}
}
func (m *GraphData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphData.Unmarshal(m, b)
}
func (m *GraphData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphData.Marshal(b, m, deterministic)
}
func (dst *GraphData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphData.Merge(dst, src)
}
func (m *GraphData) XXX_Size() int {
	return xxx_messageInfo_GraphData.Size(m)
}
func (m *GraphData) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphData.DiscardUnknown(m)
}

var xxx_messageInfo_GraphData proto.InternalMessageInfo

func (m *GraphData) GetNodes() []*GraphNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *GraphData) GetLinks() []*GraphLink {
	if m != nil {
		return m.Links
	}
	return nil
}

// JSONData - JSON data
//          - json格式的数据
type JSONData struct {
	// string for json
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONData) Reset()         { *m = JSONData{} }
func (m *JSONData) String() string { return proto.CompactTextString(m) }
func (*JSONData) ProtoMessage()    {}
func (*JSONData) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{3}
}
func (m *JSONData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONData.Unmarshal(m, b)
}
func (m *JSONData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONData.Marshal(b, m, deterministic)
}
func (dst *JSONData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONData.Merge(dst, src)
}
func (m *JSONData) XXX_Size() int {
	return xxx_messageInfo_JSONData.Size(m)
}
func (m *JSONData) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONData.DiscardUnknown(m)
}

var xxx_messageInfo_JSONData proto.InternalMessageInfo

func (m *JSONData) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

// LineSeriesData - Line series data
//                - 折线图数据
type LineSeriesData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Val                  []int32  `protobuf:"varint,2,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineSeriesData) Reset()         { *m = LineSeriesData{} }
func (m *LineSeriesData) String() string { return proto.CompactTextString(m) }
func (*LineSeriesData) ProtoMessage()    {}
func (*LineSeriesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{4}
}
func (m *LineSeriesData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineSeriesData.Unmarshal(m, b)
}
func (m *LineSeriesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineSeriesData.Marshal(b, m, deterministic)
}
func (dst *LineSeriesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineSeriesData.Merge(dst, src)
}
func (m *LineSeriesData) XXX_Size() int {
	return xxx_messageInfo_LineSeriesData.Size(m)
}
func (m *LineSeriesData) XXX_DiscardUnknown() {
	xxx_messageInfo_LineSeriesData.DiscardUnknown(m)
}

var xxx_messageInfo_LineSeriesData proto.InternalMessageInfo

func (m *LineSeriesData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LineSeriesData) GetVal() []int32 {
	if m != nil {
		return m.Val
	}
	return nil
}

// LineData - Line data
//           - 折线图完整数据
type LineData struct {
	// xAxisType
	XAxisType string `protobuf:"bytes,1,opt,name=xAxisType,proto3" json:"xAxisType,omitempty"`
	// xAxisString
	XAxisString []string `protobuf:"bytes,2,rep,name=xAxisString,proto3" json:"xAxisString,omitempty"`
	// xAxisInt32
	XAxisInt32 []int32 `protobuf:"varint,3,rep,packed,name=xAxisInt32,proto3" json:"xAxisInt32,omitempty"`
	// xAxisInt64
	XAxisInt64 []int64 `protobuf:"varint,4,rep,packed,name=xAxisInt64,proto3" json:"xAxisInt64,omitempty"`
	// yAxisType
	YAxisType string `protobuf:"bytes,10,opt,name=yAxisType,proto3" json:"yAxisType,omitempty"`
	// xAxisData
	Series               []*LineSeriesData `protobuf:"bytes,11,rep,name=series,proto3" json:"series,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LineData) Reset()         { *m = LineData{} }
func (m *LineData) String() string { return proto.CompactTextString(m) }
func (*LineData) ProtoMessage()    {}
func (*LineData) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{5}
}
func (m *LineData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineData.Unmarshal(m, b)
}
func (m *LineData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineData.Marshal(b, m, deterministic)
}
func (dst *LineData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineData.Merge(dst, src)
}
func (m *LineData) XXX_Size() int {
	return xxx_messageInfo_LineData.Size(m)
}
func (m *LineData) XXX_DiscardUnknown() {
	xxx_messageInfo_LineData.DiscardUnknown(m)
}

var xxx_messageInfo_LineData proto.InternalMessageInfo

func (m *LineData) GetXAxisType() string {
	if m != nil {
		return m.XAxisType
	}
	return ""
}

func (m *LineData) GetXAxisString() []string {
	if m != nil {
		return m.XAxisString
	}
	return nil
}

func (m *LineData) GetXAxisInt32() []int32 {
	if m != nil {
		return m.XAxisInt32
	}
	return nil
}

func (m *LineData) GetXAxisInt64() []int64 {
	if m != nil {
		return m.XAxisInt64
	}
	return nil
}

func (m *LineData) GetYAxisType() string {
	if m != nil {
		return m.YAxisType
	}
	return ""
}

func (m *LineData) GetSeries() []*LineSeriesData {
	if m != nil {
		return m.Series
	}
	return nil
}

// ViewerData - viewer data
//            - viewer数据
type ViewerData struct {
	// type
	Type ViewerType `protobuf:"varint,1,opt,name=type,proto3,enum=viewerdbpb.ViewerType" json:"type,omitempty"`
	// title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// token
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*ViewerData_Graph
	//	*ViewerData_Json
	//	*ViewerData_Line
	Data                 isViewerData_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ViewerData) Reset()         { *m = ViewerData{} }
func (m *ViewerData) String() string { return proto.CompactTextString(m) }
func (*ViewerData) ProtoMessage()    {}
func (*ViewerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_viewer_ae7f542d40794fb4, []int{6}
}
func (m *ViewerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewerData.Unmarshal(m, b)
}
func (m *ViewerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewerData.Marshal(b, m, deterministic)
}
func (dst *ViewerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewerData.Merge(dst, src)
}
func (m *ViewerData) XXX_Size() int {
	return xxx_messageInfo_ViewerData.Size(m)
}
func (m *ViewerData) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewerData.DiscardUnknown(m)
}

var xxx_messageInfo_ViewerData proto.InternalMessageInfo

type isViewerData_Data interface {
	isViewerData_Data()
}

type ViewerData_Graph struct {
	Graph *GraphData `protobuf:"bytes,1000,opt,name=graph,proto3,oneof"`
}
type ViewerData_Json struct {
	Json *JSONData `protobuf:"bytes,1001,opt,name=json,proto3,oneof"`
}
type ViewerData_Line struct {
	Line *LineData `protobuf:"bytes,1002,opt,name=line,proto3,oneof"`
}

func (*ViewerData_Graph) isViewerData_Data() {}
func (*ViewerData_Json) isViewerData_Data()  {}
func (*ViewerData_Line) isViewerData_Data()  {}

func (m *ViewerData) GetData() isViewerData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ViewerData) GetType() ViewerType {
	if m != nil {
		return m.Type
	}
	return ViewerType_EMPTY
}

func (m *ViewerData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ViewerData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ViewerData) GetGraph() *GraphData {
	if x, ok := m.GetData().(*ViewerData_Graph); ok {
		return x.Graph
	}
	return nil
}

func (m *ViewerData) GetJson() *JSONData {
	if x, ok := m.GetData().(*ViewerData_Json); ok {
		return x.Json
	}
	return nil
}

func (m *ViewerData) GetLine() *LineData {
	if x, ok := m.GetData().(*ViewerData_Line); ok {
		return x.Line
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ViewerData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ViewerData_OneofMarshaler, _ViewerData_OneofUnmarshaler, _ViewerData_OneofSizer, []interface{}{
		(*ViewerData_Graph)(nil),
		(*ViewerData_Json)(nil),
		(*ViewerData_Line)(nil),
	}
}

func _ViewerData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ViewerData)
	// data
	switch x := m.Data.(type) {
	case *ViewerData_Graph:
		b.EncodeVarint(1000<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Graph); err != nil {
			return err
		}
	case *ViewerData_Json:
		b.EncodeVarint(1001<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Json); err != nil {
			return err
		}
	case *ViewerData_Line:
		b.EncodeVarint(1002<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Line); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ViewerData.Data has unexpected type %T", x)
	}
	return nil
}

func _ViewerData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ViewerData)
	switch tag {
	case 1000: // data.graph
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GraphData)
		err := b.DecodeMessage(msg)
		m.Data = &ViewerData_Graph{msg}
		return true, err
	case 1001: // data.json
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JSONData)
		err := b.DecodeMessage(msg)
		m.Data = &ViewerData_Json{msg}
		return true, err
	case 1002: // data.line
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LineData)
		err := b.DecodeMessage(msg)
		m.Data = &ViewerData_Line{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ViewerData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ViewerData)
	// data
	switch x := m.Data.(type) {
	case *ViewerData_Graph:
		s := proto.Size(x.Graph)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ViewerData_Json:
		s := proto.Size(x.Json)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ViewerData_Line:
		s := proto.Size(x.Line)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GraphNode)(nil), "viewerdbpb.GraphNode")
	proto.RegisterType((*GraphLink)(nil), "viewerdbpb.GraphLink")
	proto.RegisterType((*GraphData)(nil), "viewerdbpb.GraphData")
	proto.RegisterType((*JSONData)(nil), "viewerdbpb.JSONData")
	proto.RegisterType((*LineSeriesData)(nil), "viewerdbpb.LineSeriesData")
	proto.RegisterType((*LineData)(nil), "viewerdbpb.LineData")
	proto.RegisterType((*ViewerData)(nil), "viewerdbpb.ViewerData")
	proto.RegisterEnum("viewerdbpb.ViewerType", ViewerType_name, ViewerType_value)
}

func init() { proto.RegisterFile("viewer.proto", fileDescriptor_viewer_ae7f542d40794fb4) }

var fileDescriptor_viewer_ae7f542d40794fb4 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x33, 0xfb, 0x27, 0x66, 0x4f, 0x24, 0x2c, 0x43, 0x95, 0xa1, 0x14, 0x59, 0xf6, 0x2a,
	0xb4, 0x90, 0x8b, 0x54, 0x72, 0x5f, 0x31, 0x34, 0x95, 0x18, 0xcb, 0xa4, 0x08, 0x5e, 0xc9, 0xa6,
	0x3b, 0xc4, 0x31, 0x71, 0x36, 0xec, 0x0c, 0xb5, 0xf1, 0x59, 0x7d, 0x00, 0xf5, 0x01, 0xbc, 0x96,
	0x73, 0x36, 0x9b, 0xac, 0x90, 0xde, 0x7d, 0xe7, 0xcb, 0x6f, 0xe6, 0x9b, 0x73, 0xce, 0x06, 0x9e,
	0x3f, 0x68, 0xf5, 0x5d, 0x95, 0x83, 0x4d, 0x59, 0xb8, 0x82, 0x43, 0x55, 0xe5, 0x8b, 0xcd, 0x22,
	0xfd, 0x0c, 0xd1, 0x75, 0x99, 0x6d, 0xbe, 0xcc, 0x8a, 0x5c, 0xf1, 0x1e, 0x78, 0x3a, 0x17, 0x2c,
	0x61, 0xfd, 0x50, 0x7a, 0x3a, 0xe7, 0x1c, 0x02, 0x93, 0x7d, 0x53, 0xc2, 0x4b, 0x58, 0x3f, 0x92,
	0xa4, 0xd1, 0xb3, 0xfa, 0x87, 0x12, 0x3e, 0x51, 0xa4, 0xf9, 0x29, 0x74, 0xee, 0x33, 0xa7, 0x96,
	0x45, 0xb9, 0x15, 0x01, 0xb1, 0xfb, 0x3a, 0x1d, 0xef, 0x02, 0xa6, 0xda, 0xac, 0x78, 0x0c, 0xbe,
	0x2d, 0xef, 0x77, 0x09, 0x28, 0xf1, 0xba, 0x5c, 0x59, 0x47, 0x11, 0xa1, 0x24, 0x7d, 0x2c, 0x22,
	0x55, 0xbb, 0x6b, 0xde, 0x66, 0x2e, 0xe3, 0x17, 0x10, 0x9a, 0x22, 0x57, 0x56, 0xb0, 0xc4, 0xef,
	0x77, 0x87, 0x2f, 0x06, 0x87, 0x86, 0x06, 0xfb, 0x6e, 0x64, 0xc5, 0x20, 0xbc, 0xd6, 0x66, 0x65,
	0x85, 0xf7, 0x04, 0x8c, 0x2f, 0x93, 0x15, 0x93, 0x9e, 0x41, 0xe7, 0xdd, 0xfc, 0xc3, 0x8c, 0x52,
	0xf0, 0xb1, 0xae, 0xa4, 0xc7, 0x46, 0x12, 0x65, 0x3a, 0x82, 0xde, 0x54, 0x1b, 0x35, 0x57, 0xa5,
	0x56, 0x96, 0x98, 0x7a, 0x42, 0xac, 0x31, 0xa1, 0x18, 0xfc, 0x87, 0x6c, 0x4d, 0x71, 0xa1, 0x44,
	0x99, 0xfe, 0x64, 0xd0, 0xc1, 0x83, 0x74, 0xe4, 0x0c, 0xa2, 0xc7, 0xab, 0x47, 0x6d, 0xef, 0xb6,
	0x9b, 0xfa, 0xdc, 0xc1, 0xe0, 0x09, 0x74, 0xa9, 0x98, 0xbb, 0x52, 0x9b, 0x25, 0x5d, 0x12, 0xc9,
	0xa6, 0xc5, 0x5f, 0x01, 0x50, 0x79, 0x63, 0xdc, 0xe5, 0x50, 0xf8, 0x94, 0xd2, 0x70, 0x9a, 0xbf,
	0x8f, 0x5e, 0x8b, 0x20, 0xf1, 0xfb, 0xbe, 0x6c, 0x38, 0x98, 0xbf, 0xdd, 0xe7, 0x43, 0x95, 0xbf,
	0x37, 0xf8, 0x10, 0xda, 0x96, 0xda, 0x13, 0x5d, 0x1a, 0xd7, 0x69, 0x73, 0x5c, 0xff, 0x37, 0x2f,
	0x77, 0x64, 0xfa, 0x97, 0x01, 0x7c, 0x24, 0x8a, 0x1a, 0x3c, 0x87, 0xc0, 0xd5, 0xbd, 0xf5, 0x86,
	0x2f, 0x9b, 0x17, 0x54, 0x14, 0x06, 0x49, 0x62, 0xf8, 0x09, 0x84, 0x4e, 0xbb, 0x75, 0xfd, 0x89,
	0x55, 0x05, 0xb9, 0xc5, 0x4a, 0x19, 0xfa, 0x02, 0xd0, 0xc5, 0x82, 0x0f, 0x20, 0x5c, 0xe2, 0xbe,
	0xc4, 0xaf, 0x67, 0x09, 0x3b, 0xba, 0x49, 0x8c, 0x9f, 0xb4, 0x64, 0x85, 0xf1, 0x0b, 0x08, 0xbe,
	0xda, 0xc2, 0x88, 0xdf, 0x15, 0x7e, 0xd2, 0xc4, 0xeb, 0x25, 0x4f, 0x5a, 0x92, 0x20, 0x84, 0xd7,
	0xda, 0x28, 0xf1, 0xe7, 0x08, 0x5c, 0xaf, 0x0e, 0x61, 0x84, 0xde, 0xb4, 0x21, 0xc8, 0x33, 0x97,
	0x9d, 0x8f, 0xea, 0xbe, 0x69, 0x74, 0x11, 0x84, 0xe3, 0xf7, 0xb7, 0x77, 0x9f, 0xe2, 0x16, 0xca,
	0x6b, 0x79, 0x75, 0x3b, 0x89, 0x19, 0xef, 0x40, 0x80, 0x61, 0xb1, 0x87, 0x6a, 0x7a, 0x33, 0x1b,
	0xc7, 0xfe, 0xa2, 0x4d, 0xff, 0xc3, 0xcb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x25, 0x48,
	0xbb, 0x97, 0x03, 0x00, 0x00,
}
