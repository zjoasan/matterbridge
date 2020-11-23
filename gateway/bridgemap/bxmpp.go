// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/zjoasan/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
