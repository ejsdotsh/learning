package reverse

import "strings"

func Reverse(input string) string {
	ch := strings.Split(input, "")
	var out []string
	for i := len(ch) - 1; i >= 0; i-- {
		out = append(out, string(ch[i]))
	}

	return strings.Join(out, "")
}
