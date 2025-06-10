package main

import (
	"github.com/injoyai/conv/cfg"
	"github.com/injoyai/tool/forward"
)

func main() {
	forward.Run(cfg.GetInt("port"))
}
