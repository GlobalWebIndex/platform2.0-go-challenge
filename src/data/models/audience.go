package models

import "strconv"

type Gender uint

const (
	Male = Gender(iota)
	Female
	genderEnd
)

func (g Gender) Valid() bool {
	return uint(g) < uint(genderEnd)
}

func (g Gender) String(plural bool) string {
	pluralLiterals := []string{"Males", "Females"}
	singularLiterals := []string{"Male", "Female"}
	i := uint(g)
	switch {
	case i < uint(genderEnd):
		if plural {
			return pluralLiterals[i]
		} else {
			return singularLiterals[i]
		}
	default:
		return strconv.Itoa(int(i))
	}
}

type AudienceInfoType uint

const (
	HoursOnSocialMedia = AudienceInfoType(iota)
	PurchasesPerMonth
	audienceInfoTypeEnd
)

func (a AudienceInfoType) Valid() bool {
	return uint(a) < uint(audienceInfoTypeEnd)
}

type Audience struct {
	ID                 uint                `gorm:"primary_key;"`
	Title              string              `gorm:"column:title;unique_index;not null;"`
	FavoritedBy        []FavoritedAudience `gorm:"foreignkey:AudienceID"`
	AudienceInfo       *AudienceInfo       `gorm:"foreignkey:AudienceID"`
	BirthCountry       string              `gorm:"column:birth_country;not null;"`
	Gender             Gender              `gorm:"column:gender;not null;"`
	AgeGroupUpperLimit uint                `gorm:"column:age_group_upper_limit;not null;"`
	AgeGroupLowerLimit uint                `gorm:"column:age_group_lower_limit;not null;"`
}

type AudienceInfo struct {
	ID               uint             `gorm:"primary_key;"`
	AudienceID       uint             `gorm:"column:audience_id;not null;"`
	AudienceInfoType AudienceInfoType `gorm:"column:audience_info_type;not null;"`
	AudienceInfoStat uint             `gorm:"column:audience_info_stat;not null;"`
}

func (audienceToCompose *Audience) ComposeAudienceLiteral() string {
	responseLiteral := "On " + audienceToCompose.BirthCountry + ", "
	responseLiteral += audienceToCompose.Gender.String(true)
	responseLiteral += " from "
	responseLiteral += strconv.FormatUint(uint64(audienceToCompose.AgeGroupLowerLimit), 10) + "-" + strconv.FormatUint(uint64(audienceToCompose.AgeGroupUpperLimit), 10)

	if audienceToCompose.AudienceInfo.AudienceInfoType == HoursOnSocialMedia {
		responseLiteral += " that spent more than " + strconv.FormatUint(uint64(audienceToCompose.AudienceInfo.AudienceInfoStat), 10) + " hours on social media daily"
	} else {
		responseLiteral += " that make more than " + strconv.FormatUint(uint64(audienceToCompose.AudienceInfo.AudienceInfoStat), 10) + " purchases per month"
	}
	return responseLiteral
}
