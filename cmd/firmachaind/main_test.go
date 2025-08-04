package main

import (
	"fmt"
	"testing"
)

func TestNewRootCmd(t *testing.T) {

	_ = NewRootCmd()

	fmt.Println("Ok")
}
