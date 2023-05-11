//go:build ignore

package main

import (
	"fmt"
	"os"

	"com.vitanexus/main/internal/custom/models"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(models.New()),
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
