// +build !nonctalk

package bridgemap

import (
	btalk "github.com/zjoasan/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
