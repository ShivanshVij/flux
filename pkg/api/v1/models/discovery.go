package models

import "github.com/shivanshvij/flux/pkg/sdcp"

type DiscoveryResponse struct {
	Discovered []sdcp.DiscoveryData `json:"discovered"`
}
