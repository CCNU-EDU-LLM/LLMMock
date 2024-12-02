package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	Email                string `json:"email"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于签名的字符串
var CustomSecret = []byte("CCNU-EDU-LLM")

// GenToken 生成JWT
func GenToken(email string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		email, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "CCNU", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

func CheckJWT(tokenString string) (*CustomClaims, error) {
	// 使用 CustomClaims 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("CCNU-EDU-LLM"), nil
	})
	if err != nil {
		return nil, err
	}

	// 检查解析后的 Claims 是否为 CustomClaims 类型
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func JWTAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				if tokenString == "" {
					return nil, errors.New("authorization header missing")
				}

				// 验证 token
				claims, err := CheckJWT(tokenString)
				if err != nil {
					return nil, errors.New("invalid or expired token")
				}

				// 将 claims 添加到上下文中
				ctx = context.WithValue(ctx, "claims", claims)
			}

			// 调用下一个 handler
			return handler(ctx, req)
		}
	}
}
