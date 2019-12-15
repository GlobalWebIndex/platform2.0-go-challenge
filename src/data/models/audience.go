package models

import "strconv"

const (
	Male   = 1
	Female = 2
)

var genderText = map[uint]string{
	Male:   "Male",
	Female: "Female",
}

var genderTextPlural = map[uint]string{
	Male:   "Males",
	Female: "Females",
}

const (
	HoursOnSocialMedia = 1
	PurchasesPerMonth  = 2
)

type Audience struct {
	ID                 uint                `gorm:"primary_key;"`
	Title              string              `gorm:"column:title;unique_index;not null;"`
	FavoritedBy        []FavoritedAudience `gorm:"foreignkey:AudienceID"`
	AudienceInfo       *AudienceInfo       `gorm:"foreignkey:AudienceID"`
	BirthCountry       string              `gorm:"column:birth_country;not null;"`
	Gender             uint                `gorm:"column:gender;not null;"`
	AgeGroupUpperLimit uint                `gorm:"column:age_group_upper_limit;not null;"`
	AgeGroupLowerLimit uint                `gorm:"column:age_group_lower_limit;not null;"`
}

type AudienceInfo struct {
	ID               uint `gorm:"primary_key;"`
	AudienceID       uint `gorm:"column:audience_id;not null;"`
	AudienceInfoType uint `gorm:"column:audience_info_type;not null;"`
	AudienceInfoStat uint `gorm:"column:audience_info_stat;not null;"`
}

func (audienceToCompose *Audience) ComposeAudienceLiteral() string {
	responseLiteral := "On " + audienceToCompose.BirthCountry + ", "
	responseLiteral += genderTextPlural[audienceToCompose.Gender]
	responseLiteral += " from "
	responseLiteral += strconv.FormatUint(uint64(audienceToCompose.AgeGroupLowerLimit), 10) + "-" + strconv.FormatUint(uint64(audienceToCompose.AgeGroupUpperLimit), 10)

	if audienceToCompose.AudienceInfo.AudienceInfoType == HoursOnSocialMedia {
		responseLiteral += " that spent more than " + strconv.FormatUint(uint64(audienceToCompose.AudienceInfo.AudienceInfoStat), 10) + " hours on social media daily"
	} else {
		responseLiteral += " that make more than " + strconv.FormatUint(uint64(audienceToCompose.AudienceInfo.AudienceInfoStat), 10) + " purchases per month"
	}
	return responseLiteral
}

func IsValidGender(genderToValidate uint) bool {
	_, ok := genderText[genderToValidate]
	return ok
}
