package states

import (
	"errors"
	"strings"
)

// ScrubToAbbr takes a name or abbr and returns an abbr
func ScrubToAbbr(nm string) (string, error) {
	var err error
	orig := nm
	if IsValidName(nm) {
		nm, err = NameToAbbr(nm)
	}
	if !IsValidAbbr(nm) {
		return "", errors.New("Invalid State: " + orig)
	}
	return nm, err
}

// IsValidName returns whether the given string is a valid name
func IsValidName(nm string) bool {
	nm = strings.Title(nm)
	switch nm {
	case Alabama:
		return true
	case Alaska:
		return true
	case Arizona:
		return true
	case Arkansas:
		return true
	case California:
		return true
	case Colorado:
		return true
	case Connecticut:
		return true
	case Delaware:
		return true
	case DistrictOfColumbia:
		return true
	case Florida:
		return true
	case Georgia:
		return true
	case Hawaii:
		return true
	case Idaho:
		return true
	case Illinois:
		return true
	case Indiana:
		return true
	case Iowa:
		return true
	case Kansas:
		return true
	case Kentucky:
		return true
	case Louisiana:
		return true
	case Maine:
		return true
	case Maryland:
		return true
	case Massachusetts:
		return true
	case Michigan:
		return true
	case Minnesota:
		return true
	case Mississippi:
		return true
	case Missouri:
		return true
	case Montana:
		return true
	case Nebraska:
		return true
	case Nevada:
		return true
	case NewHampshire:
		return true
	case NewJersey:
		return true
	case NewMexico:
		return true
	case NewYork:
		return true
	case NorthCarolina:
		return true
	case NorthDakota:
		return true
	case Ohio:
		return true
	case Oklahoma:
		return true
	case Oregon:
		return true
	case Pennsylvania:
		return true
	case PuertoRico:
		return true
	case RhodeIsland:
		return true
	case SouthCarolina:
		return true
	case SouthDakota:
		return true
	case Tennessee:
		return true
	case Texas:
		return true
	case Utah:
		return true
	case Vermont:
		return true
	case Virginia:
		return true
	case Washington:
		return true
	case WestVirginia:
		return true
	case Wisconsin:
		return true
	case Wyoming:
		return true
	}
	return false
}

// IsValidAbbr return whether it was passed a valid abbreviation
func IsValidAbbr(a string) bool {
	a = strings.ToLower(a)
	switch a {
	case AlabamaAbbr:
		return true
	case AlaskaAbbr:
		return true
	case ArizonaAbbr:
		return true
	case ArkansasAbbr:
		return true
	case CaliforniaAbbr:
		return true
	case ColoradoAbbr:
		return true
	case ConnecticutAbbr:
		return true
	case DelawareAbbr:
		return true
	case DistrictOfColumbiaAbbr:
		return true
	case FloridaAbbr:
		return true
	case GeorgiaAbbr:
		return true
	case HawaiiAbbr:
		return true
	case IdahoAbbr:
		return true
	case IllinoisAbbr:
		return true
	case IndianaAbbr:
		return true
	case IowaAbbr:
		return true
	case KansasAbbr:
		return true
	case KentuckyAbbr:
		return true
	case LouisianaAbbr:
		return true
	case MaineAbbr:
		return true
	case MarylandAbbr:
		return true
	case MassachusettsAbbr:
		return true
	case MichiganAbbr:
		return true
	case MinnesotaAbbr:
		return true
	case MississippiAbbr:
		return true
	case MissouriAbbr:
		return true
	case MontanaAbbr:
		return true
	case NebraskaAbbr:
		return true
	case NevadaAbbr:
		return true
	case NewHampshireAbbr:
		return true
	case NewJerseyAbbr:
		return true
	case NewMexicoAbbr:
		return true
	case NewYorkAbbr:
		return true
	case NorthCarolinaAbbr:
		return true
	case NorthDakotaAbbr:
		return true
	case OhioAbbr:
		return true
	case OklahomaAbbr:
		return true
	case OregonAbbr:
		return true
	case PennsylvaniaAbbr:
		return true
	case PuertoRicoAbbr:
		return true
	case RhodeIslandAbbr:
		return true
	case SouthCarolinaAbbr:
		return true
	case SouthDakotaAbbr:
		return true
	case TennesseeAbbr:
		return true
	case TexasAbbr:
		return true
	case UtahAbbr:
		return true
	case VermontAbbr:
		return true
	case VirginiaAbbr:
		return true
	case WashingtonAbbr:
		return true
	case WestVirginiaAbbr:
		return true
	case WisconsinAbbr:
		return true
	case WyomingAbbr:
		return true
	}
	return false
}

