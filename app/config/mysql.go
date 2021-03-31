package config

import (
	"log"
	"strings"
)

type MySQL struct {
	dataSourceName string
	driverName     string
}

func host() string {
	v, err := getString("MYSQL_HOST")
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func user() string {
	v, err := getString("MYSQL_USER")
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func password() string {
	v, err := getString("MYSQL_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func database() string {
	v, err := getString("MYSQL_DATABASE")
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func timeZone() string {
	tz, err := getString("MYSQL_TZ")
	if err != nil {
		return "Asia%2FTokyo"
	}

	return strings.Replace(tz, "/", "%2F", 1)
}

func (m MySQL) DriverName() string {
	return m.driverName
}

func (m MySQL) DataSourceName() string {
	return m.dataSourceName
}

func MySQLConfig() *MySQL {
	driverName := "mysql"
	dataSourceName := user() + ":" + password() + "@tcp(" + host() + ")/" + database() + "?parseTime=true&loc=" + timeZone()

	return &MySQL{
		dataSourceName: dataSourceName,
		driverName:     driverName,
	}
}
