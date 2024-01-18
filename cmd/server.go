package main

import (
	"flag"
	app "github.com/eduartua/ddd-web-service"
	"github.com/eduartua/ddd-web-service/psql"
	"github.com/eduartua/ddd-web-service/rand"
)

func main() {
	boolPtr := flag.Bool("prod", false, "Provide this flag in production. This ensures that a .config file is provided before the application starts.")
	flag.Parse()
	cfg := app.LoadConfig(*boolPtr)

	dbCfg := cfg.Database

	psqlStores, err := psql.NewStores(
		psql.WithGorm(dbCfg.ConnectionInfo()),
		psql.WithLogMode(!cfg.IsProd()),
		psql.WithUser(cfg.Vars.Pepper, cfg.Vars.HMACKey),
	)
	must(err)
	defer psqlStores.Close()
	//_ = psqlStores.DestructiveReset()
	_ = psqlStores.AutoMigrate()

	b, err := rand.Bytes(32)
	must(err)

	server := http.NewServer(
		psqlStores.User,
		b
	)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}