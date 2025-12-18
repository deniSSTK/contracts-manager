package contracthandler

import (
	"contracts-manager/internal/domain/contract"
	contractusecase "contracts-manager/internal/usecases/contract"
	"contracts-manager/internal/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	contractUC *contractusecase.Usecase
}

func NewHandler(contractUC *contractusecase.Usecase) *Handler {
	return &Handler{contractUC}
}

func (h *Handler) Create(c *gin.Context) {
	var dto contract.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.contractUC.Create(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, res)
}

func (h *Handler) Get(c *gin.Context) {
	contractID, err := context.GetUUIDFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.contractUC.GetByID(c.Request.Context(), contractID)
	if err != nil {
		context.RespondError(c, http.StatusNotFound, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, res)
}

func (h *Handler) Update(c *gin.Context) {
	contractID, err := context.GetUUIDFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	var dto contract.UpdateDTO
	if err = c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.contractUC.Update(c.Request.Context(), contractID, dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, res)
}

func (h *Handler) Delete(c *gin.Context) {
	contractID, err := context.GetUUIDFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	if err = h.contractUC.Delete(c.Request.Context(), contractID); err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondVoid(c, http.StatusOK)
}
