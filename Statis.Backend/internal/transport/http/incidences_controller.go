package http

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IncidencesController struct {
	db *gorm.DB
}

func NewIncidencesController(db *gorm.DB) *IncidencesController {
	return &IncidencesController{
		db: db,
	}
}

// GET
func (e *IncidencesController) Get(ctx *gin.Context) {

	data, ok := os.LookupEnv("INCIDENCE_DIRECTORY")
	data = strings.ReplaceAll(data, "{page}", ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	res, err := http.Get(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	ctx.Data(http.StatusOK, "text/plain", bytes)
}
