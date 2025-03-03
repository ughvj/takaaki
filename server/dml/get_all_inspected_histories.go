package dml

import (
	"fmt"
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
			fmt.Printf("%s¥n", err)
			return nil
		}
		ihs = append(ihs, ih)
	}

	return &ihs
}