// NameToAbbr converts a state name to it's abbreviation
func NameToAbbr(nm string) (string, error) {
	nm = strings.Title(nm)
	switch nm {
	case Alabama:
		return AlabamaAbbr, nil
	case Alaska:
		return AlaskaAbbr, nil
	case Arizona:
		return ArizonaAbbr, nil
	case Arkansas:
		return ArkansasAbbr, nil
	case California:
		return CaliforniaAbbr, nil
	case Colorado:
		return ColoradoAbbr, nil
	case Connecticut:
		return ConnecticutAbbr, nil
	case Delaware:
		return DelawareAbbr, nil
	case DistrictOfColumbia:
		return DistrictOfColumbiaAbbr, nil
	case Florida:
		return FloridaAbbr, nil
	case Georgia:
		return GeorgiaAbbr, nil
	case Hawaii:
		return HawaiiAbbr, nil
	case Idaho:
		return IdahoAbbr, nil
	case Illinois:
		return IllinoisAbbr, nil
	case Indiana:
		return IndianaAbbr, nil
	case Iowa:
		return IowaAbbr, nil
	case Kansas:
		return KansasAbbr, nil
	case Kentucky:
		return KentuckyAbbr, nil
	case Louisiana:
		return LouisianaAbbr, nil
	case Maine:
		return MaineAbbr, nil
	case Maryland:
		return MarylandAbbr, nil
	case Massachusetts:
		return MassachusettsAbbr, nil
	case Michigan:
		return MichiganAbbr, nil
	case Minnesota:
		return MinnesotaAbbr, nil
	case Mississippi:
		return MississippiAbbr, nil
	case Missouri:
		return MissouriAbbr, nil
	case Montana:
		return MontanaAbbr, nil
	case Nebraska:
		return NebraskaAbbr, nil
	case Nevada:
		return NevadaAbbr, nil
	case NewHampshire:
		return NewHampshireAbbr, nil
	case NewJersey:
		return NewJerseyAbbr, nil
	case NewMexico:
		return NewMexicoAbbr, nil
	case NewYork:
		return NewYorkAbbr, nil
	case NorthCarolina:
		return NorthCarolinaAbbr, nil
	case NorthDakota:
		return NorthDakotaAbbr, nil
	case Ohio:
		return OhioAbbr, nil
	case Oklahoma:
		return OklahomaAbbr, nil
	case Oregon:
		return OregonAbbr, nil
	case Pennsylvania:
		return PennsylvaniaAbbr, nil
	case PuertoRico:
		return PuertoRicoAbbr, nil
	case RhodeIsland:
		return RhodeIslandAbbr, nil
	case SouthCarolina:
		return SouthCarolinaAbbr, nil
	case SouthDakota:
		return SouthDakotaAbbr, nil
	case Tennessee:
		return TennesseeAbbr, nil
	case Texas:
		return TexasAbbr, nil
	case Utah:
		return UtahAbbr, nil
	case Vermont:
		return VermontAbbr, nil
	case Virginia:
		return VirginiaAbbr, nil
	case Washington:
		return WashingtonAbbr, nil
	case WestVirginia:
		return WestVirginiaAbbr, nil
	case Wisconsin:
		return WisconsinAbbr, nil
	case Wyoming:
		return WyomingAbbr, nil
	}
	return "", errors.New("Invalid State Name")
}

