package isbn

type Range struct {
	Prefix string
	Agency string
	Ranges [][2]string
}

var Ranges = map[string]Range{
	"0": Range{
		Prefix: "978",
		Agency: "English language",
		Ranges: [][2]string{{"00", "19"}, {"200", "227"}, {"2280", "2289"}, {"229", "638"}, {"6390", "6399"}, {"640", "647"}, {"6480000", "6489999"}, {"649", "654"}, {"6550", "6559"}, {"656", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "949999"}, {"9500000", "9999999"}}},
	"1": Range{
		Prefix: "978",
		Agency: "English language",
		Ranges: [][2]string{{"00", "09"}, {"100", "399"}, {"4000", "5499"}, {"55000", "73199"}, {"7320000", "7399999"}, {"74000", "77499"}, {"7750000", "7753999"}, {"77540", "86979"}, {"869800", "972999"}, {"9730", "9877"}, {"987800", "998999"}, {"9990000", "9999999"}}},
	"2": Range{
		Prefix: "978",
		Agency: "French language",
		Ranges: [][2]string{{"00", "19"}, {"200", "349"}, {"35000", "39999"}, {"400", "489"}, {"490000", "494999"}, {"495", "699"}, {"7000", "8399"}, {"84000", "89999"}, {"900000", "919942"}, {"9199430", "9199689"}, {"919969", "949999"}, {"9500000", "9999999"}}},
	"3": Range{
		Prefix: "978",
		Agency: "German language",
		Ranges: [][2]string{{"00", "02"}, {"030", "033"}, {"0340", "0369"}, {"03700", "03999"}, {"04", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "949999"}, {"9500000", "9539999"}, {"95400", "96999"}, {"9700000", "9899999"}, {"99000", "99499"}, {"99500", "99999"}}},
	"4": Range{
		Prefix: "978",
		Agency: "Japan",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "949999"}, {"9500000", "9999999"}}},
	"5": Range{
		Prefix: "978",
		Agency: "former U.S.S.R",
		Ranges: [][2]string{{"00000", "00499"}, {"0050", "0099"}, {"01", "19"}, {"200", "420"}, {"4210", "4299"}, {"430", "430"}, {"4310", "4399"}, {"440", "440"}, {"4410", "4499"}, {"450", "603"}, {"6040000", "6049999"}, {"605", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "909999"}, {"91000", "91999"}, {"9200", "9299"}, {"93000", "94999"}, {"9500000", "9500999"}, {"9501", "9799"}, {"98000", "98999"}, {"9900000", "9909999"}, {"9910", "9999"}}},
	"600": Range{
		Prefix: "978",
		Agency: "Iran",
		Ranges: [][2]string{{"00", "09"}, {"100", "499"}, {"5000", "8999"}, {"90000", "99999"}}},
	"601": Range{
		Prefix: "978",
		Agency: "Kazakhstan",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "7999"}, {"80000", "84999"}, {"85", "99"}}},
	"602": Range{
		Prefix: "978",
		Agency: "Indonesia",
		Ranges: [][2]string{{"00", "07"}, {"0800", "0899"}, {"0900", "1099"}, {"1100", "1199"}, {"1200", "1399"}, {"14000", "14999"}, {"1500", "1699"}, {"17000", "17999"}, {"18000", "18999"}, {"19000", "19999"}, {"200", "499"}, {"50000", "53999"}, {"5400", "5999"}, {"60000", "61999"}, {"6200", "6749"}, {"6750", "6999"}, {"70000", "74999"}, {"7500", "7999"}, {"8000", "9499"}, {"95000", "99999"}}},
	"603": Range{
		Prefix: "978",
		Agency: "Saudi Arabia",
		Ranges: [][2]string{{"00", "04"}, {"05", "49"}, {"500", "799"}, {"8000", "8999"}, {"90000", "99999"}}},
	"604": Range{
		Prefix: "978",
		Agency: "Vietnam",
		Ranges: [][2]string{{"0", "4"}, {"50", "89"}, {"900", "979"}, {"9800", "9999"}}},
	"605": Range{
		Prefix: "978",
		Agency: "Turkey",
		Ranges: [][2]string{{"01", "02"}, {"030", "039"}, {"04", "09"}, {"100", "199"}, {"2000", "2399"}, {"240", "399"}, {"4000", "5999"}, {"60000", "89999"}, {"9000", "9999"}}},
	"606": Range{
		Prefix: "978",
		Agency: "Romania",
		Ranges: [][2]string{{"000", "099"}, {"10", "49"}, {"500", "799"}, {"8000", "9099"}, {"910", "919"}, {"92000", "97499"}, {"975", "999"}}},
	"607": Range{
		Prefix: "978",
		Agency: "Mexico",
		Ranges: [][2]string{{"00", "39"}, {"400", "749"}, {"7500", "9499"}, {"95000", "99999"}}},
	"608": Range{
		Prefix: "978",
		Agency: "Macedonia",
		Ranges: [][2]string{{"0", "0"}, {"10", "19"}, {"200", "449"}, {"4500", "6499"}, {"65000", "69999"}, {"7", "9"}}},
	"609": Range{
		Prefix: "978",
		Agency: "Lithuania",
		Ranges: [][2]string{{"00", "39"}, {"400", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"611": Range{
		Prefix: "978",
		Agency: "Thailand",
		Ranges: [][2]string{}},
	"612": Range{
		Prefix: "978",
		Agency: "Peru",
		Ranges: [][2]string{{"00", "29"}, {"300", "399"}, {"4000", "4499"}, {"45000", "49999"}, {"50", "99"}}},
	"613": Range{
		Prefix: "978",
		Agency: "Mauritius",
		Ranges: [][2]string{{"0", "9"}}},
	"614": Range{
		Prefix: "978",
		Agency: "Lebanon",
		Ranges: [][2]string{{"00", "39"}, {"400", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"615": Range{
		Prefix: "978",
		Agency: "Hungary",
		Ranges: [][2]string{{"00", "09"}, {"100", "499"}, {"5000", "7999"}, {"80000", "89999"}}},
	"616": Range{
		Prefix: "978",
		Agency: "Thailand",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8999"}, {"90000", "99999"}}},
	"617": Range{
		Prefix: "978",
		Agency: "Ukraine",
		Ranges: [][2]string{{"00", "49"}, {"500", "699"}, {"7000", "8999"}, {"90000", "99999"}}},
	"618": Range{
		Prefix: "978",
		Agency: "Greece",
		Ranges: [][2]string{{"00", "19"}, {"200", "499"}, {"5000", "7999"}, {"80000", "99999"}}},
	"619": Range{
		Prefix: "978",
		Agency: "Bulgaria",
		Ranges: [][2]string{{"00", "14"}, {"150", "699"}, {"7000", "8999"}, {"90000", "99999"}}},
	"620": Range{
		Prefix: "978",
		Agency: "Mauritius",
		Ranges: [][2]string{{"0", "9"}}},
	"621": Range{
		Prefix: "978",
		Agency: "Philippines",
		Ranges: [][2]string{{"00", "29"}, {"400", "599"}, {"8000", "8999"}, {"95000", "99999"}}},
	"7": Range{
		Prefix: "978",
		Agency: "China, People's Republic",
		Ranges: [][2]string{{"00", "09"}, {"100", "499"}, {"5000", "7999"}, {"80000", "89999"}, {"900000", "999999"}}},
	"80": Range{
		Prefix: "978",
		Agency: "former Czechoslovakia",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "999999"}}},
	"81": Range{
		Prefix: "978",
		Agency: "India",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "999999"}}},
	"82": Range{
		Prefix: "978",
		Agency: "Norway",
		Ranges: [][2]string{{"00", "19"}, {"200", "689"}, {"690000", "699999"}, {"7000", "8999"}, {"90000", "98999"}, {"990000", "999999"}}},
	"83": Range{
		Prefix: "978",
		Agency: "Poland",
		Ranges: [][2]string{{"00", "19"}, {"200", "599"}, {"60000", "69999"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "999999"}}},
	"84": Range{
		Prefix: "978",
		Agency: "Spain",
		Ranges: [][2]string{{"00", "13"}, {"140", "149"}, {"15000", "19999"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"9000", "9199"}, {"920000", "923999"}, {"92400", "92999"}, {"930000", "949999"}, {"95000", "96999"}, {"9700", "9999"}}},
	"85": Range{
		Prefix: "978",
		Agency: "Brazil",
		Ranges: [][2]string{{"00", "19"}, {"200", "549"}, {"5500", "5999"}, {"60000", "69999"}, {"7000", "8499"}, {"85000", "89999"}, {"900000", "924999"}, {"92500", "94499"}, {"9450", "9599"}, {"96", "97"}, {"98000", "99999"}}},
	"86": Range{
		Prefix: "978",
		Agency: "former Yugoslavia",
		Ranges: [][2]string{{"00", "29"}, {"300", "599"}, {"6000", "7999"}, {"80000", "89999"}, {"900000", "999999"}}},
	"87": Range{
		Prefix: "978",
		Agency: "Denmark",
		Ranges: [][2]string{{"00", "29"}, {"400", "649"}, {"7000", "7999"}, {"85000", "94999"}, {"970000", "999999"}}},
	"88": Range{
		Prefix: "978",
		Agency: "Italy",
		Ranges: [][2]string{{"00", "19"}, {"200", "326"}, {"3270", "3389"}, {"339", "599"}, {"6000", "8499"}, {"85000", "89999"}, {"900000", "909999"}, {"910", "929"}, {"9300", "9399"}, {"940000", "947999"}, {"94800", "94999"}, {"95000", "99999"}}},
	"89": Range{
		Prefix: "978",
		Agency: "Korea, Republic",
		Ranges: [][2]string{{"00", "24"}, {"250", "549"}, {"5500", "8499"}, {"85000", "94999"}, {"950000", "969999"}, {"97000", "98999"}, {"990", "999"}}},
	"90": Range{
		Prefix: "978",
		Agency: "Netherlands",
		Ranges: [][2]string{{"00", "19"}, {"200", "499"}, {"5000", "6999"}, {"70000", "79999"}, {"800000", "849999"}, {"8500", "8999"}, {"90", "90"}, {"94", "94"}}},
	"91": Range{
		Prefix: "978",
		Agency: "Sweden",
		Ranges: [][2]string{{"0", "1"}, {"20", "49"}, {"500", "649"}, {"7000", "7999"}, {"85000", "94999"}, {"970000", "999999"}}},
	"92": Range{
		Prefix: "978",
		Agency: "International NGO Publishers and EU Organizations",
		Ranges: [][2]string{{"0", "5"}, {"60", "79"}, {"800", "899"}, {"9000", "9499"}, {"95000", "98999"}, {"990000", "999999"}}},
	"93": Range{
		Prefix: "978",
		Agency: "India",
		Ranges: [][2]string{{"00", "09"}, {"100", "499"}, {"5000", "7999"}, {"80000", "94999"}, {"950000", "999999"}}},
	"94": Range{
		Prefix: "978",
		Agency: "Netherlands",
		Ranges: [][2]string{{"000", "599"}, {"6000", "8999"}, {"90000", "99999"}}},
	"950": Range{
		Prefix: "978",
		Agency: "Argentina",
		Ranges: [][2]string{{"00", "49"}, {"500", "899"}, {"9000", "9899"}, {"99000", "99999"}}},
	"951": Range{
		Prefix: "978",
		Agency: "Finland",
		Ranges: [][2]string{{"0", "1"}, {"20", "54"}, {"550", "889"}, {"8900", "9499"}, {"95000", "99999"}}},
	"952": Range{
		Prefix: "978",
		Agency: "Finland",
		Ranges: [][2]string{{"00", "19"}, {"200", "499"}, {"5000", "5999"}, {"60", "65"}, {"6600", "6699"}, {"67000", "69999"}, {"7000", "7999"}, {"80", "94"}, {"9500", "9899"}, {"99000", "99999"}}},
	"953": Range{
		Prefix: "978",
		Agency: "Croatia",
		Ranges: [][2]string{{"0", "0"}, {"10", "14"}, {"150", "509"}, {"51", "54"}, {"55000", "59999"}, {"6000", "9499"}, {"95000", "99999"}}},
	"954": Range{
		Prefix: "978",
		Agency: "Bulgaria",
		Ranges: [][2]string{{"00", "28"}, {"2900", "2999"}, {"300", "799"}, {"8000", "8999"}, {"90000", "92999"}, {"9300", "9999"}}},
	"955": Range{
		Prefix: "978",
		Agency: "Sri Lanka",
		Ranges: [][2]string{{"0000", "1999"}, {"20", "37"}, {"38000", "38999"}, {"3900", "4099"}, {"41000", "43999"}, {"44000", "44999"}, {"4500", "4999"}, {"50000", "54999"}, {"550", "719"}, {"7200", "7499"}, {"7500", "7999"}, {"8000", "9499"}, {"95000", "99999"}}},
	"956": Range{
		Prefix: "978",
		Agency: "Chile",
		Ranges: [][2]string{{"00", "08"}, {"09000", "09999"}, {"10", "19"}, {"200", "599"}, {"6000", "6999"}, {"7000", "9999"}}},
	"957": Range{
		Prefix: "978",
		Agency: "Taiwan",
		Ranges: [][2]string{{"00", "02"}, {"0300", "0499"}, {"05", "19"}, {"2000", "2099"}, {"21", "27"}, {"28000", "30999"}, {"31", "43"}, {"440", "819"}, {"8200", "9699"}, {"97000", "99999"}}},
	"958": Range{
		Prefix: "978",
		Agency: "Colombia",
		Ranges: [][2]string{{"00", "53"}, {"5400", "5599"}, {"56000", "56999"}, {"57000", "59999"}, {"600", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"959": Range{
		Prefix: "978",
		Agency: "Cuba",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "99999"}}},
	"960": Range{
		Prefix: "978",
		Agency: "Greece",
		Ranges: [][2]string{{"00", "19"}, {"200", "659"}, {"6600", "6899"}, {"690", "699"}, {"7000", "8499"}, {"85000", "92999"}, {"93", "93"}, {"9400", "9799"}, {"98000", "99999"}}},
	"961": Range{
		Prefix: "978",
		Agency: "Slovenia",
		Ranges: [][2]string{{"00", "19"}, {"200", "599"}, {"6000", "8999"}, {"90000", "94999"}}},
	"962": Range{
		Prefix: "978",
		Agency: "Hong Kong, China",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "86999"}, {"8700", "8999"}, {"900", "999"}}},
	"963": Range{
		Prefix: "978",
		Agency: "Hungary",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"9000", "9999"}}},
	"964": Range{
		Prefix: "978",
		Agency: "Iran",
		Ranges: [][2]string{{"00", "14"}, {"150", "249"}, {"2500", "2999"}, {"300", "549"}, {"5500", "8999"}, {"90000", "96999"}, {"970", "989"}, {"9900", "9999"}}},
	"965": Range{
		Prefix: "978",
		Agency: "Israel",
		Ranges: [][2]string{{"00", "19"}, {"200", "599"}, {"7000", "7999"}, {"90000", "99999"}}},
	"966": Range{
		Prefix: "978",
		Agency: "Ukraine",
		Ranges: [][2]string{{"00", "12"}, {"130", "139"}, {"14", "14"}, {"1500", "1699"}, {"170", "199"}, {"2000", "2789"}, {"279", "289"}, {"2900", "2999"}, {"300", "699"}, {"7000", "8999"}, {"90000", "90999"}, {"910", "949"}, {"95000", "97999"}, {"980", "999"}}},
	"967": Range{
		Prefix: "978",
		Agency: "Malaysia",
		Ranges: [][2]string{{"00", "00"}, {"0100", "0999"}, {"10000", "19999"}, {"2000", "2499"}, {"300", "499"}, {"5000", "5999"}, {"60", "89"}, {"900", "989"}, {"9900", "9989"}, {"99900", "99999"}}},
	"968": Range{
		Prefix: "978",
		Agency: "Mexico",
		Ranges: [][2]string{{"01", "39"}, {"400", "499"}, {"5000", "7999"}, {"800", "899"}, {"9000", "9999"}}},
	"969": Range{
		Prefix: "978",
		Agency: "Pakistan",
		Ranges: [][2]string{{"0", "1"}, {"20", "22"}, {"23000", "23999"}, {"24", "39"}, {"400", "749"}, {"7500", "9999"}}},
	"970": Range{
		Prefix: "978",
		Agency: "Mexico",
		Ranges: [][2]string{{"01", "59"}, {"600", "899"}, {"9000", "9099"}, {"91000", "96999"}, {"9700", "9999"}}},
	"971": Range{
		Prefix: "978",
		Agency: "Philippines",
		Ranges: [][2]string{{"000", "015"}, {"0160", "0199"}, {"02", "02"}, {"0300", "0599"}, {"06", "49"}, {"500", "849"}, {"8500", "9099"}, {"91000", "95999"}, {"9600", "9699"}, {"97", "98"}, {"9900", "9999"}}},
	"972": Range{
		Prefix: "978",
		Agency: "Portugal",
		Ranges: [][2]string{{"0", "1"}, {"20", "54"}, {"550", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"973": Range{
		Prefix: "978",
		Agency: "Romania",
		Ranges: [][2]string{{"0", "0"}, {"100", "169"}, {"1700", "1999"}, {"20", "54"}, {"550", "759"}, {"7600", "8499"}, {"85000", "88999"}, {"8900", "9499"}, {"95000", "99999"}}},
	"974": Range{
		Prefix: "978",
		Agency: "Thailand",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8499"}, {"85000", "89999"}, {"90000", "94999"}, {"9500", "9999"}}},
	"975": Range{
		Prefix: "978",
		Agency: "Turkey",
		Ranges: [][2]string{{"00000", "01999"}, {"02", "23"}, {"2400", "2499"}, {"250", "599"}, {"6000", "9199"}, {"92000", "98999"}, {"990", "999"}}},
	"976": Range{
		Prefix: "978",
		Agency: "Caribbean Community",
		Ranges: [][2]string{{"0", "3"}, {"40", "59"}, {"600", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"977": Range{
		Prefix: "978",
		Agency: "Egypt",
		Ranges: [][2]string{{"00", "19"}, {"200", "499"}, {"5000", "6999"}, {"700", "849"}, {"85000", "89999"}, {"90", "99"}}},
	"978": Range{
		Prefix: "978",
		Agency: "Nigeria",
		Ranges: [][2]string{{"000", "199"}, {"2000", "2999"}, {"30000", "79999"}, {"8000", "8999"}, {"900", "999"}}},
	"979": Range{
		Prefix: "978",
		Agency: "Indonesia",
		Ranges: [][2]string{{"000", "099"}, {"1000", "1499"}, {"15000", "19999"}, {"20", "29"}, {"3000", "3999"}, {"400", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"980": Range{
		Prefix: "978",
		Agency: "Venezuela",
		Ranges: [][2]string{{"00", "19"}, {"200", "599"}, {"6000", "9999"}}},
	"981": Range{
		Prefix: "978",
		Agency: "Singapore",
		Ranges: [][2]string{{"00", "16"}, {"17000", "19999"}, {"200", "299"}, {"3000", "3099"}, {"310", "399"}, {"4000", "9999"}}},
	"982": Range{
		Prefix: "978",
		Agency: "South Pacific",
		Ranges: [][2]string{{"00", "09"}, {"100", "699"}, {"70", "89"}, {"9000", "9799"}, {"98000", "99999"}}},
	"983": Range{
		Prefix: "978",
		Agency: "Malaysia",
		Ranges: [][2]string{{"00", "01"}, {"020", "199"}, {"2000", "3999"}, {"40000", "44999"}, {"45", "49"}, {"50", "79"}, {"800", "899"}, {"9000", "9899"}, {"99000", "99999"}}},
	"984": Range{
		Prefix: "978",
		Agency: "Bangladesh",
		Ranges: [][2]string{{"00", "39"}, {"400", "799"}, {"8000", "8999"}, {"90000", "99999"}}},
	"985": Range{
		Prefix: "978",
		Agency: "Belarus",
		Ranges: [][2]string{{"00", "39"}, {"400", "599"}, {"6000", "8999"}, {"90000", "99999"}}},
	"986": Range{
		Prefix: "978",
		Agency: "Taiwan",
		Ranges: [][2]string{{"00", "11"}, {"120", "559"}, {"5600", "7999"}, {"80000", "99999"}}},
	"987": Range{
		Prefix: "978",
		Agency: "Argentina",
		Ranges: [][2]string{{"00", "09"}, {"1000", "1999"}, {"20000", "29999"}, {"30", "35"}, {"3600", "3999"}, {"4000", "4199"}, {"42", "43"}, {"4400", "4499"}, {"45000", "48999"}, {"4900", "4999"}, {"500", "899"}, {"9000", "9499"}, {"95000", "99999"}}},
	"988": Range{
		Prefix: "978",
		Agency: "Hong Kong, China",
		Ranges: [][2]string{{"00", "11"}, {"12000", "14999"}, {"15000", "16999"}, {"17000", "19999"}, {"200", "769"}, {"77000", "79999"}, {"8000", "9699"}, {"97000", "99999"}}},
	"989": Range{
		Prefix: "978",
		Agency: "Portugal",
		Ranges: [][2]string{{"0", "1"}, {"20", "54"}, {"550", "799"}, {"8000", "9499"}, {"95000", "99999"}}},
	"9924": Range{
		Prefix: "978",
		Agency: "Cambodia",
		Ranges: [][2]string{{"30", "39"}, {"500", "649"}, {"9000", "9999"}}},
	"9925": Range{
		Prefix: "978",
		Agency: "Cyprus",
		Ranges: [][2]string{{"0", "2"}, {"30", "54"}, {"550", "734"}, {"7350", "9999"}}},
	"9926": Range{
		Prefix: "978",
		Agency: "Bosnia and Herzegovina",
		Ranges: [][2]string{{"0", "1"}, {"20", "39"}, {"400", "799"}, {"8000", "9999"}}},
	"9927": Range{
		Prefix: "978",
		Agency: "Qatar",
		Ranges: [][2]string{{"00", "09"}, {"100", "399"}, {"4000", "4999"}}},
	"9928": Range{
		Prefix: "978",
		Agency: "Albania",
		Ranges: [][2]string{{"00", "09"}, {"100", "399"}, {"4000", "4999"}}},
	"9929": Range{
		Prefix: "978",
		Agency: "Guatemala",
		Ranges: [][2]string{{"0", "3"}, {"40", "54"}, {"550", "799"}, {"8000", "9999"}}},
	"9930": Range{
		Prefix: "978",
		Agency: "Costa Rica",
		Ranges: [][2]string{{"00", "49"}, {"500", "939"}, {"9400", "9999"}}},
	"9931": Range{
		Prefix: "978",
		Agency: "Algeria",
		Ranges: [][2]string{{"00", "29"}, {"300", "899"}, {"9000", "9999"}}},
	"9932": Range{
		Prefix: "978",
		Agency: "Lao People's Democratic Republic",
		Ranges: [][2]string{{"00", "39"}, {"400", "849"}, {"8500", "9999"}}},
	"9933": Range{
		Prefix: "978",
		Agency: "Syria",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9934": Range{
		Prefix: "978",
		Agency: "Latvia",
		Ranges: [][2]string{{"0", "0"}, {"10", "49"}, {"500", "799"}, {"8000", "9999"}}},
	"9935": Range{
		Prefix: "978",
		Agency: "Iceland",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9936": Range{
		Prefix: "978",
		Agency: "Afghanistan",
		Ranges: [][2]string{{"0", "1"}, {"20", "39"}, {"400", "799"}, {"8000", "9999"}}},
	"9937": Range{
		Prefix: "978",
		Agency: "Nepal",
		Ranges: [][2]string{{"0", "2"}, {"30", "49"}, {"500", "799"}, {"8000", "9999"}}},
	"9938": Range{
		Prefix: "978",
		Agency: "Tunisia",
		Ranges: [][2]string{{"00", "79"}, {"800", "949"}, {"9500", "9999"}}},
	"9939": Range{
		Prefix: "978",
		Agency: "Armenia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "899"}, {"9000", "9999"}}},
	"9940": Range{
		Prefix: "978",
		Agency: "Montenegro",
		Ranges: [][2]string{{"0", "1"}, {"20", "49"}, {"500", "899"}, {"9000", "9999"}}},
	"9941": Range{
		Prefix: "978",
		Agency: "Georgia",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9942": Range{
		Prefix: "978",
		Agency: "Ecuador",
		Ranges: [][2]string{{"00", "74"}, {"750", "849"}, {"8500", "8999"}, {"900", "984"}, {"9850", "9999"}}},
	"9943": Range{
		Prefix: "978",
		Agency: "Uzbekistan",
		Ranges: [][2]string{{"00", "29"}, {"300", "399"}, {"4000", "9749"}, {"975", "999"}}},
	"9944": Range{
		Prefix: "978",
		Agency: "Turkey",
		Ranges: [][2]string{{"0000", "0999"}, {"100", "499"}, {"5000", "5999"}, {"60", "69"}, {"700", "799"}, {"80", "89"}, {"900", "999"}}},
	"9945": Range{
		Prefix: "978",
		Agency: "Dominican Republic",
		Ranges: [][2]string{{"00", "00"}, {"010", "079"}, {"08", "39"}, {"400", "569"}, {"57", "57"}, {"580", "849"}, {"8500", "9999"}}},
	"9946": Range{
		Prefix: "978",
		Agency: "Korea, P.D.R.",
		Ranges: [][2]string{{"0", "1"}, {"20", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9947": Range{
		Prefix: "978",
		Agency: "Algeria",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"9948": Range{
		Prefix: "978",
		Agency: "United Arab Emirates",
		Ranges: [][2]string{{"00", "39"}, {"400", "849"}, {"8500", "9999"}}},
	"9949": Range{
		Prefix: "978",
		Agency: "Estonia",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "749"}, {"75", "89"}, {"9000", "9999"}}},
	"9950": Range{
		Prefix: "978",
		Agency: "Palestine",
		Ranges: [][2]string{{"00", "29"}, {"300", "849"}, {"8500", "9999"}}},
	"9951": Range{
		Prefix: "978",
		Agency: "Kosova",
		Ranges: [][2]string{{"00", "39"}, {"400", "849"}, {"8500", "9999"}}},
	"9952": Range{
		Prefix: "978",
		Agency: "Azerbaijan",
		Ranges: [][2]string{{"0", "1"}, {"20", "39"}, {"400", "799"}, {"8000", "9999"}}},
	"9953": Range{
		Prefix: "978",
		Agency: "Lebanon",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "599"}, {"60", "89"}, {"9000", "9299"}, {"93", "96"}, {"970", "999"}}},
	"9954": Range{
		Prefix: "978",
		Agency: "Morocco",
		Ranges: [][2]string{{"0", "1"}, {"20", "39"}, {"400", "799"}, {"8000", "9899"}, {"99", "99"}}},
	"9955": Range{
		Prefix: "978",
		Agency: "Lithuania",
		Ranges: [][2]string{{"00", "39"}, {"400", "929"}, {"9300", "9999"}}},
	"9956": Range{
		Prefix: "978",
		Agency: "Cameroon",
		Ranges: [][2]string{{"0", "0"}, {"10", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9957": Range{
		Prefix: "978",
		Agency: "Jordan",
		Ranges: [][2]string{{"00", "39"}, {"400", "649"}, {"65", "67"}, {"680", "689"}, {"690", "699"}, {"70", "84"}, {"8500", "8799"}, {"88", "99"}}},
	"9958": Range{
		Prefix: "978",
		Agency: "Bosnia and Herzegovina",
		Ranges: [][2]string{{"00", "01"}, {"020", "029"}, {"0300", "0399"}, {"040", "089"}, {"0900", "0999"}, {"10", "18"}, {"1900", "1999"}, {"20", "49"}, {"500", "899"}, {"9000", "9999"}}},
	"9959": Range{
		Prefix: "978",
		Agency: "Libya",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "949"}, {"9500", "9699"}, {"970", "979"}, {"98", "99"}}},
	"9960": Range{
		Prefix: "978",
		Agency: "Saudi Arabia",
		Ranges: [][2]string{{"00", "59"}, {"600", "899"}, {"9000", "9999"}}},
	"9961": Range{
		Prefix: "978",
		Agency: "Algeria",
		Ranges: [][2]string{{"0", "2"}, {"30", "69"}, {"700", "949"}, {"9500", "9999"}}},
	"9962": Range{
		Prefix: "978",
		Agency: "Panama",
		Ranges: [][2]string{{"00", "54"}, {"5500", "5599"}, {"56", "59"}, {"600", "849"}, {"8500", "9999"}}},
	"9963": Range{
		Prefix: "978",
		Agency: "Cyprus",
		Ranges: [][2]string{{"0", "1"}, {"2000", "2499"}, {"250", "279"}, {"2800", "2999"}, {"30", "54"}, {"550", "734"}, {"7350", "7499"}, {"7500", "9999"}}},
	"9964": Range{
		Prefix: "978",
		Agency: "Ghana",
		Ranges: [][2]string{{"0", "6"}, {"70", "94"}, {"950", "999"}}},
	"9965": Range{
		Prefix: "978",
		Agency: "Kazakhstan",
		Ranges: [][2]string{{"00", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9966": Range{
		Prefix: "978",
		Agency: "Kenya",
		Ranges: [][2]string{{"000", "149"}, {"1500", "1999"}, {"20", "69"}, {"7000", "7499"}, {"750", "959"}, {"9600", "9999"}}},
	"9967": Range{
		Prefix: "978",
		Agency: "Kyrgyz Republic",
		Ranges: [][2]string{{"00", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9968": Range{
		Prefix: "978",
		Agency: "Costa Rica",
		Ranges: [][2]string{{"00", "49"}, {"500", "939"}, {"9400", "9999"}}},
	"9970": Range{
		Prefix: "978",
		Agency: "Uganda",
		Ranges: [][2]string{{"00", "39"}, {"400", "899"}, {"9000", "9999"}}},
	"9971": Range{
		Prefix: "978",
		Agency: "Singapore",
		Ranges: [][2]string{{"0", "5"}, {"60", "89"}, {"900", "989"}, {"9900", "9999"}}},
	"9972": Range{
		Prefix: "978",
		Agency: "Peru",
		Ranges: [][2]string{{"00", "09"}, {"1", "1"}, {"200", "249"}, {"2500", "2999"}, {"30", "59"}, {"600", "899"}, {"9000", "9999"}}},
	"9973": Range{
		Prefix: "978",
		Agency: "Tunisia",
		Ranges: [][2]string{{"00", "05"}, {"060", "089"}, {"0900", "0999"}, {"10", "69"}, {"700", "969"}, {"9700", "9999"}}},
	"9974": Range{
		Prefix: "978",
		Agency: "Uruguay",
		Ranges: [][2]string{{"0", "2"}, {"30", "54"}, {"550", "749"}, {"7500", "8799"}, {"880", "909"}, {"91", "94"}, {"95", "99"}}},
	"9975": Range{
		Prefix: "978",
		Agency: "Moldova",
		Ranges: [][2]string{{"0", "0"}, {"100", "299"}, {"3000", "3999"}, {"4000", "4499"}, {"45", "89"}, {"900", "949"}, {"9500", "9999"}}},
	"9976": Range{
		Prefix: "978",
		Agency: "Tanzania",
		Ranges: [][2]string{{"0", "4"}, {"5000", "5999"}, {"60", "89"}, {"900", "989"}, {"9900", "9999"}}},
	"9977": Range{
		Prefix: "978",
		Agency: "Costa Rica",
		Ranges: [][2]string{{"00", "89"}, {"900", "989"}, {"9900", "9999"}}},
	"9978": Range{
		Prefix: "978",
		Agency: "Ecuador",
		Ranges: [][2]string{{"00", "29"}, {"300", "399"}, {"40", "94"}, {"950", "989"}, {"9900", "9999"}}},
	"9979": Range{
		Prefix: "978",
		Agency: "Iceland",
		Ranges: [][2]string{{"0", "4"}, {"50", "64"}, {"650", "659"}, {"66", "75"}, {"760", "899"}, {"9000", "9999"}}},
	"9980": Range{
		Prefix: "978",
		Agency: "Papua New Guinea",
		Ranges: [][2]string{{"0", "3"}, {"40", "89"}, {"900", "989"}, {"9900", "9999"}}},
	"9981": Range{
		Prefix: "978",
		Agency: "Morocco",
		Ranges: [][2]string{{"00", "09"}, {"100", "159"}, {"1600", "1999"}, {"20", "79"}, {"800", "949"}, {"9500", "9999"}}},
	"9982": Range{
		Prefix: "978",
		Agency: "Zambia",
		Ranges: [][2]string{{"00", "79"}, {"800", "989"}, {"9900", "9999"}}},
	"9983": Range{
		Prefix: "978",
		Agency: "Gambia",
		Ranges: [][2]string{{"80", "94"}, {"950", "989"}, {"9900", "9999"}}},
	"9984": Range{
		Prefix: "978",
		Agency: "Latvia",
		Ranges: [][2]string{{"00", "49"}, {"500", "899"}, {"9000", "9999"}}},
	"9985": Range{
		Prefix: "978",
		Agency: "Estonia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "899"}, {"9000", "9999"}}},
	"9986": Range{
		Prefix: "978",
		Agency: "Lithuania",
		Ranges: [][2]string{{"00", "39"}, {"400", "899"}, {"9000", "9399"}, {"940", "969"}, {"97", "99"}}},
	"9987": Range{
		Prefix: "978",
		Agency: "Tanzania",
		Ranges: [][2]string{{"00", "39"}, {"400", "879"}, {"8800", "9999"}}},
	"9988": Range{
		Prefix: "978",
		Agency: "Ghana",
		Ranges: [][2]string{{"0", "2"}, {"30", "54"}, {"550", "749"}, {"7500", "9999"}}},
	"9989": Range{
		Prefix: "978",
		Agency: "Macedonia",
		Ranges: [][2]string{{"0", "0"}, {"100", "199"}, {"2000", "2999"}, {"30", "59"}, {"600", "949"}, {"9500", "9999"}}},
	"99901": Range{
		Prefix: "978",
		Agency: "Bahrain",
		Ranges: [][2]string{{"00", "49"}, {"500", "799"}, {"80", "99"}}},
	"99902": Range{
		Prefix: "978",
		Agency: "Reserved Agency",
		Ranges: [][2]string{}},
	"99903": Range{
		Prefix: "978",
		Agency: "Mauritius",
		Ranges: [][2]string{{"0", "1"}, {"20", "89"}, {"900", "999"}}},
	"99904": Range{
		Prefix: "978",
		Agency: "Curaçao",
		Ranges: [][2]string{{"0", "5"}, {"60", "89"}, {"900", "999"}}},
	"99905": Range{
		Prefix: "978",
		Agency: "Bolivia",
		Ranges: [][2]string{{"0", "3"}, {"40", "79"}, {"800", "999"}}},
	"99906": Range{
		Prefix: "978",
		Agency: "Kuwait",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "699"}, {"70", "89"}, {"90", "94"}, {"950", "999"}}},
	"99908": Range{
		Prefix: "978",
		Agency: "Malawi",
		Ranges: [][2]string{{"0", "0"}, {"10", "89"}, {"900", "999"}}},
	"99909": Range{
		Prefix: "978",
		Agency: "Malta",
		Ranges: [][2]string{{"0", "3"}, {"40", "94"}, {"950", "999"}}},
	"99910": Range{
		Prefix: "978",
		Agency: "Sierra Leone",
		Ranges: [][2]string{{"0", "2"}, {"30", "89"}, {"900", "999"}}},
	"99911": Range{
		Prefix: "978",
		Agency: "Lesotho",
		Ranges: [][2]string{{"00", "59"}, {"600", "999"}}},
	"99912": Range{
		Prefix: "978",
		Agency: "Botswana",
		Ranges: [][2]string{{"0", "3"}, {"400", "599"}, {"60", "89"}, {"900", "999"}}},
	"99913": Range{
		Prefix: "978",
		Agency: "Andorra",
		Ranges: [][2]string{{"0", "2"}, {"30", "35"}, {"600", "604"}}},
	"99914": Range{
		Prefix: "978",
		Agency: "International NGO Publishers",
		Ranges: [][2]string{{"0", "4"}, {"50", "69"}, {"7", "7"}, {"80", "89"}, {"900", "999"}}},
	"99915": Range{
		Prefix: "978",
		Agency: "Maldives",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99916": Range{
		Prefix: "978",
		Agency: "Namibia",
		Ranges: [][2]string{{"0", "2"}, {"30", "69"}, {"700", "999"}}},
	"99917": Range{
		Prefix: "978",
		Agency: "Brunei Darussalam",
		Ranges: [][2]string{{"0", "2"}, {"30", "89"}, {"900", "999"}}},
	"99918": Range{
		Prefix: "978",
		Agency: "Faroe Islands",
		Ranges: [][2]string{{"0", "3"}, {"40", "79"}, {"800", "999"}}},
	"99919": Range{
		Prefix: "978",
		Agency: "Benin",
		Ranges: [][2]string{{"0", "2"}, {"300", "399"}, {"40", "69"}, {"70", "79"}, {"800", "849"}, {"850", "899"}, {"900", "999"}}},
	"99920": Range{
		Prefix: "978",
		Agency: "Andorra",
		Ranges: [][2]string{{"0", "4"}, {"50", "89"}, {"900", "999"}}},
	"99921": Range{
		Prefix: "978",
		Agency: "Qatar",
		Ranges: [][2]string{{"0", "1"}, {"20", "69"}, {"700", "799"}, {"8", "8"}, {"90", "99"}}},
	"99922": Range{
		Prefix: "978",
		Agency: "Guatemala",
		Ranges: [][2]string{{"0", "3"}, {"40", "69"}, {"700", "999"}}},
	"99923": Range{
		Prefix: "978",
		Agency: "El Salvador",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"99924": Range{
		Prefix: "978",
		Agency: "Nicaragua",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"99925": Range{
		Prefix: "978",
		Agency: "Paraguay",
		Ranges: [][2]string{{"0", "3"}, {"40", "79"}, {"800", "999"}}},
	"99926": Range{
		Prefix: "978",
		Agency: "Honduras",
		Ranges: [][2]string{{"0", "0"}, {"10", "59"}, {"600", "869"}, {"87", "89"}, {"90", "99"}}},
	"99927": Range{
		Prefix: "978",
		Agency: "Albania",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "999"}}},
	"99928": Range{
		Prefix: "978",
		Agency: "Georgia",
		Ranges: [][2]string{{"0", "0"}, {"10", "79"}, {"800", "999"}}},
	"99929": Range{
		Prefix: "978",
		Agency: "Mongolia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99930": Range{
		Prefix: "978",
		Agency: "Armenia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99931": Range{
		Prefix: "978",
		Agency: "Seychelles",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99932": Range{
		Prefix: "978",
		Agency: "Malta",
		Ranges: [][2]string{{"0", "0"}, {"10", "59"}, {"600", "699"}, {"7", "7"}, {"80", "99"}}},
	"99933": Range{
		Prefix: "978",
		Agency: "Nepal",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "999"}}},
	"99934": Range{
		Prefix: "978",
		Agency: "Dominican Republic",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"99935": Range{
		Prefix: "978",
		Agency: "Haiti",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "699"}, {"7", "8"}, {"90", "99"}}},
	"99936": Range{
		Prefix: "978",
		Agency: "Bhutan",
		Ranges: [][2]string{{"0", "0"}, {"10", "59"}, {"600", "999"}}},
	"99937": Range{
		Prefix: "978",
		Agency: "Macau",
		Ranges: [][2]string{{"0", "1"}, {"20", "59"}, {"600", "999"}}},
	"99938": Range{
		Prefix: "978",
		Agency: "Srpska, Republic of",
		Ranges: [][2]string{{"0", "1"}, {"20", "59"}, {"600", "899"}, {"90", "99"}}},
	"99939": Range{
		Prefix: "978",
		Agency: "Guatemala",
		Ranges: [][2]string{{"0", "5"}, {"60", "89"}, {"900", "999"}}},
	"99940": Range{
		Prefix: "978",
		Agency: "Georgia",
		Ranges: [][2]string{{"0", "0"}, {"10", "69"}, {"700", "999"}}},
	"99941": Range{
		Prefix: "978",
		Agency: "Armenia",
		Ranges: [][2]string{{"0", "2"}, {"30", "79"}, {"800", "999"}}},
	"99942": Range{
		Prefix: "978",
		Agency: "Sudan",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99943": Range{
		Prefix: "978",
		Agency: "Albania",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "999"}}},
	"99944": Range{
		Prefix: "978",
		Agency: "Ethiopia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99945": Range{
		Prefix: "978",
		Agency: "Namibia",
		Ranges: [][2]string{{"0", "5"}, {"60", "89"}, {"900", "999"}}},
	"99946": Range{
		Prefix: "978",
		Agency: "Nepal",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "999"}}},
	"99947": Range{
		Prefix: "978",
		Agency: "Tajikistan",
		Ranges: [][2]string{{"0", "2"}, {"30", "69"}, {"700", "999"}}},
	"99948": Range{
		Prefix: "978",
		Agency: "Eritrea",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99949": Range{
		Prefix: "978",
		Agency: "Mauritius",
		Ranges: [][2]string{{"0", "1"}, {"20", "89"}, {"900", "999"}}},
	"99950": Range{
		Prefix: "978",
		Agency: "Cambodia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99951": Range{
		Prefix: "978",
		Agency: "Reserved Agency",
		Ranges: [][2]string{}},
	"99952": Range{
		Prefix: "978",
		Agency: "Mali",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99953": Range{
		Prefix: "978",
		Agency: "Paraguay",
		Ranges: [][2]string{{"0", "2"}, {"30", "79"}, {"800", "939"}, {"94", "99"}}},
	"99954": Range{
		Prefix: "978",
		Agency: "Bolivia",
		Ranges: [][2]string{{"0", "2"}, {"30", "69"}, {"700", "879"}, {"88", "99"}}},
	"99955": Range{
		Prefix: "978",
		Agency: "Srpska, Republic of",
		Ranges: [][2]string{{"0", "1"}, {"20", "59"}, {"600", "799"}, {"80", "99"}}},
	"99956": Range{
		Prefix: "978",
		Agency: "Albania",
		Ranges: [][2]string{{"00", "59"}, {"600", "859"}, {"86", "99"}}},
	"99957": Range{
		Prefix: "978",
		Agency: "Malta",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"99958": Range{
		Prefix: "978",
		Agency: "Bahrain",
		Ranges: [][2]string{{"0", "4"}, {"50", "93"}, {"940", "949"}, {"950", "999"}}},
	"99959": Range{
		Prefix: "978",
		Agency: "Luxembourg",
		Ranges: [][2]string{{"0", "2"}, {"30", "59"}, {"600", "999"}}},
	"99960": Range{
		Prefix: "978",
		Agency: "Malawi",
		Ranges: [][2]string{{"0", "0"}, {"10", "94"}, {"950", "999"}}},
	"99961": Range{
		Prefix: "978",
		Agency: "El Salvador",
		Ranges: [][2]string{{"0", "2"}, {"300", "399"}, {"40", "89"}, {"900", "999"}}},
	"99962": Range{
		Prefix: "978",
		Agency: "Mongolia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99963": Range{
		Prefix: "978",
		Agency: "Cambodia",
		Ranges: [][2]string{{"00", "49"}, {"500", "919"}, {"92", "99"}}},
	"99964": Range{
		Prefix: "978",
		Agency: "Nicaragua",
		Ranges: [][2]string{{"0", "1"}, {"20", "79"}, {"800", "999"}}},
	"99965": Range{
		Prefix: "978",
		Agency: "Macau",
		Ranges: [][2]string{{"0", "2"}, {"300", "399"}, {"40", "62"}, {"630", "999"}}},
	"99966": Range{
		Prefix: "978",
		Agency: "Kuwait",
		Ranges: [][2]string{{"0", "2"}, {"30", "69"}, {"700", "799"}, {"80", "94"}}},
	"99967": Range{
		Prefix: "978",
		Agency: "Paraguay",
		Ranges: [][2]string{{"0", "1"}, {"20", "59"}, {"600", "899"}}},
	"99968": Range{
		Prefix: "978",
		Agency: "Botswana",
		Ranges: [][2]string{{"0", "3"}, {"400", "599"}, {"60", "89"}, {"900", "999"}}},
	"99969": Range{
		Prefix: "978",
		Agency: "Oman",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99970": Range{
		Prefix: "978",
		Agency: "Haiti",
		Ranges: [][2]string{{"0", "4"}, {"50", "89"}, {"900", "999"}}},
	"99971": Range{
		Prefix: "978",
		Agency: "Myanmar",
		Ranges: [][2]string{{"0", "5"}, {"60", "84"}, {"850", "999"}}},
	"99972": Range{
		Prefix: "978",
		Agency: "Faroe Islands",
		Ranges: [][2]string{{"0", "4"}, {"50", "89"}, {"900", "999"}}},
	"99973": Range{
		Prefix: "978",
		Agency: "Mongolia",
		Ranges: [][2]string{{"0", "3"}, {"40", "79"}, {"800", "999"}}},
	"99974": Range{
		Prefix: "978",
		Agency: "Bolivia",
		Ranges: [][2]string{{"40", "79"}, {"800", "999"}}},
	"99975": Range{
		Prefix: "978",
		Agency: "Tajikistan",
		Ranges: [][2]string{{"0", "3"}, {"40", "79"}, {"800", "999"}}},
	"99976": Range{
		Prefix: "978",
		Agency: "Srpska, Republic of",
		Ranges: [][2]string{{"0", "1"}, {"20", "59"}, {"600", "799"}}},
	"99977": Range{
		Prefix: "978",
		Agency: "Rwanda",
		Ranges: [][2]string{{"0", "1"}, {"40", "69"}, {"700", "799"}}},
	"99978": Range{
		Prefix: "978",
		Agency: "Mongolia",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"99979": Range{
		Prefix: "978",
		Agency: "Honduras",
		Ranges: [][2]string{{"0", "4"}, {"50", "79"}, {"800", "999"}}},
	"10": Range{
		Prefix: "979",
		Agency: "France",
		Ranges: [][2]string{{"00", "19"}, {"200", "699"}, {"7000", "8999"}, {"90000", "97599"}, {"976000", "999999"}}},
	"11": Range{
		Prefix: "979",
		Agency: "Korea, Republic",
		Ranges: [][2]string{{"00", "24"}, {"250", "549"}, {"5500", "8499"}, {"85000", "94999"}, {"950000", "999999"}}},
	"12": Range{
		Prefix: "979",
		Agency: "Italy",
		Ranges: [][2]string{{"200", "200"}}},
}