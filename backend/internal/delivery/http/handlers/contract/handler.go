package contracthandler

import (
	filehandler "contracts-manager/internal/delivery/http/file"
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
	contractID, err := context.GetIdFromParam(c)
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
	contractID, err := context.GetIdFromParam(c)
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
	contractID, err := context.GetIdFromParam(c)
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

func (h *Handler) AddPerson(c *gin.Context) {
	var dto contract.AddPersonDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.contractUC.AddPerson(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, res)
}

func (h *Handler) RemovePerson(c *gin.Context) {
	contractID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	personID, err := context.GetUUIDFromParam(c, "personId")
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	if err = h.contractUC.RemovePerson(c.Request.Context(), contractID, personID); err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondVoid(c, http.StatusOK)
}

func (h *Handler) GetPersons(c *gin.Context) {
	contractID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	persons, err := h.contractUC.GetPersons(c.Request.Context(), contractID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, persons)
}

func (h *Handler) GetContractsByPerson(c *gin.Context) {
	personID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	contracts, err := h.contractUC.GetContractsByPerson(c.Request.Context(), personID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, contracts)
}

func (h *Handler) Import(c *gin.Context) {
	filehandler.Import(c, h.contractUC)
}

func (h *Handler) Export(c *gin.Context) {
	filehandler.Export(c, h.contractUC)
}
