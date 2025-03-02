package dml

import (
	"github.com/ughvj/takaaki/drivers"
	"github.com/ughvj/takaaki/types"
)

func GetAllStudents(db drivers.DBConnector) *types.Students {
	loadedDML, err := Loader.Get("get_all_students")

	rows, err := db.Use().Query(loadedDML)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var ss types.Students
	for rows.Next() {
		var s types.Student
		err := rows.Scan(s.Refs()...)
		if err != nil {
			panic(err.Error())
		}
		ss = append(ss, s)
	}

	return &ss
}
