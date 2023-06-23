package sliceutil

import "reflect"

func StringToInterface(input []string) []interface{} {
	s := make([]interface{}, len(input))
	for i, v := range input {
		s[i] = v
	}
	return s
}

func IntToInterface(input []int) []interface{} {
	s := make([]interface{}, len(input))
	for i, v := range input {
		s[i] = v
	}
	return s
}

func InterfaceToString(input []interface{}) []string {
	list := make([]string, 0)
	if reflect.TypeOf(input).Kind() == reflect.Slice {
		for i := 0; i < len(input); i++ {
			list = append(list, input[i].(string))
		}
	}
	return list
}

func InterfaceToInt(input []interface{}) []int {
	list := make([]int, 0)
	if reflect.TypeOf(input).Kind() == reflect.Slice {
		for i := 0; i < len(input); i++ {
			list = append(list, input[i].(int))
		}
	}
	return list
}

func ReverseInterface(lst []interface{}) []interface{} {
	n := len(lst)
	for i := 0; i < n/2; i++ {
		lst[i], lst[n-i-1] = lst[n-i-1], lst[i]
	}
	return lst
}

func CountItemInterface(lst []interface{}) map[interface{}]int {
	frequency := make(map[interface{}]int)
	for _, v := range lst {
		frequency[v]++
	}
	return frequency
}

func DEDuplicateItemInterface(input []interface{}) []interface{} {
	output := make([]interface{}, 0)
	inputSet := make(map[interface{}]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)
		}
		inputSet[v] = struct{}{}
	}
	return output
}
