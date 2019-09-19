package country

type Country struct {
	CurrencyName 		string `json:"currency_name"`
	Language 			string `json:"languages"`
	Name				string `json:"country"`
	CountryId			string `json:"country_id"`
	Phone 				string `json:"phone"`
	PostalCodeFormat	string `json:"postal_code_format"`
	Capital				string `json:"capital"`
	Iso 				string `json:"iso"`
	CurrencySymbol		string `json:"currency_symbol"`
	CurrencyCode		string `json:"currency_code"`
} 