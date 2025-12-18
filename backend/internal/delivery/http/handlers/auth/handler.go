package authhandler

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/token"
	authusecase "contracts-manager/internal/usecases/auth"
	"contracts-manager/internal/utils"
	"contracts-manager/internal/utils/context"
	"contracts-manager/internal/utils/cookie"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	jwtProvider *token.JWTProvider
	authUC      *authusecase.Usecase
}

func NewHandler(
	authUC *authusecase.Usecase,
	jwtProvider *token.JWTProvider,
) *Handler {
	return &Handler{
		authUC:      authUC,
		jwtProvider: jwtProvider,
	}
}

func (h *Handler) tokenGeneration(c *gin.Context, userID uuid.UUID) {
	refreshToken, err := h.jwtProvider.GenerateRefreshToken(userID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	cookie.SetCookie(c, cookie.RefreshToken, refreshToken, utils.Week)

	accessToken, err := h.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, gin.H{"accessToken": accessToken})
}

func (h *Handler) Login(c *gin.Context) {
	var dto auth.LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	userID, err := h.authUC.Login(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	h.tokenGeneration(c, userID)
}

func (h *Handler) Signup(c *gin.Context) {
	var dto auth.SignupDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	userID, err := h.authUC.Signup(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	h.tokenGeneration(c, userID)
}

func (h *Handler) Logout(c *gin.Context) {
	cookie.ClearCookie(c, cookie.RefreshToken)
	context.RespondVoid(c, http.StatusOK)
}

func (h *Handler) RefreshAccessToken(c *gin.Context) {
	refreshToken, err := cookie.GetCookie(c, cookie.RefreshToken)
	if err != nil {
		context.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	accessToken, err := h.jwtProvider.RefreshAccessToken(refreshToken)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, gin.H{"accessToken": accessToken})
	return
}

func (h *Handler) GetAuthUser(c *gin.Context) {
	authUser, err := context.GetAuthUser(c)
	if err != nil {
		context.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, authUser)
}
