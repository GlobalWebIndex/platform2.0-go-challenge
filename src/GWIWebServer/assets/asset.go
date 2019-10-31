package assets

import (
	"bytes"
)

// AssetInterface wraps asset structs
type AssetInterface interface {
	GetAsset() *Asset
	GetDetails() string
}

// Asset is an item available in GWI platform
type Asset struct {
	assetID string
	AssetType
	description string
}

// SetDescription sets description for any Asset
func (asset *Asset) SetDescription(description string) {
	asset.description = description
}

// GetDescription returns description of any Asset
func (asset *Asset) GetDescription() string {
	return asset.description
}

// GetDetails returns common fields info of any Asset
func (asset *Asset) GetDetails() string {
	var out bytes.Buffer
	out.WriteString("AssetID: '" + asset.assetID + "'\n")
	out.WriteString("Type: '" + asset.AssetType.String() + "'\n")
	out.WriteString("Description: '" + asset.description + "'\n")
	return out.String()
}
