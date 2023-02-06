package factory

import (
	mhsDelivery "be13/ca/features/mahasiswa/delivery"
	mhsRepo "be13/ca/features/mahasiswa/repository"
	mhsService "be13/ca/features/mahasiswa/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitFactory(e *echo.Echo, db *sql.DB) {
	mhsRepoFactory := mhsRepo.NewRaw(db)
	mhsServiceFactory := mhsService.New(mhsRepoFactory)
	mhsDelivery.New(mhsServiceFactory, e)

}
