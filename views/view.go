package views

import "html/template"

func NewView (layout string ,files ...string) *View {
	files = append(files, "views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml")

// converting files slice in to multiple strings arg 
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View {
		Template: t,
	}
}

type View struct {
	Template * template.Template
	Layout string
} 