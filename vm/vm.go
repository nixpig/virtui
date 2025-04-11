package vm

import (
	"fmt"

	"github.com/charmbracelet/log"
	"libvirt.org/go/libvirt"
)

var (
	UNKNOWN_STATE_PRESENTABLE   = "Unknown"
	NO_STATE_PRESENTABLE        = "None"
	RUNNING_STATE_PRESENTABLE   = "Running"
	BLOCKED_STATE_PRESENTABLE   = "Blocked"
	PAUSED_STATE_PRESENTABLE    = "Paused"
	SHUTDOWN_STATE_PRESENTABLE  = "Shutdown"
	CRASHED_STATE_PRESENTABLE   = "Crashed"
	SUSPENDED_STATE_PRESENTABLE = "Suspended"
	SHUTOFF_STATE_PRESENTABLE   = "Shutoff"
)

var presentableState = map[libvirt.DomainState]string{
	libvirt.DOMAIN_NOSTATE:     NO_STATE_PRESENTABLE,
	libvirt.DOMAIN_RUNNING:     RUNNING_STATE_PRESENTABLE,
	libvirt.DOMAIN_BLOCKED:     BLOCKED_STATE_PRESENTABLE,
	libvirt.DOMAIN_PAUSED:      PAUSED_STATE_PRESENTABLE,
	libvirt.DOMAIN_SHUTDOWN:    SHUTDOWN_STATE_PRESENTABLE,
	libvirt.DOMAIN_CRASHED:     CRASHED_STATE_PRESENTABLE,
	libvirt.DOMAIN_PMSUSPENDED: SUSPENDED_STATE_PRESENTABLE,
	libvirt.DOMAIN_SHUTOFF:     SHUTOFF_STATE_PRESENTABLE,
}

type VM struct {
	*libvirt.Domain
}

type Event struct {
	VM    *VM
	Event libvirt.DomainEventType
}

func New() *VM {
	return &VM{}
}

func FromDomain(domain *libvirt.Domain) *VM {
	return &VM{domain}
}

func GetAll(conn *libvirt.Connect) []VM {
	flags := libvirt.ConnectListAllDomainsFlags(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	domains, err := conn.ListAllDomains(flags)
	if err != nil {
		log.Error("list all domains", "err", err)
	}

	vms := make([]VM, len(domains))
	for i, d := range domains {
		vms[i] = VM{&d}
	}

	return vms
}

func (v *VM) GetPresentableName() string {
	n, err := v.Domain.GetName()
	if err != nil {
		log.Error("get domain name", "err", err)
	}

	return n
}

func (v *VM) GetPresentableID() string {
	id := "-"
	s, _, err := v.Domain.GetState()
	if err != nil {
		log.Error("get domain state", "err", err)
	}

	if err == nil && s != libvirt.DOMAIN_SHUTOFF {
		u, _ := v.Domain.GetID()
		id = fmt.Sprintf("%d", u)
	}

	return id
}

func (v *VM) GetPresentableUUID() string {
	uuid, err := v.Domain.GetUUIDString()
	if err != nil {
		log.Error("get domain uuid", "err", err)
	}

	return uuid
}

func (v *VM) GetPresentableState() string {
	s, _, err := v.Domain.GetState()
	if err != nil {
		log.Error("get domain state", "err", err)
		return UNKNOWN_STATE_PRESENTABLE
	}

	state, ok := presentableState[s]
	if !ok {
		log.Debug("unknown domain state", "state", s)
		return UNKNOWN_STATE_PRESENTABLE
	}

	return state
}

func (v *VM) Run() error {
	if err := v.Domain.Create(); err != nil {
		return err
	}

	return nil
}

func (v *VM) PauseResume() error {
	s, _, err := v.Domain.GetState()
	if err != nil {
		return err
	}

	if s == libvirt.DOMAIN_PAUSED {
		if err := v.Domain.Resume(); err != nil {
			return err
		}
	} else {
		if err := v.Domain.Suspend(); err != nil {
			return err
		}
	}

	return nil
}

func (v *VM) Shutdown() error {
	if err := v.Domain.Shutdown(); err != nil {
		return err
	}

	return nil
}

func (v *VM) Reboot() error {
	if err := v.Domain.Reboot(0); err != nil {
		return err
	}

	return nil
}

func (v *VM) ForceReset() error {
	if err := v.Domain.Reset(0); err != nil {
		return err
	}

	return nil
}

func (v *VM) SaveRestore() error {
	s, _, err := v.Domain.GetState()
	if err != nil {
		return err
	}

	if s == libvirt.DOMAIN_RUNNING || s == libvirt.DOMAIN_PAUSED {
		if err := v.Domain.ManagedSave(0); err != nil {
			return err
		}
	} else {
		if err := v.Domain.Create(); err != nil {
			return err
		}
	}

	return nil
}

func (v *VM) ForceOff() error {
	return fmt.Errorf("not implemented")
}

func (v *VM) Delete() error {
	if err := v.Domain.Destroy(); err != nil {
		return err
	}

	return nil
}
