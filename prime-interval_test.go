package main

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Printf("hei")
	expected := true
	actual := isPrime(10)

	if actual != expected {
		t.Errorf("failed")
	}
}
