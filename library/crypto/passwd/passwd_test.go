package passwd

import (
	"fmt"
	"testing"
)

func TestNew(T *testing.T) {
	passwd, err := New("84a4s654d65as", ")((*)(m89hj987")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("passwd: ", passwd)
}
