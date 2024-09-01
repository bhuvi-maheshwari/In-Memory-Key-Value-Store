package service

import (
	"fmt"
	"strings"
)

type Pair struct {
	Key   string
	Value interface{}
}

type Value []Pair

// Prints the value of the key
func (v Value) String() string {
	var sb strings.Builder
	for _, val := range v {
		s := fmt.Sprintf("%v: %v, ", val.Key, val.Value)
		sb.WriteString(s)
	}
	sb.WriteString("\n")
	return sb.String()
}
