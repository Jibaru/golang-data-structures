package tests

import (
	"testing"
)

func printError(funcName string, t *testing.T, expected, got any) {
	t.Fatalf(`%v: Expected %v, got %v`, funcName, expected, got)
}
