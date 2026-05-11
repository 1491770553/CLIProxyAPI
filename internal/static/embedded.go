package static

import (
	_ "embed"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed management.html
var embeddedManagementHTML string

// ServeManagementPanel serves the embedded management panel HTML.
func ServeManagementPanel(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Header("Cache-Control", "no-cache")
	reader := strings.NewReader(embeddedManagementHTML)
	c.DataFromReader(http.StatusOK, int64(len(embeddedManagementHTML)), "text/html; charset=utf-8", reader, nil)
}
