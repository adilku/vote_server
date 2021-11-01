package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/adilku/vote_server/internal/app/walletserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/walletserver.toml", "path to config file")
}

func main()  {
	flag.Parse()

	config := walletserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := walletserver.Start(config); err != nil {
		log.Fatal(err)
	}
	//log.Println("Server started at ", config.BindAddr)
}


