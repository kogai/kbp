package list62

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	HigherRate           = "H"
	TaxPaidAtSourceItaly = "P"
	LowerRate            = "R"
	StandardRate         = "S"
	ZeroRated            = "Z"
)

var (
	sortedCodes = []string{"H", "P", "R", "S", "Z"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		HigherRate:           {"Higher rate", "Specifies that tax is applied at a higher rate than standard.", "Høyere sats", "Mva. som er høyere en standard sats. Brukes ikke i Norge"},
		TaxPaidAtSourceItaly: {"Tax paid at source (Italy)", "Under Italian tax rules, VAT on books may be paid at source by the publisher, and subsequent transactions through the supply chain are tax-exempt.", "Moms betalt av forlaget (Italia)", "Etter Italienske skatteregler kan omsetningsavgift på bøker betales av forlaget. Alle transaksjoner videre i verdikjeden er uten ytterligere moms."},
		LowerRate:            {"Lower rate", "Specifies that tax is applied at a lower rate than standard.", "Lav sats", "I Norge: matmoms satsen"},
		StandardRate:         {"Standard rate", "", "Høy sats", "I Norge: normal moms"},
		ZeroRated:            {"Zero-rated", "", "Uten moms", ""},
	}
)

// Items returns alle the possible items in this list, with labels and description in the requested language.
func Items(lang codes.Language) (res []codes.Item) {
	for _, code := range sortedCodes {
		if lang == codes.Norwegian {
			res = append(res, codes.Item{Code: code, Label: all[code].labelNo, Notes: all[code].notesNo})
		} else {
			res = append(res, codes.Item{Code: code, Label: all[code].labelEn, Notes: all[code].notesEn})
		}
	}
	return res
}

// Item return the Item for the given code and language, or an error if it doesn't exist.
func Item(code string, lang codes.Language) (codes.Item, error) {
	item, ok := all[code]
	if !ok {
		return codes.Item{}, fmt.Errorf("no item with code: %q", code)
	}
	if lang == codes.Norwegian {
		return codes.Item{Code: code, Label: item.labelNo, Notes: item.notesNo}, nil
	}
	return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}, nil
}

// MustItem returns the Item for the given code. It will panic if it doesn't exist.
func MustItem(code string, lang codes.Language) codes.Item {
	item, ok := all[code]
	if !ok {
		panic(fmt.Errorf("no item with code: %q", code))
	}
	if lang == codes.Norwegian {
		return codes.Item{Code: code, Label: item.labelNo, Notes: item.notesNo}
	}
	return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}

}
