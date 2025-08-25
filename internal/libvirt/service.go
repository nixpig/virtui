package libvirt

import (
	"fmt"
	"log"

	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

type Service interface {
	ConnectionDetails() (ConnectionDetails, error)
	ListAllDomains() ([]Domain, error)
	DomainState(uuid string) (string, error)
	DomainMemoryStats(uuid string) (DomainMemoryStats, error)
	DomainDiskStats(uuid string) ([]DomainDiskStats, error)
	DomainInterfaceStats(uuid string) ([]DomainInterfaceStats, error)
	DomainCPUStats(uuid string) (DomainCPUStats, error)
	DomainBlockJobInfo(uuid string) (DomainBlockJobInfo, error)
	DomainXML(uuid string) (string, error)
	DomainCreate(xml string) error
	DomainDefine(xml string) error
	DomainStart(uuid string) error
	DomainShutdown(uuid string) error
	DomainReboot(uuid string) error
	DomainDestroy(uuid string) error
	DomainSuspend(uuid string) error
	DomainResume(uuid string) error
	DomainUndefine(uuid string) error
	DomainMigrate(uuid string, destURI string) error
	DomainEventLifecycleRegister(
		cb func(DomainEvent),
	) (int, error)
	DomainEventLifecycleDeregister(callbackID int) error
}

type ConnectionDetails struct {
	Hostname   string
	LibVersion string
	Version    string
	Type       string
	URI        string
}

type DomainMemoryStats struct {
	Total  uint64
	Used   uint64
	Actual uint64
}

type DomainDiskStats struct {
	Device string
	Read   uint64
	Write  uint64
}

type DomainInterfaceStats struct {
	Device  string
	RxBytes uint64
	TxBytes uint64
}

type DomainCPUStats struct {
	Time   uint64
	System uint64
	User   uint64
}

type DomainBlockJobInfo struct {
	Type      libvirt.DomainBlockJobType
	Bandwidth uint64
	Cur       uint64
	End       uint64
}

type service struct {
	conn Connection
}

func NewService(conn Connection) Service {
	return &service{conn: conn}
}

func (s *service) ConnectionDetails() (ConnectionDetails, error) {
	if s.conn == nil {
		return ConnectionDetails{}, fmt.Errorf("not connected to libvirt")
	}

	hostname, err := s.conn.GetHostname()
	if err != nil {
		return ConnectionDetails{}, fmt.Errorf(
			"failed to get hostname: %w",
			err,
		)
	}

	libVersion, err := s.conn.GetLibVersion()
	if err != nil {
		return ConnectionDetails{}, fmt.Errorf(
			"failed to get libvirt version: %w",
			err,
		)
	}

	version, err := s.conn.GetVersion()
	if err != nil {
		return ConnectionDetails{}, fmt.Errorf("failed to get version: %w", err)
	}

	connType, err := s.conn.GetType()
	if err != nil {
		return ConnectionDetails{}, fmt.Errorf("failed to get type: %w", err)
	}

	uri, err := s.conn.GetURI()
	if err != nil {
		return ConnectionDetails{}, fmt.Errorf("failed to get URI: %w", err)
	}

	return ConnectionDetails{
		Hostname:   hostname,
		LibVersion: Version(libVersion),
		Version:    Version(version),
		Type:       connType,
		URI:        uri,
	}, nil
}

func (s *service) ListAllDomains() ([]Domain, error) {
	if s.conn == nil {
		return nil, fmt.Errorf("not connected to libvirt")
	}

	domains, err := s.conn.ListAllDomains(
		libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list all domains: %w", err)
	}

	var result []Domain
	for _, d := range domains {
		domain, err := ToDomainStruct(&d)
		if err != nil {
			log.Printf("failed to convert domain to struct: %v", err)
			continue
		}
		result = append(result, domain)
	}

	return result, nil
}

func (s *service) DomainState(uuid string) (string, error) {
	if s.conn == nil {
		return "", fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return "", fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	state, _, err := domain.GetState()
	if err != nil {
		return "", fmt.Errorf("failed to get domain state: %w", err)
	}

	return FromState(state), nil
}

func (s *service) DomainMemoryStats(uuid string) (DomainMemoryStats, error) {
	if s.conn == nil {
		return DomainMemoryStats{}, fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return DomainMemoryStats{}, fmt.Errorf(
			"failed to lookup domain by UUID: %w",
			err,
		)
	}
	defer domain.Free()

	stats, err := domain.MemoryStats(uint32(libvirt.DOMAIN_MEMORY_STAT_NR), 0)
	if err != nil {
		return DomainMemoryStats{}, fmt.Errorf(
			"failed to get domain memory stats: %w",
			err,
		)
	}

	var total, used, actual uint64
	for _, stat := range stats {
		switch libvirt.DomainMemoryStatTags(stat.Tag) {
		case libvirt.DOMAIN_MEMORY_STAT_UNUSED:
			used = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_AVAILABLE:
			total = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_ACTUAL_BALLOON:
			actual = stat.Val
		}
	}

	return DomainMemoryStats{
		Total:  total,
		Used:   total - used,
		Actual: actual,
	}, nil
}

func (s *service) DomainDiskStats(uuid string) ([]DomainDiskStats, error) {
	if s.conn == nil {
		return nil, fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	xmlDesc, err := domain.GetXMLDesc(0)
	if err != nil {
		return nil, fmt.Errorf("failed to get domain XML description: %w", err)
	}

	var domainXML libvirtxml.Domain
	if err := domainXML.Unmarshal(xmlDesc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal domain XML: %w", err)
	}

	var diskStats []DomainDiskStats
	for _, disk := range domainXML.Devices.Disks {
		if disk.Source != nil && disk.Target != nil {
			stats, err := domain.BlockStats(disk.Target.Dev)
			if err != nil {
				log.Printf(
					"failed to get block stats for device %s: %v",
					disk.Target.Dev,
					err,
				)
				continue
			}
			diskStats = append(diskStats, DomainDiskStats{
				Device: disk.Target.Dev,
				Read:   uint64(stats.RdBytes),
				Write:  uint64(stats.WrBytes),
			})
		}
	}

	return diskStats, nil
}

func (s *service) DomainInterfaceStats(
	uuid string,
) ([]DomainInterfaceStats, error) {
	if s.conn == nil {
		return nil, fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	xmlDesc, err := domain.GetXMLDesc(0)
	if err != nil {
		return nil, fmt.Errorf("failed to get domain XML description: %w", err)
	}

	var domainXML libvirtxml.Domain
	if err := domainXML.Unmarshal(xmlDesc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal domain XML: %w", err)
	}

	var interfaceStats []DomainInterfaceStats
	for _, iface := range domainXML.Devices.Interfaces {
		if iface.Target != nil {
			stats, err := domain.InterfaceStats(iface.Target.Dev)
			if err != nil {
				log.Printf(
					"failed to get interface stats for device %s: %v",
					iface.Target.Dev,
					err,
				)
				continue
			}
			interfaceStats = append(interfaceStats, DomainInterfaceStats{
				Device:  iface.Target.Dev,
				RxBytes: uint64(stats.RxBytes),
				TxBytes: uint64(stats.TxBytes),
			})
		}
	}

	return interfaceStats, nil
}

func (s *service) DomainCPUStats(uuid string) (DomainCPUStats, error) {
	if s.conn == nil {
		return DomainCPUStats{}, fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return DomainCPUStats{}, fmt.Errorf(
			"failed to lookup domain by UUID: %w",
			err,
		)
	}
	defer domain.Free()

	stats, err := domain.GetCPUStats(0, 0, 1)
	if err != nil {
		return DomainCPUStats{}, fmt.Errorf(
			"failed to get domain CPU stats: %w",
			err,
		)
	}

	if len(stats) == 0 {
		return DomainCPUStats{}, fmt.Errorf("no CPU stats available")
	}

	return DomainCPUStats{
		Time:   stats[0].CpuTime,
		System: stats[0].SystemTime,
		User:   stats[0].UserTime,
	}, nil
}

func (s *service) DomainBlockJobInfo(uuid string) (DomainBlockJobInfo, error) {
	if s.conn == nil {
		return DomainBlockJobInfo{}, fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return DomainBlockJobInfo{}, fmt.Errorf(
			"failed to lookup domain by UUID: %w",
			err,
		)
	}
	defer domain.Free()

	info, err := domain.GetBlockJobInfo("", 0)
	if err != nil {
		return DomainBlockJobInfo{}, fmt.Errorf(
			"failed to get domain block job info: %w",
			err,
		)
	}

	return DomainBlockJobInfo{
		Type:      info.Type,
		Bandwidth: info.Bandwidth,
		Cur:       info.Cur,
		End:       info.End,
	}, nil
}

func (s *service) DomainXML(uuid string) (string, error) {
	if s.conn == nil {
		return "", fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return "", fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	xml, err := domain.GetXMLDesc(0)
	if err != nil {
		return "", fmt.Errorf("failed to get domain XML description: %w", err)
	}

	return xml, nil
}

func (s *service) DomainCreate(xml string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.DefineDomainFlags(
		xml,
		libvirt.DOMAIN_DEFINE_VALIDATE,
	)
	if err != nil {
		return fmt.Errorf("failed to define domain: %w", err)
	}
	defer domain.Free()

	if err := domain.Create(); err != nil {
		return fmt.Errorf("failed to create domain: %w", err)
	}

	return nil
}

func (s *service) DomainDefine(xml string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	_, err := s.conn.DefineDomainFlags(
		xml,
		libvirt.DOMAIN_DEFINE_VALIDATE,
	)
	if err != nil {
		return fmt.Errorf("failed to define domain: %w", err)
	}

	return nil
}

func (s *service) DomainStart(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Create(); err != nil {
		return fmt.Errorf("failed to start domain: %w", err)
	}

	return nil
}

func (s *service) DomainShutdown(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown domain: %w", err)
	}

	return nil
}

func (s *service) DomainReboot(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT); err != nil {
		return fmt.Errorf("failed to reboot domain: %w", err)
	}

	return nil
}

func (s *service) DomainDestroy(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Destroy(); err != nil {
		return fmt.Errorf("failed to destroy domain: %w", err)
	}

	return nil
}

func (s *service) DomainSuspend(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Suspend(); err != nil {
		return fmt.Errorf("failed to suspend domain: %w", err)
	}

	return nil
}

func (s *service) DomainResume(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Resume(); err != nil {
		return fmt.Errorf("failed to resume domain: %w", err)
	}

	return nil
}

func (s *service) DomainUndefine(uuid string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if err := domain.Undefine(); err != nil {
		return fmt.Errorf("failed to undefine domain: %w", err)
	}

	return nil
}

func (s *service) DomainMigrate(uuid string, destURI string) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return fmt.Errorf("failed to lookup domain by UUID: %w", err)
	}
	defer domain.Free()

	if _, err := domain.Migrate(s.conn.(*connection).Connect, 0, destURI, "", 0); err != nil {
		return fmt.Errorf("failed to migrate domain: %w", err)
	}

	return nil
}

func (s *service) DomainEventLifecycleRegister(
	cb func(DomainEvent),
) (int, error) {
	if s.conn == nil {
		return 0, fmt.Errorf("not connected to libvirt")
	}

	callbackID, err := s.conn.DomainEventLifecycleRegister(cb)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to register domain event lifecycle callback: %w",
			err,
		)
	}

	return callbackID, nil
}

func (s *service) DomainEventLifecycleDeregister(callbackID int) error {
	if s.conn == nil {
		return fmt.Errorf("not connected to libvirt")
	}

	if err := s.conn.DomainEventLifecycleDeregister(callbackID); err != nil {
		return fmt.Errorf(
			"failed to deregister domain event lifecycle callback: %w",
			err,
		)
	}

	return nil
}
