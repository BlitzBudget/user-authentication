package locale

func ConvertToSymbol(currencyName *string) *string {
	currencyNameToSymbol := currencyNameToSymbol()
	var currencySymbol string

	for currencyNameInMap, symbol := range currencyNameToSymbol {
		if currencyNameInMap == *currencyName {
			currencySymbol = symbol
			continue
		}
	}

	return &currencySymbol
}

func currencyNameToSymbol() map[string]string {
	return map[string]string{
		"Afghan Afghani":             "؋",
		"Albanian Lek":               "L",
		"Algerian Dinar":             "دج",
		"US Dollar":                  "$",
		"Angolan Kwanza":             "Kz",
		"East Caribbean Dollar":      "EC$",
		"Argentine Peso":             "ARS$",
		"Armenian Dram":              "֏",
		"Aruban Florin":              "ƒ",
		"Australian Dollar":          "AU$",
		"Azerbaijan Manat":           "₼",
		"Bahamian Dollar":            "B$",
		"Bahraini Dinar":             ".د.ب",
		"Bangladeshi Taka":           "Tk",
		"Barbados Dollar":            "Bds$",
		"Belarusian Ruble":           "Br",
		"Belize Dollar":              "BZ$",
		"CFA Franc":                  "CFA",
		"Bermudian Dollar":           "BD$",
		"Bhutanese Ngultrum":         "Nu.",
		"Boliviano":                  "Bs",
		"Convertible Mark":           "KM",
		"Pula":                       "P",
		"Norwegian Krone":            "kr",
		"Brazilian Real":             "R$",
		"Brunei Dollar":              "B$",
		"Bulgarian Lev":              "Лв.",
		"CFA Franc BCEAO":            "CFA",
		"Burundi Franc":              "FBu",
		"Cape Verde Escudo":          "Esc",
		"Cambodian Riel":             "៛",
		"CFA Franc BEAC":             "FCFA",
		"Canadian Dollar":            "C$",
		"Cayman Islands Dollar":      "CI$",
		"Chilean Peso":               "CLP$",
		"Renminbi (Yuan)":            "¥",
		"Colombian Peso":             "COL$",
		"Comoro Franc":               "CF",
		"Franc Congolais":            "FC",
		"New Zealand Dollar":         "NZ$",
		"Costa Rican Colón":          "₡",
		"Croatian Kuna":              "kn",
		"Cuban Peso":                 "$MN",
		"Antillean Guilder":          "NAƒ",
		"Czech Koruna":               "Kč",
		"Danish Kroner":              "Kr.",
		"Djibouti Franc":             "Fdj",
		"Dominican Peso":             "RD$",
		"Egyptian Pound":             "E£",
		"Euro":                       "€",
		"Nakfa":                      "Nfk",
		"Swazi lilangeni":            "L",
		"Ethiopian Birr":             "Br",
		"Falkland Islands pound":     "FK£",
		"Faroese króna":              "Kr.",
		"Fiji Dollar":                "FJ$",
		"Dalasi":                     "D",
		"Lari":                       "ლ",
		"Cedi":                       "GH₵",
		"Gibraltar Pound":            "GIP£",
		"Danish Krone":               "Kr.",
		"Quetzal":                    "GTQ",
		"Guernsey pound":             "GGP£",
		"Guinea Franc":               "FG",
		"Guyana Dollar":              "GY$",
		"Haitian Gourde":             "G",
		"Lempira":                    "L",
		"Hong Kong Dollar":           "HK$",
		"Forint":                     "Ft",
		"Iceland Krona":              "Íkr",
		"Indian Rupee":               "₹",
		"Rupiah":                     "Rp",
		"Iranian Rial":               "﷼",
		"Iraqi Dinar":                "ع.د",
		"Manx pound":                 "IMP£",
		"New Israeli Sheqel":         "₪",
		"Jamaican Dollar":            "J$",
		"Yen":                        "円",
		"Jersey pound":               "JEP£",
		"Jordanian Dinar":            "د.أ",
		"Tenge":                      "₸",
		"Kenyan Shilling":            "K",
		"North Korean Won":           "₩",
		"Won":                        "₩",
		"Kuwaiti Dinar":              "د.ك",
		"Som":                        "Лв",
		"Kip":                        "₭",
		"Lebanese Pound":             "ل.ل.",
		"Loti":                       "L",
		"Liberian Dollar":            "L$",
		"Lybian Dinar":               "ل.د",
		"Macanese Pataca":            "MOP$",
		"Malagasy Franc":             "Ar",
		"Kwacha":                     "MK",
		"Malaysian Ringgit":          "RM",
		"Rufiyaa":                    ".ރ",
		"Ouguiya":                    "UM",
		"Mauritius Rupee":            "₨",
		"Mexican Peso":               "Mex$",
		"Moldovan Leu":               "L",
		"Mongolia Tughrik":           "₮",
		"Mozambican Metical":         "MT",
		"Burmese Kyat":               "K",
		"Namibia Dollar":             "N$",
		"Nepalese Rupee":             "रू",
		"CFP Franc":                  "F",
		"Nicaraguan Córdoba":         "C$",
		"West African CFA Franc":     "CFA",
		"Nigerian Naira":             "₦",
		"Macedonian Denar":           "Ден",
		"Omani Rial":                 "﷼.",
		"Pakistan Rupee":             "Rp",
		"Panamanian Balboa":          "B/.",
		"Papua New Guinean Kina":     "K",
		"Paraguayan Guarani":         "₲",
		"Peruvian (Nuevo) Sol":       "S/",
		"Philippine Peso":            "₱",
		"Polish Zloty (złoty)":       "zł",
		"Qatari Rial":                "ر.ق",
		"Leu":                        "lei",
		"Russian Ruble":              "₽",
		"Rwanda Franc":               "RF",
		"Saint Helena Pound":         "SHP£",
		"Tala":                       "WS$",
		"Dobra":                      "Db",
		"Saudi Riyal":                "﷼",
		"Serbian Dinar":              "din",
		"Seychelles Rupee":           "SRe",
		"Leone":                      "Le",
		"Singapore Dollar":           "S$",
		"Solomon Islands Dollar":     "Si$",
		"Somali Shilling":            "Sh.so.",
		"Rand":                       "R",
		"South Sudanese pound":       "ج.س.",
		"Sri Lanka Rupee":            "ரூ",
		"Sudanese Pound":             "ج.س.",
		"Suriname Guilder":           "Sr$",
		"Lilangeni":                  "kr.",
		"Swedish Krona":              "kr",
		"Swiss Franc":                "Fr",
		"Syrian Pound":               "£S",
		"New Taiwan Dollar":          "NT$",
		"Somoni":                     "ЅM",
		"Tanzanian Shilling":         "TSh",
		"Thai Baht":                  "฿",
		"Paʻanga":                    "T$",
		"Trinidad and Tobago Dollar": "TT$",
		"Tunisian Dinar":             "TD",
		"Turkish Lira":               "₺",
		"Manat":                      "T",
		"Tuvaluan Dollar":            "TV$",
		"Uganda Shilling":            "USh",
		"Ukrainian Hryvnia":          "₴",
		"UAE Dirham":                 "د.إ",
		"Pound Sterling":             "£",
		"Uruguayan Peso":             "$U",
		"Uzbekistan Sum":             "so\"m",
		"Vanuatu Vatu":               "VT",
		"Venezuelan Bolívar":         "Bs",
		"Vietnamese Dong":            "₫",
		"Moroccan Dirham":            "MAD",
		"Yemeni Rial":                "﷼.",
		"Zambian Kwacha":             "ZK",
		"Zimbabwe Dollar":            "Z$",
	}
}
