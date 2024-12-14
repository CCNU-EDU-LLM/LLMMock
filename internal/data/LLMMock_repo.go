package data

import (
	"LLMMock/internal/biz"
	"LLMMock/internal/middleware/auth"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type LLMMock struct {
	data *Data
	log  *log.Helper
}

func NewLLMMockRepo(data *Data, logger log.Logger) biz.LLMMockRepo {
	return &LLMMock{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (L LLMMock) TestConnect(ctx context.Context) (*biz.TestConnect, error) {
	return &biz.TestConnect{
		Message: "Welcome to the Chat API! Use POST /chat to interact.",
	}, nil
}

func (L LLMMock) SendMessage(ctx context.Context, msg *biz.SendMessage) (*biz.SessionMessage, *biz.SessionMessage, error) {
	return &biz.SessionMessage{
			Role:    "user",
			Content: msg.Message,
		}, &biz.SessionMessage{
			Role:    msg.ModelName,
			Content: "Message, I am a robot",
		}, nil
}

func (L LLMMock) GetHistory(ctx context.Context, sessionId string) (*biz.GetHistoryResponse, error) {
	return &biz.GetHistoryResponse{
		UserMessages: []biz.SessionMessage{
			{
				Role:    "user",
				Content: "Message, I am a user",
			},
			{
				Role:    "user",
				Content: "Message, I am a user again",
			},
		},
		BotMessages: []biz.SessionMessage{
			{
				Role:    "robot",
				Content: "Message, I am a robot",
			},
			{
				Role:    "robot",
				Content: "Message, I am a robot again",
			},
		},
	}, nil
}

func (L LLMMock) DeleteHistory(ctx context.Context, sessionId string) (*biz.DeleteHistoryResponse, error) {
	claims, ok := ctx.Value("claims").(*auth.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("claims not found or invalid type")
	}

	// middleware/auth/auth.go
	// 从 context 中获取 CustomClaims
	// type CustomClaims struct {
	// 		Email                string `json:"email"`
	//	    jwt.RegisteredClaims
	// }

	// 使用 CustomClaims 的数据
	m := map[string]string{
		"email": claims.Email,
		"iss":   claims.Issuer,
	}

	return &biz.DeleteHistoryResponse{
		Message: "user: " + m["email"] + "sessionId: " + sessionId + " deleted successfully",
	}, nil
}
