package domain

import (
	"github.com/chack93/go_base/internal/domain/session"
	"github.com/chack93/go_base/internal/service/database"
)

func Init() error {
	database.Get().AutoMigrate(
		&session.Session{},
	)

	return nil
}
