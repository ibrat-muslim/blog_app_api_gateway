package v1

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog_app_api_gateway/api/models"
	"github.com/ibrat-muslim/blog_app_api_gateway/config"
	grpcPkg "github.com/ibrat-muslim/blog_app_api_gateway/pkg/grpc_client"
	"github.com/sirupsen/logrus"
)

var (
	ErrWrongEmailOrPass = errors.New("wrong email or password")
	ErrUserNotVerified  = errors.New("user not verified")
	ErrEmailExists      = errors.New("email already exists")
	ErrIncorrectCode    = errors.New("incorrect verification code")
	ErrCodeExpired      = errors.New("verification code has been expired")
	ErrNotAllowed       = errors.New("method not allowed")
	ErrWeakPassword     = errors.New("password must contain at least one small letter, one capital letter, one number and one symbol")
)

type handlerV1 struct {
	cfg        *config.Config
	grpcClient grpcPkg.GrpcClientI
	logger     *logrus.Logger
}

type HandlerV1Options struct {
	Cfg        *config.Config
	GrpcClient grpcPkg.GrpcClientI
	Logger     *logrus.Logger
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:        options.Cfg,
		grpcClient: options.GrpcClient,
		logger:     options.Logger,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}

func validateGetAllParamsRequest(ctx *gin.Context) (*models.GetAllParamsRequest, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)

	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllParamsRequest{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: ctx.Query("search"),
	}, nil
}
