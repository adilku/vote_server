package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/adilku/vote_server/internal/app/voteserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/voteserver.toml", "path to config file")
}

func main()  {

	flag.Parse()

	config := voteserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := voteserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}


