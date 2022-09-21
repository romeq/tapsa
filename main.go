package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/romeq/usva/api"
	"github.com/romeq/usva/arguments"
	"github.com/romeq/usva/config"
	"github.com/romeq/usva/dbengine"
	"github.com/romeq/usva/utils"
)

type Options config.Config

func parseOpts(cfg config.Config, args arguments.Arguments) Options {
	return Options{
		Server: config.Server{
			Address: utils.StringOr(args.Config.Server.Address, cfg.Server.Address),
			Port:    utils.IntOr(args.Config.Server.Port, cfg.Server.Port),
		},
		Database: config.Database{
			Database: utils.StringOr(args.Config.Database.Database, cfg.Database.Database),
			Host:     utils.StringOr(args.Config.Database.Host, cfg.Database.Host),
			Port:     utils.IntOr(args.Config.Database.Port, cfg.Database.Port),
			User:     utils.StringOr(args.Config.Database.User, cfg.Database.User),
			Password: utils.StringOr(args.Config.Database.Password, cfg.Database.Password),
		},
	}
}

func (o *Options) getaddr() string {
	return fmt.Sprintf("%s:%d", o.Server.Address, o.Server.Port)
}

func setuprouter(cfg config.Config) *gin.Engine {
	if !cfg.Server.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Server.AllowedOrigins,
		AllowMethods: []string{"POST", "GET", "DELETE"},
		AllowHeaders: []string{"Authorization"},
	}))

	r.Use(gin.Recovery())
	if !cfg.Server.HideRequests {
		r.Use(requestLogger)
	}

	api.SetupRoutes(r, &cfg)
	utils.Check(r.SetTrustedProxies(cfg.Server.TrustedProxies))

	return r
}

func setuplogger() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

func setLogWriter(file string) *os.File {
	if file == "" {
		return nil
	}

	fhandle, err := os.OpenFile(file, os.O_WRONLY, 0644)
	utils.Check(err)

	log.SetOutput(fhandle)
	return fhandle
}

func requestLogger(ctx *gin.Context) {
	t := time.Now()
	ctx.Next()
	c := time.Since(t).Milliseconds()

	log.Printf("request: [%s] %s %s %d (took %dms) \n",
		ctx.RemoteIP(), ctx.Request.Method, ctx.Request.URL, ctx.Writer.Status(), c)
}

func main() {
	setuplogger()

	// arguments
	args := arguments.Parse()
	defer setLogWriter(args.LogOutput).Close()

	// config file
	cfgHandle, err := os.Open(args.ConfigFile)
	utils.Check(err)
	cfg := config.ParseFromFile(cfgHandle)

	// runtime options
	opts := parseOpts(cfg, args)
	dbengine.Init(
		uint16(opts.Database.Port),
		opts.Database.Host,
		opts.Database.Database,
		opts.Database.User,
		opts.Database.Password,
	)
	defer dbengine.DbConnection.Close()

	log.Println("Starting server at", opts.getaddr())

	// start server
	r := setuprouter(cfg)
	tlssettings := cfg.Server.TLS
	if tlssettings.Enabled {
		utils.Check(r.RunTLS(opts.getaddr(), tlssettings.KeyFile, tlssettings.CertFile))
	} else {
		utils.Check(r.Run(opts.getaddr()))
	}
}
