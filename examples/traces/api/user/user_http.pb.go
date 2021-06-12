// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc3

package v1

import (
	context "context"
	transport "github.com/go-kratos/kratos/v2/transport"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(transport.Transporter)
var _ = binding.EncodeVars

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/user/get/message/{count}", _User_GetMyMessages0_HTTP_Handler(srv))
}

func _User_GetMyMessages0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyMessagesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/api.user.v1.User/GetMyMessages")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyMessages(ctx, req.(*GetMyMessagesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyMessagesReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	GetMyMessages(ctx context.Context, req *GetMyMessagesRequest, opts ...http.CallOption) (rsp *GetMyMessagesReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetMyMessages(ctx context.Context, in *GetMyMessagesRequest, opts ...http.CallOption) (*GetMyMessagesReply, error) {
	var out GetMyMessagesReply
	path := binding.EncodeVars("/v1/user/get/message/{count}", in, true)
	opts = append(opts, http.Operation("/api.user.v1.User/GetMyMessages"))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}