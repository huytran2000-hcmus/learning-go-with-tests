package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleTag       = "Title: "
	descriptionTag = "Description: "
	tagsTag        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(r io.Reader) (Post, error) {
	scanner := bufio.NewScanner(r)

	readMetaLine := func(tag string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tag)
	}

	title := readMetaLine(titleTag)
	description := readMetaLine(descriptionTag)
	tags := strings.Split(readMetaLine(tagsTag), ", ")

	body := ReadBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        string(body),
	}, nil
}

func ReadBody(scanner *bufio.Scanner) string {
	scanner.Scan() // Ignore --- line

	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