// AbbrToName converts a state abbreviation to it's name
func AbbrToName(nm string) (string, error) {
	nm = strings.ToLower(nm)
	switch nm {
	case AlabamaAbbr:
		return Alabama, nil
	case AlaskaAbbr:
		return Alaska, nil
	case ArizonaAbbr:
		return Arizona, nil
	case ArkansasAbbr:
		return Arkansas, nil
	case CaliforniaAbbr:
		return California, nil
	case ColoradoAbbr:
		return Colorado, nil
	case ConnecticutAbbr:
		return Connecticut, nil
	case DelawareAbbr:
		return Delaware, nil
	case DistrictOfColumbiaAbbr:
		return DistrictOfColumbia, nil
	case FloridaAbbr:
		return Florida, nil
	case GeorgiaAbbr:
		return Georgia, nil
	case HawaiiAbbr:
		return Hawaii, nil
	case IdahoAbbr:
		return Idaho, nil
	case IllinoisAbbr:
		return Illinois, nil
	case IndianaAbbr:
		return Indiana, nil
	case IowaAbbr:
		return Iowa, nil
	case KansasAbbr:
		return Kansas, nil
	case KentuckyAbbr:
		return Kentucky, nil
	case LouisianaAbbr:
		return Louisiana, nil
	case MaineAbbr:
		return Maine, nil
	case MarylandAbbr:
		return Maryland, nil
	case MassachusettsAbbr:
		return Massachusetts, nil
	case MichiganAbbr:
		return Michigan, nil
	case MinnesotaAbbr:
		return Minnesota, nil
	case MississippiAbbr:
		return Mississippi, nil
	case MissouriAbbr:
		return Missouri, nil
	case MontanaAbbr:
		return Montana, nil
	case NebraskaAbbr:
		return Nebraska, nil
	case NevadaAbbr:
		return Nevada, nil
	case NewHampshireAbbr:
		return NewHampshire, nil
	case NewJerseyAbbr:
		return NewJersey, nil
	case NewMexicoAbbr:
		return NewMexico, nil
	case NewYorkAbbr:
		return NewYork, nil
	case NorthCarolinaAbbr:
		return NorthCarolina, nil
	case NorthDakotaAbbr:
		return NorthDakota, nil
	case OhioAbbr:
		return Ohio, nil
	case OklahomaAbbr:
		return Oklahoma, nil
	case OregonAbbr:
		return Oregon, nil
	case PennsylvaniaAbbr:
		return Pennsylvania, nil
	case PuertoRicoAbbr:
		return PuertoRico, nil
	case RhodeIslandAbbr:
		return RhodeIsland, nil
	case SouthCarolinaAbbr:
		return SouthCarolina, nil
	case SouthDakotaAbbr:
		return SouthDakota, nil
	case TennesseeAbbr:
		return Tennessee, nil
	case TexasAbbr:
		return Texas, nil
	case UtahAbbr:
		return Utah, nil
	case VermontAbbr:
		return Vermont, nil
	case VirginiaAbbr:
		return Virginia, nil
	case WashingtonAbbr:
		return Washington, nil
	case WestVirginiaAbbr:
		return WestVirginia, nil
	case WisconsinAbbr:
		return Wisconsin, nil
	case WyomingAbbr:
		return Wyoming, nil
	}
	return "", errors.New("Invalid State Abbreviation")
}

