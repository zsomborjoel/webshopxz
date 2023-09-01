package webpage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GetTemplates(pagePath string) (templates *template.Template, err error) {
	templatePath := os.Getenv("TEMPLATE_PATH")
	fullPath := templatePath + pagePath

	var allFiles []string
	files, _ := ioutil.ReadDir(fullPath)
	for _, file := range files {
		filename := file.Name()
		fmt.Println(filename)
		if strings.HasSuffix(filename, ".html") {
			filePath := filepath.Join(fullPath, filename)
			allFiles = append(allFiles, filePath)
		}
	}

	fmt.Println(allFiles)

	return template.New("").ParseFiles(allFiles...)
}
