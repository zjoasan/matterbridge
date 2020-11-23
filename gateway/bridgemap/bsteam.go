// +build !nosteam

package bridgemap

import (
	bsteam "github.com/zjoasan/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
