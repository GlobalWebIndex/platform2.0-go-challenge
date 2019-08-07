package model

import (
	"encoding/json"
	"gwi/utils"
)

// Asset is the struct for the asset model.
type Asset struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	*Chart      `json:"chart,omitempty"`
	*Insight    `json:"insight,omitempty"`
	*Audience   `json:"audience,omitempty"`
}

func (a *Asset) fieldsValid() bool {
	cnt := 0
	if a.Chart != nil {
		cnt++
	}
	if a.Insight != nil {
		cnt++
	}
	if a.Audience != nil {
		cnt++
	}
	return cnt <= 1
}

// IsValid is a helper function validating that only one of Chart, Insight and Audience has Content
func (a *Asset) IsValid() error {
	if a.fieldsValid() {
		return nil
	}
	return utils.ErrInvalidAsset
}

// UnmarshalJSON is a custom json unmarhaler taking into account the mutual exclusion between Chart, Insight and Audience
func (a *Asset) UnmarshalJSON(data []byte) error {
	type asset Asset
	if err := json.Unmarshal(data, (*asset)(a)); err != nil {
		return err
	}
	return a.IsValid()
}

// Chart is the struct for the chart model.
type Chart struct {
	Title      string      `json:"title"`
	AxisTitles [2]string   `json:"axis_titles"`
	Data       interface{} `json:"data"`
}

// Insight is the struct for the insight model.
type Insight struct {
	Insight string `json:"insight"`
}

// Audience is the struct for the audience model.
type Audience struct {
	Characteristics Characteristics `json:"characteristics"`
}
