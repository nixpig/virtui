# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt. 

## 🎯 Initial project goals

Create a domain with CPU, memory, network and disk, from an ISO.


## Notes

1. Install packages: `pacman -Sy fuse3 libvirt radvd qemu-base virt-manager dnsmasq`
1. Enable and start service: `sudo systemctl enable --now libvirtd.service`
1. Uncomment in: `/etc/libvirt/libvirtd.conf`
    - `unix_sock_group = "libvirt"`
    - `unix_sock_rw_perms = "0770"`
1. Set `firewall_backend=iptables` in `/etc/libvirt/network.conf`
1. Add user to group: `sudo usermod -aG libvirt $USER`
1. Enable polling and stats for memory etc... via virt-manager.


## Commands from list

- Select domain - ↓/↑ j/k
- Open          - O
- Run           - ENTER
    - Pause         - P
    - Shutdown      - S
    - Reboot        - R
    - Force Reset   - O
    - Force Off     - F
    - Save          - V
    - Close menu    - ESC
- Migrate       - M
- Delete        - D
- Clone         - C
- Help          - H
- New           - N
- Quit          - Q/ESC

