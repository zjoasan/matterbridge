// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/zjoasan/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
