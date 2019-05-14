package base64

import (
	"fmt"
	"testing"
)

func TestBase64(T *testing.T) {
	encodestring := EncodeToString("犇猋骉麤毳淼掱焱垚赑燚翯寗鎏")
	fmt.Println("base64 编码: ", encodestring)
	decodeString, err := DecodeToString(encodestring)
	fmt.Println("base64 解码: ", decodeString, "错误信息:", err)
}
