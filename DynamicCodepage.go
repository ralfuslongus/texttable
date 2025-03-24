package texttable

type DynamicCodepage struct {
	rune_byte_map map[rune]byte
	byte_rune_map map[byte]rune
	top           int
}

const (
	// alle runes, die nicht mehr gemapped werden können (>MAX_TOP) werden als 0-bytes
	// encoded. Beim decoden werden alle 0-bytes zu dem rune UNMAPPED decoded
	UNMAPPED rune = '?'
	MAX      int  = 255
)

// TODO: Keine DEFAULT_CODEPAGE mehr anbieten, Tables müssen ihre Zellen mit einer eigenen
//
//	Codepage versorgen, oder so
var (
	DEFAULT_CODEPAGE = NewDynamicCodepage()
)

func NewDynamicCodepage() *DynamicCodepage {
	cp := DynamicCodepage{
		rune_byte_map: make(map[rune]byte),
		byte_rune_map: make(map[byte]rune),
	}
	return &cp
}

func (cp *DynamicCodepage) Encode(s string) []byte {
	bytes := make([]byte, 0, len(s))
	for _, r := range s {
		b := cp.encodeRune(r)

		bytes = append(bytes, b)
		// println("rune:", r, "byte:", b)
		// fmt.Printf("buffer:%v\n", buffer)
	}
	// println("encoded", s, "to mapped bytes")
	return bytes
}
func (cp *DynamicCodepage) encodeRune(r rune) byte {
	b, exists := cp.rune_byte_map[r]
	if !exists {
		if cp.top >= MAX {
			// if TOP == 0 {
			// panic("TOP overflow, can not use more than 256 different runes")
			// println("TOP overflow, can not use more than 256 different runes")
			b = 0
			// println("encoded unmapped rune", r, "to byte", b)
		} else {
			cp.top++
			b = byte(cp.top)
			// println("encoded rune", r, "to byte", b)
			cp.rune_byte_map[r] = b
			cp.byte_rune_map[b] = r
			// println("new mapping from:", string(r), "to", b)
		}
	}
	return b
}

func (cp *DynamicCodepage) Decode(bytes []byte) string {
	runes := make([]rune, 0, len(bytes))
	for _, b := range bytes {
		r := cp.ecodeByte(b)
		// println("decoded rune", r, "from mapped byte", b)
		runes = append(runes, r)
	}
	s := string(runes)
	// println("decoded", s, "from mapped bytes")
	return s
}

func (cp *DynamicCodepage) ecodeByte(b byte) rune {
	if b <= 0 || b > byte(MAX) {
		// println("decoded unmapped byte", b, "to rune", UNMAPPED)
		return rune(UNMAPPED)
	} else {
		r := cp.byte_rune_map[b]
		// println("decoded byte", b, "to rune", r)
		return r
	}
}
