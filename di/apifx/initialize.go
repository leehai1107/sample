package apifx

import (
	"github.com/leehai1107/bipbip/service/bid/delivery/http"
	"github.com/leehai1107/bipbip/service/bid/repository"
	"github.com/leehai1107/bipbip/service/bid/usecase"
	"github.com/leehai1107/bipbip/sql"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideRouter,
	provideHandler,
	provideRepo,
	provideUsecase,
	provideMigration,
)

func provideRouter(handler http.IHandler) http.Router {
	return http.NewRouter(handler)
}

func provideMigration(db *gorm.DB) sql.IMigration {
	return sql.NewMigration(db)
}

func provideHandler(usecase usecase.IUserUsecase) http.IHandler {
	return http.NewHandler(usecase)
}

func provideRepo(db *gorm.DB) repository.IUserRepo {
	return repository.NewUserRepo(db)
}

func provideUsecase(repo repository.IUserRepo) usecase.IUserUsecase {
	return usecase.NewUserUsecase(repo)
}
