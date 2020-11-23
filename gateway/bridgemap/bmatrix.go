// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/zjoasan/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
