// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.6
// source: LLMMock/v1/LLM.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLLMMockDeleteHistory = "/LLMMock.v1.LLMMock/DeleteHistory"
const OperationLLMMockGetHistory = "/LLMMock.v1.LLMMock/GetHistory"
const OperationLLMMockSayHello = "/LLMMock.v1.LLMMock/SayHello"
const OperationLLMMockSendMessage = "/LLMMock.v1.LLMMock/SendMessage"

type LLMMockHTTPServer interface {
	DeleteHistory(context.Context, *DeleteHistoryRequest) (*DeleteHistoryReply, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryReply, error)
	// SayHello根目录，验证服务是否正常运行，无需参数
	SayHello(context.Context, *RootRequest) (*RootReply, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageReply, error)
}

func RegisterLLMMockHTTPServer(s *http.Server, srv LLMMockHTTPServer) {
	r := s.Route("/")
	r.GET("/", _LLMMock_SayHello0_HTTP_Handler(srv))
	r.POST("/chat", _LLMMock_SendMessage0_HTTP_Handler(srv))
	r.GET("/history", _LLMMock_GetHistory0_HTTP_Handler(srv))
	r.POST("/clear_history", _LLMMock_DeleteHistory0_HTTP_Handler(srv))
}

func _LLMMock_SayHello0_HTTP_Handler(srv LLMMockHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RootRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLLMMockSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*RootRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RootReply)
		return ctx.Result(200, reply)
	}
}

func _LLMMock_SendMessage0_HTTP_Handler(srv LLMMockHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLLMMockSendMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMessage(ctx, req.(*SendMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageReply)
		return ctx.Result(200, reply)
	}
}

func _LLMMock_GetHistory0_HTTP_Handler(srv LLMMockHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHistoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLLMMockGetHistory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHistory(ctx, req.(*GetHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHistoryReply)
		return ctx.Result(200, reply)
	}
}

func _LLMMock_DeleteHistory0_HTTP_Handler(srv LLMMockHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteHistoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLLMMockDeleteHistory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteHistory(ctx, req.(*DeleteHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteHistoryReply)
		return ctx.Result(200, reply)
	}
}

type LLMMockHTTPClient interface {
	DeleteHistory(ctx context.Context, req *DeleteHistoryRequest, opts ...http.CallOption) (rsp *DeleteHistoryReply, err error)
	GetHistory(ctx context.Context, req *GetHistoryRequest, opts ...http.CallOption) (rsp *GetHistoryReply, err error)
	SayHello(ctx context.Context, req *RootRequest, opts ...http.CallOption) (rsp *RootReply, err error)
	SendMessage(ctx context.Context, req *SendMessageRequest, opts ...http.CallOption) (rsp *SendMessageReply, err error)
}

type LLMMockHTTPClientImpl struct {
	cc *http.Client
}

func NewLLMMockHTTPClient(client *http.Client) LLMMockHTTPClient {
	return &LLMMockHTTPClientImpl{client}
}

func (c *LLMMockHTTPClientImpl) DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...http.CallOption) (*DeleteHistoryReply, error) {
	var out DeleteHistoryReply
	pattern := "/clear_history"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLLMMockDeleteHistory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LLMMockHTTPClientImpl) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...http.CallOption) (*GetHistoryReply, error) {
	var out GetHistoryReply
	pattern := "/history"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLLMMockGetHistory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LLMMockHTTPClientImpl) SayHello(ctx context.Context, in *RootRequest, opts ...http.CallOption) (*RootReply, error) {
	var out RootReply
	pattern := "/"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLLMMockSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LLMMockHTTPClientImpl) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...http.CallOption) (*SendMessageReply, error) {
	var out SendMessageReply
	pattern := "/chat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLLMMockSendMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
