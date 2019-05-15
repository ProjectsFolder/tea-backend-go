package main

import "github.com/op/go-logging"
import conf "./config"

var log = logging.MustGetLogger("tea")

func main() {
	config := conf.ReadConfig()
}
