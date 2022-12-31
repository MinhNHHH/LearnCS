package stringbytes

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func StringPrint() {
	fmt.Println(basename("a/b/c.go")) // "c"
}

// basename ver1
// func basename(s string) string {
// 	// Discard last '/' and everything before.
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '/' {
// 			s = s[i+1:]
// 			break
// 		}
// 	}
// 	// Preserve everything before last '.'.
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '.' {
// 			s = s[:i]
// 			break
// 		}
// 	}
// 	return s
// }

// intsToString is like fmt.Sprintf(values) but adds commas.
func IntsToString(values string) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteString(Comma(values))
	buf.WriteByte(']')
	return buf.String()
}

func Comma(s string) string {

	var buf bytes.Buffer
	l := len(s)
	index := 0

	if l <= 3 {
		return s
	}
	r := l % 3 // remainder modulo 3 to get initial digits

	// insert comma after initial digits
	if r >= 1 {
		buf.WriteString(s[:r])
		buf.WriteString(",")
	}
	for i := r; i < l; i += 3 {

		if string(s[i]) == "." {
			buf.WriteString(s[index:])
			break
		}
		buf.WriteString(s[i : i+3])
		index = i + 3
		if i+3 < l && string(s[i+3]) != "." {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	fmt.Println("sssss", slash)
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func SortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func CheckAnagrams(s1 string, s2 string) bool {
	if SortString(s1) == SortString(s2) {
		return true
	}
	return false
}
