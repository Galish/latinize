package latinize

import "strings"

func Latinize(input string) string {
	r := strings.NewReplacer(characters...)
	return r.Replace(input)
}
