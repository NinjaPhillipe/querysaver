package utils

import (
	"strconv"
	"strings"
)

const DEFAULT_ARRAY_SIZE = 1000

func ConcatArrayInt(nbs []int, sep string) string {

	strNumbers := make([]string, len(nbs))

	for i, num := range nbs {
		strNumbers[i] = strconv.Itoa(num)
	}
	return strings.Join(strNumbers, sep)
}
