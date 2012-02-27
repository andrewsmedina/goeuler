package main

import (
    "testing"
)

func TestString(t *testing.T) {
	if sumOfMultiplesOf3And5Bellow(10) != 23 {
        t.Errorf("Assertion Error")
	}
}
