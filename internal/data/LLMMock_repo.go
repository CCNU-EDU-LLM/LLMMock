package data

import (
	"LLMMock/internal/biz"
	"context"

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
	return &biz.DeleteHistoryResponse{
		Message: "History successful deleted",
	}, nil
}
