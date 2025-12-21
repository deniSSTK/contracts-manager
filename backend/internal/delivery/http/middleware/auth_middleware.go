package middleware

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/token"
	authusecase "contracts-manager/internal/usecases/auth"
	contractusecase "contracts-manager/internal/usecases/contract"
	"contracts-manager/internal/utils/context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtProvider *token.JWTProvider
	authUC      *authusecase.Usecase
	contractUC  *contractusecase.Usecase
}

func NewAuthMiddleware(
	jwtProvider *token.JWTProvider,
	authUC *authusecase.Usecase,
	contractUC *contractusecase.Usecase,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtProvider,
		authUC,
		contractUC,
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

func (md *AuthMiddleware) MemberOrAdmin() gin.HandlerFunc {
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

		if authUser.Type == models.UserTypeAdmin {
			c.Next()
			return
		}

		contractID, err := context.GetIdFromParam(c)
		if err != nil {
			context.RespondError(c, http.StatusBadRequest, errors.New("invalid contract_id"))
			return
		}

		ctx := c.Request.Context()

		user, err := md.authUC.GetByID(ctx, authUser.ID)
		if err != nil {
			context.RespondError(c, http.StatusUnauthorized, err)
			return
		}

		if user.PersonID == nil {
			context.RespondError(c, http.StatusUnauthorized, http.ErrBodyNotAllowed)
			return
		}

		isMember, err := md.contractUC.IsContractMember(c.Request.Context(), *user.PersonID, contractID)
		if err != nil {
			context.RespondError(c, http.StatusInternalServerError, err)
			return
		}

		if !isMember {
			context.RespondError(c, http.StatusForbidden, auth.ErrRoleNotAllowed)
			return
		}

		c.Next()
	}
}
