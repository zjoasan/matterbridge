// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/zjoasan/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
