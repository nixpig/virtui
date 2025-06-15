# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt.

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


## 🤝 Contributing


## 📦️ Dependencies

You'll need a working [libvirt](https://libvirt.org/) installation to make use of this software. Below are instructions to get to that on Arch; I'm sure you can figure it out for other distros.

```
pacman -Sy fuse3 libvirt radvd qemu-base virt-manager dnsmasq
```

```
sudo systemctl enable --now libvirtd.service
```

```
sudo usermod -aG libvirt $USER
```

```
# /etc/libvirt/libvirtd.conf
unix_sock_group = "libvirt"
unix_sock_rw_perms = "0770"
```

```
# /etc/libvirt/network.conf
firewall_backend=iptables
```

If you want to use LXC, you'll also need: `yay -Sy virt-bootstrap-git`

---

`sudo virt-install --print-xml --connect qemu:///system --name test-vm --ram 2048 --vcpus 2 --disk path=/var/lib/libvirt/images/linux2022.qcow2 --location /var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso,kernel=casper/vmlinuz,initrd=casper/initrd --os-variant ubuntu24.04 --graphics none --extra-args='console=ttyS0'`

---
