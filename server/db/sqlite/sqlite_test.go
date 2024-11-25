package sqlite

import (
	"github.com/gsxhnd/jaha/server/db/database"
	"github.com/gsxhnd/jaha/utils"
)

func getMockDB() (database.Driver, error) {
	var logger = utils.NewLogger(&utils.Config{
		Mode: "dev",
		Log: utils.LogConfig{
			Level: "debug",
		},
	})

	var mockSqliteDB, err = NewSqliteDB("../../../data/jaha.db", logger)
	if err != nil {
		return nil, err
	}

	return mockSqliteDB, nil
}
