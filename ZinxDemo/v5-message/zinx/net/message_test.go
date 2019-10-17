package net

import (
	"fmt"
	"testing"
)

func TestNewMessage(t *testing.T) {
	data := []byte("hello")
	fmt.Println(NewMessage(5,0,data))

}
