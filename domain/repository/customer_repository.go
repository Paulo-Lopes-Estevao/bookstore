package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type ICustomerRepository interface {
	repository.IRepository[aggregate.Customer]
}
