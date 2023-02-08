package blogpost

import (
	"io/fs"
)

func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range dir {
		post, err := getPost(fileSys, entry.Name())
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSys fs.FS, filename string) (post Post, err error) {
	postFile, err := fileSys.Open(filename)
	defer func() {
		closeErr := postFile.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()
	if err != nil {
		return Post{}, err
	}

	return newPost(postFile)
}
