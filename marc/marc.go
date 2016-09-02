package marc

import "bytes"

type Flavor int

const (
	unknownFlavor Flavor = iota
	MARC21
	NORMARC
	UNIMARC
)

type Format int

const (
	unknownFormat Format = iota
	MARC
	LineMARC
	MARCXML
)

// String returns a string representation of a Format.
func (f Format) String() string {
	switch f {
	case unknownFormat:
		return "Unknown MARC format"
	case MARC:
		return "Standard MARC (ISO2709)"
	case LineMARC:
		return "Line-MARC"
	case MARCXML:
		return "MarcXchange (ISO25577)"
	default:
		panic("unreachable")
	}
}

// Record represents a MARC record.
type Record struct {
	leader  []byte
	cfields map[ControlTag][]byte
	dfields map[DataTag][]DataField
}

var emptyLeader = []byte("                        ") // 24 spaces

// NewRecord returns a new MARC Record.
func NewRecord() *Record {
	leader := make([]byte, 24)
	copy(leader, emptyLeader)
	return &Record{
		leader:  leader,
		cfields: make(map[ControlTag][]byte),
		dfields: make(map[DataTag][]DataField),
	}
}

var controltags = []ControlTag{Tag001, Tag002, Tag003, Tag004, Tag005, Tag006, Tag007, Tag008}

var datatags = []DataTag{
	Tag010, Tag020, Tag041, Tag111, Tag245, Tag260, Tag300, Tag440, Tag508, Tag546, Tag700, Tag911,
}

func (r *Record) String() string {
	var b bytes.Buffer
	b.WriteString("000  ")
	b.Write(r.leader)
	b.WriteRune('\n')
	for _, tag := range controltags {
		if val, ok := r.cfields[tag]; ok {
			b.WriteString(controlfields2str[tag])
			b.WriteString("  ")
			b.Write(val)
			b.WriteRune('\n')
		}
	}
	for _, tag := range datatags {
		if dfs, ok := r.dfields[tag]; ok {
			for _, df := range dfs {
				b.WriteString(datafield2str[tag])
				b.WriteRune(df.Indicator1)
				b.WriteRune(df.Indicator2)
				// TODO sort codes alphabetically
				indented := true
				for code, vals := range df.subfields {
					for _, v := range vals {
						if !indented {
							b.WriteString("\n     ")
						}
						b.WriteRune('$')
						b.WriteRune(code)
						b.WriteRune(' ')
						b.WriteString(v)
						indented = false
					}
				}
			}
			b.WriteRune('\n')
		}
	}
	return b.String()
}

type leaderPos int

const (
	LeaderRecordStatus leaderPos = iota + 5
	LeaderRecordType
	LeaderBibliograhicLevel
	LeaderControlType
	LeaderCharacterEncoding
	LeaderIndicatorCount
	LeaderSubfieldCodeCount
	_ // pos 12-16: "Base adress of data"
	_
	_
	_
	_
	LeaderEncodingLevel
)

// SetLeaderPos sets the Record leader position to the given value.
//
// It is not possible to set the values for "Record length" or
// "Base address of data"; they are computed automatically when
// serializing a Record.
func (r *Record) SetLeaderPos(pos leaderPos, val byte) *Record {
	r.leader[int(pos)] = val
	return r
}

func (r *Record) AddControlField(f ControlField) *Record {
	r.cfields[f.Tag] = f.value
	return r
}

func (r *Record) AddDataField(f DataField) *Record {
	r.dfields[f.Tag] = append(r.dfields[f.Tag], f)
	return r
}

func NewDataField(tag DataTag) DataField {
	return DataField{
		Tag:        tag,
		Indicator1: ' ',
		Indicator2: ' ',
		subfields:  make(map[rune][]string),
	}
}

func NewDataFieldWithIndicators(tag DataTag, ind1, ind2 rune) DataField {
	return DataField{
		Tag:        tag,
		Indicator1: ind1,
		Indicator2: ind2,
		subfields:  make(map[rune][]string),
	}
}

type DataField struct {
	Tag        DataTag
	Indicator1 rune
	Indicator2 rune
	subfields  map[rune][]string
}

func (df DataField) Add(code rune, value string) DataField {
	df.subfields[code] = append(df.subfields[code], value)
	return df
}

type ControlField struct {
	Tag   ControlTag
	value []byte
}

func NewControlField(tag ControlTag) ControlField {
	return ControlField{Tag: tag}
}

func (f ControlField) Set(s string) ControlField {
	f.value = []byte(s)
	return f
}

func (f ControlField) SetPos(pos int, s string) ControlField {
	if l := pos + len(s); len(f.value) < l {
		b := make([]byte, l)
		for i := range b {
			b[i] = ' '
		}
		copy(b, f.value)
		f.value = b
	}
	copy(f.value[pos:], s)
	return f
}

func (r *Record) Eq(other *Record) bool {
	if !bytes.Equal(r.leader, other.leader) {
		return false
	}
	for tag, a := range r.cfields {
		b, ok := other.cfields[tag]
		if !ok {
			return false
		}
		if !bytes.Equal(a, b) {
			return false
		}
	}
	for tag, _ := range r.dfields {
		if _, ok := other.dfields[tag]; !ok {
			return false
		}
		// TODO check eq of of dfields[tag] (need to sort them)
	}
	return true
}
