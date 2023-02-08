package factory

import (
	mhsDelivery "be13/ca/features/mahasiswa/delivery"
	mhsRepo "be13/ca/features/mahasiswa/repository"
	mhsService "be13/ca/features/mahasiswa/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitFactory(e *gin.Engine, db *sql.DB) {
	mhsRepoFactory := mhsRepo.NewRaw(db)
	mhsServiceFactory := mhsService.New(mhsRepoFactory)
	mhsDelivery.New(mhsServiceFactory, e)

}
