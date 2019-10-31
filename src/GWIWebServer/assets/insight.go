package assets

import "bytes"

// Insight structure
type Insight struct {
	Asset
	text string
}

// CreateNewInsight creates a new insight asset
func CreateNewInsight(assetID string, description string, text string) Insight {
	return Insight{
		Asset: Asset{
			assetID:     assetID,
			AssetType:   InsightType,
			description: description},
		text: text}
}

// GetAsset returns pointer of basic asset structure
func (insight Insight) GetAsset() *Asset {
	return &insight.Asset
}

// GetDetails returns unique fields info of insight
func (insight Insight) GetDetails() string {
	var out bytes.Buffer
	out.WriteString(insight.Asset.GetDetails())
	out.WriteString("\tText: '" + insight.text + "'\n")
	return out.String()
}
