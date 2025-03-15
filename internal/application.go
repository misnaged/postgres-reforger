package internal

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/misnaged/scriptorium/logger"
	"os"
	"os/signal"
	"postgres-reforger/internal/repository"
	"postgres-reforger/internal/server"
	"postgres-reforger/internal/service"
	"syscall"

	"postgres-reforger/config"
)

type App struct {
	// application configuration
	config *config.Scheme

	http server.IServer
	db   *pg.DB
	Srv  service.IService
}

// NewApplication create new App instance
func NewApplication() (app *App, err error) {
	return &App{
		config: &config.Scheme{},
	}, nil
}

// Init initialize application and all necessary instances
func (app *App) Init() error {

	if err := app.initDb(app.config); err != nil {
		return fmt.Errorf("application database initialisation: %w", err)
	}
	if err := app.initService(); err != nil {
		return fmt.Errorf("application database initialisation: %w", err)
	}
	app.initServer()

	return nil
}

func (app *App) initService() error {
	srvc, err := repository.NewRepository(app.db)
	if err != nil {
		return fmt.Errorf("failed to create new repo: %w", err)
	}
	app.Srv = service.NewService(app.Config(), srvc)
	logger.Log().Info("application service layer initialized")
	return nil
}

// Config return App config Scheme
func (app *App) Config() *config.Scheme {
	return app.config
}
func (app *App) initDb(cfg *config.Scheme) error {
	opts := pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Db.Host, cfg.Db.Port),
		User:     cfg.Db.User,
		Password: cfg.Db.Password,
		Database: cfg.Db.Name,
	}

	app.db = pg.Connect(&opts)

	if _, err := app.db.Exec("SELECT 1"); err != nil {
		return fmt.Errorf("db.Exec failed: %w", err)
	}

	logger.Log().Infof("Connected to Postgresql database %s on %s", opts.Database, opts.Addr)
	return nil
}

func (app *App) initServer() {
	app.http = server.NewServer(app.Srv)
	app.http.Route()
}
func (app *App) Serve() error {
	go app.http.Serve()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	return nil
}
func (app *App) PrepareDb() error {
	if err := app.Srv.PrepareChimeraCharacters(); err != nil {
		return fmt.Errorf("PrepareChimeraCharacters failed: %w", err)
	}
	return nil
}
