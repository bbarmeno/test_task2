package main

import (
	"acc_balance"
	"acc_balance/config"
	"acc_balance/storage"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

func main() {

	cfg, err := config.SetupConfig()
	if err != nil {
		log.Fatal("failed to setup config: ", err)
	}

	db, err := storage.NewSql(cfg.SqlCfg)
	if err != nil {
		log.Fatal("failed to establish database connection: ", err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	if err := acc_balance.Setup(db, app); err != nil {
		log.Fatal("failed to setup handlers: ", err)
	}

	errgrp, ctx := errgroup.WithContext(context.Background())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	errgrp.Go(func() error {

		if err := app.Listen(":8080"); err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	errgrp.Go(func() error {
		<-ctx.Done()

		return app.ShutdownWithTimeout(5 * time.Second)
	})

	if err := errgrp.Wait(); err != nil {
		log.Fatal(err)
	}
}
