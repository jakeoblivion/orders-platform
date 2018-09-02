package main

type Country struct {
	countryCode string
	countryName string
}

func ebayCountryCodeConverter(countryCodeFromEbay string) string {
	var countries = []Country{
		Country{
			countryCode: "AA",
			countryName: "NOTE: APO/FPO was defined in eBay list previously but they are not defined in ISO 3166 This country will remain on eBay country code list for backward compatibility",
		},
		Country{
			countryCode: "AD",
			countryName: "Andorra",
		},
		Country{
			countryCode: "AE",
			countryName: "United Arab Emirates",
		},
		Country{
			countryCode: "AF",
			countryName: "Afghanistan",
		},
		Country{
			countryCode: "AG",
			countryName: "Antigua and Barbuda",
		},
		Country{
			countryCode: "AI",
			countryName: "Anguilla",
		},
		Country{
			countryCode: "AL",
			countryName: "Albania",
		},
		Country{
			countryCode: "AM",
			countryName: "Armenia",
		},
		Country{
			countryCode: "AN",
			countryName: "Netherlands Antilles",
		},
		Country{
			countryCode: "AO",
			countryName: "Angola",
		},
		Country{
			countryCode: "AQ",
			countryName: "Antarctica",
		},
		Country{
			countryCode: "AR",
			countryName: "Argentina",
		},
		Country{
			countryCode: "AS",
			countryName: "American Samoa",
		},
		Country{
			countryCode: "AT",
			countryName: "Austria",
		},
		Country{
			countryCode: "AU",
			countryName: "Australia",
		},
		Country{
			countryCode: "AW",
			countryName: "Aruba",
		},
		Country{
			countryCode: "AZ",
			countryName: "Azerbaijan",
		},
		Country{
			countryCode: "BA",
			countryName: "Bosnia and Herzegovina",
		},
		Country{
			countryCode: "BB",
			countryName: "Barbados",
		},
		Country{
			countryCode: "BD",
			countryName: "Bangladesh",
		},
		Country{
			countryCode: "BE",
			countryName: "Belgium",
		},
		Country{
			countryCode: "BF",
			countryName: "Burkina Faso",
		},
		Country{
			countryCode: "BG",
			countryName: "Bulgaria",
		},
		Country{
			countryCode: "BH",
			countryName: "Bahrain",
		},
		Country{
			countryCode: "BI",
			countryName: "Burundi",
		},
		Country{
			countryCode: "BJ",
			countryName: "Benin",
		},
		Country{
			countryCode: "BM",
			countryName: "Bermuda",
		},
		Country{
			countryCode: "BN",
			countryName: "Brunei Darussalam",
		},
		Country{
			countryCode: "BO",
			countryName: "Bolivia",
		},
		Country{
			countryCode: "BR",
			countryName: "Brazil",
		},
		Country{
			countryCode: "BS",
			countryName: "Bahamas",
		},
		Country{
			countryCode: "BT",
			countryName: "Bhutan",
		},
		Country{
			countryCode: "BV",
			countryName: "Bouvet Island",
		},
		Country{
			countryCode: "BW",
			countryName: "Botswana",
		},
		Country{
			countryCode: "BY",
			countryName: "Belarus",
		},
		Country{
			countryCode: "BZ",
			countryName: "Belize",
		},
		Country{
			countryCode: "CA",
			countryName: "Canada",
		},
		Country{
			countryCode: "CC",
			countryName: "Cocos (Keeling) Islands",
		},
		Country{
			countryCode: "CD",
			countryName: "Congo, The Democratic Republic of the",
		},
		Country{
			countryCode: "CF",
			countryName: "Central African Republic",
		},
		Country{
			countryCode: "CG",
			countryName: "Congo",
		},
		Country{
			countryCode: "CH",
			countryName: "Switzerland",
		},
		Country{
			countryCode: "CI",
			countryName: "Cote d'Ivoire",
		},
		Country{
			countryCode: "CK",
			countryName: "Cook Islands",
		},
		Country{
			countryCode: "CL",
			countryName: "Chile",
		},
		Country{
			countryCode: "CM",
			countryName: "Cameroon",
		},
		Country{
			countryCode: "CN",
			countryName: "China",
		},
		Country{
			countryCode: "CO",
			countryName: "Colombia",
		},
		Country{
			countryCode: "CR",
			countryName: "Costa Rica",
		},
		Country{
			countryCode: "CU",
			countryName: "Cuba",
		},
		Country{
			countryCode: "Cu",
			countryName: "tomCode	Reserved for internal or future use",
		},
		Country{
			countryCode: "CV",
			countryName: "Cape Verde",
		},
		Country{
			countryCode: "CX",
			countryName: "Christmas Island",
		},
		Country{
			countryCode: "CY",
			countryName: "Cyprus",
		},
		Country{
			countryCode: "CZ",
			countryName: "Czech Republic",
		},
		Country{
			countryCode: "DE",
			countryName: "Germany",
		},
		Country{
			countryCode: "DJ",
			countryName: "Djibouti",
		},
		Country{
			countryCode: "DK",
			countryName: "Denmark",
		},
		Country{
			countryCode: "DM",
			countryName: "Dominica",
		},
		Country{
			countryCode: "DO",
			countryName: "Dominican Republic",
		},
		Country{
			countryCode: "DZ",
			countryName: "Algeria",
		},
		Country{
			countryCode: "EC",
			countryName: "Ecuador",
		},
		Country{
			countryCode: "EE",
			countryName: "Estonia",
		},
		Country{
			countryCode: "EG",
			countryName: "Egypt",
		},
		Country{
			countryCode: "EH",
			countryName: "Western Sahara",
		},
		Country{
			countryCode: "ER",
			countryName: "Eritrea",
		},
		Country{
			countryCode: "ES",
			countryName: "Spain",
		},
		Country{
			countryCode: "ET",
			countryName: "Ethiopia",
		},
		Country{
			countryCode: "FI",
			countryName: "Finland",
		},
		Country{
			countryCode: "FJ",
			countryName: "Fiji",
		},
		Country{
			countryCode: "FK",
			countryName: "Falkland Islands (Malvinas)",
		},
		Country{
			countryCode: "FM",
			countryName: "Federated States of Micronesia",
		},
		Country{
			countryCode: "FO",
			countryName: "Faroe Islands",
		},
		Country{
			countryCode: "FR",
			countryName: "France",
		},
		Country{
			countryCode: "GA",
			countryName: "Gabon",
		},
		Country{
			countryCode: "GB",
			countryName: "United Kingdom",
		},
		Country{
			countryCode: "GD",
			countryName: "Grenada",
		},
		Country{
			countryCode: "GE",
			countryName: "Georgia",
		},
		Country{
			countryCode: "GF",
			countryName: "French Guiana",
		},
		Country{
			countryCode: "GG",
			countryName: "Guernsey",
		},
		Country{
			countryCode: "GH",
			countryName: "Ghana",
		},
		Country{
			countryCode: "GI",
			countryName: "Gibraltar",
		},
		Country{
			countryCode: "GL",
			countryName: "Greenland",
		},
		Country{
			countryCode: "GM",
			countryName: "Gambia",
		},
		Country{
			countryCode: "GN",
			countryName: "Guinea",
		},
		Country{
			countryCode: "GP",
			countryName: "Guadeloupe",
		},
		Country{
			countryCode: "GQ",
			countryName: "Equatorial Guinea",
		},
		Country{
			countryCode: "GR",
			countryName: "Greece",
		},
		Country{
			countryCode: "GS",
			countryName: "South Georgia and the South Sandwich Islands",
		},
		Country{
			countryCode: "GT",
			countryName: "Guatemala",
		},
		Country{
			countryCode: "GU",
			countryName: "Guam",
		},
		Country{
			countryCode: "GW",
			countryName: "Guinea-Bissau",
		},
		Country{
			countryCode: "GY",
			countryName: "Guyana",
		},
		Country{
			countryCode: "HK",
			countryName: "Hong Kong",
		},
		Country{
			countryCode: "HM",
			countryName: "Heard Island and McDonald Islands",
		},
		Country{
			countryCode: "HN",
			countryName: "Honduras",
		},
		Country{
			countryCode: "HR",
			countryName: "Croatia",
		},
		Country{
			countryCode: "HT",
			countryName: "Haiti",
		},
		Country{
			countryCode: "HU",
			countryName: "Hungary",
		},
		Country{
			countryCode: "ID",
			countryName: "Indonesia",
		},
		Country{
			countryCode: "IE",
			countryName: "Ireland",
		},
		Country{
			countryCode: "IL",
			countryName: "Israel",
		},
		Country{
			countryCode: "IN",
			countryName: "India",
		},
		Country{
			countryCode: "IO",
			countryName: "British Indian Ocean Territory",
		},
		Country{
			countryCode: "IQ",
			countryName: "Iraq",
		},
		Country{
			countryCode: "IR",
			countryName: "Islamic Republic of Iran",
		},
		Country{
			countryCode: "IS",
			countryName: "Iceland",
		},
		Country{
			countryCode: "IT",
			countryName: "Italy",
		},
		Country{
			countryCode: "JE",
			countryName: "Jersey",
		},
		Country{
			countryCode: "JM",
			countryName: "Jamaica",
		},
		Country{
			countryCode: "JO",
			countryName: "Jordan",
		},
		Country{
			countryCode: "JP",
			countryName: "Japan",
		},
		Country{
			countryCode: "KE",
			countryName: "Kenya",
		},
		Country{
			countryCode: "KG",
			countryName: "Kyrgyzstan",
		},
		Country{
			countryCode: "KH",
			countryName: "Cambodia",
		},
		Country{
			countryCode: "KI",
			countryName: "Kiribati",
		},
		Country{
			countryCode: "KM",
			countryName: "Comoros",
		},
		Country{
			countryCode: "KN",
			countryName: "Saint Kitts and Nevis",
		},
		Country{
			countryCode: "KP",
			countryName: "Democratic People's Republic of Korea",
		},
		Country{
			countryCode: "KR",
			countryName: "Republic of Korea",
		},
		Country{
			countryCode: "KW",
			countryName: "Kuwait",
		},
		Country{
			countryCode: "KY",
			countryName: "Cayman Islands",
		},
		Country{
			countryCode: "KZ",
			countryName: "Kazakhstan",
		},
		Country{
			countryCode: "LA",
			countryName: "Lao People's Democratic Republic",
		},
		Country{
			countryCode: "LB",
			countryName: "Lebanon",
		},
		Country{
			countryCode: "LC",
			countryName: "Saint Lucia",
		},
		Country{
			countryCode: "LI",
			countryName: "Liechtenstein",
		},
		Country{
			countryCode: "LK",
			countryName: "Sri Lanka",
		},
		Country{
			countryCode: "LR",
			countryName: "Liberia",
		},
		Country{
			countryCode: "LS",
			countryName: "Lesotho",
		},
		Country{
			countryCode: "LT",
			countryName: "Lithuania",
		},
		Country{
			countryCode: "LU",
			countryName: "Luxembourg",
		},
		Country{
			countryCode: "LV",
			countryName: "Latvia",
		},
		Country{
			countryCode: "LY",
			countryName: "Libyan Arab Jamahiriya",
		},
		Country{
			countryCode: "MA",
			countryName: "Morocco",
		},
		Country{
			countryCode: "MC",
			countryName: "Monaco",
		},
		Country{
			countryCode: "MD",
			countryName: "Republic of Moldova",
		},
		Country{
			countryCode: "ME",
			countryName: "Montenegro",
		},
		Country{
			countryCode: "MG",
			countryName: "Madagascar",
		},
		Country{
			countryCode: "MH",
			countryName: "Marshall Islands",
		},
		Country{
			countryCode: "MK",
			countryName: "The Former Yugoslav Republic of Macedonia",
		},
		Country{
			countryCode: "ML",
			countryName: "Mali",
		},
		Country{
			countryCode: "MM",
			countryName: "Myanmar",
		},
		Country{
			countryCode: "MN",
			countryName: "Mongolia",
		},
		Country{
			countryCode: "MO",
			countryName: "Macao",
		},
		Country{
			countryCode: "MP",
			countryName: "Northern Mariana Islands",
		},
		Country{
			countryCode: "MQ",
			countryName: "Martinique",
		},
		Country{
			countryCode: "MR",
			countryName: "Mauritania",
		},
		Country{
			countryCode: "MS",
			countryName: "Montserrat",
		},
		Country{
			countryCode: "MT",
			countryName: "Malta",
		},
		Country{
			countryCode: "MU",
			countryName: "Mauritius",
		},
		Country{
			countryCode: "MV",
			countryName: "Maldives",
		},
		Country{
			countryCode: "MW",
			countryName: "Malawi",
		},
		Country{
			countryCode: "MX",
			countryName: "Mexico",
		},
		Country{
			countryCode: "MY",
			countryName: "Malaysia",
		},
		Country{
			countryCode: "MZ",
			countryName: "Mozambique",
		},
		Country{
			countryCode: "NA",
			countryName: "Namibia",
		},
		Country{
			countryCode: "NC",
			countryName: "New Caledonia",
		},
		Country{
			countryCode: "NE",
			countryName: "Niger",
		},
		Country{
			countryCode: "NF",
			countryName: "Norfolk Island",
		},
		Country{
			countryCode: "NG",
			countryName: "Nigeria",
		},
		Country{
			countryCode: "NI",
			countryName: "Nicaragua",
		},
		Country{
			countryCode: "NL",
			countryName: "Netherlands",
		},
		Country{
			countryCode: "NO",
			countryName: "Norway",
		},
		Country{
			countryCode: "NP",
			countryName: "Nepal",
		},
		Country{
			countryCode: "NR",
			countryName: "Nauru",
		},
		Country{
			countryCode: "NU",
			countryName: "Niue",
		},
		Country{
			countryCode: "NZ",
			countryName: "New Zealand",
		},
		Country{
			countryCode: "OM",
			countryName: "Oman",
		},
		Country{
			countryCode: "PA",
			countryName: "Panama",
		},
		Country{
			countryCode: "PE",
			countryName: "Peru",
		},
		Country{
			countryCode: "PF",
			countryName: "French Polynesia Includes Tahiti",
		},
		Country{
			countryCode: "PG",
			countryName: "Papua New Guinea",
		},
		Country{
			countryCode: "PH",
			countryName: "Philippines",
		},
		Country{
			countryCode: "PK",
			countryName: "Pakistan",
		},
		Country{
			countryCode: "PL",
			countryName: "Poland",
		},
		Country{
			countryCode: "PM",
			countryName: "Saint Pierre and Miquelon",
		},
		Country{
			countryCode: "PN",
			countryName: "Pitcairn",
		},
		Country{
			countryCode: "PR",
			countryName: "Puerto Rico",
		},
		Country{
			countryCode: "PS",
			countryName: "Palestinian territory, Occupied",
		},
		Country{
			countryCode: "PT",
			countryName: "Portugal",
		},
		Country{
			countryCode: "PW",
			countryName: "Palau",
		},
		Country{
			countryCode: "PY",
			countryName: "Paraguay",
		},
		Country{
			countryCode: "QA",
			countryName: "Qatar",
		},
		Country{
			countryCode: "QM",
			countryName: "NOTE : Guernsey was defined in eBay list previously but they are not defined in ISO 3166 This country will remain on eBay country list for backward compatibility",
		},
		Country{
			countryCode: "QN",
			countryName: "NOTE: Jan Mayen was defined in eBay list previously but they are not defined in ISO 3166 This country will remain on eBay country list for backward compatibility",
		},
		Country{
			countryCode: "QO",
			countryName: "NOTE: Jersey was defined in eBay list previously but they are not defined in ISO 3166 This country will remain on eBay country list for backward compatibility",
		},
		Country{
			countryCode: "RE",
			countryName: "Reunion",
		},
		Country{
			countryCode: "RO",
			countryName: "Romania",
		},
		Country{
			countryCode: "RS",
			countryName: "Serbia",
		},
		Country{
			countryCode: "RU",
			countryName: "Russian Federation",
		},
		Country{
			countryCode: "RW",
			countryName: "Rwanda",
		},
		Country{
			countryCode: "SA",
			countryName: "Saudi Arabia",
		},
		Country{
			countryCode: "SB",
			countryName: "Solomon Islands",
		},
		Country{
			countryCode: "SC",
			countryName: "Seychelles",
		},
		Country{
			countryCode: "SD",
			countryName: "Sudan",
		},
		Country{
			countryCode: "SE",
			countryName: "Sweden",
		},
		Country{
			countryCode: "SG",
			countryName: "Singapore",
		},
		Country{
			countryCode: "SH",
			countryName: "Saint Helena",
		},
		Country{
			countryCode: "SI",
			countryName: "Slovenia",
		},
		Country{
			countryCode: "SJ",
			countryName: "Svalbard and Jan Mayen",
		},
		Country{
			countryCode: "SK",
			countryName: "Slovakia",
		},
		Country{
			countryCode: "SL",
			countryName: "Sierra Leone",
		},
		Country{
			countryCode: "SM",
			countryName: "San Marino",
		},
		Country{
			countryCode: "SN",
			countryName: "Senegal",
		},
		Country{
			countryCode: "SO",
			countryName: "Somalia",
		},
		Country{
			countryCode: "SR",
			countryName: "Suriname",
		},
		Country{
			countryCode: "ST",
			countryName: "Sao Tome and Principe",
		},
		Country{
			countryCode: "SV",
			countryName: "El Salvador",
		},
		Country{
			countryCode: "SY",
			countryName: "Syrian Arab Republic",
		},
		Country{
			countryCode: "SZ",
			countryName: "Swaziland",
		},
		Country{
			countryCode: "TC",
			countryName: "Turks and Caicos Islands",
		},
		Country{
			countryCode: "TD",
			countryName: "Chad",
		},
		Country{
			countryCode: "TF",
			countryName: "French Southern Territories",
		},
		Country{
			countryCode: "TG",
			countryName: "Togo",
		},
		Country{
			countryCode: "TH",
			countryName: "Thailand",
		},
		Country{
			countryCode: "TJ",
			countryName: "Tajikistan",
		},
		Country{
			countryCode: "TK",
			countryName: "Tokelau",
		},
		Country{
			countryCode: "TM",
			countryName: "Turkmenistan",
		},
		Country{
			countryCode: "TN",
			countryName: "Tunisia",
		},
		Country{
			countryCode: "TO",
			countryName: "Tonga",
		},
		Country{
			countryCode: "TP",
			countryName: "No longer in use",
		},
		Country{
			countryCode: "TR",
			countryName: "Turkey",
		},
		Country{
			countryCode: "TT",
			countryName: "Trinidad and Tobago",
		},
		Country{
			countryCode: "TV",
			countryName: "Tuvalu",
		},
		Country{
			countryCode: "TW",
			countryName: "Taiwan, Province of China",
		},
		Country{
			countryCode: "TZ",
			countryName: "Tanzania, United Republic of",
		},
		Country{
			countryCode: "UA",
			countryName: "Ukraine",
		},
		Country{
			countryCode: "UG",
			countryName: "Uganda",
		},
		Country{
			countryCode: "UM",
			countryName: "NOTE: United States Minor Outlying Islands was defined in the eBay list previously but is no longer a viable option This country will remain on eBay country list for backward compatibility Use 'US' instead",
		},
		Country{
			countryCode: "US",
			countryName: "United States",
		},
		Country{
			countryCode: "UY",
			countryName: "Uruguay",
		},
		Country{
			countryCode: "UZ",
			countryName: "Uzbekistan",
		},
		Country{
			countryCode: "VA",
			countryName: "Holy See (Vatican City state)",
		},
		Country{
			countryCode: "VC",
			countryName: "Saint Vincent and the Grenadines",
		},
		Country{
			countryCode: "VE",
			countryName: "Venezuela",
		},
		Country{
			countryCode: "VG",
			countryName: "Virgin Islands, British",
		},
		Country{
			countryCode: "VI",
			countryName: "Virgin Islands, US",
		},
		Country{
			countryCode: "VN",
			countryName: "Vietnam",
		},
		Country{
			countryCode: "VU",
			countryName: "Vanuatu",
		},
		Country{
			countryCode: "WF",
			countryName: "Wallis and Futuna",
		},
		Country{
			countryCode: "WS",
			countryName: "Samoa",
		},
		Country{
			countryCode: "YE",
			countryName: "Yemen",
		},
		Country{
			countryCode: "YT",
			countryName: "Mayotte",
		},
		Country{
			countryCode: "YU",
			countryName: "No longer in use See RS for Serbia and ME for Montenegro",
		},
		Country{
			countryCode: "ZA",
			countryName: "South Africa",
		},
		Country{
			countryCode: "ZM",
			countryName: "Zambia",
		},
		Country{
			countryCode: "ZW",
			countryName: "Zimbabwe",
		},
		Country{
			countryCode: "ZZ",
			countryName: "Unknown country",
		},
	}

	var countryName = ""

	for _, country := range countries {
		if country.countryCode == countryCodeFromEbay {
			countryName = country.countryName
			return country.countryName
		}
	}

	return countryName
}
