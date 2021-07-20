package main

import (
	"fmt"
	"testing"
)

func TestStringFunc_ToString(t *testing.T) {

	stringFunc := StringFunc{
		MsgName: "user",
	}

	toString := stringFunc.ToString()

	fmt.Println(toString)

}
