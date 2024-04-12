package structs

type BinData struct {
	Bin      string `json:"bin,omitempty" db:"bin-id"`
	Brand    string `json:"brand" db:"brand"`
	Type     string `json:"type" db:"type"`
	Category string `json:"category" db:"category"`
	Issuer   string `json:"issuer" db:"issuer"`
	Alpha2   string `json:"alpha2" db:"alpha_2"`
	Alpha3   string `json:"alpha_3" db:"alpha_3"`
	Country  string `json:"country" db:"country"`
}

type SaveBinData struct {
	Number   Number  `json:"number"`
	Scheme   string  `json:"scheme,omitempty"`
	Type     string  `json:"type,omitempty"`
	Category string  `json:"category,omitempty"`
	Country  Country `json:"country"`
	Bank     Bank    `json:"bank"`
	Success  bool    `json:"success,omitempty"`
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
