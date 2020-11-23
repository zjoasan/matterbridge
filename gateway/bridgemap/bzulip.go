// +build !nozulip

package bridgemap

import (
	bzulip "github.com/zjoasan/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
