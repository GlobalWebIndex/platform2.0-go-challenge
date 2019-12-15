package serializers

import "gwi-challenge/data/models"

type InsightSerializer struct {
	*models.Insight
}

type InsightResponse struct {
	Title   string `json:"title"`
	Insight string `json:"insight"`
}

func NewInsightSerializer() InsightSerializer {
	return InsightSerializer{}
}

func (serializer *InsightSerializer) Response() *InsightResponse {
	response := InsightResponse{
		Title:   serializer.Title,
		Insight: serializer.Text,
	}
	return &response
}
