package viewer

import (
	"context"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvisviewer/basedef"
)

// Viewer - viewer
type Viewer struct {
	node jarviscore.JarvisNode
	serv *HTTPServer
}

// NewViewer - new viewer
func NewViewer(cfg *Config) (*Viewer, error) {
	jcfg, err := jarviscore.LoadConfig(cfg.JarvisNodeCfg)
	if err != nil {
		return nil, ErrJarvisNodeConfigFile
	}

	jarviscore.InitJarvisCore(jcfg)
	defer jarviscore.ReleaseJarvisCore()

	v := &Viewer{}

	serv, err := newHTTPServer(cfg.BindAddr)
	if err != nil {
		return nil, err
	}

	v.serv = serv

	node := jarviscore.NewNode(jcfg)
	node.RegCtrl(CtrlTypeViewer, &CtrlJarvisViewer{})
	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	v.node = node

	return v, nil
}

// Start - Start viewer
//		 - This will block the current goroutine
func (v *Viewer) Start(ctx context.Context) error {

	go v.serv.start(ctx)
	v.node.Start(ctx)

	return nil
}
