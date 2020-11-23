// +build !nomumble

package bridgemap

import (
	bmumble "github.com/zjoasan/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
