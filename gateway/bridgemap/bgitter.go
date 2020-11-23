// +build !nogitter

package bridgemap

import (
	bgitter "github.com/zjoasan/matterbridge/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
