// +build !noirc

package bridgemap

import (
	birc "github.com/zjoasan/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
