package middleware

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/token"
	authusecase "contracts-manager/internal/usecases/auth"
	"contracts-manager/internal/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtProvider *token.JWTProvider
	authUC      *authusecase.Usecase
}

func NewAuthMiddleware(
	jwtProvider *token.JWTProvider,
	authUC *authusecase.Usecase,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtProvider,
		authUC,
	}
}

func (md *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := context.GetAccessToken(c)
		if err != nil {
			context.RespondError(c, http.StatusUnauthorized, err)
			return
		}

		userID, err := md.jwtProvider.ParseUserID(tokenStr)
		if err != nil {
			context.RespondError(c, http.StatusUnauthorized, err)
			return
		}

		authUser, err := md.authUC.GetAuthUser(c.Request.Context(), userID)
		if err != nil {
			context.RespondError(c, http.StatusUnauthorized, err)
		}

		context.SetContextValue(c, context.AuthUser, authUser)

		c.Next()
	}
}

func (md *AuthMiddleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		md.Middleware()(c)
		if c.IsAborted() {
			return
		}

		authUser, err := context.GetAuthUser(c)
		if err != nil {
			context.RespondError(c, http.StatusUnauthorized, err)
			return
		}

		if authUser.Type != models.UserTypeAdmin {
			context.RespondError(c, http.StatusForbidden, auth.ErrRoleNotAllowed)
			return
		}

		c.Next()
	}
}
