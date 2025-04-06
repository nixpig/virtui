package main

import (
	"fmt"
	"os"

	"libvirt.org/go/libvirt"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("new connection: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("connected")

	n, _ := conn.NumOfDomains()
	fmt.Println("num domains: ", n)

	d, _ := conn.LookupDomainById(1)
	i, _ := d.GetUUIDString()
	fmt.Println(i)

	// doms, err := conn.GetAllDomainStats()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(doms)
}
