package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)


const TokenExpireDuration = time.Hour * 2   // JWT的过期时间

var mySecret = []byte("夏天夏天悄悄过去")


// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成JWT
func GenToken(userID int64, username, password string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID: userID,
		Username: username,
		Password: password,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "bluebell", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

// parseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	var mc = MyClaims{}
	token , err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return &mc, nil
	}
	return nil, errors.New("invalid token")
}





