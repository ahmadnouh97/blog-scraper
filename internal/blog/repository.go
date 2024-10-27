package blog

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddBlog(blog *Blog) (int64, error) {
	query := "INSERT INTO blogs (id, title, content, created_at) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, blog.ID, blog.Title, blog.Content, blog.CreatedAt)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *Repository) GetBlogs() ([]*Blog, error) {
	query := "SELECT id, title, content, created_at FROM blogs"
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blogs []*Blog

	for rows.Next() {
		blog := &Blog{}
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CreatedAt)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}
