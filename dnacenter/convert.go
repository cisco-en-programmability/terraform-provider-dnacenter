package dnacenter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func interfaceToFloat64Ptr(item interface{}) *float64 {
	nItem := interfaceToString(item)
	nnItem, err := strconv.ParseFloat(nItem, 64)
	if err != nil {
		return nil
	}
	return &nnItem
}

func mapInterfaceToMapString(m map[string]interface{}) map[string]string {
	new_m := map[string]string{}
	for k, v := range m {
		new_m[k] = interfaceToString(v)
	}
	return new_m
}

func interfaceToIntPtr(item interface{}) *int {
	nItem := interfaceToString(item)
	nnItem, err := strconv.Atoi(nItem)
	if err != nil {
		return nil
	}
	return &nnItem
}

func boolPtrToString(item *bool) string {
	if item == nil {
		return ""
	}
	if *item {
		return "true"
	}
	return "false"
}

func interfaceToBoolPtr(item interface{}) *bool {
	nItem := interfaceToString(item)
	if nItem != "" {
		nItemBool := nItem == "true"
		return &nItemBool
	}
	return nil
}

func getResourceItems(item interface{}) *[]map[string]interface{} {
	vItems, ok1 := item.([]interface{})
	if !ok1 {
		return nil
	}
	if len(vItems) <= 0 {
		return nil
	}
	vvItems := []map[string]interface{}{}
	for _, vItem := range vItems {
		vvItem, ok2 := vItem.(map[string]interface{})
		if !ok2 {
			continue
		}
		vvItems = append(vvItems, vvItem)
	}
	return &vvItems
}

func getResourceItem(item interface{}) *map[string]interface{} {
	vItems, ok1 := item.([]interface{})
	if !ok1 {
		return nil
	}
	if len(vItems) <= 0 {
		return nil
	}
	vItem := vItems[0]
	vvItem, ok2 := vItem.(map[string]interface{})
	if !ok2 {
		return nil
	}
	return &vvItem
}

func interfaceToSliceString(v interface{}) []string {
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	newValue := []string{}
	for _, i := range value {
		newValue = append(newValue, interfaceToString(i))
	}
	return newValue
}

func interfaceToSliceInt(v interface{}) *[]int {
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	newValue := []int{}
	for _, i := range value {
		value := interfaceToIntPtr(i)
		if value != nil {
			newValue = append(newValue, *value)
		}
	}
	return &newValue
}

func interfaceToString(v interface{}) string {
	return fmt.Sprint(v)
}

func responseInterfaceToSliceString(v interface{}) []string {
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	newValue := []string{}
	for _, i := range value {
		newValue = append(newValue, responseInterfaceToString(i))
	}
	return newValue
}

func responseInterfaceToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return fmt.Sprint(string(b))
}

func stringToFloat64Ptr(v string) *float64 {
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		return &s
	}
	return nil
}

func stringToIntPtr(v string) *int {
	if s, err := strconv.Atoi(v); err == nil {
		return &s
	}
	return nil
}

func stringToBooleanPtr(v string) *bool {
	if s, err := strconv.ParseBool(v); err == nil {
		return &s
	}
	return nil
}
