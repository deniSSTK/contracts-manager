package authhandler

import (
	gocontext "context"
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/config"
	"contracts-manager/internal/infrastructure/logger"
	"contracts-manager/internal/infrastructure/token"
	authusecase "contracts-manager/internal/usecases/auth"
	"contracts-manager/internal/utils"
	"contracts-manager/internal/utils/context"
	"contracts-manager/internal/utils/cookie"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	authUC      *authusecase.Usecase
	jwtProvider *token.JWTProvider
}

func NewHandler(
	authUC *authusecase.Usecase,
	jwtProvider *token.JWTProvider,
	cfg *config.Config,
	log *logger.Logger,
) *Handler {
	adminId, err := authUC.Signup(gocontext.Background(), cfg.Admin)

	if err != nil {
		log.Errorf(errors.New("failed to create admin"), err)
	} else {
		log.Info("created admin", zap.String("adminId", adminId.String()))
	}

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

	accessTokenInfo, err := h.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, accessTokenInfo)
}

func (h *Handler) Login(c *gin.Context) {
	var dto auth.LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	userID, err := h.authUC.Login(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
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

	accessTokenInfo, err := h.jwtProvider.RefreshAccessToken(refreshToken)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, accessTokenInfo)
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

func (h *Handler) Get(c *gin.Context) {
	userId, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.authUC.GetByID(c.Request.Context(), userId)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, user)
}

func (h *Handler) List(c *gin.Context) {
	username := c.Query("username")
	email := c.Query("email")
	userType := c.Query("type")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	filter := auth.Filter{
		Username: nil,
		Email:    nil,
		Type:     nil,
		Page:     page,
		Limit:    limit,
	}

	if username != "" {
		filter.Username = &username
	}
	if email != "" {
		filter.Email = &email
	}
	if userType != "" {
		filter.Type = &userType
	}

	result, err := h.authUC.List(c.Request.Context(), filter)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, result)
}

func (h *Handler) Update(c *gin.Context) {
	userID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	var dto auth.UpdateDTO
	if err = c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.authUC.Update(c.Request.Context(), userID, dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, res)
}

func (h *Handler) GetUserContracts(c *gin.Context) {
	authUser, err := context.GetAuthUser(c)
	if err != nil {
		context.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	contracts, err := h.authUC.GetContractsByID(c.Request.Context(), authUser.ID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, contracts)
}
