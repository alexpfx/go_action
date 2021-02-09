package util

import "strings"

func SplitSep(str, sep string) [] string{
	if strings.TrimSpace(sep) == ""{
		return strings.Fields(str)
	}
	
	return strings.Split(str, sep)
	
}