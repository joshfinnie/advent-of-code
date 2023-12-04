package utils

import "strconv"

func ConvertStringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return result
}

func IntInSlice(i int, s []int) bool {
	for _, a := range s {
        if a == i {
            return true
        }
    }
    return false
}
