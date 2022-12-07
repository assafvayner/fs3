package service

import (
	"context"
	"errors"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (handler *AuthServiceHandler) NewUser(ctx context.Context, req *authservice.NewUserRequest) (*authservice.NewUserReply, error) {
	handler.VerifyRedisClient()

	reply := &authservice.NewUserReply{
		Username: req.GetUsername(),
	}

	usernameKey := GetKeyFromUsername(req.GetUsername())

	userExistsCopies, err := handler.RedisClient.Exists(context.Background(), usernameKey).Result()
	if err != nil {
		handler.Logger.Printf("NewUser: Failed to check if user %s exists already; err: %s\n", req.GetUsername(), err)
		return internalError(reply)
	}
	if userExistsCopies != 0 {
		reply.Status = &authservice.NewUserReply_Status{
			Success: false,
			Message: "user already exists",
		}
		return reply, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		handler.Logger.Printf("NewUser: Failed to hash password for new user %s exists already; err: %s\n", req.GetUsername(), err)
		return internalError(reply)
	}

	setResult := handler.RedisClient.Set(context.Background(), usernameKey, string(passwordHash), 0)
	if setResult.Err() != nil {
		handler.Logger.Printf("NewUser: Failed to insert user: %s into db\n", req.GetUsername())
		return internalError(reply)
	}

	handler.Logger.Printf("NewUser: Succeeded to add new user: %s\n", req.GetUsername())
	reply.Status = &authservice.NewUserReply_Status{
		Success: true,
		Message: "OK",
	}
	return reply, nil
}

func internalError(reply *authservice.NewUserReply) (*authservice.NewUserReply, error) {
	reply.Status = &authservice.NewUserReply_Status{
		Success: false,
		Message: "internal error",
	}
	return reply, errors.New("internal error")
}
