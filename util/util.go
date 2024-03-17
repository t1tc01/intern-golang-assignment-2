package util

import (
	"context"
	"strings"
)

var Ctx = context.Background()

func ConvertMapToString(myMap map[string][]string) string {
	var sb strings.Builder
	sb.WriteString("{")
	for key, values := range myMap {
		sb.WriteString(key + "=")
		sb.WriteString(strings.Join(values, ","))
		sb.WriteString(", ")
	}
	sb.WriteString("}")
	return sb.String()
}
