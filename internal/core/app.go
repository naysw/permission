package core

import (
	"github.com/naysw/permission/internal/db"
	"github.com/naysw/permission/internal/db/sqlite"
	"github.com/naysw/permission/internal/usecase"
)

// type App interface {
// 	PolicyUsecase() usecase.PolicyUsecase
// 	RoleUsecase() usecase.RoleUsecase
// }

type App struct {
	roleUsecase   *usecase.RoleUsecase
	policyUsecase *usecase.PolicyUsecase
}

func (a App) RoleUsecase() *usecase.RoleUsecase {
	return a.roleUsecase
}

func (a App) PolicyUsecase() *usecase.PolicyUsecase {
	return a.policyUsecase
}

func NewApp() *App {
	d, err := db.NewSQLiteConnection()
	if err != nil {
		panic(err)
	}

	if err := sqlite.Migrate(d); err != nil {
		panic(err)
	}

	prp := sqlite.NewPolicyRepo(d)
	rrp := sqlite.NewRoleRepo(d)

	return &App{
		policyUsecase: usecase.NewPolicyUsecase(prp),
		roleUsecase:   usecase.NewRoleUsecase(rrp, prp),
	}
}
