package grpc

import (
	"context"

	userpb "github.com/Vishnevyy/project-protos/proto/user"
	"github.com/Vishnevyy/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(s *user.Service) *Handler { return &Handler{svc: s} }

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.Create(req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{User: &userpb.User{Id: u.ID, Email: u.Email}}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	u, err := h.svc.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: u.ID, Email: u.Email}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	us, err := h.svc.List()
	if err != nil {
		return nil, err
	}
	items := make([]*userpb.User, 0, len(us))
	for _, it := range us {
		items = append(items, &userpb.User{Id: it.ID, Email: it.Email})
	}
	return &userpb.ListUsersResponse{Items: items}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.User, error) {
	u, err := h.svc.Update(req.Id, req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: u.ID, Email: u.Email}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.svc.Delete(req.Id); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Ok: true}, nil
}
