package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarvissh/basedef"
	"github.com/zhs007/jarvisviewer/viewer"
)

func main() {
	fmt.Print("jarvis viewer start...\n")
	fmt.Print("jarvis viewer version is " + basedef.VERSION + "\n")

	cfg, err := viewer.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Print("load config.yaml fail!")

		return
	}

	jv, err := viewer.NewViewer(cfg)
	if err != nil {
		fmt.Print(err.Error())

		return
	}

	jv.Start(context.Background())

	// jarviscore.InitJarvisCore(cfg)
	// defer jarviscore.ReleaseJarvisCore()

	// node := jarviscore.NewNode(cfg)
	// node.RegCtrl(viewer.CtrlTypeViewer, &viewer.CtrlJarvisViewer{})
	// node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	// node.Start(context.Background())

	fmt.Print("jarvis viewer end.\n")
}
