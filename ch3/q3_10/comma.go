package q3_10

import "bytes"

func Comma(s string) string {
	var buf bytes.Buffer

	x := len(s) % 3
	if x == 0 {
		x = 3
	}

	buf.WriteString(s[:x])
	for ; x < len(s); x += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[x : x+3])
	}
	return buf.String()
}