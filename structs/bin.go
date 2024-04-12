package structs

type SaveBinData struct {
	Number      *Number  `json:"number,omitempty"`
	Iin         string   `json:"bin-id" db:"bin-id"`
	Scheme      string   `json:"scheme,omitempty" db:"brand"`
	Type        string   `json:"type,omitempty" db:"type"`
	Category    string   `json:"category,omitempty" db:"category"`
	Issuer      string   `json:"issuer,omitempty" db:"issuer"`
	Alpha2      string   `json:"alpha2,omitempty" db:"alpha_2"`
	Alpha3      string   `json:"alpha3,omitempty" db:"alpha_3"`
	CountryName string   `json:"country_name,omitempty" db:"country"`
	CountryInfo *Country `json:"country_info,omitempty"`
	Bank        *Bank    `json:"bank,omitempty"`
	Success     bool     `json:"success,omitempty"`
}

type Country struct {
	Alpha2 string `json:"alpha2,omitempty"`
	Alpha3 string `json:"alpha3,omitempty"`
	Name   string `json:"name,omitempty"`
	Emoji  string `json:"emoji,omitempty"`
}

type Bank struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Url   string `json:"url,omitempty"`
}

type Number struct {
	Iin    string `json:"iin,omitempty"`
	Length int    `json:"length,omitempty"`
	Luhn   bool   `json:"luhn,omitempty"`
}
