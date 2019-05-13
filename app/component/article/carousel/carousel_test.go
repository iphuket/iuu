package carousel

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/iphuket/pkt/app/component/article/model"
)

func TestGet(T *testing.T) {
	c, err := get("class_uuid", "case_uuid")
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	if c == nil {
		fmt.Println("no data ")
		return
	}
	fmt.Println("success ", c)
}
func TestPut(T *testing.T) {
	carousel := new(model.Carousel)
	carousel.UUID = uuid.New().String() // c.Request.FormValue("c")
	carousel.CaseUUID = "case_uuid"
	carousel.ClassUUID = "class_uuid"
	carousel.Name = "test"
	carousel.Desc = "desc"
	carousel.Source = "ssss"
	carousel.Picture = "sadasdasda"
	carousel.UserUUID = "sadasda"
	c, err := put(carousel)
	fmt.Println(c, err)
}
func TestDelete(T *testing.T) {
	fmt.Println(delete("017dfcb1-f73e-4c91-b47a-73093f135f28"))
}
