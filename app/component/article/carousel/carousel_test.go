package carousel

import (
	"fmt"
	"testing"
)

func TestGetCarousel(T *testing.T) {
	c, err := get("sss", "sss")
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	fmt.Println("success ", c)
}
