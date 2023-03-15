package renderer_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/huytran2000-hcmus/learn-go-with-tests/files_and_templating/blogpost/domain"
	"github.com/huytran2000-hcmus/learn-go-with-tests/files_and_templating/blogpost/renderer"

	approvals "github.com/approvals/go-approval-tests"
)

func TestMain(t *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(t.Run())
}

func TestPostRenderer_RenderPost(t *testing.T) {
	post := domain.Post{
		Title: "TDD suck",
		Body: `
# Why TDD is used wrong
## What is TDD
TDD(test-development driven) is a software development technique of writing test first before you write code.
`,
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	rdr, err := renderer.New()
	assertNoError(t, err)

	buf := bytes.Buffer{}
	err = rdr.RenderPost(&buf, post)
	assertNoError(t, err)

	approvals.VerifyString(t, buf.String())
}

func TestPostRenderer_RenderIndex(t *testing.T) {
	posts := []domain.Post{{Title: "Part 1"}, {Title: "Part 2"}}
	rdr, err := renderer.New()
	assertNoError(t, err)

	buf := bytes.Buffer{}
	err = rdr.RenderIndex(&buf, posts)
	assertNoError(t, err)

	got := buf.String()
	approvals.VerifyString(t, got)
}

func BenchmarkRender(b *testing.B) {
	aPost := domain.Post{
		Title:       "hello world",
		Body:        "This is a body",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	rdr, err := renderer.New()
	assertNoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rdr.RenderPost(io.Discard, aPost)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatalf("didn't expect an error %q", err)
	}
}
