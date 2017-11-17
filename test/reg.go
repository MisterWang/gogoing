package test

import (
	"regexp"
)

func Test_reg(exp string,str string)[]string {
	re:=regexp.MustCompile(exp)

	return re.FindStringSubmatch(str)
}