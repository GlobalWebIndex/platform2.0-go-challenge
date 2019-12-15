package serializers

import (
	"gwi-challenge/data/models"
	"sync"
)

type AssetsSerializer struct {
	models.Assets
}

type AssetsResponse struct {
	Charts    []ChartResponse    `json:"charts"`
	Insights  []InsightResponse  `json:"insights"`
	Audiences []AudienceResponse `json:"audiences"`
}

func NewAssetSerializer() AssetsSerializer {
	return AssetsSerializer{
		Assets: models.Assets{
			Charts:    &[]models.Chart{},
			Insights:  &[]models.Insight{},
			Audiences: &[]models.Audience{},
		},
	}
}

func (assetsSerializer *AssetsSerializer) Response() *AssetsResponse {
	response := AssetsResponse{}

	var wg sync.WaitGroup

	if len(*assetsSerializer.Charts) > 0 {
		wg.Add(1)
		go buildChartResponses(assetsSerializer, &response, &wg)
	}
	if len(*assetsSerializer.Insights) > 0 {
		wg.Add(1)
		go buildInsightResponses(assetsSerializer, &response, &wg)
	}
	if len(*assetsSerializer.Charts) > 0 {
		wg.Add(1)
		go buildAudienceResponses(assetsSerializer, &response, &wg)
	}
	wg.Wait()
	return &response
}

func buildChartResponses(assetsSerializer *AssetsSerializer, response *AssetsResponse, wg *sync.WaitGroup) {
	for _, chart := range *assetsSerializer.Charts {
		serializer := NewChartSerializer()
		serializer.Chart = &chart
		response.Charts = append(response.Charts, *serializer.Response())
	}
	wg.Done()
}

func buildInsightResponses(assetsSerializer *AssetsSerializer, response *AssetsResponse, wg *sync.WaitGroup) {
	for _, insight := range *assetsSerializer.Insights {
		serializer := NewInsightSerializer()
		serializer.Insight = &insight
		response.Insights = append(response.Insights, *serializer.Response())
	}
	wg.Done()
}

func buildAudienceResponses(assetsSerializer *AssetsSerializer, response *AssetsResponse, wg *sync.WaitGroup) {
	for _, audience := range *assetsSerializer.Audiences {
		serializer := NewAudienceSerializer()
		serializer.Audience = &audience
		response.Audiences = append(response.Audiences, *serializer.Response())
	}
	wg.Done()
}
