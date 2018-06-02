package util

import (
	"testing"
	"fmt"
)

var length = 24
var num = 10

func TestRandId(t *testing.T) {
	for i := 0; i < num; i++ {
		fmt.Println(RandId(length))
	}
}

func TestSecureRandId(t *testing.T) {
	for i := 0; i < num; i++ {
		fmt.Println(SecureRandId(length))
	}
}

func TestSecureRandIdOrPanic(t *testing.T) {
	for i := 0; i < num; i++ {
		fmt.Println(SecureRandIdOrPanic(length))
	}
}

func TestRandNumId(t *testing.T) {
	for i := 0; i < num; i++ {
		fmt.Println(RandNumId(length))
	}
}