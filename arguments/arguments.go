package arguments

import (
	"flag"

	"github.com/romeq/usva/config"
)

type Arguments struct {
	Config     config.Config
	ConfigFile string
	LogOutput  string
}

func Parse() Arguments {
	args := Arguments{}

	flag.StringVar(&args.Config.Server.Address, "a", "127.0.0.1", "server address")
	flag.StringVar(&args.ConfigFile, "c", "/etc/usva/tapsa.toml", "config location")
	flag.StringVar(&args.LogOutput, "l", "", "log location")
	flag.IntVar(&args.Config.Server.Port, "p", 8080, "server port")

	flag.Parse()
	return args
}
