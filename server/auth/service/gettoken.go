package service

import (
	"context"
	"errors"
	"fmt"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"golang.org/x/crypto/bcrypt"
)

func (handler *AuthServiceHandler) GetToken(ctx context.Context, req *authservice.GetNewTokenRequest) (*authservice.GetNewTokenReply, error) {
	handler.VerifyRedisClient()
	reply := &authservice.GetNewTokenReply{
		Username: req.GetUsername(),
	}

	usernameKey := GetKeyFromUsername(req.GetUsername())

	passwordHash, err := handler.RedisClient.Get(context.Background(), usernameKey).Result()
	if err != nil {
		handler.Logger.Printf("GetToken: Error from redis for username: %s, err: %s\n", req.Username, err)
		reply.Status = &authservice.GetNewTokenReply_Status{
			Success: false,
			Message: "internal error",
		}
		return reply, errors.New("internal error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.GetPassword()))
	if err != nil {
		handler.Logger.Printf("GetToken: Error from bcrypt compare password for username: %s, err: %s\n", req.Username, err)
		reply.Status = &authservice.GetNewTokenReply_Status{
			Success: false,
			Message: fmt.Sprintf("could not authorize user %s", req.GetUsername()),
		}
		return reply, errors.New("Authorization failed")
	}

	token, err := GetToken(req.GetUsername())
	if err != nil {
		handler.Logger.Printf("GetToken: failed to get token for authorized user: %s, err: %s\n", req.Username, err)
		reply.Status = &authservice.GetNewTokenReply_Status{
			Success: false,
			Message: "internal error",
		}
		return reply, errors.New("internal error")
	}

	handler.Logger.Printf("GetToken: Successfully produced token for user: %s\n", req.Username)
	reply.Status = &authservice.GetNewTokenReply_Status{
		Success: true,
		Message: "OK",
	}
	reply.Token = token
	return reply, nil
}
