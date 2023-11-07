package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/truecaller/config"
	"github.com/keshu12345/truecaller/machingPrefixes"
	"github.com/keshu12345/truecaller/server/router"
	"go.uber.org/fx"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()
	log.New(os.Stdout, "", 0)
	app := fx.New(
		config.NewFxModule(*configDirPath, ""),
		machingPrefixes.Module,
		router.Module,
	)
	app.Run()
}
