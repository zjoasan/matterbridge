// +build !nowhatsapp

package bridgemap

import (
	bwhatsapp "github.com/zjoasan/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
