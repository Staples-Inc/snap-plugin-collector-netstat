package main

import (
	"os"

	"github.com/intelsdi-x/snap-plugin-collector-netstat/netstat"
	"github.com/intelsdi-x/snap/control/plugin"
)

func main() {
	plugin.Start(netstat.Meta(), netstat.New(), os.Args[1])
}
