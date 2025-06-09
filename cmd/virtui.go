package main

import (
	"fmt"
	"net/url"
	"os"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/api"
	"github.com/nixpig/virtui/vm/pool"
	"github.com/nixpig/virtui/vm/volume"
)

func main() {
	u, err := url.Parse(string(libvirt.QEMUSystem))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 0. Connecto to hypervisor

	conn, err := libvirt.ConnectToURI(u)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	a := api.NewAPI(conn)

	// 1. Get pools

	pools, err := pool.List(conn)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defaultPool := pools[slices.IndexFunc(pools, func(p pool.Pool) bool {
		return p.Name == "default"
	})]

	fmt.Println(" -- 🍱 default pool:")
	fmt.Println(defaultPool)
	fmt.Println("")

	// 2. Get volumes

	v, err := defaultPool.GetVolumes(conn)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Println(" -- 💽 volumes:")
	for _, x := range v {
		fmt.Printf("%+v\n", x)
	}
	fmt.Println("")

	// 3. Create volume in pool

	fmt.Println(" -- 🚀 create volume")
	vol := volume.NewWithDefaults("FOO-VOLUME")

	vx, err := a.CreateVolume(vol, &defaultPool)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Println("created:", vx)

	// 4. Get networks

	// 5. Create network and attach to bridge

	// 6. Get VMs

	// 7. Create VM

	// 8. Start VM

}
