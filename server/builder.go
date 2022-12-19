package server

func buildQueryFilter(query map[string][]string) (result map[string]string) {
	result = make(map[string]string)
	for k, v := range query {
		switch k {
		case "name":
			result[k] = v[0]
		case "options":
			result[k] = v[0]
		case "type":
			result[k] = v[0]
		case "sort_by":
			result[k] = v[0]
		case "order_by":
			result[k] = v[0]
		}
	}

	return result
}
