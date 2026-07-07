package common

import (
	"strings"
)

type TranslationFactory struct{}

func (t TranslationFactory) GetTable(key string) (table string) {
	key = strings.ToLower(key)

	switch key {
	case "t_asv", "asv":
		table = "t_asv"
	case "t_bbe", "bbe":
		table = "t_bbe"
	case "t_dby", "dby":
		table = "t_dby"
	case "t_kjv", "kjv":
		table = "t_kjv"
	case "t_wbt", "wbt":
		table = "t_wbt"
	case "t_web", "web":
		table = "t_web"
	case "t_ylt", "ylt":
		table = "t_ylt"
	case "t_esv", "esv":
		table = "t_esv"
	case "t_niv", "niv":
		table = "t_niv"
	case "t_nlt", "nlt":
		table = "t_nlt"
	case "t_nlt_2015", "nlt2015":
		table = "t_nlt_2015"
	default:
		table = "t_kjv"
	}

	return
}

func (t TranslationFactory) GetIndex(key string) (index string) {
	key = strings.ToLower(key)

	switch key {
	case "t_asv", "asv":
		index = "asv"
	case "t_bbe", "bbe":
		index = "bbe"
	case "t_dby", "dby":
		index = "dby"
	case "t_kjv", "kjv":
		index = "kjv"
	case "t_wbt", "wbt":
		index = "wbt"
	case "t_web", "web":
		index = "web"
	case "t_ylt", "ylt":
		index = "ylt"
	case "t_esv", "esv":
		index = "esv"
	case "t_niv", "niv":
		index = "niv"
	case "t_nlt", "nlt":
		index = "nlt"
	case "t_nlt_2015", "nlt2015":
		index = "nlt2015"
	default:
		index = "kjv"
	}

	return
}
