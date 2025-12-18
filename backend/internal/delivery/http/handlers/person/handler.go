package personhandler

import (
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

	personID, err := h.personUC.Create(c.Request.Context(), dto)
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	context.RespondWithValue(c, http.StatusCreated, gin.H{"personID": personID})
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
	limitStr := c.DefaultQuery("limit", "20")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	filter := person.PersonFilter{
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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		context.RespondError(c, http.StatusBadRequest, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		context.RespondError(c, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	var imported int
	var errors []string

	switch fileHeader.Header.Get("Content-Type") {
	case "application/json":
		imported, errors = h.personUC.ImportJSON(c.Request.Context(), file)
	case "text/csv", "application/vnd.ms-excel":
		imported, errors = h.personUC.ImportCSV(c.Request.Context(), file)
	default:
		context.RespondError(c, http.StatusBadRequest, ErrUnsupportedFileType)
		return
	}

	context.RespondWithValue(c, http.StatusOK, gin.H{
		"imported": imported,
		"errors":   errors,
	})
}

func (h *Handler) Export(c *gin.Context) {
	format := c.DefaultQuery("format", "csv")

	ctx := c.Request.Context()

	switch format {
	case "csv":
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", `attachment; filename="persons.csv"`)

		err := h.personUC.ExportCSV(ctx, c.Writer)
		if err != nil {
			context.RespondError(c, http.StatusInternalServerError, err)
		}

	case "json":
		c.Header("Content-Type", "application/json")
		c.Header("Content-Disposition", `attachment; filename="persons.json"`)

		err := h.personUC.ExportJSON(ctx, c.Writer)
		if err != nil {
			context.RespondError(c, http.StatusInternalServerError, err)
		}

	default:
		context.RespondError(c, http.StatusBadRequest, ErrUnsupportedFileType)
	}
}
