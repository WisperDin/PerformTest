package main

import (
	."./module"
	"./conf"
)

func main(){
	conf.Init("./test.toml")
	Login("59991979-2","123456","admin","web")
}
