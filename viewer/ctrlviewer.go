package viewer

import (
	pb "github.com/zhs007/jarviscore/proto"
)

const (
	// CtrlTypeViewer - jarvisviewer ctrltype
	CtrlTypeViewer = "jarvisviewer"
)

// CtrlJarvisViewer -
type CtrlJarvisViewer struct {
}

// Run -
func (ctrl *CtrlJarvisViewer) Run(ci *pb.CtrlInfo) ([]byte, error) {
	return nil, nil
}
