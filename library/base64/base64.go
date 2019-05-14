package base64

import "encoding/base64"

// EncodeToString 编码 base64
func EncodeToString(str string) (encodeString string) {
	input := []byte(str)
	encodeString = base64.StdEncoding.EncodeToString(input)
	return
}

// DecodeToString 解码 base64
func DecodeToString(str string) (decodeString string, err error) {
	encodingByte, err := base64.StdEncoding.DecodeString(str)
	decodeString = string(encodingByte)
	return
}
