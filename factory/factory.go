package factory

import (
	userDelivery "be13/ca/features/user/delivery"
	userRepo "be13/ca/features/user/repository"
	userService "be13/ca/features/user/service"

	bookDelivery "be13/ca/features/book/delivery"
	bookRepo "be13/ca/features/book/repository"
	bookService "be13/ca/features/book/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	bookRepoFactory := bookRepo.New(db)
	bookServiceFactory := bookService.New(bookRepoFactory)
	bookDelivery.New(bookServiceFactory, e)

}
