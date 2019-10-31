package assets

// AssetType type
type AssetType int8

// AssetType values
const (
	ChartType    = 0
	InsightType  = 1
	AudienceType = 2
)

func (assetType AssetType) String() string {
	switch assetType {
	case ChartType:
		return "Chart"
	case InsightType:
		return "Insight"
	case AudienceType:
		return "Audience"
	default:
		return "AssetTypeError"
	}
}

// Gender type
type Gender int8

// Gender values
const (
	GenderEmpty = 0
	Male        = 1
	Female      = 2
)

func (gender Gender) String() string {
	switch gender {
	case GenderEmpty:
		return "Not Specified"
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "GenderError"
	}
}

// AgeGroup type
type AgeGroup int8

// AgeGroup values
const (
	AgeGroupEmpty = 0
	From0to12     = 1
	From13to17    = 2
	From18to24    = 3
	From25to34    = 4
	From35to44    = 5
	From45to54    = 6
	From55to64    = 7
	From65plus    = 8
)

func (ageGroup AgeGroup) String() string {
	switch ageGroup {
	case AgeGroupEmpty:
		return "Not Specified"
	case From0to12:
		return "0-12"
	case From13to17:
		return "13-17"
	case From18to24:
		return "18-24"
	case From25to34:
		return "25-34"
	case From35to44:
		return "35-44"
	case From45to54:
		return "45-54"
	case From55to64:
		return "55-64"
	case From65plus:
		return "65+"
	default:
		return "AgeGroupError"
	}
}
