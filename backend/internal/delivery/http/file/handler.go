package fileheader

import (
	gocontext "context"
	"contracts-manager/internal/utils/context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileUsecase interface {
	ImportJSON(ctx gocontext.Context, reader io.Reader) (int, []string)
	ImportCSV(ctx gocontext.Context, reader io.Reader) (int, []string)

	ExportCSV(ctx gocontext.Context, w io.Writer) error
	ExportJSON(ctx gocontext.Context, w io.Writer) error
}

func Import(
	c *gin.Context,
	usecase FileUsecase,
) {
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
		imported, errors = usecase.ImportJSON(c.Request.Context(), file)
	case "text/csv", "application/vnd.ms-excel":
		imported, errors = usecase.ImportCSV(c.Request.Context(), file)
	default:
		context.RespondError(c, http.StatusBadRequest, ErrUnsupportedFileType)
		return
	}

	context.RespondWithValue(c, http.StatusOK, gin.H{
		"imported": imported,
		"errors":   errors,
	})
}

func Export(
	c *gin.Context,
	usecase FileUsecase,
) {
	format := c.DefaultQuery("format", "csv")

	ctx := c.Request.Context()

	switch format {
	case "csv":
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", `attachment; filename="persons.csv"`)

		err := usecase.ExportCSV(ctx, c.Writer)
		if err != nil {
			context.RespondError(c, http.StatusInternalServerError, err)
		}

	case "json":
		c.Header("Content-Type", "application/json")
		c.Header("Content-Disposition", `attachment; filename="persons.json"`)

		err := usecase.ExportJSON(ctx, c.Writer)
		if err != nil {
			context.RespondError(c, http.StatusInternalServerError, err)
		}

	default:
		context.RespondError(c, http.StatusBadRequest, ErrUnsupportedFileType)
	}
}
