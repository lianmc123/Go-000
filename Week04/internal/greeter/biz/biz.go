package biz

import (
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloItem struct {
	Message string
	Source  string
}

type HelloRepo interface {
	SaveHello(db *sql.DB, hello *HelloItem) error
}

type Biz struct {
	repo HelloRepo
	db   *sql.DB
}

func NewBiz(db *sql.DB, dao HelloRepo) *Biz {
	return &Biz{db: db, repo: dao}
}

func (b *Biz) Hello(message string) (string, error) {
	if err := b.repo.SaveHello(b.db, &HelloItem{
		Message: message,
		Source:  "Hello",
	}); err != nil {
		return "", status.Errorf(codes.Internal, "Hello Biz.repo.SaveHello err")
	}
	return "Hello " + message, nil
}

func (b *Biz) HelloAgain(message string) (string, error) {
	if err := b.repo.SaveHello(b.db, &HelloItem{
		Message: message,
		Source:  "HelloAgain",
	}); err != nil {
		return "", status.Errorf(codes.Internal, "HelloAgain Biz.repo.SaveHello err")
	}
	return "HelloAgain " + message, nil
}
