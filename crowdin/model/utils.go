package model

import (
	"fmt"
	"strings"
)

// JoinIntSlice is a helper function that joins the elements
// of the int slice into a single string.
func JoinIntSlice(s []int) string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(res, ",")
}
