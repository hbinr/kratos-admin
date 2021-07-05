// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc7

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

type UserHTTPServer interface {
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserReply, error)
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	ListUser(context.Context, *ListUserReq) (*ListUserReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _User_Register0_HTTP_Handler(srv))
	r.POST("/v1/login", _User_Login0_HTTP_Handler(srv))
	r.PUT("/v1/user", _User_UpdateUser0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/v1/users", _User_ListUser0_HTTP_Handler(srv))
	r.DELETE("/v1/user/{user_id}", _User_DeleteUser0_HTTP_Handler(srv))
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/ListUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.service.v1.User/DeleteUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	DeleteUser(ctx context.Context, req *DeleteUserReq, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListUser(ctx context.Context, req *ListUserReq, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserReq, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/v1/user/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.service.v1.User/DeleteUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/v1/user/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.service.v1.User/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserReq, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.service.v1.User/ListUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.service.v1.User/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.service.v1.User/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.service.v1.User/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
