// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bcsproject.proto

package bcsproject

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BCSProject service

func NewBCSProjectEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "BCSProject.CreateProject",
			Path:    []string{"/bcsproject/v1/projects"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BCSProject.GetProject",
			Path:    []string{"/bcsproject/v1/projects/{projectIDOrCode}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BCSProject.UpdateProject",
			Path:    []string{"/bcsproject/v1/projects/{projectID}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BCSProject.DeleteProject",
			Path:    []string{"/bcsproject/v1/projects/{projectID}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BCSProject.ListProjects",
			Path:    []string{"/bcsproject/v1/projects"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for BCSProject service

type BCSProjectService interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*ProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...client.CallOption) (*ProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*ProjectResponse, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...client.CallOption) (*ProjectResponse, error)
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error)
}

type bCSProjectService struct {
	c    client.Client
	name string
}

func NewBCSProjectService(name string, c client.Client) BCSProjectService {
	return &bCSProjectService{
		c:    c,
		name: name,
	}
}

func (c *bCSProjectService) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*ProjectResponse, error) {
	req := c.c.NewRequest(c.name, "BCSProject.CreateProject", in)
	out := new(ProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bCSProjectService) GetProject(ctx context.Context, in *GetProjectRequest, opts ...client.CallOption) (*ProjectResponse, error) {
	req := c.c.NewRequest(c.name, "BCSProject.GetProject", in)
	out := new(ProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bCSProjectService) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*ProjectResponse, error) {
	req := c.c.NewRequest(c.name, "BCSProject.UpdateProject", in)
	out := new(ProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bCSProjectService) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...client.CallOption) (*ProjectResponse, error) {
	req := c.c.NewRequest(c.name, "BCSProject.DeleteProject", in)
	out := new(ProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bCSProjectService) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error) {
	req := c.c.NewRequest(c.name, "BCSProject.ListProjects", in)
	out := new(ListProjectsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BCSProject service

type BCSProjectHandler interface {
	CreateProject(context.Context, *CreateProjectRequest, *ProjectResponse) error
	GetProject(context.Context, *GetProjectRequest, *ProjectResponse) error
	UpdateProject(context.Context, *UpdateProjectRequest, *ProjectResponse) error
	DeleteProject(context.Context, *DeleteProjectRequest, *ProjectResponse) error
	ListProjects(context.Context, *ListProjectsRequest, *ListProjectsResponse) error
}

func RegisterBCSProjectHandler(s server.Server, hdlr BCSProjectHandler, opts ...server.HandlerOption) error {
	type bCSProject interface {
		CreateProject(ctx context.Context, in *CreateProjectRequest, out *ProjectResponse) error
		GetProject(ctx context.Context, in *GetProjectRequest, out *ProjectResponse) error
		UpdateProject(ctx context.Context, in *UpdateProjectRequest, out *ProjectResponse) error
		DeleteProject(ctx context.Context, in *DeleteProjectRequest, out *ProjectResponse) error
		ListProjects(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error
	}
	type BCSProject struct {
		bCSProject
	}
	h := &bCSProjectHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BCSProject.CreateProject",
		Path:    []string{"/bcsproject/v1/projects"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BCSProject.GetProject",
		Path:    []string{"/bcsproject/v1/projects/{projectIDOrCode}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BCSProject.UpdateProject",
		Path:    []string{"/bcsproject/v1/projects/{projectID}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BCSProject.DeleteProject",
		Path:    []string{"/bcsproject/v1/projects/{projectID}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BCSProject.ListProjects",
		Path:    []string{"/bcsproject/v1/projects"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&BCSProject{h}, opts...))
}

type bCSProjectHandler struct {
	BCSProjectHandler
}

func (h *bCSProjectHandler) CreateProject(ctx context.Context, in *CreateProjectRequest, out *ProjectResponse) error {
	return h.BCSProjectHandler.CreateProject(ctx, in, out)
}

func (h *bCSProjectHandler) GetProject(ctx context.Context, in *GetProjectRequest, out *ProjectResponse) error {
	return h.BCSProjectHandler.GetProject(ctx, in, out)
}

func (h *bCSProjectHandler) UpdateProject(ctx context.Context, in *UpdateProjectRequest, out *ProjectResponse) error {
	return h.BCSProjectHandler.UpdateProject(ctx, in, out)
}

func (h *bCSProjectHandler) DeleteProject(ctx context.Context, in *DeleteProjectRequest, out *ProjectResponse) error {
	return h.BCSProjectHandler.DeleteProject(ctx, in, out)
}

func (h *bCSProjectHandler) ListProjects(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error {
	return h.BCSProjectHandler.ListProjects(ctx, in, out)
}