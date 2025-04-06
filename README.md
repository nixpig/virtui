# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt. 

## 🎯 Initial project goals

Create a domain with CPU, memory, network and disk, from an ISO.


## Notes

1. Install packages: `pacman -Sy fuse3 libvirt radvd qemu-base virt-manager dnsmasq`
1. Enable and start: `sudo systemctl enable --now libvirtd.service`
1. Uncomment in: `/etc/libvirt/libvirtd.conf`
    - `unix_sock_group = "libvirt"`
    - `unix_sock_rw_perms = "0770"`
1. Add user to group: `sudo usermod -aG libvirt $USER`

