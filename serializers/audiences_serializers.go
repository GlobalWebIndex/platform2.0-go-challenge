package serializers

import "gwi-challenge/data/models"

type AudienceSerializer struct {
	*models.Audience
}

type AudienceResponse struct {
	Title           string `json:"title"`
	AudienceLiteral string `json:"audience"`
}

func NewAudienceSerializer() AudienceSerializer {
	return AudienceSerializer{}
}

func (serializer *AudienceSerializer) Response() *AudienceResponse {
	response := AudienceResponse{
		Title:           serializer.Title,
		AudienceLiteral: serializer.Audience.ComposeAudienceLiteral(),
	}
	return &response
}
