package dnacenter

func convertSliceInterfaceToSliceString(interfaceSlice []interface{}) []string {
	stringSlice := []string{}
	for _, v := range interfaceSlice {
		stringSlice = append(stringSlice, v.(string))
	}
	return stringSlice
}
