package dml

import (
	"github.com/ughvj/takaaki/drivers"
	"github.com/ughvj/takaaki/types"
)

func GetAllInspectedHistories(db drivers.DBConnector) *types.InspectedHistories {
	loadedDML, err := Loader.Get("get_all_inspected_histories")

	rows, err := db.Use().Query(loadedDML)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var ihs types.InspectedHistories
	for rows.Next() {
		var ih types.InspectedHistory
		err := rows.Scan(ih.Refs()...)
		if err != nil {
			panic(err.Error())
		}
		ihs = append(ihs, ih)
	}

	return &ihs
}
