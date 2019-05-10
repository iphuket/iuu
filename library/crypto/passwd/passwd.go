package passwd

import (
	"github.com/iphuket/iuu/library/crypto"
)

// New passwd
func New(x, y string) (passwd string, err error) {
	z := x + "x+y" + y
	passwd, err = crypto.Md5Encrypt(z)
	return passwd, err
}
