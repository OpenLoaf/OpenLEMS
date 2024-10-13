package internal_service_sqlite

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
)

type SServiceSqlite struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}
