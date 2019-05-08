package carousel

import (
	"fmt"
	"testing"
)

func TestGet(T *testing.T) {
	c, err := get("sss", "sss")
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	if c.UUID == "" {
		fmt.Println("no data ")
		return
	}
	fmt.Println("success ", c.UUID)
}
