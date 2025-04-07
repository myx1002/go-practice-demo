package book

import (
	"context"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
)

type sBook struct {
}

func (s *sBook) GetList(ctx context.Context) (books []entity.Book, err error) {
	err = dao.Book.Ctx(ctx).Scan(&books)
	return
}

func (s *sBook) Add(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sBook) Update(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sBook) Delete(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

func init() {
	service.RegisterIBook(&sBook{})
}
