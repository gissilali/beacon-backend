package api

import (
	"beacon.silali.com/internal/api/config"
	"beacon.silali.com/internal/api/core"
	"fmt"
)

func StartServer(cfg *config.Config, ctx *core.AppContext) error {
	e := RegisterRoutes(ctx)

	err := e.Start(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}

	return nil
}
