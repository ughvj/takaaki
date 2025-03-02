package processing

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/ughvj/takaaki/drivers"
	"github.com/ughvj/takaaki/dml"
	"github.com/ughvj/takaaki/types"
	"github.com/ughvj/takaaki/config"
	"github.com/ughvj/takaaki/external"
)

func GetInspectionTargetsDryrun(c echo.Context) error {
	return c.JSON(http.StatusOK, SubtractNEAERfromIHS(
		types.GetStaticNyutaikunEntranceAndExitResult(),
		types.GetStaticInspectedHistoryCase2(),
		types.GetStaticStudents(),
	))
}

func GetInspectionTargets(c echo.Context) error {
	conf := config.Loader.Get()
	api := external.NewNyutaikunApi(conf.NyutaikunAccessToken)
	neaer := api.RequestEntranceAndExit()

	db, err := drivers.NewMysqlDriver()
	if err != nil {
		panic(err)
	}
	defer db.Use().Close()

	ihs := dml.GetAllInspectedHistories(db)
	ss := dml.GetAllStudents(db)

	return c.JSON(http.StatusOK, SubtractNEAERfromIHS(neaer, ihs, ss))
}

func SubtractNEAERfromIHS(
	neaer *types.NyutaikunEntranceAndExitResult,
	ihs *types.InspectedHistories,
	ss *types.Students,
) *types.Students {
	targets := types.Students{}

	for _, neae := range neaer.Data {
		cnt := 0
		for _, ih := range *ihs {
			if neae.UserId == ih.NyutaikunUserId {
				break
			}
			cnt++
		}
		if len(*ihs) == cnt {
			targets = append(targets, *(ss.FindByNyutaikunUserId(neae.UserId)))
		}
	}

	return &targets
}
