# 🖥️ virtui

An interactive Terminal UI (TUI) for managing QEMU/KVM virtual machines via libvirt. 

## 🎯 Initial project goals

- [ ] Connect to default `qemu:///system` hypervisor
- [ ] Create a volume in file-based storage pool
- [ ] Create a virtual network over bridge device
- [ ] Create an Ubuntu 24.04 VM from ISO image
- [ ] Boot VM + forward console to host stdio

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

- Open          - ENTER
    - xxx

- Start         - T
- Pause/Resume  - P
- Shutdown      - S
- Reboot        - R
- Reset         - O
- Off           - F
- Save/Restore  - V

- Migrate       - M
    - Close wizard  - ESC

- Delete        - D
    - Confirm

- Clone         - C
    - Close wizard  - ESC

- Help          - H/?
    - Close help    - ESC

- New           - N
    - Close wizard  - ESC

- Quit          - Q


## States

libvirt.DOMAIN_NOSTATE:     "None"
libvirt.DOMAIN_RUNNING:     "Running"
libvirt.DOMAIN_BLOCKED:     "Blocked"
libvirt.DOMAIN_PAUSED:      "Paused"
libvirt.DOMAIN_SHUTDOWN:    "Shutdown"
libvirt.DOMAIN_CRASHED:     "Crashed"
libvirt.DOMAIN_PMSUSPENDED: "Suspended"
libvirt.DOMAIN_SHUTOFF:     "Shutoff"

If it's `x` then it can be `yz`: 

Running -> Pause, Save, Migrate, Delete, Reboot, Shutdown, Force Reset, Force Off
Paused -> Resume, Save, Migrate, Delete, Reboot, Shutdown, Force Reset, Force Off

Shutoff -> Run, Clone, Delete

Saved -> Restore, Clone, Delete


---

`sudo virt-install --print-xml --connect qemu:///system --name test-vm --ram 2048 --vcpus 2 --disk path=/var/lib/libvirt/images/linux2022.qcow2 --location /var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso,kernel=casper/vmlinuz,initrd=casper/initrd --os-variant ubuntu24.04 --graphics none --extra-args='console=ttyS0'`


---


