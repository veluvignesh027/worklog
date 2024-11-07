package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var UserStoryCounter int = 1000
var DBInstances DatabaseInstancesModal

type DatabaseInstancesModal struct {
	UserDatabaseInstance       *sql.DB
	UserStoryDatabaseInstance  *sql.DB
	WorkLogDatabaseInstance    *sql.DB
	SubWorkLogDatabaseInstance *sql.DB
	SessionDatabaseInstance    *sql.DB
}

func ConnectAllDatabase() error {

	var err error
	userConn := os.Getenv("USER_DB_CONN")
	userStoryConn := os.Getenv("USER_STORY_DB_CONN")
	workLogConn := os.Getenv("WORK_LOG_DB_CONN")
	subWorkLogConn := os.Getenv("SUB_WL_DB_CONN")
	sessionConn := os.Getenv("SESSION_DB_CONN")

	if err = Connect2DB(DBInstances.UserDatabaseInstance, userConn); err != nil {
		return err
	}
	if err = Connect2DB(DBInstances.UserStoryDatabaseInstance, userStoryConn); err != nil {
		return err
	}
	if err = Connect2DB(DBInstances.WorkLogDatabaseInstance, workLogConn); err != nil {
		return err
	}
	if err = Connect2DB(DBInstances.SubWorkLogDatabaseInstance, subWorkLogConn); err != nil {
		return err
	}
	if err = Connect2DB(DBInstances.SessionDatabaseInstance, sessionConn); err != nil {
		return err
	}
	return nil
}

func Connect2DB(db *sql.DB, connstr string) error {
	var err error
	if db == nil {
		db, err = sql.Open("postgres", connstr)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = db.Ping()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func CloseAllDatabase() error {
	if err := DBInstances.UserDatabaseInstance.Close(); err != nil {
		return err
	}
	if err := DBInstances.UserStoryDatabaseInstance.Close(); err != nil {
		return err
	}
	if err := DBInstances.WorkLogDatabaseInstance.Close(); err != nil {
		return err
	}
	if err := DBInstances.SubWorkLogDatabaseInstance.Close(); err != nil {
		return err
	}
	return nil
}
