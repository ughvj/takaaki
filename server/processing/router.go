package processing

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ughvj/takaaki/config"
)

func Init() *echo.Echo {
	conf := config.Loader.Get()
	fmt.Printf("%+v\n", conf)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowOrigins,
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	if conf.Dryrun {
		fmt.Printf("Running on dryrun.")
		e.GET("/inspection_targets", GetInspectionTargetsDryrun)
	} else {
		e.GET("/inspection_targets", GetInspectionTargets)
	}

	return e;
}