// State Name Constants
const (
	Alabama            = "Alabama"
	Alaska             = "Alaska"
	Arizona            = "Arizona"
	Arkansas           = "Arkansas"
	California         = "California"
	Colorado           = "Colorado"
	Connecticut        = "Connecticut"
	Delaware           = "Delaware"
	DistrictOfColumbia = "District of Columbia"
	Florida            = "Florida"
	Georgia            = "Georgia"
	Hawaii             = "Hawaii"
	Idaho              = "Idaho"
	Illinois           = "Illinois"
	Indiana            = "Indiana"
	Iowa               = "Iowa"
	Kansas             = "Kansas"
	Kentucky           = "Kentucky"
	Louisiana          = "Louisiana"
	Maine              = "Maine"
	Maryland           = "Maryland"
	Massachusetts      = "Massachusetts"
	Michigan           = "Michigan"
	Minnesota          = "Minnesota"
	Mississippi        = "Mississippi"
	Missouri           = "Missouri"
	Montana            = "Montana"
	Nebraska           = "Nebraska"
	Nevada             = "Nevada"
	NewHampshire       = "New Hampshire"
	NewJersey          = "New Jersey"
	NewMexico          = "New Mexico"
	NewYork            = "New York"
	NorthCarolina      = "North Carolina"
	NorthDakota        = "North Dakota"
	Ohio               = "Ohio"
	Oklahoma           = "Oklahoma"
	Oregon             = "Oregon"
	Pennsylvania       = "Pennsylvania"
	PuertoRico         = "Puerto Rico"
	RhodeIsland        = "Rhode Island"
	SouthCarolina      = "South Carolina"
	SouthDakota        = "South Dakota"
	Tennessee          = "Tennessee"
	Texas              = "Texas"
	Utah               = "Utah"
	Vermont            = "Vermont"
	Virginia           = "Virginia"
	Washington         = "Washington"
	WestVirginia       = "West Virginia"
	Wisconsin          = "Wisconsin"
	Wyoming            = "Wyoming"
)

// State Abbreviation Constants
const (
	AlabamaAbbr            = "al"
	AlaskaAbbr             = "ak"
	ArizonaAbbr            = "az"
	ArkansasAbbr           = "ar"
	CaliforniaAbbr         = "ca"
	ColoradoAbbr           = "co"
	ConnecticutAbbr        = "ct"
	DelawareAbbr           = "de"
	DistrictOfColumbiaAbbr = "dc"
	FloridaAbbr            = "fl"
	GeorgiaAbbr            = "ga"
	HawaiiAbbr             = "hi"
	IdahoAbbr              = "id"
	IllinoisAbbr           = "il"
	IndianaAbbr            = "in"
	IowaAbbr               = "io"
	KansasAbbr             = "ks"
	KentuckyAbbr           = "ky"
	LouisianaAbbr          = "la"
	MaineAbbr              = "me"
	MarylandAbbr           = "md"
	MassachusettsAbbr      = "ma"
	MichiganAbbr           = "mi"
	MinnesotaAbbr          = "mn"
	MississippiAbbr        = "ms"
	MissouriAbbr           = "mo"
	MontanaAbbr            = "mt"
	NebraskaAbbr           = "ne"
	NevadaAbbr             = "nv"
	NewHampshireAbbr       = "nh"
	NewJerseyAbbr          = "nj"
	NewMexicoAbbr          = "nm"
	NewYorkAbbr            = "ny"
	NorthCarolinaAbbr      = "nc"
	NorthDakotaAbbr        = "nd"
	OhioAbbr               = "oh"
	OklahomaAbbr           = "ok"
	OregonAbbr             = "or"
	PennsylvaniaAbbr       = "pa"
	PuertoRicoAbbr         = "pr"
	RhodeIslandAbbr        = "ri"
	SouthCarolinaAbbr      = "sc"
	SouthDakotaAbbr        = "sd"
	TennesseeAbbr          = "tn"
	TexasAbbr              = "tx"
	UtahAbbr               = "ut"
	VermontAbbr            = "vt"
	VirginiaAbbr           = "va"
	WashingtonAbbr         = "wa"
	WestVirginiaAbbr       = "wv"
	WisconsinAbbr          = "wi"
	WyomingAbbr            = "wy"
)
