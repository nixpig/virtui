# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt.

## 🎯 Initial project goals

- [x] Connect to `qemu:///system` hypervisor
- [ ] Create a filesystem directory storage pool
- [ ] Create storage volume in storage pool
- [ ] Create virtual network and connect to bridge
- [ ] Create an Ubuntu VM from ISO (default settings)
- [ ] Boot VM + connect console to host stdio

## 👑 Longer-term goals

- [ ] Feature parity with GUI tools like <a href="https://virt-manager.org/" target="_blank">virt-manager</a>

## Notes

1. Install packages: `pacman -Sy fuse3 libvirt radvd qemu-base virt-manager dnsmasq`
1. Enable and start service: `sudo systemctl enable --now libvirtd.service`
1. Uncomment in: `/etc/libvirt/libvirtd.conf`
   - `unix_sock_group = "libvirt"`
   - `unix_sock_rw_perms = "0770"`
1. Set `firewall_backend=iptables` in `/etc/libvirt/network.conf`
1. Add user to group: `sudo usermod -aG libvirt $USER`
1. Enable polling and stats for memory etc... via virt-manager.

For LXC OS directory tree creation - `yay -Sy virt-bootstrap-git`

## Commands from list

- Select domain - ↓/↑ j/k

- Open - ENTER

  - xxx

- Start - T
- Pause/Resume - P
- Shutdown - S
- Reboot - R
- Reset - O
- Off - F
- Save/Restore - V

- Migrate - M

  - Close wizard - ESC

- Delete - D

  - Confirm

- Clone - C

  - Close wizard - ESC

- Help - H/?

  - Close help - ESC

- New - N

  - Close wizard - ESC

- Quit - CTRL+C

## States

libvirt.DOMAIN_NOSTATE: "None"
libvirt.DOMAIN_RUNNING: "Running"
libvirt.DOMAIN_BLOCKED: "Blocked"
libvirt.DOMAIN_PAUSED: "Paused"
libvirt.DOMAIN_SHUTDOWN: "Shutdown"
libvirt.DOMAIN_CRASHED: "Crashed"
libvirt.DOMAIN_PMSUSPENDED: "Suspended"
libvirt.DOMAIN_SHUTOFF: "Shutoff"

If it's `x` then it can be `yz`:

Running -> Pause, Save, Migrate, Delete, Reboot, Shutdown, Force Reset, Force Off
Paused -> Resume, Save, Migrate, Delete, Reboot, Shutdown, Force Reset, Force Off

Shutoff -> Run, Clone, Delete

Saved -> Restore, Clone, Delete

---

`sudo virt-install --print-xml --connect qemu:///system --name test-vm --ram 2048 --vcpus 2 --disk path=/var/lib/libvirt/images/linux2022.qcow2 --location /var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso,kernel=casper/vmlinuz,initrd=casper/initrd --os-variant ubuntu24.04 --graphics none --extra-args='console=ttyS0'`

---

### Connect to QEMU

```go
	uri, err := url.Parse("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := libvirt.ConnectToURI(uri)
	defer conn.ConnectClose()
```

### Create storage volume

```go
	vol := volume.NewWithDefaults("default-vm.qcow2")
	volXML, err := vol.ToXML()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := conn.StoragePoolLookupByName("default")
	if err != nil {
		log.Fatal("get storage pool: ", err.Error())
	}

	if _, err := conn.StorageVolCreateXML(pool, string(volXML), 0); err != nil {
		log.Fatal("create storage volume: " + err.Error())
	}
```

### Create and start network

```go
	vnet := network.NewWithDefaults("default1")
	vnetXML, err := vnet.ToXML()
	if err != nil {
		log.Fatal("network to xml: " + err.Error())
	}

	if n, err := conn.NetworkDefineXML(string(vnetXML)); err != nil {
		log.Fatal("define network: " + err.Error())
	} else {
		if err := conn.NetworkCreate(n); err != nil {
			log.Fatal("start network: " + err.Error())
		}
	}

```

### Create and start VM

```go
	dom := domain.NewWithDefaults("network-test-vm")

	domXML, err := dom.ToXML()
	if err != nil {
		log.Fatal("domain to xml: " + err.Error())
	}

	if d, err := conn.DomainDefineXML(string(domXML)); err != nil {
		log.Fatal("define domain: " + err.Error())
	} else {
		if err := conn.DomainCreate(d); err != nil {
			log.Fatal("start domain: " + err.Error())
		}
	}
```
