package service

import (
	v1 "LLMMock/api/LLMMock/v1"
	"LLMMock/internal/biz"
	"context"
)

// LLMMockService is a greeter service.
type LLMMockService struct {
	v1.UnimplementedLLMMockServer

	uc *biz.LLMMockUsecase
}

// NewGreeterService new a greeter service.
func NewLLMMockService(uc *biz.LLMMockUsecase) *LLMMockService {
	return &LLMMockService{uc: uc}
}

// SayHello implements LLMMock.GreeterServer.
func (s *LLMMockService) SayHello(ctx context.Context, in *v1.RootRequest) (*v1.RootReply, error) {
	u, err := s.uc.TestConnect(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.RootReply{
		Message: u.Message,
	}, nil
}

func (s *LLMMockService) SendMessage(ctx context.Context, in *v1.SendMessageRequest) (*v1.SendMessageReply, error) {
	u, r, err := s.uc.SendMessage(ctx, &biz.SendMessage{
		SessionID: in.SessionId,
		Message:   in.Message,
		ModelName: in.ModelName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SendMessageReply{
		Response: r.Content,
		SessionMessages: &v1.SessionMessage{
			Role:    u.Role,
			Content: u.Content,
		},
		BotMessages: &v1.SessionMessage{
			Role:    r.Role,
			Content: r.Content,
		},
	}, nil
}

func (s *LLMMockService) GetHistory(ctx context.Context, in *v1.GetHistoryRequest) (*v1.GetHistoryReply, error) {
	u, err := s.uc.GetHistory(ctx, in.SessionId)
	if err != nil {
		return nil, err
	}
	// 转换 UserMessages
	userMessages := make([]*v1.SessionMessage, len(u.UserMessages))
	for i, msg := range u.UserMessages {
		userMessages[i] = &v1.SessionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	// 转换 BotMessages
	botMessages := make([]*v1.SessionMessage, len(u.BotMessages))
	for i, msg := range u.BotMessages {
		botMessages[i] = &v1.SessionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	return &v1.GetHistoryReply{
		SessionMessages: userMessages,
		BotMessages:     botMessages,
	}, nil
}

func (s *LLMMockService) DeleteHistory(ctx context.Context, in *v1.DeleteHistoryRequest) (*v1.DeleteHistoryReply, error) {
	u, err := s.uc.DeleteHistory(ctx, in.SessionId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteHistoryReply{
		Message: u.Message,
	}, nil
}
