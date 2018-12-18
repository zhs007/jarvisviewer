package viewerdb

import (
	"fmt"
)

const prefixKeyViewerData = "vd:"

func makeViewerDataKey(token string) string {
	return fmt.Sprintf("%v%v", prefixKeyViewerData, token)
}
