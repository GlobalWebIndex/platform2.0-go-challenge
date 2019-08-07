package model

import (
	"bytes"
	"encoding/json"
)

// Characteristics is a helper model for the contents of an Audience
type Characteristics struct {
	Gender       Gender  `json:"gender,omitempty"`
	BirthCountry string  `json:"birth_country,omitempty"`
	AgeGroup     [2]uint `json:"age_group,omitempty"`
	HoursOnSM    uint    `json:"hours_on_sm,omitempty"`
	LastMoPurch  uint    `json:"last_month_purchases,omitempty"`
}

// Gender is a helper type for gender enumeration
type Gender int

const (
	// Male Gender enum
	Male Gender = iota + 1
	// Female Gender enum
	Female
	// Other Gender enum
	Other
)

func (g Gender) String() string {
	var toString = map[Gender]string{
		Male:   "Male",
		Female: "Female",
		Other:  "Other",
	}
	return toString[g]
}

// MarshalJSON marshals the enum as a quoted json string
func (g Gender) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(g.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (g *Gender) UnmarshalJSON(b []byte) error {
	var toID = map[string]Gender{
		"Male":   Male,
		"Female": Female,
		"Other":  Other,
	}
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*g = toID[j]
	return nil
}
