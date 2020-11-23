// +build !noapi

package bridgemap

import (
	"github.com/zjoasan/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
