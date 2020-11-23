// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/zjoasan/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
