package service

import "strings"

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
