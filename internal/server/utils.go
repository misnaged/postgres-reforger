package server

import (
	"fmt"
	"strings"
)

func FormatPositionString(in string) (out string) {
	inArr := strings.Split(in, ", ")
	inArr[0] = strings.TrimPrefix(inArr[0], "<")
	inArr[len(inArr)-1] = strings.TrimSuffix(inArr[len(inArr)-1], ">")
	out = strings.Join(inArr, " ")
	return
}

func FormatPrefabSymbols(in string, isRounded bool) (out string) {
	inArr := strings.Split(in, "/")
	inArr2 := strings.Split(inArr[0], "")
	for i := range inArr2 {
		if !isRounded {
			if inArr2[i] == "{" {

				inArr2[i] = "("
			}
			if inArr2[i] == "}" {

				inArr2[i] = ")"
			}
		} else {
			if inArr2[i] == "(" {

				inArr2[i] = "{"
			}
			if inArr2[i] == ")" {

				inArr2[i] = "}"
			}

		}
	}
	t := strings.Join(inArr2, "")
	inArr[0] = t
	out = strings.Join(inArr, "/")
	return
}
func Aad(in []string) []string {
	out := strings.Join(in, " ")
	fmt.Printf("\n\n out: %s \n\n", out)
	ii := strings.Split(out, "")
	for i := range ii {
		if ii[i] == `"` {

		}
	}
	return in
}
