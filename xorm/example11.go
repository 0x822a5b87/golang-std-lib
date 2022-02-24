package main

import "fmt"

func example11() {
	querySql := "select * from user limit 1"
	results, _ := engine.Query(querySql)
	for _, record := range results {
		for key, val := range record {
			fmt.Println(key, string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, _ := engine.Exec(updateSql, "xxx-sql", 1)
	fmt.Println(res.RowsAffected())
}
