package splat

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Post struct {
	Id      string
	Title   string
	Content string
}

type Store interface {
	GetPost(ctx context.Context, id string) (Post, error)
	ListPosts(ctx context.Context) ([]Post, error)
	CreatePost(ctx context.Context, post Post) error
	UpdatePost(ctx context.Context, post Post) error
	DeletePost(ctx context.Context, id string) error
}

type Service struct {
	log   *zap.Logger
	Store Store
}

func NewService(log *zap.Logger, store Store) *Service {
	return &Service{
		log:   log,
		Store: store,
	}
}

func (s *Service) GetPost(id string) (Post, error) {
	s.log.Info("getting post with id: ", zap.String("id", id))

	post, err := s.Store.GetPost(context.Background(), id)
	if err != nil {
		return Post{}, fmt.Errorf("error getting post: %w", err)
	}

	return post, nil
}

func (s *Service) ListPosts() ([]Post, error) {
	s.log.Info("listing posts")

	posts, err := s.Store.ListPosts(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting posts: %w", err)
	}

	return posts, nil
}

func (s *Service) CreatePost(post Post) error {
	s.log.Info("creating post: ", zap.Any("post", post))

	if err := s.Store.CreatePost(context.Background(), post); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (s *Service) UpdatePost(post Post) error {
	s.log.Info("updating post: ", zap.Any("post", post))

	if err := s.Store.UpdatePost(context.Background(), post); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}

	return nil
}

func (s *Service) DeletePost(id string) error {
	s.log.Info("deleting post with id: ", zap.String("id", id))

	if err := s.Store.DeletePost(context.Background(), id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	return nil
}
