# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt.

[screenshot]

Why...?

### 🎯 Initial project goals

- [x] Connect to `qemu:///system` hypervisor
- [ ] Create a filesystem directory storage pool
- [ ] Create storage volume in storage pool
- [ ] Create virtual network and connect to bridge
- [ ] Create an Ubuntu VM from ISO (default settings)
- [ ] Boot VM + connect console to host stdio

### 👑 Longer-term goals

- [ ] Feature parity with GUI tools like <a href="https://virt-manager.org/" target="_blank">virt-manager</a>


## 🚀 Quick start


## 👩‍💻 Usage


## 📦️ Dependencies

You'll need a working [libvirt](https://libvirt.org/) installation to make use of this software. Below are instructions to get that on Arch (feel free to submit a PR for instructions for other distros!).

### Arch

1. Install dependencies
```
pacman -Sy fuse3 libvirt radvd qemu-base virt-manager dnsmasq
```

2. Enable the `libvirtd` service
```
sudo systemctl enable --now libvirtd.service
```

3. Add user to the `libvirt` group
```
sudo usermod -aG libvirt $USER
```

4. Configure socket group owner and permissions
```sh
# /etc/libvirt/libvirtd.conf
unix_sock_group = "libvirt"
unix_sock_rw_perms = "0770"
```

5. Configure backend to work with `iptables-nft`
```sh
# /etc/libvirt/network.conf
firewall_backend = "iptables"
```

If you want to use LXC, you'll also need to install `virt-bootstrap-git` from the AUR.



## 🤝 Contributing

