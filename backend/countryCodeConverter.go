package main

type Country struct {
	countryCodeId string
	countryCode   string
	countryName   string
}

func countryCodeConverter(countryCode string) string {
	var countries = []Country{
		Country{
			countryCode: "AA",
			countryName: "NOTE: APO/FPO was defined in eBay list previously but they are not defined in ISO 3166 This country will remain on eBay country code list for backward compatibility",
		},
		Country{
			countryCodeId: "306",
			countryCode:   "AX",
			countryName:   "Åland Islands",
		},
		Country{
			countryCodeId: "57",
			countryCode:   "AL",
			countryName:   "Albania",
		},
		Country{
			countryCodeId: "95",
			countryCode:   "DZ",
			countryName:   "Algeria",
		},
		Country{
			countryCodeId: "250",
			countryCode:   "AS",
			countryName:   "American Samoa",
		},
		Country{
			countryCodeId: "228",
			countryCode:   "AD",
			countryName:   "Andorra",
		},
		Country{
			countryCodeId: "56",
			countryCode:   "AO",
			countryName:   "Angola",
		},
		Country{
			countryCodeId: "251",
			countryCode:   "AI",
			countryName:   "Anguilla",
		},
		Country{
			countryCodeId: "252",
			countryCode:   "AG",
			countryName:   "Antigua and Barbuda",
		},
		Country{
			countryCodeId: "59",
			countryCode:   "AR",
			countryName:   "Argentina",
		},
		Country{
			countryCodeId: "60",
			countryCode:   "AM",
			countryName:   "Armenia",
		},
		Country{
			countryCodeId: "253",
			countryCode:   "AW",
			countryName:   "Aruba",
		},
		Country{
			countryCodeId: "61",
			countryCode:   "AU",
			countryName:   "Australia",
		},
		Country{
			countryCodeId: "62",
			countryCode:   "AT",
			countryName:   "Austria",
		},
		Country{
			countryCodeId: "63",
			countryCode:   "AZ",
			countryName:   "Azerbaijan",
		},
		Country{
			countryCodeId: "229",
			countryCode:   "BS",
			countryName:   "Bahamas",
		},
		Country{
			countryCodeId: "232",
			countryCode:   "BH",
			countryName:   "Bahrain",
		},
		Country{
			countryCodeId: "68",
			countryCode:   "BD",
			countryName:   "Bangladesh",
		},
		Country{
			countryCodeId: "237",
			countryCode:   "BB",
			countryName:   "Barbados",
		},
		Country{
			countryCodeId: "71",
			countryCode:   "BY",
			countryName:   "Belarus",
		},
		Country{
			countryCodeId: "65",
			countryCode:   "BE",
			countryName:   "Belgium",
		},
		Country{
			countryCodeId: "72",
			countryCode:   "BZ",
			countryName:   "Belize",
		},
		Country{
			countryCodeId: "66",
			countryCode:   "BJ",
			countryName:   "Benin",
		},
		Country{
			countryCodeId: "225",
			countryCode:   "BM",
			countryName:   "Bermuda",
		},
		Country{
			countryCodeId: "76",
			countryCode:   "BT",
			countryName:   "Bhutan",
		},
		Country{
			countryCodeId: "73",
			countryCode:   "BO",
			countryName:   "Bolivia",
		},
		Country{
			countryCodeId: "70",
			countryCode:   "BA",
			countryName:   "Bosnia and Herzegovina",
		},
		Country{
			countryCodeId: "77",
			countryCode:   "BW",
			countryName:   "Botswana",
		},
		Country{
			countryCodeId: "254",
			countryCode:   "BV",
			countryName:   "Bouvet Island",
		},
		Country{
			countryCodeId: "74",
			countryCode:   "BR",
			countryName:   "Brazil",
		},
		Country{
			countryCodeId: "255",
			countryCode:   "IO",
			countryName:   "British Indian Ocean Territory",
		},
		Country{
			countryCodeId: "231",
			countryCode:   "VG",
			countryName:   "British Virgin Islands",
		},
		Country{
			countryCodeId: "75",
			countryCode:   "BN",
			countryName:   "Brunei",
		},
		Country{
			countryCodeId: "69",
			countryCode:   "BG",
			countryName:   "Bulgaria",
		},
		Country{
			countryCodeId: "67",
			countryCode:   "BF",
			countryName:   "Burkina Faso",
		},
		Country{
			countryCodeId: "64",
			countryCode:   "BI",
			countryName:   "Burundi",
		},
		Country{
			countryCodeId: "135",
			countryCode:   "KH",
			countryName:   "Cambodia",
		},
		Country{
			countryCodeId: "84",
			countryCode:   "CM",
			countryName:   "Cameroon",
		},
		Country{
			countryCodeId: "79",
			countryCode:   "CA",
			countryName:   "Canada",
		},
		Country{
			countryCodeId: "222",
			countryCode:   "CV",
			countryName:   "Cape Verde",
		},
		Country{
			countryCodeId: "247",
			countryCode:   "KY",
			countryName:   "Cayman Islands",
		},
		Country{
			countryCodeId: "78",
			countryCode:   "CF",
			countryName:   "Central African Republic",
		},
		Country{
			countryCodeId: "196",
			countryCode:   "TD",
			countryName:   "Chad",
		},
		Country{
			countryCodeId: "81",
			countryCode:   "CL",
			countryName:   "Chile",
		},
		Country{
			countryCodeId: "82",
			countryCode:   "CN",
			countryName:   "China",
		},
		Country{
			countryCodeId: "257",
			countryCode:   "CX",
			countryName:   "Christmas Island",
		},
		Country{
			countryCodeId: "258",
			countryCode:   "CC",
			countryName:   "Cocos (Keeling) Islands",
		},
		Country{
			countryCodeId: "86",
			countryCode:   "CO",
			countryName:   "Colombia",
		},
		Country{
			countryCodeId: "259",
			countryCode:   "KM",
			countryName:   "Comoros",
		},
		Country{
			countryCodeId: "85",
			countryCode:   "CG",
			countryName:   "Congo, Republic of",
		},
		Country{
			countryCodeId: "260",
			countryCode:   "CK",
			countryName:   "Cook Islands",
		},
		Country{
			countryCodeId: "87",
			countryCode:   "CR",
			countryName:   "Costa Rica",
		},
		Country{
			countryCodeId: "118",
			countryCode:   "HR",
			countryName:   "Croatia",
		},
		Country{
			countryCodeId: "338",
			countryCode:   "CW",
			countryName:   "Curaçao",
		},
		Country{
			countryCodeId: "89",
			countryCode:   "CY",
			countryName:   "Cyprus",
		},
		Country{
			countryCodeId: "90",
			countryCode:   "CZ",
			countryName:   "Czech Republic",
		},
		Country{
			countryCodeId: "93",
			countryCode:   "DK",
			countryName:   "Denmark",
		},
		Country{
			countryCodeId: "92",
			countryCode:   "DJ",
			countryName:   "Djibouti",
		},
		Country{
			countryCodeId: "261",
			countryCode:   "DM",
			countryName:   "Dominica",
		},
		Country{
			countryCodeId: "94",
			countryCode:   "DO",
			countryName:   "Dominican Republic",
		},
		Country{
			countryCodeId: "96",
			countryCode:   "EC",
			countryName:   "Ecuador",
		},
		Country{
			countryCodeId: "97",
			countryCode:   "EG",
			countryName:   "Egypt",
		},
		Country{
			countryCodeId: "187",
			countryCode:   "SV",
			countryName:   "El Salvador",
		},
		Country{
			countryCodeId: "111",
			countryCode:   "GQ",
			countryName:   "Equatorial Guinea",
		},
		Country{
			countryCodeId: "98",
			countryCode:   "ER",
			countryName:   "Eritrea",
		},
		Country{
			countryCodeId: "100",
			countryCode:   "EE",
			countryName:   "Estonia",
		},
		Country{
			countryCodeId: "101",
			countryCode:   "ET",
			countryName:   "Ethiopia",
		},
		Country{
			countryCodeId: "262",
			countryCode:   "FK",
			countryName:   "Falkland Islands (Malvinas)",
		},
		Country{
			countryCodeId: "241",
			countryCode:   "FO",
			countryName:   "Faroe Islands",
		},
		Country{
			countryCodeId: "234",
			countryCode:   "FJ",
			countryName:   "Fiji",
		},
		Country{
			countryCodeId: "102",
			countryCode:   "FI",
			countryName:   "Finland",
		},
		Country{
			countryCodeId: "103",
			countryCode:   "FR",
			countryName:   "France",
		},
		Country{
			countryCodeId: "115",
			countryCode:   "GF",
			countryName:   "French Guiana",
		},
		Country{
			countryCodeId: "263",
			countryCode:   "PF",
			countryName:   "French Polynesia",
		},
		Country{
			countryCodeId: "264",
			countryCode:   "TF",
			countryName:   "French Southern Territories",
		},
		Country{
			countryCodeId: "104",
			countryCode:   "GA",
			countryName:   "Gabon",
		},
		Country{
			countryCodeId: "109",
			countryCode:   "GM",
			countryName:   "Gambia",
		},
		Country{
			countryCodeId: "106",
			countryCode:   "GE",
			countryName:   "Georgia",
		},
		Country{
			countryCodeId: "91",
			countryCode:   "DE",
			countryName:   "Germany",
		},
		Country{
			countryCodeId: "107",
			countryCode:   "GH",
			countryName:   "Ghana",
		},
		Country{
			countryCodeId: "226",
			countryCode:   "GI",
			countryName:   "Gibraltar",
		},
		Country{
			countryCodeId: "112",
			countryCode:   "GR",
			countryName:   "Greece",
		},
		Country{
			countryCodeId: "113",
			countryCode:   "GL",
			countryName:   "Greenland",
		},
		Country{
			countryCodeId: "245",
			countryCode:   "GD",
			countryName:   "Grenada",
		},
		Country{
			countryCodeId: "265",
			countryCode:   "GP",
			countryName:   "Guadeloupe",
		},
		Country{
			countryCodeId: "266",
			countryCode:   "GU",
			countryName:   "Guam",
		},
		Country{
			countryCodeId: "114",
			countryCode:   "GT",
			countryName:   "Guatemala",
		},
		Country{
			countryCodeId: "305",
			countryCode:   "GG",
			countryName:   "Guernsey",
		},
		Country{
			countryCodeId: "108",
			countryCode:   "GN",
			countryName:   "Guinea",
		},
		Country{
			countryCodeId: "110",
			countryCode:   "GW",
			countryName:   "Guinea-Bissau",
		},
		Country{
			countryCodeId: "116",
			countryCode:   "GY",
			countryName:   "Guyana",
		},
		Country{
			countryCodeId: "119",
			countryCode:   "HT",
			countryName:   "Haiti",
		},
		Country{
			countryCodeId: "267",
			countryCode:   "HM",
			countryName:   "Heard Island and McDonald Islands",
		},
		Country{
			countryCodeId: "268",
			countryCode:   "VA",
			countryName:   "Holy See (Vatican City State)",
		},
		Country{
			countryCodeId: "117",
			countryCode:   "HN",
			countryName:   "Honduras",
		},
		Country{
			countryCodeId: "219",
			countryCode:   "HK",
			countryName:   "Hong Kong",
		},
		Country{
			countryCodeId: "120",
			countryCode:   "HU",
			countryName:   "Hungary",
		},
		Country{
			countryCodeId: "126",
			countryCode:   "IS",
			countryName:   "Iceland",
		},
		Country{
			countryCodeId: "122",
			countryCode:   "IN",
			countryName:   "India",
		},
		Country{
			countryCodeId: "121",
			countryCode:   "ID",
			countryName:   "Indonesia",
		},
		Country{
			countryCodeId: "125",
			countryCode:   "IQ",
			countryName:   "Iraq",
		},
		Country{
			countryCodeId: "123",
			countryCode:   "IE",
			countryName:   "Ireland",
		},
		Country{
			countryCodeId: "269",
			countryCode:   "IM",
			countryName:   "Isle of Man",
		},
		Country{
			countryCodeId: "127",
			countryCode:   "IL",
			countryName:   "Israel",
		},
		Country{
			countryCodeId: "128",
			countryCode:   "IT",
			countryName:   "Italy",
		},
		Country{
			countryCodeId: "83",
			countryCode:   "IC",
			countryName:   "Ivory Coast",
		},
		Country{
			countryCodeId: "129",
			countryCode:   "JM",
			countryName:   "Jamaica",
		},
		Country{
			countryCodeId: "131",
			countryCode:   "JP",
			countryName:   "Japan",
		},
		Country{
			countryCodeId: "307",
			countryCode:   "JE",
			countryName:   "Jersey",
		},
		Country{
			countryCodeId: "130",
			countryCode:   "JO",
			countryName:   "Jordan",
		},
		Country{
			countryCodeId: "132",
			countryCode:   "KZ",
			countryName:   "Kazakhstan",
		},
		Country{
			countryCodeId: "133",
			countryCode:   "KE",
			countryName:   "Kenya",
		},
		Country{
			countryCodeId: "270",
			countryCode:   "KI",
			countryName:   "Kiribati",
		},
		Country{
			countryCodeId: "271",
			countryCode:   "KV",
			countryName:   "Kosovo",
		},
		Country{
			countryCodeId: "137",
			countryCode:   "KW",
			countryName:   "Kuwait",
		},
		Country{
			countryCodeId: "134",
			countryCode:   "KG",
			countryName:   "Kyrgyzstan",
		},
		Country{
			countryCodeId: "138",
			countryCode:   "LA",
			countryName:   "Laos",
		},
		Country{
			countryCodeId: "146",
			countryCode:   "LV",
			countryName:   "Latvia",
		},
		Country{
			countryCodeId: "139",
			countryCode:   "LB",
			countryName:   "Lebanon",
		},
		Country{
			countryCodeId: "143",
			countryCode:   "LS",
			countryName:   "Lesotho",
		},
		Country{
			countryCodeId: "140",
			countryCode:   "LR",
			countryName:   "Liberia",
		},
		Country{
			countryCodeId: "141",
			countryCode:   "LY",
			countryName:   "Libya",
		},
		Country{
			countryCodeId: "272",
			countryCode:   "LI",
			countryName:   "Liechtenstein",
		},
		Country{
			countryCodeId: "144",
			countryCode:   "LT",
			countryName:   "Lithuania",
		},
		Country{
			countryCodeId: "145",
			countryCode:   "LU",
			countryName:   "Luxembourg",
		},
		Country{
			countryCodeId: "273",
			countryCode:   "MO",
			countryName:   "Macao",
		},
		Country{
			countryCodeId: "151",
			countryCode:   "MK",
			countryName:   "Macedonia",
		},
		Country{
			countryCodeId: "149",
			countryCode:   "MG",
			countryName:   "Madagascar",
		},
		Country{
			countryCodeId: "158",
			countryCode:   "MW",
			countryName:   "Malawi",
		},
		Country{
			countryCodeId: "159",
			countryCode:   "MY",
			countryName:   "Malaysia",
		},
		Country{
			countryCodeId: "238",
			countryCode:   "MV",
			countryName:   "Maldives",
		},
		Country{
			countryCodeId: "152",
			countryCode:   "ML",
			countryName:   "Mali",
		},
		Country{
			countryCodeId: "227",
			countryCode:   "MT",
			countryName:   "Malta",
		},
		Country{
			countryCodeId: "274",
			countryCode:   "MH",
			countryName:   "Marshall Islands",
		},
		Country{
			countryCodeId: "275",
			countryCode:   "MQ",
			countryName:   "Martinique",
		},
		Country{
			countryCodeId: "157",
			countryCode:   "MR",
			countryName:   "Mauritania",
		},
		Country{
			countryCodeId: "239",
			countryCode:   "MU",
			countryName:   "Mauritius",
		},
		Country{
			countryCodeId: "276",
			countryCode:   "YT",
			countryName:   "Mayotte",
		},
		Country{
			countryCodeId: "150",
			countryCode:   "MX",
			countryName:   "Mexico",
		},
		Country{
			countryCodeId: "277",
			countryCode:   "FM",
			countryName:   "Micronesia, Federated States of",
		},
		Country{
			countryCodeId: "148",
			countryCode:   "MD",
			countryName:   "Moldova",
		},
		Country{
			countryCodeId: "278",
			countryCode:   "MC",
			countryName:   "Monaco",
		},
		Country{
			countryCodeId: "154",
			countryCode:   "MN",
			countryName:   "Mongolia",
		},
		Country{
			countryCodeId: "155",
			countryCode:   "ME",
			countryName:   "Montenegro",
		},
		Country{
			countryCodeId: "279",
			countryCode:   "MS",
			countryName:   "Montserrat",
		},
		Country{
			countryCodeId: "147",
			countryCode:   "MA",
			countryName:   "Morocco",
		},
		Country{
			countryCodeId: "156",
			countryCode:   "MZ",
			countryName:   "Mozambique",
		},
		Country{
			countryCodeId: "153",
			countryCode:   "MM",
			countryName:   "Myanmar (Burma)",
		},
		Country{
			countryCodeId: "160",
			countryCode:   "NA",
			countryName:   "Namibia",
		},
		Country{
			countryCodeId: "280",
			countryCode:   "NR",
			countryName:   "Nauru",
		},
		Country{
			countryCodeId: "166",
			countryCode:   "NP",
			countryName:   "Nepal",
		},
		Country{
			countryCodeId: "243",
			countryCode:   "AN",
			countryName:   "Netherlands Antilles",
		},
		Country{
			countryCodeId: "233",
			countryCode:   "NC",
			countryName:   "New Caledonia",
		},
		Country{
			countryCodeId: "167",
			countryCode:   "NZ",
			countryName:   "New Zealand",
		},
		Country{
			countryCodeId: "163",
			countryCode:   "NI",
			countryName:   "Nicaragua",
		},
		Country{
			countryCodeId: "161",
			countryCode:   "NE",
			countryName:   "Niger",
		},
		Country{
			countryCodeId: "162",
			countryCode:   "NG",
			countryName:   "Nigeria",
		},
		Country{
			countryCodeId: "281",
			countryCode:   "NU",
			countryName:   "Niue",
		},
		Country{
			countryCodeId: "282",
			countryCode:   "NF",
			countryName:   "Norfolk Island",
		},
		Country{
			countryCodeId: "283",
			countryCode:   "MP",
			countryName:   "Northern Mariana Islands",
		},
		Country{
			countryCodeId: "165",
			countryCode:   "NO",
			countryName:   "Norway",
		},
		Country{
			countryCodeId: "168",
			countryCode:   "OM",
			countryName:   "Oman",
		},
		Country{
			countryCodeId: "169",
			countryCode:   "PK",
			countryName:   "Pakistan",
		},
		Country{
			countryCodeId: "284",
			countryCode:   "PW",
			countryName:   "Palau",
		},
		Country{
			countryCodeId: "285",
			countryCode:   "PS",
			countryName:   "Palestine, State of",
		},
		Country{
			countryCodeId: "170",
			countryCode:   "PA",
			countryName:   "Panama",
		},
		Country{
			countryCodeId: "173",
			countryCode:   "PG",
			countryName:   "Papua New Guinea",
		},
		Country{
			countryCodeId: "178",
			countryCode:   "PY",
			countryName:   "Paraguay",
		},
		Country{
			countryCodeId: "171",
			countryCode:   "PE",
			countryName:   "Peru",
		},
		Country{
			countryCodeId: "172",
			countryCode:   "PH",
			countryName:   "Philippines",
		},
		Country{
			countryCodeId: "174",
			countryCode:   "PL",
			countryName:   "Poland",
		},
		Country{
			countryCodeId: "177",
			countryCode:   "PT",
			countryName:   "Portugal",
		},
		Country{
			countryCodeId: "175",
			countryCode:   "PR",
			countryName:   "Puerto Rico",
		},
		Country{
			countryCodeId: "179",
			countryCode:   "QA",
			countryName:   "Qatar",
		},
		Country{
			countryCodeId: "304",
			countryCode:   "RE",
			countryName:   "Reunion",
		},
		Country{
			countryCodeId: "180",
			countryCode:   "RO",
			countryName:   "Romania",
		},
		Country{
			countryCodeId: "181",
			countryCode:   "RU",
			countryName:   "Russia",
		},
		Country{
			countryCodeId: "182",
			countryCode:   "RW",
			countryName:   "Rwanda",
		},
		Country{
			countryCodeId: "308",
			countryCode:   "BL",
			countryName:   "Saint Barthélemy",
		},
		Country{
			countryCodeId: "286",
			countryCode:   "SH",
			countryName:   "Saint Helena",
		},
		Country{
			countryCodeId: "287",
			countryCode:   "KN",
			countryName:   "Saint Kitts and Nevis",
		},
		Country{
			countryCodeId: "244",
			countryCode:   "LC",
			countryName:   "Saint Lucia",
		},
		Country{
			countryCodeId: "288",
			countryCode:   "MF",
			countryName:   "Saint Martin (French part)",
		},
		Country{
			countryCodeId: "289",
			countryCode:   "PM",
			countryName:   "Saint Pierre and Miquelon",
		},
		Country{
			countryCodeId: "249",
			countryCode:   "VC",
			countryName:   "Saint Vincent and the Grenadines",
		},
		Country{
			countryCodeId: "290",
			countryCode:   "WS",
			countryName:   "Samoa",
		},
		Country{
			countryCodeId: "291",
			countryCode:   "SM",
			countryName:   "San Marino",
		},
		Country{
			countryCodeId: "292",
			countryCode:   "ST",
			countryName:   "Sao Tome and Principe",
		},
		Country{
			countryCodeId: "183",
			countryCode:   "SA",
			countryName:   "Saudi Arabia",
		},
		Country{
			countryCodeId: "185",
			countryCode:   "SN",
			countryName:   "Senegal",
		},
		Country{
			countryCodeId: "189",
			countryCode:   "RS",
			countryName:   "Serbia",
		},
		Country{
			countryCodeId: "293",
			countryCode:   "SC",
			countryName:   "Seychelles",
		},
		Country{
			countryCodeId: "186",
			countryCode:   "SL",
			countryName:   "Sierra Leone",
		},
		Country{
			countryCodeId: "220",
			countryCode:   "SG",
			countryName:   "Singapore",
		},
		Country{
			countryCodeId: "337",
			countryCode:   "SX",
			countryName:   "Sint Maarten (Dutch part)",
		},
		Country{
			countryCodeId: "191",
			countryCode:   "SK",
			countryName:   "Slovakia",
		},
		Country{
			countryCodeId: "192",
			countryCode:   "SI",
			countryName:   "Slovenia",
		},
		Country{
			countryCodeId: "242",
			countryCode:   "SB",
			countryName:   "Solomon Islands",
		},
		Country{
			countryCodeId: "188",
			countryCode:   "SO",
			countryName:   "Somalia",
		},
		Country{
			countryCodeId: "215",
			countryCode:   "ZA",
			countryName:   "South Africa",
		},
		Country{
			countryCodeId: "294",
			countryCode:   "GS",
			countryName:   "South Georgia and the South Sandwich Islands",
		},
		Country{
			countryCodeId: "136",
			countryCode:   "KR",
			countryName:   "South Korea",
		},
		Country{
			countryCodeId: "339",
			countryCode:   "SS",
			countryName:   "South Sudan",
		},
		Country{
			countryCodeId: "99",
			countryCode:   "ES",
			countryName:   "Spain",
		},
		Country{
			countryCodeId: "142",
			countryCode:   "LK",
			countryName:   "Sri Lanka",
		},
		Country{
			countryCodeId: "184",
			countryCode:   "SD",
			countryName:   "Sudan",
		},
		Country{
			countryCodeId: "190",
			countryCode:   "SR",
			countryName:   "Suriname",
		},
		Country{
			countryCodeId: "295",
			countryCode:   "SJ",
			countryName:   "Svalbard and Jan Mayen",
		},
		Country{
			countryCodeId: "194",
			countryCode:   "SZ",
			countryName:   "Swaziland",
		},
		Country{
			countryCodeId: "193",
			countryCode:   "SE",
			countryName:   "Sweden",
		},
		Country{
			countryCodeId: "80",
			countryCode:   "CH",
			countryName:   "Switzerland",
		},
		Country{
			countryCodeId: "204",
			countryCode:   "TW",
			countryName:   "Taiwan",
		},
		Country{
			countryCodeId: "199",
			countryCode:   "TJ",
			countryName:   "Tajikistan",
		},
		Country{
			countryCodeId: "205",
			countryCode:   "TZ",
			countryName:   "Tanzania",
		},
		Country{
			countryCodeId: "198",
			countryCode:   "TH",
			countryName:   "Thailand",
		},
		Country{
			countryCodeId: "164",
			countryCode:   "NL",
			countryName:   "The Netherlands",
		},
		Country{
			countryCodeId: "296",
			countryCode:   "TL",
			countryName:   "Timor-Leste",
		},
		Country{
			countryCodeId: "197",
			countryCode:   "TG",
			countryName:   "Togo",
		},
		Country{
			countryCodeId: "297",
			countryCode:   "TK",
			countryName:   "Tokelau",
		},
		Country{
			countryCodeId: "298",
			countryCode:   "TO",
			countryName:   "Tonga",
		},
		Country{
			countryCodeId: "201",
			countryCode:   "TT",
			countryName:   "Trinidad",
		},
		Country{
			countryCodeId: "202",
			countryCode:   "TN",
			countryName:   "Tunisia",
		},
		Country{
			countryCodeId: "203",
			countryCode:   "TR",
			countryName:   "Turkey",
		},
		Country{
			countryCodeId: "200",
			countryCode:   "TM",
			countryName:   "Turkmenistan",
		},
		Country{
			countryCodeId: "299",
			countryCode:   "TC",
			countryName:   "Turks and Caicos Islands",
		},
		Country{
			countryCodeId: "300",
			countryCode:   "TV",
			countryName:   "Tuvalu",
		},
		Country{
			countryCodeId: "206",
			countryCode:   "UG",
			countryName:   "Uganda",
		},
		Country{
			countryCodeId: "207",
			countryCode:   "UA",
			countryName:   "Ukraine",
		},
		Country{
			countryCodeId: "58",
			countryCode:   "AE",
			countryName:   "United Arab Emirates",
		},
		Country{
			countryCodeId: "105",
			countryCode:   "GB",
			countryName:   "United Kingdom",
		},
		Country{
			countryCodeId: "209",
			countryCode:   "US",
			countryName:   "United States",
		},
		Country{
			countryCodeId: "302",
			countryCode:   "UM",
			countryName:   "United States Minor Outlying Islands",
		},
		Country{
			countryCodeId: "208",
			countryCode:   "UY",
			countryName:   "Uruguay",
		},
		Country{
			countryCodeId: "248",
			countryCode:   "VI",
			countryName:   "U.S. Virgin Islands",
		},
		Country{
			countryCodeId: "210",
			countryCode:   "UZ",
			countryName:   "Uzbekistan",
		},
		Country{
			countryCodeId: "221",
			countryCode:   "VU",
			countryName:   "Vanuatu",
		},
		Country{
			countryCodeId: "211",
			countryCode:   "VE",
			countryName:   "Venezuela",
		},
		Country{
			countryCodeId: "212",
			countryCode:   "VN",
			countryName:   "Vietnam",
		},
		Country{
			countryCodeId: "224",
			countryCode:   "WF",
			countryName:   "Wallis and Futuna",
		},
		Country{
			countryCodeId: "213",
			countryCode:   "EH",
			countryName:   "Western Sahara",
		},
		Country{
			countryCodeId: "214",
			countryCode:   "YE",
			countryName:   "Yemen",
		},
		Country{
			countryCodeId: "216",
			countryCode:   "CD",
			countryName:   "Zaire (Democratic Republic of Congo)",
		},
		Country{
			countryCodeId: "217",
			countryCode:   "ZM",
			countryName:   "Zambia",
		},
		Country{
			countryCodeId: "218",
			countryCode:   "ZW",
			countryName:   "Zimbabwe",
		},
		Country{
			countryCode: "ZZ",
			countryName: "Unknown country",
		},
	}

	for _, country := range countries {
		if country.countryCode == countryCode || country.countryCodeId == countryCode {
			return country.countryName
		}
	}
	return ""
}
