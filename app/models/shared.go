package models

import (
	"fmt"
	"math"

	"gwiapi/app/helpers"
)

func InsightTitle(c Insight, cd InsightDetails) string {
	res := ""
	// Result Percentage
	if cd.FiltererSample == 0 {
		return "Insuficient sample for these criteria"
	}
	percentage := uint64(cd.FiltererSample / cd.Sample * 100)
	res += fmt.Sprintf("%d %% of ", percentage)

	// Filter Section
	res += CommonDetailsTitle(c.Gender, c.CountryCode, c.AgeFrom, c.AgeTo)

	// Result Filter
	switch c.HoursComparator {
	case ">=":
		res += fmt.Sprintf(" spend %s hour(s) or more ", string(c.HoursReference))
	case ">":
		res += fmt.Sprintf(" spend more than %s hour(s)", string(c.HoursReference))
	case "<=":
		res += fmt.Sprintf(" spend %s hour(s) or less ", string(c.HoursReference))
	case "<":
		res += fmt.Sprintf(" spend less than %s hour(s)", string(c.HoursReference))
	default:
	}

	res = res + " on social media daily"

	return res
}
func AudienceTitle(c Audience, cd AudienceDetails) string {
	res := ""

	// Filter Section
	res += CommonDetailsTitle(c.Gender, c.CountryCode, c.AgeFrom, c.AgeTo)

	// Result
	if uint64(math.Floor(cd.Result)) == 0 {
		res += "spent less than 1 hour on social media daily"
	}
	if uint64(math.Floor(cd.Result)) >= 1 {
		res += fmt.Sprintf("spent more than %d hours on social media daily", uint64(math.Floor(cd.Result)))
	}
	return res
}

func CommonSqlFilterBuilder(gender, countryCode string, ageFrom, ageTo uint8) string {
	// This is a bad idea (values must be passed to ORM for sanitization)
	// SQL Injection
	res := "1 = 1"
	if gender != "" {
		res += fmt.Sprintf(" and gender = %s", gender)
	}
	if ageFrom != 0 {
		res += fmt.Sprintf(" and age >= %d", ageFrom)
	}
	if ageTo != 0 {
		res += fmt.Sprintf(" and age <= %d", ageTo)
	}
	if countryCode != "" {
		res += fmt.Sprintf(" and country_code = %s", countryCode)
	}

	return res
}

func CommonDetailsTitle(gender, countryCode string, ageFrom, ageTo uint8) string {
	res := ""
	if gender != "" {
		res += helpers.StringTernary(gender == "f", "Females", "Males")
	}
	if gender == "" {
		res += "People"
	}

	if ageFrom != 0 && ageTo != 0 {
		res += fmt.Sprintf(" at the age %d-%d ", ageFrom, ageTo)
	}
	if ageFrom > 0 && ageTo == 0 {
		res += fmt.Sprintf(" older than %d ", ageFrom)
	}
	if ageFrom == 0 && ageTo > 0 {
		res += fmt.Sprintf(" younger than %d ", ageTo)
	}

	if countryCode != "" {
		res += fmt.Sprintf(" from %s ", countryCode)
	}
	return res
}

func FindAsset(assets []Asset, id uint32) int {
	for i, a := range assets {
		if a.ID == id {
			return i
		}
	}
	return 0
}
