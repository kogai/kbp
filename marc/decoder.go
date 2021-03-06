package marc

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

func Unmarshal(b []byte, f Format, rec *Record) error {
	switch f {
	case MARCXML:
		dec := Decoder{
			xmlDec: xml.NewDecoder(bytes.NewBuffer(b)),
		}
		for {
			t, err := dec.xmlDec.Token()
			if err != nil {
				return err
			}
			switch elem := t.(type) {
			case xml.SyntaxError:
				return errors.New(elem.Error())
			case xml.StartElement:
				if elem.Name.Local == "record" {
					return dec.decodeMARCXMLRecord(rec)
				}
			}
		}

	default:
		panic("Unmarshal: TODO")
	}
}

// Decoder can decode MARC records from a stream, in one of the supported formats:
// MARCXML (ISO25577), LineMARC or Standard MARC (ISO2709)
type Decoder struct {
	input  *bufio.Reader
	xmlDec *xml.Decoder
	format Format
	lineN  int
}

// NewDecoder returns a new Decoder for the given stream and format.
func NewDecoder(r io.Reader, f Format) *Decoder {
	switch f {
	case MARCXML:
		return &Decoder{
			xmlDec: xml.NewDecoder(r),
			format: f,
		}
	case MARC, LineMARC:
	default:
		panic("Cannot decode unknown MARC Format")
	}

	return &Decoder{
		input:  bufio.NewReader(r),
		format: f,
	}
}

// DecodeAll consumes the input stream and returns all decoded records.
// If there is an error, it will return, together with the succesfully
// parsed MARC records up til then.
func (d *Decoder) DecodeAll() ([]*Record, error) {
	res := make([]*Record, 0)
	for r, err := d.Decode(); err != io.EOF; r, err = d.Decode() {
		if err != nil {
			return res, err
		}
		res = append(res, r)
	}
	return res, nil
}

// Decode decodes and returns a single MARC Record, or and error.
func (d *Decoder) Decode() (*Record, error) {
	switch d.format {
	case LineMARC:
		return d.decodeLineMARC()
	case MARCXML:
		return d.decodeMARCXML()
	case MARC:
		panic("TODO decode ISOMARC")
	default:
		panic("Cannot decode unknown MARC Format")
	}
}

// LineMARC-specific constants
const (
	linemarcRT  = 0x5E // ^ (record terminator)
	linemarcFS  = 0x2A // * (field separator)
	linemarcSFS = 0x24 // $ (subfield separator)
)

func (d *Decoder) decodeMARCXML() (*Record, error) {
	for {
		t, err := d.xmlDec.Token()
		if err != nil {
			return nil, err
		}
		switch elem := t.(type) {
		case xml.SyntaxError:
			return nil, errors.New(elem.Error())
		case xml.StartElement:
			if elem.Name.Local == "record" {
				r := NewRecord()
				err := d.decodeMARCXMLRecord(r)
				return r, err
			}
		}
	}
}

func (d *Decoder) decodeMARCXMLRecord(r *Record) error {
outer:
	for {
		t, err := d.xmlDec.Token()
		if err != nil {
			return err
		}
		switch elem := t.(type) {
		case xml.SyntaxError:
			return errors.New(elem.Error())
		case xml.StartElement:
			switch elem.Name.Local {
			case "leader":
				for t, err := d.xmlDec.Token(); err == nil; t, _ = d.xmlDec.Token() {
					switch elem := t.(type) {
					case xml.CharData:
						copy(r.leader, elem)
					case xml.EndElement:
						continue outer
					}
				}
			case "controlfield":
				var tag ControlTag
				for _, a := range elem.Attr {
					if a.Name.Local == "tag" {
						var err error
						tag, err = controlTagFromString(a.Value)
						if err != nil {
							return err
						}
					}
				}
				cf := NewControlField(tag)
			controlfield:
				for t, err := d.xmlDec.Token(); err == nil; t, err = d.xmlDec.Token() {
					switch elem := t.(type) {
					case xml.CharData:
						cf.value = make([]byte, len(elem))
						copy(cf.value, elem)
					case xml.EndElement:
						break controlfield
					}
				}
				r.AddControlField(cf)
			case "datafield":
				var tag DataTag
				var i1, i2 rune
				for _, a := range elem.Attr {
					switch a.Name.Local {
					case "tag":
						var err error
						tag, err = dataTagFromString(a.Value)
						if err != nil {
							return err
						}
					case "ind1":
						i1 = rune(a.Value[0])
					case "ind2":
						i2 = rune(a.Value[0])
					}
				}
				df := NewDataFieldWithIndicators(tag, i1, i2)
			datafield:
				for t, err := d.xmlDec.Token(); err == nil; t, err = d.xmlDec.Token() {
					switch elem := t.(type) {
					case xml.StartElement:
						if elem.Name.Local == "subfield" {
							var code rune
							for _, a := range elem.Attr {
								if a.Name.Local == "code" {
									code, _ = utf8.DecodeRuneInString(a.Value)
								}
							}
						subfield:
							for t, err := d.xmlDec.Token(); err == nil; t, err = d.xmlDec.Token() {
								switch elem := t.(type) {
								case xml.CharData:
									df.Add(code, string(elem))
								case xml.EndElement:
									break subfield
								}
							}
						}
					case xml.EndElement:
						break datafield
					}
				}
				r.AddDataField(df)
			}
		case xml.EndElement:
			if elem.Name.Local == "record" {
				return nil
			}
		}
	}
}

