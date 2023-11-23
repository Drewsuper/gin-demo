package middleware

import (
	"errors"
	"gin-new/app/config"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}

func ParseToken(token string) (*MyClaims, error) {
	mySecret := []byte(config.Config.JWT.Secret)
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	} else {
		claims, errOne := jwtToken.Claims.(*MyClaims)
		if errOne && jwtToken.Valid {
			return claims, nil
		}
		return nil, errors.New("token无法解析")
	}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Abort()
			ctx.JSON(400, types.CommonRps{
				Code: 400,
				Mes:  "no auth request",
			})
			return
		}
		_, err := ParseToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.JSON(400, types.CommonRps{
				Code: 400,
				Mes:  "no auth request",
			})
			return
		}
		ctx.Next()
	}
}

func NewToken(data any) (string, error) {
	claims := MyClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}
