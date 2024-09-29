package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nishchay-veer/simplebank/token"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizarionPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2{
			err := errors.New("invalid authorization header format ")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("invalid authorization type")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return 
		}

		accessToken := fields[1]
		payload , err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}

		c.Set(authorizarionPayloadKey, payload)
		c.Next()
	}
}