package gofi

import "fmt"

func Add(argList []string, arg string, value interface{}) []string {
	bp, ok := value.(bool)
	if ok {
		if bp {
			return append(argList, arg)
		}
		return argList
	}

	st, ok := value.(string)
	argList = append(argList, arg)
	if ok && st != "" {
		argList = append(argList, fmt.Sprintf("%s", st))
		return argList
	}

	argList = append(argList, fmt.Sprintf("%v", value))

	return argList
}
