package hexdump

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

func prefix(num int) string {
	return fmt.Sprintf("%016x", num)
}

func Dump(r io.Reader) (io.Writer, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := &strings.Builder{}
	var chars []byte
	var ct int
	for i, b := range data {
		if i == 0 || i%16 == 0 {
			if i > 0 {
				s.WriteString("|")
				s.WriteString(string(chars))
				s.WriteString("|")
				s.WriteString("\n")
			}
			ct = 0
			chars = make([]byte, 16)
			s.Grow(80)
			s.WriteString(prefix(i))
			s.WriteString(" ")
		}
		s.WriteString(fmt.Sprintf("%02x ", b))

		if unicode.IsPrint(rune(b)) {
			chars[ct] = b
		} else {
			chars[ct] = byte('.')
		}
		ct++

		if i%16 != 0 && i == len(data)-1 {
			s.WriteString("|")
			s.WriteString(string(chars))
			s.WriteString("|")
		}
	}
	return s, nil
}
