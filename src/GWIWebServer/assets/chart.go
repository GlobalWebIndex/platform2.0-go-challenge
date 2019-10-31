package assets

import "bytes"

// XY structure for chart points
type XY struct{ X, Y float64 }

// Chart structure
type Chart struct {
	Asset
	title      string
	xAxisTitle string
	yAxisTitle string
	data       []XY
}

// CreateNewChart creates a new chart asset
func CreateNewChart(assetID string, description string, title string, xAxisTitle string, yAxisTitle string) Chart {
	return Chart{
		Asset: Asset{
			assetID:     assetID,
			AssetType:   ChartType,
			description: description},
		title:      title,
		xAxisTitle: xAxisTitle,
		yAxisTitle: yAxisTitle}
}

// GetAsset returns pointer of basic asset structure
func (chart Chart) GetAsset() *Asset {
	return &chart.Asset
}

// GetDetails returns unique fields info of chart
func (chart Chart) GetDetails() string {
	var out bytes.Buffer
	out.WriteString(chart.Asset.GetDetails())
	out.WriteString("\tTitle: '" + chart.title + "'\n")
	out.WriteString("\tX Axis Title: '" + chart.xAxisTitle + "'\n")
	out.WriteString("\tY Axis Title: '" + chart.yAxisTitle + "'\n")
	return out.String()
}
