package repo

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/skurtz97/splat/internal/splat"
)

func (cp *PostgresPool) GetPost(ctx context.Context, id string) (splat.Post, error) {
	var post splat.Post
	row := cp.pool.QueryRow(ctx, `SELECT id, title, content FROM posts WHERE id = $1`, id)
	if err := row.Scan(&post.Id, &post.Title, &post.Content); err != nil {
		return splat.Post{}, fmt.Errorf("error getting post: %w", err)
	}

	return post, nil
}

func (cp *PostgresPool) ListPosts(ctx context.Context) ([]splat.Post, error) {
	var posts []splat.Post
	rows, err := cp.pool.Query(ctx, `SELECT id, title, content FROM posts`)
	if err != nil {
		return nil, fmt.Errorf("error getting posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post splat.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content); err != nil {
			return nil, fmt.Errorf("error getting posts: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (cp *PostgresPool) CreatePost(ctx context.Context, post splat.Post) error {
	post.Id = uuid.NewV4().String()
	_, err := cp.pool.Exec(ctx, `INSERT INTO posts (id, title, content) VALUES ($1, $2, $3)`, post.Id, post.Title, post.Content)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (cp *PostgresPool) UpdatePost(ctx context.Context, post splat.Post) error {
	_, err := cp.pool.Exec(ctx, `UPDATE posts SET title = $1, content = $2 WHERE id = $3`, post.Title, post.Content, post.Id)
	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}

	return nil
}

func (cp *PostgresPool) DeletePost(ctx context.Context, id string) error {
	_, err := cp.pool.Exec(ctx, `DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	return nil
}
