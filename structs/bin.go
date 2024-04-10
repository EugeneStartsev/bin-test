package structs

type Bin1 struct {
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"country-code,omitempty"`
	CardBrand    string `json:"card-brand,omitempty"`
	CityIP       string `json:"ip-city,omitempty"`
	IsCommercial bool   `json:"is-commercial,omitempty"`
	BinNumber    string `json:"bin-number,omitempty"`
	Issuer       string `json:"issuer,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	CardCategory string `json:"card-category,omitempty"`
	CurrencyCode string `json:"currency-code,omitempty"`
	CardType     string `json:"card-type,omitempty"`
}

type Bin2 struct {
	Success bool `json:"success,omitempty"`
	Code    int  `json:"code,omitempty"`
	Bin     Bin
}

type Bin struct {
	Valid    bool    `json:"valid,omitempty"`
	Number   int     `json:"number,omitempty"`
	Length   int     `json:"length,omitempty"`
	Brand    string  `json:"brand,omitempty"`
	Type     string  `json:"type,omitempty"`
	Level    string  `json:"level,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Issuer   Issuer  `json:"issuer"`
	Country  Country `json:"country"`
}

type Issuer struct {
	Name    string `json:"name,omitempty"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type Country struct {
	Name     string `json:"name,omitempty"`
	Numeric  string `json:"numeric,omitempty"`
	Capital  string `json:"capital,omitempty"`
	Currency string `json:"currency,omitempty"`
	Region   string `json:"region,omitempty"`
}
