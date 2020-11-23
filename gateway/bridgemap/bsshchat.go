// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/zjoasan/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
