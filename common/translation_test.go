package common

import "testing"

func TestGetTable(t *testing.T) {
	f := TranslationFactory{}

	cases := []struct {
		input string
		want  string
	}{
		{"asv", "t_asv"},
		{"ASV", "t_asv"},
		{"t_asv", "t_asv"},
		{"bbe", "t_bbe"},
		{"BBE", "t_bbe"},
		{"t_bbe", "t_bbe"},
		{"dby", "t_dby"},
		{"DBY", "t_dby"},
		{"t_dby", "t_dby"},
		{"kjv", "t_kjv"},
		{"KJV", "t_kjv"},
		{"t_kjv", "t_kjv"},
		{"wbt", "t_wbt"},
		{"WBT", "t_wbt"},
		{"t_wbt", "t_wbt"},
		{"web", "t_web"},
		{"WEB", "t_web"},
		{"t_web", "t_web"},
		{"ylt", "t_ylt"},
		{"YLT", "t_ylt"},
		{"t_ylt", "t_ylt"},
		{"esv", "t_esv"},
		{"ESV", "t_esv"},
		{"t_esv", "t_esv"},
		{"niv", "t_niv"},
		{"NIV", "t_niv"},
		{"t_niv", "t_niv"},
		{"nlt", "t_nlt"},
		{"NLT", "t_nlt"},
		{"t_nlt", "t_nlt"},
		{"nlt2015", "t_nlt_2015"},
		{"NLT2015", "t_nlt_2015"},
		{"t_nlt_2015", "t_nlt_2015"},
		{"", "t_kjv"},
		{"unknown", "t_kjv"},
	}

	for _, c := range cases {
		got := f.GetTable(c.input)
		if got != c.want {
			t.Errorf("GetTable(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

func TestGetIndex(t *testing.T) {
	f := TranslationFactory{}

	cases := []struct {
		input string
		want  string
	}{
		{"asv", "asv"},
		{"ASV", "asv"},
		{"t_asv", "asv"},
		{"bbe", "bbe"},
		{"BBE", "bbe"},
		{"t_bbe", "bbe"},
		{"dby", "dby"},
		{"DBY", "dby"},
		{"t_dby", "dby"},
		{"kjv", "kjv"},
		{"KJV", "kjv"},
		{"t_kjv", "kjv"},
		{"wbt", "wbt"},
		{"WBT", "wbt"},
		{"t_wbt", "wbt"},
		{"web", "web"},
		{"WEB", "web"},
		{"t_web", "web"},
		{"ylt", "ylt"},
		{"YLT", "ylt"},
		{"t_ylt", "ylt"},
		{"esv", "esv"},
		{"ESV", "esv"},
		{"t_esv", "esv"},
		{"niv", "niv"},
		{"NIV", "niv"},
		{"t_niv", "niv"},
		{"nlt", "nlt"},
		{"NLT", "nlt"},
		{"t_nlt", "nlt"},
		{"nlt2015", "nlt2015"},
		{"NLT2015", "nlt2015"},
		{"t_nlt_2015", "nlt2015"},
		{"", "kjv"},
		{"unknown", "kjv"},
	}

	for _, c := range cases {
		got := f.GetIndex(c.input)
		if got != c.want {
			t.Errorf("GetIndex(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}
