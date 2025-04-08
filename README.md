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
