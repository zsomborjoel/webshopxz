package templaterenderer

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

var AllTemplate *template.Template

func Init() {
	rootDir := os.Getenv("TEMPLATE_PATH")
	if rootDir == "" {
		log.Fatal().Msg("TEMPLATE_PATH environment variable is not set")
	}

	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			log.Debug().Msg(path)
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal().Stack().Msg("Error loading templates")
	}

	AllTemplate = root
}

func Render(wr io.Writer, name string, data any) {
	AllTemplate.ExecuteTemplate(wr, name, data)
}