func (d *Decoder) decodeLineMARC() (r *Record, err error) {
	r = NewRecord()
decodeRecord:
	for {
		line, err := d.input.ReadBytes('\n')
		if err != nil && len(line) == 0 {
			return r, err
		}
		d.lineN++

		switch line[0] {
		case linemarcRT:
			break decodeRecord
		case '\n':
			continue decodeRecord // could be extra newline between records
		case linemarcFS:
			// OK
		default:
			err = fmt.Errorf("%d:0: expected '*', found %q", d.lineN, string(line))
			break decodeRecord
		}

		// strip newline
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		if len(line) < 4 || !isAllDigits(line[1:4]) {
			err = fmt.Errorf(
				"%d:0: expected a valid control field or data field number, found %q",
				d.lineN, string(line))
			break
		}

		p := 1 // position in line

		if bytes.HasPrefix(line[p:], []byte("00")) {
			if (line[p+2]) == '0' {
				// Parse leader
				copy(r.leader, line[p+3:])
			} else {
				// Parse Control field
				cField, err2 := parseControlField(line[p:])
				if err != nil {
					err = err2
					break
				}
				r.AddControlField(&cField)
			}
		} else {
			// Parse Data field
			dField, err2 := parseDataField(line[p:])
			if err2 != nil {
				err = err2
				break decodeRecord
			}
			if len(dField.subfields) > 0 {
				r.AddDataField(&dField)
			}
		}

	}
	return r, err
}

// DetectFormat tries to detect the MARC encoding of the given byte slice. It
// detects one of LineMARC/MARC/MARCXML, otherwise unknown.
func DetectFormat(data []byte) Format {
	// Find the first non-whitespace byte
	i := 0
	for ; i < len(data) && isWS(data[i]); i++ {
	}

	switch data[i] {
	case '<':
		return MARCXML
	case '*': // TODO, also '^' ?
		return LineMARC
	default:
		if data[i] >= '0' && data[i] <= '9' {
			return MARC
		}
		return unknownFormat
	}
}

// Parsing helper functions:

func parseControlField(b []byte) (ControlField, error) {
	// We can asume that len(b) >= 3, and that b[0:3] is numeric,
	// and that b[3] != '0'.
	f := ControlField{
		Tag:   ControlTag(byteToInt(b[:3])),
		value: b[3:],
	}
	return f, nil
}

func parseDataField(b []byte) (DataField, error) {
	// We can asume that len(b) >= 3, and that b[0:3] is numeric.
	f := DataField{
		Tag:       DataTag(byteToInt(b[:3])),
		subfields: make(map[rune][]string),
	}
	if len(b) >= 5 {
		f.Indicator1 = rune(b[3])
		f.Indicator2 = rune(b[4])
	}

	p := 4 // position in b
subfields:
	for p+2 < len(b) {
		p += 2       // dollarsign + code
		code := b[p] // TODO use DecodeRune; Bibsys uses multibyte char like 'ø' as subfield code
		p++
		start := p
		for {
			ch, w := utf8.DecodeRune(b[p:])
			p += w

			if ch == linemarcSFS {
				f.Add(rune(code), string(b[start:p-1]))
				p -= 2
				continue subfields
			}

			if w == 0 { // eof
				if p > start {
					f.Add(rune(code), string(b[start:p]))
					break subfields
				}
			}
		}
	}

	return f, nil
}

func isWS(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAllDigits(b []byte) bool {
	for _, c := range b {
		if !isDigit(c) {
			return false
		}
	}
	return true
}

func byteToInt(in []byte) (x int) {
	for _, c := range in {
		x = x*10 + int(c-'0')
	}
	return x
}
