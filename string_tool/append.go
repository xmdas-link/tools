package string_tool

import "fmt"

func AppendId(ids []string, id interface{}) []string {

	var (
		checkId = fmt.Sprint(id)
		flag    = true
	)

	for i := range ids {
		if ids[i] == checkId {
			flag = false
		}
	}

	if flag {
		return append(ids, checkId)
	}

	return ids

}
