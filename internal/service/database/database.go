package database

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var database *Database

func Get() *Database {
	return database
}

func New() *Database {
	database = &Database{}
	return database
}

func (s *Database) Init() error {
	dbUrl, err := url.Parse(viper.GetString("database.url"))
	if err != nil {
		logrus.Errorf("database.url config invalid: %s, err: %v", dbUrl, err)
		return err
	}
	if viper.GetBool("database.usetestdb") {
		dbUrl.Path += "_test"
	}

	if err := ensureAppTableExists(*dbUrl); err != nil {
		logrus.Errorf("create app table failed, err: %v", err)
		return err
	}

	db, err := gorm.Open(postgres.Open(dbUrl.String()), &gorm.Config{})
	if err != nil {
		logrus.Errorf("open db connection failed, err: %v", err)
		return err
	}
	s.db = db

	if err := db.AutoMigrate(); err != nil {
		logrus.Errorf("auto-migrate db failed, err: %v", err)
		return err
	}

	logrus.Infof("connected to host: %s%s", dbUrl.Host, dbUrl.Path)
	return nil
}

func ensureAppTableExists(dbUrl url.URL) error {
	appTable := strings.Trim(dbUrl.Path, "/")
	dbUrl.Path = "postgres"
	db, err := gorm.Open(postgres.Open(dbUrl.String()), &gorm.Config{})
	if err != nil {
		logrus.Errorf("open db connection failed, err: %v", err)
		return err
	}

	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", appTable)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		logrus.Errorf("query for %s failed, err: %v", appTable, rs.Error)
		return rs.Error
	}

	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", appTable)
		if rs := db.Exec(stmt); rs.Error != nil {
			logrus.Errorf("create table %s failed, err: %v", appTable, rs.Error)
			return rs.Error
		}

		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			logrus.Errorf("close connection failed, err: %v", err)
			return err
		}
		logrus.Infof("app table: %s created", appTable)
	} else {
		logrus.Debugf("app table: %s exists", appTable)
	}

	return nil
}
