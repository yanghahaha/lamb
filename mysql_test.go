package database

import (
	"encoding/json"
	"testing"
)

func Test_query(t *testing.T) {
	GetDb().Config(
		"w",
		"pwd",
		"sss",
		"3306",
		"aaa",
		"utf8")
	sql := "select id, url from ts_comment order by id desc limit 2"
	rows, err := GetDb().Query(sql)
	if err != nil {
		println(err.Error())
		t.Error(err.Error())
	} else {
		strings, _ := json.Marshal(rows)
		println(string(strings))
	}
}
