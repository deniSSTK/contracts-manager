package personhandler

import (
	filehandler "contracts-manager/internal/delivery/http/file"
	"contracts-manager/internal/domain/person"
	personusecase "contracts-manager/internal/usecases/person"
	"contracts-manager/internal/utils/context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	personUC *personusecase.Usecase
}

func NewHandler(personUC *personusecase.Usecase) *Handler {
	return &Handler{
		personUC: personUC,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var dto person.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}
	if err := h.personUC.Create(c.Request.Context(), dto); err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondVoid(c, http.StatusCreated)
}

func (h *Handler) Get(c *gin.Context) {
	personID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	targetPerson, err := h.personUC.GetByID(c.Request.Context(), personID)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, targetPerson)
}

func (h *Handler) Update(c *gin.Context) {
	var dto person.UpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	personID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.personUC.Update(c.Request.Context(), personID, dto)

	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, res)
}

func (h *Handler) Delete(c *gin.Context) {
	personID, err := context.GetIdFromParam(c)
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	if err = h.personUC.Delete(c.Request.Context(), personID); err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondVoid(c, http.StatusOK)
}

func (h *Handler) List(c *gin.Context) {
	name := c.Query("name")
	personType := c.Query("type")
	code := c.Query("code")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	filter := person.Filter{
		Name:  nil,
		Type:  nil,
		Code:  nil,
		Page:  page,
		Limit: limit,
	}

	if name != "" {
		filter.Name = &name
	}
	if personType != "" {
		filter.Type = &personType
	}
	if code != "" {
		filter.Code = &code
	}

	result, err := h.personUC.List(c.Request.Context(), filter)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusOK, result)
}

func (h *Handler) Import(c *gin.Context) {
	filehandler.Import(c, h.personUC)
}

func (h *Handler) Export(c *gin.Context) {
	filehandler.Export(c, h.personUC)
}
