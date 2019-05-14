package jwt

import (
	"fmt"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

// Payload jwt.Payload
type Payload struct {
	jwt.Payload
	Issuer         string   // 发布者是谁
	Subject        string   // 该JWT所面向的用户，用于处理特定应用，不是常用的字段
	Audience       []string // 受众 []string //jwt.Audience{"https://minis.app"},
	ExpirationTime int64    // 过期时间 now.Add(30 * time.Minute).Unix()
	NotBefore      int64    // 生效时间 now.Add(30 * time.Minute).Unix(),
	IssuedAt       int64    // 签发时间 now.Unix(),
	JWTID          string   // jwt id 编号   // user_id 认证 (加密)
	EnInfo         EnInfo   // 加密信息
}

// EnInfo ... 加密的信息
type EnInfo struct {
	UserID string `json:"user_id,omitempty"` // 用户ID
	IP     string `json:"ip,omitempty"`      // 用户IP
}

// NewToken JWT 发放token
func NewToken(pa Payload, secret string) (token []byte, err error) {
	hs256 := jwt.NewHMAC(jwt.SHA256, []byte(secret))
	h := jwt.Header{Algorithm: "hs256"}
	p := Payload{
		Payload: jwt.Payload{
			Issuer:         pa.Issuer,
			Subject:        pa.Subject,
			Audience:       pa.Audience,
			ExpirationTime: pa.ExpirationTime,
			NotBefore:      pa.NotBefore, // now.Add(30 * time.Minute).Unix(),
			IssuedAt:       pa.IssuedAt,
			JWTID:          pa.JWTID,
		},
		EnInfo: pa.EnInfo,
	}
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
	audValidator := jwt.AudienceValidator(audience)
	err = p.Validate(iatValidator, expValidator, audValidator)
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
