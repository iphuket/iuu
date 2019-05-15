package jwt

import (
	"fmt"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

// Payload jwt.Payload
type Payload struct {
	jwt.Payload
	EnInfo EnInfo // 加密信息
}

// EnInfo ... 加密的信息
type EnInfo struct {
	UserID string `json:"user_id,omitempty"` // 用户ID
	IP     string `json:"ip,omitempty"`      // 用户IP
}

// NewToken JWT 发放token
func NewToken(p *Payload, secret string) (token []byte, err error) {
	hs256 := jwt.NewHMAC(jwt.SHA256, []byte(secret))
	h := jwt.Header{Algorithm: "hs256"}
	token, err = jwt.Sign(h, p, hs256)
	return
}

// Chcek jwt
func Chcek(secret, token string, audience ...string) (eninfo EnInfo, err error) {
	now := time.Now()
	eninfo = EnInfo{}

	hs256 := jwt.NewHMAC(jwt.SHA256, []byte(secret))
	tokenToByte := []byte(token)
	raw, err := jwt.Parse(tokenToByte)
	if err != nil {
		return eninfo, err
	}
	err = raw.Verify(hs256)
	if err != nil {
		return eninfo, err
	}
	var (
		_ jwt.Header
		p Payload
	)
	_, err = raw.Decode(&p)
	if err != nil {
		return eninfo, err
	}
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now, true)
	// audValidator := jwt.AudienceValidator(audience)
	err = p.Validate(iatValidator, expValidator)
	if err != nil {
		return eninfo, err
	}
	eninfo = p.EnInfo
	return eninfo, err
}

// IPChcek 检查IP是否变化
func IPChcek(nowip, lastip string) error {
	if nowip != lastip {
		return fmt.Errorf("ip validation error")
	}
	return nil
}
