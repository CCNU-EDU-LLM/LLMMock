package biz

import (
	"context"

	v1 "LLMMock/api/LLMMock/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type TestConnect struct {
	Message string `json:"message"`
}

type SendMessage struct {
	SessionID string `json:"session_id"` // 会话ID
	Message   string `json:"message"`    // 消息内容
	ModelName string `json:"model_name"` // 模型名称
}

type SessionMessage struct {
	Role    string `json:"role"`    // user or bot
	Content string `json:"content"` // 消息内容
}

type GetHistoryResponse struct {
	UserMessages []SessionMessage `json:"user_messages"` // 用户消息
	BotMessages  []SessionMessage `json:"bot_messages"`  // 机器人消息
}

// LLMMockRepo is a repo.
type LLMMockRepo interface {
	TestConnect(ctx context.Context) (*TestConnect, error)
	SendMessage(ctx context.Context, msg *SendMessage) (*SessionMessage, *SessionMessage, error)
	GetHistory(ctx context.Context, sessionID string) (*GetHistoryResponse, error)
	DeleteHistory(ctx context.Context, sessionID string) (*DeleteHistoryResponse, error)
}

type DeleteHistoryResponse struct {
	Message string `json:"message"` // 删除结果消息
}

// LLMMockUsecase is a Greeter usecase.
type LLMMockUsecase struct {
	repo LLMMockRepo
	log  *log.Helper
}

func NewLLMMockUsecase(repo LLMMockRepo, logger log.Logger) *LLMMockUsecase {
	return &LLMMockUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LLMMockUsecase) TestConnect(ctx context.Context) (*TestConnect, error) {
	return uc.repo.TestConnect(ctx)
}

func (uc *LLMMockUsecase) SendMessage(ctx context.Context, msg *SendMessage) (*SessionMessage, *SessionMessage, error) {
	return uc.repo.SendMessage(ctx, msg)
}

func (uc *LLMMockUsecase) GetHistory(ctx context.Context, sessionID string) (*GetHistoryResponse, error) {
	return uc.repo.GetHistory(ctx, sessionID)
}

func (uc *LLMMockUsecase) DeleteHistory(ctx context.Context, sessionID string) (*DeleteHistoryResponse, error) {
	return uc.repo.DeleteHistory(ctx, sessionID)
}
