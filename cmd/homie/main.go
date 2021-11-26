package main

import (
	"flag"
	"log"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
	"github.com/SkurkovPavel/Homie/internal/app/homie"
)

var (
	configPath string
)

func init() {

	flag.StringVar(&configPath, "confog-path", "configs/homeService.toml", "path to cnfig file")

}

func main() {

	flag.Parse()

	config := homie.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	//logger := logrus.New()
	//router := mux.NewRouter()
	//dbConfig := storage.NewConfig()

	service := homie.NewService(
		config,
		logrus.New(),
		mux.NewRouter(),
	)

	if err := service.Start(); err != nil {
		log.Fatal(err)
	}
}
