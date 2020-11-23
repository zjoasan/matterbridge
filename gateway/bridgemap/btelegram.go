// +build !notelegram

package bridgemap

import (
	btelegram "github.com/zjoasan/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
