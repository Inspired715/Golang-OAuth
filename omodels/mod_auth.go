package omodels

import (
	"fmt"
	"projects/config"
)

type Blog struct {
	Id          int64
	Title       string
	Description string
}

func GetBlogList() ([]Blog, error) {

	var blogs []Blog

	rows, err := config.DB.Query("SELECT * FROM blog")
	if err != nil {
		return nil, fmt.Errorf("Blog list error")
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog

		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Description); err != nil {
			return nil, fmt.Errorf("%s", "Parsing error of blog")
		}

		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", "Parsing error of blog")
	}

	return blogs, nil
}
