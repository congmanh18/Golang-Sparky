package db

import (
	"errors"
	"fmt"
)

type Connection struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (c Connection) validate() error {
	var errs []error
	if c.Host == "" {
		return append(errs, fmt.Errorf("Host required")) 
	}
	if c.User == "" {
		return append(errs, fmt.Errorf("User required"))
	}
	if c.Password == "" {
		return append(errs, fmt.Errorf("Password required")) 
	}
	if c.DBName == "" {
		return append(errs, fmt.Errorf("DBName required"))
	}
	if c.Port == "" {
		return append(errs, fmt.Errorf("Port required"))
	}
}

func (c Connection) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s")"host=localhost user"
}
