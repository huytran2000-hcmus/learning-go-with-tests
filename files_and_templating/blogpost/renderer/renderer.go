package renderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	mdParser "github.com/gomarkdown/markdown/parser"
	"github.com/huytran2000-hcmus/learn-go-with-tests/files_and_templating/blogpost/domain"
	"github.com/huytran2000-hcmus/learn-go-with-tests/files_and_templating/blogpost/renderer/viewmodel"
)

//go:embed "templates/*"
var postTemplate embed.FS

var (
	blogTemplate  = "blog.gohtml"
	indexTemplate = "indexes.gohtml"
)

type PostRenderer struct {
	template *template.Template
	parser   *mdParser.Parser
}

func New() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*gohtml")
	if err != nil {
		return nil, err
	}

	exts := mdParser.CommonExtensions | mdParser.AutoHeadingIDs | mdParser.Autolink
	parser := mdParser.NewWithExtensions(exts)
	return &PostRenderer{template: templ, parser: parser}, nil
}

func (p *PostRenderer) RenderPost(w io.Writer, post domain.Post) error {
	vmPost := NewPostVM(post, p.parser)
	return p.template.ExecuteTemplate(w, blogTemplate, vmPost)
}

func (p *PostRenderer) RenderIndex(w io.Writer, posts []domain.Post) error {
	var vmPosts []viewmodel.Post
	for _, post := range posts {
		vmModel := NewPostVM(post, p.parser)
		vmPosts = append(vmPosts, vmModel)
	}

	return p.template.ExecuteTemplate(w, indexTemplate, vmPosts)
}

func NewPostVM(post domain.Post, parser *mdParser.Parser) viewmodel.Post {
	md := markdown.NormalizeNewlines([]byte(post.Body))
	body := markdown.ToHTML(md, parser, nil)
	return viewmodel.Post{
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
		HTMLBody:    template.HTML(body),
	}
}
