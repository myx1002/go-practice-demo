package service

import (
	"context"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
)

// 1.定义接口
type IBook interface {
	GetList(ctx context.Context) (books []entity.Book, err error)
	Add(ctx context.Context, book do.Book) (err error)
	Update(ctx context.Context, book do.Book) (err error)
	Delete(ctx context.Context) (err error)
}

// 2.定义接口变量
var localIBook IBook

// 3.定义一个获取接口实力的函数
func Book() IBook {
	if localIBook == nil {
		panic("implement not found for interface IBook")
	}
	return localIBook
}

// 4.定义一个接口实现注册的方法
func RegisterIBook(i IBook) {
	localIBook = i
}
