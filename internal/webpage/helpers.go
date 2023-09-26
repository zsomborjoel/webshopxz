package webpage

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func GetTemplates(pagePath string) (templates *template.Template, err error) {
	allFiles := GetTemplateFiles(pagePath)

	return template.New("").ParseFiles(allFiles...)
}

func GetTemplateFiles(pagePath string) ([]string) {
	templatePath := os.Getenv("TEMPLATE_PATH")
	fullPath := templatePath + pagePath

	var allFiles []string
	files, _ := ioutil.ReadDir(fullPath)
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			filePath := filepath.Join(fullPath, filename)
			allFiles = append(allFiles, filePath)
		}
	}

	return allFiles
}

func IsHTMXRequest(c *gin.Context) bool {
	htmx := c.Request.Header.Get(common.HTMXRequest)
	return htmx != ""
}