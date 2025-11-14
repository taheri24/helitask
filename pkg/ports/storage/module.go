package storage

import (
	"go.uber.org/fx"
)

var Module fx.Option

func init() {

	Module = fx.Module("repoPostgres", fx.Provide(NewTodoRepository))

}
