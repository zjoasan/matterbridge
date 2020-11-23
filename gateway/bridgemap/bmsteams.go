// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/zjoasan/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
