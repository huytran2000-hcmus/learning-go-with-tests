package domain_test

import (
	"testing"
	"testing/fstest"

	"github.com/huytran2000-hcmus/learn-go-with-tests/files_and_templating/blogpost/domain"

	"github.com/google/go-cmp/cmp"
)

func TestNewBlogPostsFromFS(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`)},
		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Rust is gay`)},
	}

	posts, err := domain.NewPostsFromFS(fs)
	assertNoError(t, err)

	assertPostLength(t, posts, fs)

	got := posts[0]
	want := domain.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}
	assertPost(t, got, want)
}

func assertPost(t *testing.T, got domain.Post, want domain.Post) {
	t.Helper()
	if !cmp.Equal(got, want) {
		t.Errorf("diff -want +got\n%s", cmp.Diff(want, got))
	}
}

func assertPostLength(t *testing.T, posts []domain.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("fs = %v\n"+
			"NewPostFromFs(fs) = %v\n"+
			"got %d posts, want %d posts", fs, posts, len(posts), len(fs))
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Didn't expect an error but get one")
	}
}
