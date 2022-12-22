package interceptor

import (
	"context"
	"errors"
	"log"

	"github.com/assafvayner/fs3/server/shared/jwtutils"
	"google.golang.org/grpc"
)

type HasToken interface {
	GetToken() string
}

func GetAuthInterceptor(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqCasted, ok := req.(HasToken)
		if !ok {
			logger.Printf("cannot get token from request")
			return nil, errors.New("misunderstood interface of request")
		}
		tokenString := reqCasted.GetToken()
		if tokenString == "" {
			// global access case
			return handler(ctx, req)
		}

		err := validateToken(reqCasted.GetToken(), logger)
		if err != nil {
			return nil, err
		}
		// request is valid
		logger.Println("request with token validated")
		return handler(ctx, req)
	}
}

func validateToken(tokenString string, logger *log.Logger) error {
	token, err := jwtutils.ParseToken(tokenString, true)
	if err != nil {
		logger.Printf("Primary;Copy; Failed to parse token: %s, error: %s\n", tokenString, err)
		return errors.New("error parsing token")
	}
	claims, err := jwtutils.GetClaims(token)
	if err != nil {
		logger.Printf(
			"Primary;Copy; Failed to get valid claims on token: %s, err: %s\n",
			tokenString,
			err,
		)
		return errors.New("error getting claims from token")
	}
	// should check request is within the right amount of time too?
	if err = claims.Valid(); err != nil {
		logger.Printf(
			"Primary;Copy; Token claims not valid, token: %s, err: %s\n",
			tokenString,
			err,
		)
		return errors.New("error getting claims from token")
	}
	return nil
}
