package postmodel

import (
	"go-web-native-crud/config"
	"go-web-native-crud/entities"
)

func GetAll() []entities.Post {
	rows, err := config.DB.Query(`SELECT * FROM posts`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var posts []entities.Post

	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.User_id,
			&post.Category,
			&post.Image,
			&post.Slug,
			&post.Excerpt,
			&post.Body,
			&post.PublishedAt,
			&post.CreatedAt,
			&post.UpdatedAt); err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	return posts
}
