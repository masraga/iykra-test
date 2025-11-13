package main

import (
	"github.com/masraga/iykra-test/cmd/server"
	"github.com/masraga/iykra-test/config"
	"github.com/spf13/viper"
)

func init(){
	config.LoadConfig()
}

func main() {
	host := viper.GetString("HOST")
	port := viper.GetInt("PORT")
	server := server.NewServer(host, port)
	server.Start()
}