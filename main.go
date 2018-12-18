package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvissh/basedef"
	"github.com/zhs007/jarvisviewer/viewer"
)

func main() {
	fmt.Print("jarvis viewer start...")
	fmt.Print("jarvis viewer version is " + basedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Print("load config.yaml fail!")

		return
	}

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	node := jarviscore.NewNode(cfg)
	node.RegCtrl(viewer.CtrlTypeViewer, &viewer.CtrlJarvisViewer{})
	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	node.Start(context.Background())

	fmt.Print("jarvis viewer end.")
}
