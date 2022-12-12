package service

import (
	"context"
	"errors"
	"fmt"

	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/jwtutils"
	"golang.org/x/crypto/bcrypt"
)

func (handler *AuthServiceHandler) GetToken(ctx context.Context, req *authservice.GetNewTokenRequest) (*authservice.GetNewTokenReply, error) {
	reply := &authservice.GetNewTokenReply{
		Username: req.GetUsername(),
	}

	if err := handler.authenticateUser(req.GetUsername(), req.GetPassword(), req.GetPreviousToken()); err != nil {
		reply.Status = &authservice.GetNewTokenReply_Status{
			Success: false,
			Message: fmt.Sprintf("could not authorize user %s", req.GetUsername()),
		}
		return reply, fmt.Errorf("Authorization failed: %s", err)
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

func (handler *AuthServiceHandler) authenticateUser(username, password, prevToken string) error {
	if prevToken != "" {
		// user previous token validity to authorize user
		usernameFromToken, err := jwtutils.GetUsernameFromToken(prevToken)
		if err == nil && usernameFromToken == username {
			handler.Logger.Printf("authorized user %s given previous token %s\n", username, prevToken)
			return nil
		}
	}

	handler.VerifyRedisClient()
	usernameKey := GetKeyFromUsername(username)

	passwordHash, err := handler.RedisClient.Get(context.Background(), usernameKey).Result()
	if err != nil {
		handler.Logger.Printf("GetToken: Error from redis for username: %s, err: %s\n", username, err)
		return errors.New("internal error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		handler.Logger.Printf("GetToken: Error from bcrypt compare password for username: %s, err: %s\n", username, err)
		return errors.New("password does not match")
	}

	return nil
}
