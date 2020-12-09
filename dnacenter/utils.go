package dnacenter

import (
	"encoding/json"
)

func convertSliceInterfaceToSliceString(interfaceSlice []interface{}) []string {
	stringSlice := []string{}
	if interfaceSlice != nil {
		for _, v := range interfaceSlice {
			if v != nil {
				stringSlice = append(stringSlice, v.(string))
			}
		}
	}
	return stringSlice
}

func convertSliceInterfaceToString(interfaceSlice []interface{}) string {
	b, err := json.Marshal(interfaceSlice)
	if err != nil {
		return ""
	}
	return string(b)
}

func convertInterfaceToString(interfaceElem interface{}) string {
	b, err := json.Marshal(interfaceElem)
	if err != nil {
		return ""
	}
	return string(b)
}

func hasSameSliceString(aSlice, bSlice []string) bool {
	lenA := len(aSlice)
	lenB := len(bSlice)
	if lenA != lenB {
		return false
	}
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i, ai := range aSlice {
		for j, bi := range bSlice {
			_, ok1 := m1[i]
			_, ok2 := m2[j]
			if ai == bi && !ok1 && !ok2 {
				m1[i] = j
				m2[j] = i
				break
			}
		}
	}
	return lenA == len(m1)
}
