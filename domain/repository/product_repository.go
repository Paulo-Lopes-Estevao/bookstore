package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IProductRepository interface {
	repository.Repository[aggregate.Product]
	CreateCategory(name string) error
}
