package postgres

import (
	"bytes"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"time"
)

func main() {
	logger := &bytes.Buffer{}
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username("postgres").
		Password("xiaoluo.618").
		Database("embedded").
		Version("v12").
		RuntimePath("/tmp").
		BinaryRepositoryURL("https://repo.local/central.proxy").
		Port(9876).
		StartTimeout(45 * time.Second).
		Logger(logger))
	if err := postgres.Start(); err != nil {
		panic(err)
	}

	// Do test logic

	if err := postgres.Stop(); err != nil {
		panic(err)
	}
}
