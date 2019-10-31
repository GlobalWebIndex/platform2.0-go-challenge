package assets

import (
	"bytes"
	"fmt"
)

// Audience structure
type Audience struct {
	Asset
	Gender
	birthCountry string
	AgeGroup
	hoursOnSocialMedia uint8
	lastMonthPurchases uint32
}

// CreateNewAudience creates a new audience asset
func CreateNewAudience(assetID string, description string, gender Gender, birthCountry string, ageGroup AgeGroup, hoursOnSocialMedia uint8, lastMonthPurchases uint32) Audience {
	return Audience{
		Asset: Asset{
			assetID:     assetID,
			AssetType:   AudienceType,
			description: description},
		Gender:             gender,
		birthCountry:       birthCountry,
		AgeGroup:           ageGroup,
		hoursOnSocialMedia: hoursOnSocialMedia,
		lastMonthPurchases: lastMonthPurchases}
}

// GetAsset returns pointer of basic asset structure
func (audience Audience) GetAsset() *Asset {
	return &audience.Asset
}

// GetDetails returns unique fields info of audience
func (audience Audience) GetDetails() string {
	var out bytes.Buffer
	out.WriteString(audience.Asset.GetDetails())
	out.WriteString("\tGender: '" + audience.Gender.String() + "'\n")
	out.WriteString("\tBirth Country: '" + audience.birthCountry + "'\n")
	out.WriteString("\tAge Group: '" + audience.AgeGroup.String() + "'\n")
	out.WriteString("\tHours On Social Media: '" + fmt.Sprint(audience.hoursOnSocialMedia) + "'\n")
	out.WriteString("\tLast Month Purchases: '" + fmt.Sprint(audience.lastMonthPurchases) + "'\n")
	return out.String()
}
