package serializers

import "gwi-challenge/data/models"

type ChartSerializer struct {
	*models.Chart
}

type ChartResponse struct {
	ID    uint                 `json:"-"`
	Title string               `json:"title"`
	XAxes string               `json:"x-axes"`
	YAxes string               `json:"y-axes"`
	Data  []ChartPointResponse `json:"data"`
}

type ChartPointResponse struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func NewChartSerializer() ChartSerializer {
	return ChartSerializer{}
}

func (serializer *ChartSerializer) Response() *ChartResponse {
	response := ChartResponse{
		ID:    serializer.ID,
		Title: serializer.Title,
		YAxes: serializer.YAxes,
		XAxes: serializer.XAxes,
		Data:  []ChartPointResponse{},
	}

	for _, point := range serializer.Data {
		response.Data = append(response.Data, ChartPointResponse{X: point.X, Y: point.Y})
	}

	return &response
}
