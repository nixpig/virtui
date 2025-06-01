package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/api"
)

func main() {
	u, err := url.Parse(string(libvirt.QEMUSystem))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	conn, err := libvirt.ConnectToURI(u)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	a := api.NewAPI(conn)

	pools, err := a.GetPools()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	for _, p := range pools {
		fmt.Printf("%+v\n", p)
	}

}
