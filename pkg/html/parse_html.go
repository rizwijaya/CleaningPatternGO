package html

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func Render(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func ManualRender(tmpDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("product", tmpDir+"product/product_index.html", tmpDir+"layouts/base.html")
	return r
}
