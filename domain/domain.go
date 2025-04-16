package domain

import "encoding/xml"

type Acceleration struct {
	Accel2d string `xml:"accel2d,attr"`
	Accel3d string `xml:"accel3d,attr"`
}

type Access struct {
	Mode string `xml:"mode,attr"`
}

type ACPI struct {
	Index *int   `xml:"index,attr"`
	Table *Table `xml:"table"`
}

type APIC struct{}

type ActivePCRBanks struct {
	Sha256 struct{} `xml:"sha256"`
}

type Adapter struct {
	Name string `xml:"name,attr"`
}

type Address struct {
	Base          string `xml:"base,attr,omitempty"`
	Bus           string `xml:"bus,attr,omitempty"`
	Controller    string `xml:"controller,attr"`
	CSSID         string `xml:"cssid,attr,omitempty"`
	DevNo         string `xml:"devno,attr,omitempty"`
	Domain        string `xml:"domain,attr,omitempty"`
	Function      string `xml:"function,attr,omitempty"`
	IOBase        int    `xml:"iobase,attr,omitempty"`
	Multifunction string `xml:"multifunction,attr,omitempty"`
	Port          *int   `xml:"port,attr"`
	Reg           string `xml:"reg,attr,omitempty"`
	Slot          string `xml:"slot,attr,omitempty"`
	SSID          string `xml:"ssid,attr,omitempty"`
	Target        string `xml:"target,attr,omitempty"`
	Type          string `xml:"type,attr,omitempty"`
	Unit          string `xml:"unit,attr,omitempty"`
	UUID          string `xml:"uuid,attr,omitempty"`
}

type AIA struct {
	Value string `xml:"value,attr,omitempty"`
}

type Alias struct {
	Name string `xml:"name,attr,omitempty"`
}

type AlignSize struct {
	Unit     string `xml:"unit,attr,omitempty"`
	CharData string `xml:",chardata"`
}

type Allocation struct {
	Mode    string `xml:"mode,attr,omitempty"`
	Threads int    `xml:"threads,attr"`
}

type AsyncTeardown struct {
	Enabled string `xml:"enabled,attr,omitempty"`
}

type Audio struct {
	Driver      string  `xml:"driver,attr,omitempty"`
	DSPPolicy   *int    `xml:"dspPolicy,attr"`
	Exclusive   string  `xml:"exclusive,attr,omitempty"`
	ID          int     `xml:"id,attr"`
	Path        string  `xml:"path,attr,omitempty"`
	RuntimeDir  string  `xml:"runtimeDir,attr,omitempty"`
	ServerName  string  `xml:"serverName,attr,omitempty"`
	TimerPeriod *int    `xml:"timerPeriod,attr"`
	TryMMap     string  `xml:"tryMMap,attr,omitempty"`
	Type        string  `xml:"type,attr,omitempty"`
	Input       *Input  `xml:"input"`
	Output      *Output `xml:"output"`
}

type Auth struct {
	Username string `xml:"username,attr"`
	Secret   Secret `xml:"secret"`
}

type Backend struct {
	Debug          *int            `xml:"debug,attr"`
	LogFile        string          `xml:"logFile,attr,omitempty"`
	Model          string          `xml:"model,attr,omitempty"`
	Queues         *int            `xml:"queues,attr"`
	Tap            string          `xml:"tap,attr,omitempty"`
	Type           string          `xml:"type,attr,omitempty"`
	Version        *float64        `xml:"version,attr"`
	VHost          string          `xml:"vhost,attr,omitempty"`
	CharData       string          `xml:",chardata"`
	ActivePcrBanks *ActivePCRBanks `xml:"active_pcr_banks"`
	Device         *Device         `xml:"device"`
	Encryption     *Encryption     `xml:"encryption"`
	Profile        *Profile        `xml:"profile"`
	Source         *Source         `xml:"source"`
}

type BackendDomain struct {
	Name string `xml:"name,attr,omitempty"`
}

type BackingStore struct {
	Type         string        `xml:"type,attr,omitempty"`
	BackingStore *BackingStore `xml:"backingStore"`
	Format       *Format       `xml:"format"`
	Source       *Source       `xml:"source"`
}

type Bandwidth struct {
	Initiator *int     `xml:"initiator,attr"`
	Target    *int     `xml:"target,attr"`
	Type      string   `xml:"type,attr,omitempty"`
	Unit      string   `xml:"unit,attr,omitempty"`
	Value     *int     `xml:"value,attr"`
	Inbound   Inbound  `xml:"inbound"`
	Outbound  Outbound `xml:"outbound"`
}

type BaseBoard struct {
	Entries []Entry `xml:"entry"`
}

type Binary struct {
	Path       string     `xml:"path,attr,omitempty"`
	XAttr      string     `xml:"xattr,attr,omitempty"`
	Cache      Cache      `xml:"cache"`
	Lock       Lock       `xml:"lock"`
	Sandbox    Sandbox    `xml:"sandbox"`
	ThreadPool ThreadPool `xml:"thread_pool"`
}

type BIOS struct {
	RebootTimeout *int    `xml:"rebootTimeout,attr"`
	UseSerial     *string `xml:"useserial,attr"`
	Entry         Entry   `xml:"entry"`
}

type BlkioTune struct {
	Device *Device `xml:"device"`
	Weight int     `xml:"weight"`
}

type Block struct {
	Unit     *string `xml:"unit,attr"`
	CharData string  `xml:",chardata"`
}

type BlockIO struct {
	DiscardGranularity int `xml:"discard_granularity,attr"`
	LogicalBlockSize   int `xml:"logical_block_size,attr"`
	PhysicalBlockSize  int `xml:"physical_block_size,attr"`
}

type Boot struct {
	Dev   string `xml:"dev,attr,omitempty"`
	Order *int   `xml:"order,attr"`
}

type BootMenu struct {
	Enable  string `xml:"enable,attr"`
	Timeout int    `xml:"timeout,attr"`
}

type Cache struct {
	Associativity *string `xml:"associativity,attr"`
	ID            *int    `xml:"id,attr"`
	Level         *int    `xml:"level,attr"`
	Mode          *string `xml:"mode,attr"`
	Policy        *string `xml:"policy,attr"`
	Type          *string `xml:"type,attr"`
	Unit          *string `xml:"unit,attr"`
	Line          *Line   `xml:"line"`
	Size          *Size   `xml:"size"`
	SizeAttr      *int    `xml:"size,attr"`
}

type CacheTune struct {
	VCPUs   string  `xml:"vcpus,attr"`
	Cache   Cache   `xml:"cache"`
	Monitor Monitor `xml:"monitor"`
}

type Catchup struct {
	Limit     int `xml:"limit,attr"`
	Slew      int `xml:"slew,attr"`
	Threshold int `xml:"threshold,attr"`
}

type CCFAssist struct {
	State string `xml:"state,attr"`
}

type Cell struct {
	CPUs      string     `xml:"cpus,attr"`
	Discard   *string    `xml:"discard,attr"`
	ID        int        `xml:"id,attr"`
	MemAccess *string    `xml:"memAccess,attr"`
	Memory    int        `xml:"memory,attr"`
	Unit      string     `xml:"unit,attr"`
	Cache     *Cache     `xml:"cache"`
	Distances *Distances `xml:"distances"`
}

type CFPC struct {
	Value string `xml:"value,attr"`
}

type Channel struct {
	Mode    string   `xml:"mode,attr,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
	Source  *Source  `xml:"source"`
	Target  *Target  `xml:"target"`
	Address *Address `xml:"address"`
}

type Chassis struct {
	Entry []Entry `xml:"entry"`
}

type CID struct {
	Address int    `xml:"address,attr"`
	Auto    string `xml:"auto,attr"`
}

type Cipher struct {
	Name  string `xml:"name,attr"`
	State string `xml:"state,attr"`
}

type Clipboard struct {
	CopyPaste string `xml:"copypaste,attr"`
}

type Clock struct {
	Offset string  `xml:"offset,attr,omitempty"`
	Sync   string  `xml:"sync,attr,omitempty"`
	Timers []Timer `xml:"timer"`
}

type Coalesce struct {
	Rx Rx `xml:"rx"`
}

type Codec struct {
	Type string `xml:"type,attr"`
}

type Config struct {
	File string `xml:"file,attr"`
}

type Console struct {
	Type   string  `xml:"type,attr"`
	Log    *Log    `xml:"log"`
	Source *Source `xml:"source"`
	Target *Target `xml:"target"`
}

type Controller struct {
	Index            int      `xml:"index,attr"`
	MaxEventChannels *int     `xml:"maxEventChannels,attr"`
	MaxGrantFrames   *int     `xml:"maxGrantFrames,attr"`
	Model            string   `xml:"model,attr,omitempty"`
	Ports            *int     `xml:"ports,attr"`
	Type             string   `xml:"type,attr,omitempty"`
	Vectors          *int     `xml:"vectors,attr"`
	Address          *Address `xml:"address"`
	Driver           *Driver  `xml:"driver"`
	Master           *Master  `xml:"master"`
	Target           *Target  `xml:"target"`
}

type Cookie struct {
	Name     string `xml:"name,attr"`
	CharData string `xml:",chardata"`
}

type Cookies struct {
	Cookie Cookie `xml:"cookie"`
}

type CPU struct {
	Match       string       `xml:"match,attr,omitempty"`
	Migratable  string       `xml:"migratable,attr,omitempty"`
	Mode        string       `xml:"mode,attr,omitempty"`
	Cache       *Cache       `xml:"cache"`
	Check       string       `xml:"check,attr,omitempty"`
	Feature     *Feature     `xml:"feature"`
	Maxphysaddr *MaxPhysAddr `xml:"maxphysaddr"`
	Model       *Model       `xml:"model"`
	Numa        *Numa        `xml:"numa"`
	Topology    *Topology    `xml:"topology"`
	Vendor      *Vendor      `xml:"vendor"`
}

type CPUTune struct {
	CacheTune      []CacheTune   `xml:"cachetune"`
	EmulatorPeriod int           `xml:"emulator_period"`
	EmulatorQuota  int           `xml:"emulator_quota"`
	EmulatorPIN    EmulatorPIN   `xml:"emulatorpin"`
	GlobalPeriod   int           `xml:"global_period"`
	GlobalQuota    int           `xml:"global_quota"`
	IothreadPeriod int           `xml:"iothread_period"`
	IothreadQuota  int           `xml:"iothread_quota"`
	IOThreadPIN    []IOThreadPIN `xml:"iothreadpin"`
	IOThreadSched  IOThreadSched `xml:"iothreadsched"`
	MemoryTune     MemoryTune    `xml:"memorytune"`
	Period         int           `xml:"period"`
	Quota          int           `xml:"quota"`
	Shares         int           `xml:"shares"`
	VCPUPins       []VCPIPin     `xml:"vcpupin"`
	VCPUSched      VCPUSched     `xml:"vcpusched"`
}

type Crypto struct {
	Model   string  `xml:"model,attr"`
	Type    string  `xml:"type,attr"`
	Backend Backend `xml:"backend"`
}

type Current struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type CurrentMemory struct {
	Unit     string `xml:"unit,attr,omitempty"`
	CharData string `xml:",chardata"`
}

type DataStore struct {
	Type   string `xml:"type,attr"`
	Format Format `xml:"format"`
	Source Source `xml:"source"`
}

type DefaultIOThread struct {
	ThreadPoolMax int `xml:"thread_pool_max,attr"`
	ThreadPoolMin int `xml:"thread_pool_min,attr"`
}

type Device struct {
	Alias         *string   `xml:"alias,attr"`
	Frontend      *Frontend `xml:"frontend"`
	Path          string    `xml:"path"`
	ReadBytesSec  int       `xml:"read_bytes_sec,omitzero"`
	ReadIOPSSec   int       `xml:"read_iops_sec,omitzero"`
	Weight        int       `xml:"weight,omitzero"`
	WriteBytesSec int       `xml:"write_bytes_sec,omitzero"`
	WriteIOPSSec  int       `xml:"write_iops_sec,omitzero"`
}

type Devices struct {
	Audio       []Audio      `xml:"audio"`
	Channels    []Channel    `xml:"channel"`
	Consoles    []Console    `xml:"console"`
	Controllers []Controller `xml:"controller"`
	Crypto      *Crypto      `xml:"crypto"`
	Disks       []Disk       `xml:"disk"`
	Emulator    string       `xml:"emulator"`
	Filesystem  []Filesystem `xml:"filesystem"`
	Graphics    []Graphics   `xml:"graphics"`
	HostDev     []HostDev    `xml:"hostdev"`
	Hub         *Hub         `xml:"hub"`
	Input       []Input      `xml:"input"`
	Interfaces  []Interface  `xml:"interface"`
	IOMMU       *IOMMU       `xml:"iommu"`
	Lease       *Lease       `xml:"lease"`
	MemBalloon  *MemBalloon  `xml:"memballoon"`
	Memory      []Memory     `xml:"memory"`
	NVRam       *NVRAM       `xml:"nvram"`
	Panic       []Panic      `xml:"panic"`
	Parallel    []Parallel   `xml:"parallel"`
	Pstore      *PStore      `xml:"pstore"`
	RedirDev    []RedirDev   `xml:"redirdev"`
	RedirFilter *RedirFilter `xml:"redirfilter"`
	RNG         *RNG         `xml:"rng"`
	Serials     []Serial     `xml:"serial"`
	SHMem       []SHMem      `xml:"shmem"`
	Smartcard   []Smartcard  `xml:"smartcard"`
	Sound       []Sound      `xml:"sound"`
	TPM         []Tpm        `xml:"tpm"`
	Video       *Video       `xml:"video"`
	VSock       *VSock       `xml:"vsock"`
	Watchdog    *Watchdog    `xml:"watchdog"`
}

type Direct struct {
	State string `xml:"state,attr"`
}

type DirtyRing struct {
	Size  int    `xml:"size,attr"`
	State string `xml:"state,attr"`
}

type Disk struct {
	Device          string           `xml:"device,attr"`
	Snapshot        *string          `xml:"snapshot,attr"`
	Type            string           `xml:"type,attr"`
	Address         *Address         `xml:"address"`
	Alias           *Alias           `xml:"alias"`
	BackingStore    *BackingStore    `xml:"backingStore"`
	BlockIO         *BlockIO         `xml:"blockio"`
	Boot            *Boot            `xml:"boot"`
	Driver          *Driver          `xml:"driver"`
	Encryption      *Encryption      `xml:"encryption"`
	Geometry        *Geometry        `xml:"geometry"`
	IOTune          *IOTune          `xml:"iotune"`
	Readonly        *Readonly        `xml:"readonly"`
	Serial          *Serial          `xml:"serial"`
	Shareable       *struct{}        `xml:"shareable"`
	Source          *Source          `xml:"source"`
	Target          *Target          `xml:"target"`
	ThrottleFilters *ThrottleFilters `xml:"throttlefilters"`
	Transient       *struct{}        `xml:"transient"`
}

type Readonly struct{}

type Distances struct {
	Siblings []Sibling `xml:"sibling"`
}

type Domain struct {
	XMLName         xml.Name         `xml:"domain"`
	BlkioTune       *BlkioTune       `xml:"blkiotune"`
	Clock           *Clock           `xml:"clock"`
	Cpu             *CPU             `xml:"cpu"`
	CPUTune         *CPUTune         `xml:"cputune"`
	CurrentMemory   *CurrentMemory   `xml:"currentMemory"`
	DefaultIOThread *DefaultIOThread `xml:"defaultiothread"`
	Description     *string          `xml:"description"`
	Devices         *Devices         `xml:"devices"`
	Features        *Features        `xml:"features"`
	GenID           *string          `xml:"genid"`
	IOThreadIDs     *IOThreadIDs     `xml:"iothreadids"`
	IOThreads       *IOThread        `xml:"iothreads"`
	Keywrap         *Keywrap         `xml:"keywrap"`
	LaunchSecurity  *LaunchSecurity  `xml:"launchSecurity"`
	MaxMemory       *MaxMemory       `xml:"maxMemory"`
	Memory          *Memory          `xml:"memory"`
	MemoryBacking   *MemoryBacking   `xml:"memoryBacking"`
	MemTune         *Memtune         `xml:"memtune"`
	Metadata        *Metadata        `xml:"metadata"`
	Name            string           `xml:"name"`
	NumaTune        *Numatune        `xml:"numatune"`
	OnCrash         string           `xml:"on_crash,omitempty"`
	OnLockFailure   *string          `xml:"on_lockfailure"`
	OnPoweroff      string           `xml:"on_poweroff,omitempty"`
	OnReboot        string           `xml:"on_reboot,omitempty"`
	OS              *OS              `xml:"os"`
	Override        *Override        `xml:"override"`
	Perf            *Perf            `xml:"perf"`
	PM              *PM              `xml:"pm"`
	Resources       []Resource       `xml:"resource"`
	SecLabels       []SecLabel       `xml:"seclabel"`
	Sysinfo         []Sysinfo        `xml:"sysinfo"`
	ThrottleGroups  *ThrottleGroups  `xml:"throttlegroups"`
	Title           string           `xml:"title,omitempty"`
	Type            string           `xml:"type,attr,omitempty"`
	UUID            string           `xml:"uuid,omitempty"`
	VCPU            *VCPU            `xml:"vcpu"`
	VCPUs           *VCPUs           `xml:"vcpus"`
}

type DownScript struct {
	Path string `xml:"path,attr"`
}

type Driver struct {
	ATS           *string        `xml:"ats,attr"`
	Cache         *string        `xml:"cache,attr"`
	Discard       string         `xml:"discard,attr,omitempty"`
	EventIDX      *string        `xml:"event_idx,attr"`
	Format        *string        `xml:"format,attr"`
	IntRemap      *string        `xml:"intremap,attr"`
	IO            *string        `xml:"io,attr"`
	IOEventFD     *string        `xml:"ioeventfd,attr"`
	IOMMU         *string        `xml:"iommu,attr"`
	IOThread      *int           `xml:"iothread,attr"`
	Name          string         `xml:"name,attr"`
	Queue         *int           `xml:"queue,attr"`
	QueueSize     *int           `xml:"queue_size,attr"`
	Queues        *int           `xml:"queues,attr"`
	RxQueueSize   *int           `xml:"rx_queue_size,attr"`
	TxQueueSize   *int           `xml:"tx_queue_size,attr"`
	Txmode        *string        `xml:"txmode,attr"`
	Type          string         `xml:"type,attr"`
	WrPolicy      *string        `xml:"wrpolicy,attr"`
	Guest         *Guest         `xml:"guest"`
	Host          *Host          `xml:"host"`
	IOThreads     *IOThread      `xml:"iothreads"`
	MetadataCache *MetadataCache `xml:"metadata_cache"`
}

type E820Host struct {
	State string `xml:"state,attr"`
}

type EMSRBitmap struct {
	State string `xml:"state,attr"`
}

type EmulatorPIN struct {
	CPUSet string `xml:"cpuset,attr"`
}

type Encryption struct {
	Secret   *string `xml:"secret,attr"`
	Type     string  `xml:"type,attr"`
	CharData string  `xml:",chardata"`
}

type Entry struct {
	File     string  `xml:"file,attr"`
	Name     *string `xml:"name,attr"`
	CharData string  `xml:",chardata"`
}

type Event struct {
	Enabled string `xml:"enabled,attr"`
	Name    string `xml:"name,attr"`
}

type EVMCS struct {
	State string `xml:"state,attr"`
}

type Extended struct {
	State string `xml:"state,attr"`
}

type Feature struct {
	Name   string `xml:"name,attr"`
	Policy string `xml:"policy,attr"`
}

type Features struct {
	ACPI          *ACPI          `xml:"acpi"`
	AIA           *AIA           `xml:"aia"`
	APIC          *APIC          `xml:"apic"`
	AsyncTeardown *AsyncTeardown `xml:"async-teardown"`
	CCFAssist     *CCFAssist     `xml:"ccf-assist"`
	CFPC          *CFPC          `xml:"cfpc"`
	GIC           *GIC           `xml:"gic"`
	HAP           *struct{}      `xml:"hap"`
	HPT           *HPT           `xml:"hpt"`
	HTM           *HTM           `xml:"htm"`
	HyperV        *HyperV        `xml:"hyperv"`
	IBS           *IBS           `xml:"ibs"`
	IOAPIC        *IOAPIC        `xml:"ioapic"`
	KVM           *KVN           `xml:"kvm"`
	MSRS          *MSRS          `xml:"msrs"`
	PAE           *struct{}      `xml:"pae"`
	PrivNet       *struct{}      `xml:"privnet"`
	PS2           *PS2           `xml:"ps2"`
	PVSPINLock    *PVSPINLock    `xml:"pvspinlock"`
	RAS           *RAS           `xml:"ras"`
	SBBC          *SBBC          `xml:"sbbc"`
	SMM           *SMM           `xml:"smm"`
	TCG           *TCG           `xml:"tcg"`
	VMCoreInfo    *VMCoreInfo    `xml:"vmcoreinfo"`
	Xen           *Xen           `xml:"xen"`
}

type FibreChannel struct {
	AppID string `xml:"appid,attr"`
}

type Filesystem struct {
	Accessmode *string   `xml:"accessmode,attr"`
	DMode      *int      `xml:"dmode,attr"`
	FMode      *int      `xml:"fmode,attr"`
	Multidevs  *string   `xml:"multidevs,attr"`
	Type       string    `xml:"type,attr"`
	Binary     *Binary   `xml:"binary"`
	Driver     Driver    `xml:"driver"`
	IDMap      *IDMAP    `xml:"idmap"`
	Readonly   *struct{} `xml:"readonly"`
	Source     Source    `xml:"source"`
	Target     Target    `xml:"target"`
}

type FileTransfer struct {
	Enable string `xml:"enable,attr"`
}

type FilterRef struct {
	Filter     string      `xml:"filter,attr"`
	Patemeters []Parameter `xml:"parameter"`
}

type Format struct {
	Type          string         `xml:"type,attr"`
	MetadataCache *MetadataCache `xml:"metadata_cache"`
}

type Frames struct {
	Max int `xml:"max,attr"`
}

type Frequencies struct {
	State string `xml:"state,attr"`
}

type Frontend struct {
	Property []Property `xml:"property"`
}

type Geometry struct {
	Cyls  int    `xml:"cyls,attr"`
	Heads int    `xml:"heads,attr"`
	Secs  int    `xml:"secs,attr"`
	Trans string `xml:"trans,attr"`
}

type GIC struct {
	Version int `xml:"version,attr"`
}

type GID struct {
	Count  int `xml:"count,attr"`
	Start  int `xml:"start,attr"`
	Target int `xml:"target,attr"`
}

type GL struct {
	Enable     string `xml:"enable,attr"`
	RenderNode string `xml:"rendernode,attr"`
}

type Graphics struct {
	Autoport     string        `xml:"autoport,attr,omitempty"`
	Display      *string       `xml:"display,attr"`
	Fullscreen   *string       `xml:"fullscreen,attr"`
	Keymap       *string       `xml:"keymap,attr"`
	MultiUser    *string       `xml:"multiUser,attr"`
	Port         int           `xml:"port,attr"`
	SharePolicy  *string       `xml:"sharePolicy,attr"`
	TLSPort      *int          `xml:"tlsPort,attr"`
	Type         string        `xml:"type,attr"`
	Audio        *Audio        `xml:"audio"`
	Channel      []Channel     `xml:"channel"`
	Clipboard    *Clipboard    `xml:"clipboard"`
	FileTransfer *FileTransfer `xml:"filetransfer"`
	GL           *GL           `xml:"gl"`
	Image        *Image        `xml:"image"`
	Listen       *Listen       `xml:"listen"`
	Mouse        *Mouse        `xml:"mouse"`
	Streaming    *Streaming    `xml:"streaming"`
}

type Guest struct {
	CSum string  `xml:"csum,attr"`
	Dev  *string `xml:"dev,attr"`
	ECN  string  `xml:"ecn,attr"`
	TSO4 string  `xml:"tso4,attr"`
	TSO6 string  `xml:"tso6,attr"`
	UFO  string  `xml:"ufo,attr"`
}

type HardLimit struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Hidden struct {
	State string `xml:"state,attr"`
}

type HintDedicated struct {
	State string `xml:"state,attr"`
}

type Host struct {
	CSum     *string `xml:"csum,attr"`
	ECN      *string `xml:"ecn,attr"`
	GSO      *string `xml:"gso,attr"`
	MrgRxBuf *string `xml:"mrg_rxbuf,attr"`
	Name     string  `xml:"name,attr"`
	Port     *int    `xml:"port,attr"`
	TSO4     *string `xml:"tso4,attr"`
	TSO6     *string `xml:"tso6,attr"`
	UFO      *string `xml:"ufo,attr"`
}

type HostDev struct {
	Managed  *string   `xml:"managed,attr"`
	Mode     string    `xml:"mode,attr"`
	Model    *string   `xml:"model,attr"`
	RawIO    *string   `xml:"rawio,attr"`
	Type     string    `xml:"type,attr"`
	Address  *Address  `xml:"address"`
	Boot     Boot      `xml:"boot"`
	IP       *IP       `xml:"ip"`
	Readonly *struct{} `xml:"readonly"`
	ROM      *ROM      `xml:"rom"`
	Route    []Route   `xml:"route"`
	Source   Source    `xml:"source"`
	Teaming  *Teaming  `xml:"teaming"`
}

type HPT struct {
	Resizing    string      `xml:"resizing,attr"`
	MaxPagesize MaxPagesize `xml:"maxpagesize"`
}

type HTM struct {
	State string `xml:"state,attr"`
}

type Hub struct {
	Type string `xml:"type,attr"`
}

type Hugepages struct {
	Pages []Page `xml:"page"`
}

type HyperV struct {
	Mode            string          `xml:"mode,attr"`
	EMSRBitmap      EMSRBitmap      `xml:"emsr_bitmap"`
	EVMCS           EVMCS           `xml:"evmcs"`
	Frequencies     Frequencies     `xml:"frequencies"`
	IPI             IPI             `xml:"ipi"`
	Reenlightenment Reenlightenment `xml:"reenlightenment"`
	Relaxed         Relaxed         `xml:"relaxed"`
	Reset           Reset           `xml:"reset"`
	Runtime         Runtime         `xml:"runtime"`
	Spinlocks       Spinlocks       `xml:"spinlocks"`
	STimer          STimer          `xml:"stimer"`
	Synic           Synic           `xml:"synic"`
	TLBFlush        TLBFlush        `xml:"tlbflush"`
	Vapic           VAPIC           `xml:"vapic"`
	VendorID        VendorID        `xml:"vendor_id"`
	VPIndex         VPIndex         `xml:"vpindex"`
	XMMInput        XMMInput        `xml:"xmm_input"`
}

type IBS struct {
	Value string `xml:"value,attr"`
}

type Identity struct {
	Group string `xml:"group,attr"`
	User  string `xml:"user,attr"`
}

type IDMAP struct {
	GID GID `xml:"gid"`
	UID UID `xml:"uid"`
}

type Image struct {
	Compression string `xml:"compression,attr"`
}

type Inbound struct {
	Average int `xml:"average,attr"`
	Burst   int `xml:"burst,attr"`
	Floor   int `xml:"floor,attr"`
	Peak    int `xml:"peak,attr"`
}

type InitEnv struct {
	Name     string `xml:"name,attr"`
	CharData string `xml:",chardata"`
}

type Initiator struct {
	IQN IQN `xml:"iqn"`
}

type Input struct {
	BufferCount   *int      `xml:"bufferCount,attr"`
	BufferLength  *int      `xml:"bufferLength,attr"`
	Bus           *string   `xml:"bus,attr"`
	ClientName    *string   `xml:"clientName,attr"`
	ConnectPorts  *string   `xml:"connectPorts,attr"`
	Dev           *string   `xml:"dev,attr"`
	ExactName     *string   `xml:"exactName,attr"`
	FixedSettings *string   `xml:"fixedSettings,attr"`
	Latency       *int      `xml:"latency,attr"`
	MixingEngine  *string   `xml:"mixingEngine,attr"`
	Name          *string   `xml:"name,attr"`
	ServerName    *string   `xml:"serverName,attr"`
	StreamName    *string   `xml:"streamName,attr"`
	TryPoll       *string   `xml:"tryPoll,attr"`
	Type          string    `xml:"type,attr"`
	Voices        *int      `xml:"voices,attr"`
	Settings      *Settings `xml:"settings"`
	Source        Source    `xml:"source"`
}

type Interconnects struct {
	Bandwidth Bandwidth `xml:"bandwidth"`
	Latency   []Latency `xml:"latency"`
}

type Interface struct {
	TrustGuestRxFilters *string        `xml:"trustGuestRxFilters,attr"`
	Type                string         `xml:"type,attr,omitempty"`
	CharData            string         `xml:",chardata"`
	Address             *Address       `xml:"address"`
	ACPI                *ACPI          `xml:"acpi"`
	Alias               *Alias         `xml:"alias"`
	Backend             *Backend       `xml:"backend"`
	BackendDomain       *BackendDomain `xml:"backenddomain"`
	Bandwidth           *Bandwidth     `xml:"bandwidth"`
	Boot                *Boot          `xml:"boot"`
	Coalesce            *Coalesce      `xml:"coalesce"`
	Downscript          *DownScript    `xml:"downscript"`
	Driver              *Driver        `xml:"driver"`
	FilterRef           *FilterRef     `xml:"filterref"`
	Guest               *Guest         `xml:"guest"`
	IP                  []IP           `xml:"ip"`
	Link                *Link          `xml:"link"`
	Mac                 *Mac           `xml:"mac"`
	Model               *Model         `xml:"model"`
	MTU                 *MTU           `xml:"mtu"`
	Port                *Port          `xml:"port"`
	PortForward         []PortForward  `xml:"portForward"`
	ROM                 *ROM           `xml:"rom"`
	Route               []Route        `xml:"route"`
	Script              *Script        `xml:"script"`
	Source              *Source        `xml:"source"`
	Target              *Target        `xml:"target"`
	Teaming             *Teaming       `xml:"teaming"`
	Tune                *Tune          `xml:"tune"`
	Virtualport         *VirtualPort   `xml:"virtualport"`
	VLAN                *VLAN          `xml:"vlan"`
}

type IOAPIC struct {
	Driver string `xml:"driver,attr"`
}

type IOMMU struct {
	Model  string `xml:"model,attr"`
	Driver Driver `xml:"driver"`
}

type IOThread struct {
	ID            int     `xml:"id,attr"`
	ThreadPoolMax int     `xml:"thread_pool_max,attr"`
	ThreadPoolMin int     `xml:"thread_pool_min,attr"`
	Poll          Poll    `xml:"poll"`
	Queue         []Queue `xml:"queue"`
	CharData      string  `xml:",chardata"`
}

type IOThreadIDs struct {
	IOThreads []IOThread `xml:"iothread"`
}

type IOThreadPIN struct {
	CPUSet   string `xml:"cpuset,attr"`
	IOThread int    `xml:"iothread,attr"`
}

type IOThreadSched struct {
	IOThreads int    `xml:"iothreads,attr"`
	Scheduler string `xml:"scheduler,attr"`
}

type IOTune struct {
	ReadIOPSSec   int `xml:"read_iops_sec"`
	TotalBytesSec int `xml:"total_bytes_sec"`
	WriteIOPSSec  int `xml:"write_iops_sec"`
}

type IP struct {
	Address string  `xml:"address,attr"`
	Family  string  `xml:"family,attr"`
	Peer    *string `xml:"peer,attr"`
	Prefix  *int    `xml:"prefix,attr"`
}

type IPI struct {
	State string `xml:"state,attr"`
}

type IQN struct {
	Name string `xml:"name,attr"`
}

type Keywrap struct {
	Cipher Cipher `xml:"cipher"`
}

type KVN struct {
	DirtyRing     DirtyRing     `xml:"dirty-ring"`
	Hidden        Hidden        `xml:"hidden"`
	HintDedicated HintDedicated `xml:"hint-dedicated"`
	PollControl   PollControl   `xml:"poll-control"`
	PVIPI         PVIPI         `xml:"pv-ipi"`
}

type Label struct {
	CharData string `xml:",chardata"`
	Size     Size   `xml:"size"`
}

type Latency struct {
	Cache     int    `xml:"cache,attr"`
	Initiator int    `xml:"initiator,attr"`
	Target    int    `xml:"target,attr"`
	Type      string `xml:"type,attr"`
	Value     int    `xml:"value,attr"`
}

type LaunchSecurity struct {
	AuthorKey               string  `xml:"authorKey,attr"`
	KernelHashes            *string `xml:"kernelHashes,attr"`
	Type                    string  `xml:"type,attr"`
	VCEK                    string  `xml:"vcek,attr"`
	CBitPos                 int     `xml:"cbitpos"`
	DHCert                  *string `xml:"dhCert"`
	GuestVisibleWorkarounds string  `xml:"guestVisibleWorkarounds"`
	HostData                string  `xml:"hostData"`
	IDAuth                  string  `xml:"idAuth"`
	IDBlock                 string  `xml:"idBlock"`
	Policy                  string  `xml:"policy"`
	ReducedPhysBits         int     `xml:"reducedPhysBits"`
	Session                 *string `xml:"session"`
}

type Lease struct {
	Key       string `xml:"key"`
	Lockspace string `xml:"lockspace"`
	Target    Target `xml:"target"`
}

type LibOSInfo struct {
	LibOSInfo string `xml:"xmlns:libosinfo,attr"`
	OS        *OS    `xml:"libosinfo:os"`
}

type Line struct {
	Unit  string `xml:"unit,attr"`
	Value int    `xml:"value,attr"`
}

type Link struct {
	State string `xml:"state,attr"`
}

type Listen struct {
	Address *string `xml:"address,attr"`
	Network string  `xml:"network,attr"`
	Type    string  `xml:"type,attr"`
}

type Loader struct {
	Readonly  string  `xml:"readonly,attr"`
	Secure    string  `xml:"secure,attr"`
	Stateless *string `xml:"stateless,attr"`
	Type      string  `xml:"type,attr"`
	CharData  string  `xml:",chardata"`
}

type Local struct {
	Address string `xml:"address,attr"`
	Port    int    `xml:"port,attr"`
}

type Lock struct {
	Flock string `xml:"flock,attr"`
	Posix string `xml:"posix,attr"`
}

type Log struct {
	Append string `xml:"append,attr,omitempty"`
	File   string `xml:"file,attr,omitempty"`
}

type Mac struct {
	Address string `xml:"address,attr"`
}

type Master struct {
	Startport int `xml:"startport,attr"`
}

type MaxMemory struct {
	Slots    int    `xml:"slots,attr"`
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type MaxSize struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type MaxPagesize struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type MaxPhysAddr struct {
	Bits  int    `xml:"bits,attr"`
	Limit *int   `xml:"limit,attr"`
	Mode  string `xml:"mode,attr"`
}

type MemBalloon struct {
	Model   string   `xml:"model,attr"`
	Address *Address `xml:"address"`
	Driver  *Driver  `xml:"driver"`
	Stats   *Stats   `xml:"stats"`
}

type MemNode struct {
	CellID  int    `xml:"cellid,attr"`
	Mode    string `xml:"mode,attr"`
	Nodeset int    `xml:"nodeset,attr"`
}

type Memory struct {
	Access   *string `xml:"access,attr"`
	Discard  *string `xml:"discard,attr"`
	Mode     *string `xml:"mode,attr"`
	Model    *string `xml:"model,attr"`
	Nodeset  *string `xml:"nodeset,attr"`
	Unit     string  `xml:"unit,attr,omitempty"`
	CharData string  `xml:",chardata"`
	Source   *Source `xml:"source"`
	Target   *Target `xml:"target"`
	UUID     *string `xml:"uuid"`
}

type MemoryBacking struct {
	Access       Access     `xml:"access"`
	Allocation   Allocation `xml:"allocation"`
	Discard      struct{}   `xml:"discard"`
	Hugepages    Hugepages  `xml:"hugepages"`
	Locked       struct{}   `xml:"locked"`
	Nosharepages struct{}   `xml:"nosharepages"`
	Source       Source     `xml:"source"`
}

type MemoryTune struct {
	VCPUs string `xml:"vcpus,attr"`
	Node  Node   `xml:"node"`
}

type Memtune struct {
	HardLimit     HardLimit     `xml:"hard_limit"`
	MinGuarantee  MinGuarantee  `xml:"min_guarantee"`
	SoftLimit     SoftLimit     `xml:"soft_limit"`
	SwapHardLimit SwapHardLimit `xml:"swap_hard_limit"`
}

type Metadata struct {
	LibOSInfo *LibOSInfo `xml:"libosinfo:libosinfo"`
}

type MetadataCache struct {
	MaxSize MaxSize `xml:"max_size"`
}

type MinGuarantee struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Model struct {
	Fallback     *string       `xml:"fallback,attr"`
	Heads        int           `xml:"heads,attr"`
	Name         string        `xml:"name,attr,omitempty"`
	Type         string        `xml:"type,attr"`
	VRAM         int           `xml:"vram,attr"`
	CharData     string        `xml:",chardata"`
	Acceleration *Acceleration `xml:"acceleration"`
	Primary      string        `xml:"primary,attr,omitempty"`
}

type Monitor struct {
	Level int    `xml:"level,attr"`
	VCPUs string `xml:"vcpus,attr"`
}

type Mouse struct {
	Mode string `xml:"mode,attr"`
}

type MSI struct {
	IOEventFD string `xml:"ioeventfd,attr"`
	Vectors   int    `xml:"vectors,attr"`
}

type MSRS struct {
	Unknown string `xml:"unknown,attr"`
}

type MTU struct {
	Size int `xml:"size,attr"`
}

type Node struct {
	Bandwidth int    `xml:"bandwidth,attr"`
	ID        int    `xml:"id,attr"`
	CharData  string `xml:",chardata"`
}

type Numa struct {
	Cell          []Cell         `xml:"cell"`
	Interconnects *Interconnects `xml:"interconnects"`
}

type Numatune struct {
	MemNodes []MemNode `xml:"memnode"`
	Memory   Memory    `xml:"memory"`
}

type NVRAM struct {
	Template *string  `xml:"template,attr"`
	Type     string   `xml:"type,attr"`
	CharData string   `xml:",chardata"`
	Address  *Address `xml:"address"`
	Source   Source   `xml:"source"`
}

type OEMStrings struct {
	Entries []Entry `xml:"entry"`
}

type OS struct {
	Firmware       *string   `xml:"firmware,attr"`
	ID             string    `xml:"id,attr,omitempty"`
	ACPI           *ACPI     `xml:"acpi"`
	BIOS           *BIOS     `xml:"bios"`
	Boot           *Boot     `xml:"boot"`
	Bootloader     *string   `xml:"bootloader"`
	BootloaderArgs *string   `xml:"bootloader_args"`
	BootMenu       *BootMenu `xml:"bootmenu"`
	Cmdline        string    `xml:"cmdline,omitempty"`
	DTB            *string   `xml:"dtb"`
	IDMap          *IDMAP    `xml:"idmap"`
	Init           *string   `xml:"init"`
	InitArgs       []string  `xml:"initarg"`
	InitDir        *string   `xml:"initdir"`
	InitEnv        *InitEnv  `xml:"initenv"`
	InitGroup      *int      `xml:"initgroup"`
	InitRD         string    `xml:"initrd,omitempty"`
	InitUser       *string   `xml:"inituser"`
	Kernel         string    `xml:"kernel,omitempty"`
	Loader         *Loader   `xml:"loader"`
	NVRAM          *NVRAM    `xml:"nvram"`
	Shim           *string   `xml:"shim"`
	SMBIOS         *SMBIOS   `xml:"smbios"`
	Type           *Type     `xml:"type"`
}

type Outbound struct {
	Average int `xml:"average,attr"`
	Burst   int `xml:"burst,attr"`
	Peak    int `xml:"peak,attr"`
}

type Output struct {
	BufferCount   *int     `xml:"bufferCount,attr"`
	BufferLength  int      `xml:"bufferLength,attr"`
	ClientName    *string  `xml:"clientName,attr"`
	ConnectPorts  *string  `xml:"connectPorts,attr"`
	Dev           *string  `xml:"dev,attr"`
	ExactName     *string  `xml:"exactName,attr"`
	FixedSettings string   `xml:"fixedSettings,attr"`
	Latency       *int     `xml:"latency,attr"`
	MixingEngine  string   `xml:"mixingEngine,attr"`
	Name          *string  `xml:"name,attr"`
	ServerName    *string  `xml:"serverName,attr"`
	StreamName    *string  `xml:"streamName,attr"`
	TryPoll       *string  `xml:"tryPoll,attr"`
	Voices        int      `xml:"voices,attr"`
	Settings      Settings `xml:"settings"`
}

type Override struct {
	Device Device `xml:"device"`
}

type Page struct {
	Nodeset string `xml:"nodeset,attr"`
	Size    int    `xml:"size,attr"`
	Unit    string `xml:"unit,attr"`
}

type Pagesize struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Panic struct {
	Model   string  `xml:"model,attr"`
	Address Address `xml:"address"`
}

type Parallel struct {
	Type   string `xml:"type,attr"`
	Source Source `xml:"source"`
	Target Target `xml:"target"`
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Parameters struct {
	InstanceID    *string `xml:"instanceid,attr"`
	InterfaceID   *string `xml:"interfaceid,attr"`
	ManagerID     *int    `xml:"managerid,attr"`
	ProfileID     *string `xml:"profileid,attr"`
	TypeID        *int    `xml:"typeid,attr"`
	TypeIDVersion *int    `xml:"typeidversion,attr"`
}

type Passthrough struct {
	Mode  string `xml:"mode,attr"`
	State string `xml:"state,attr"`
}

type Perf struct {
	Events []Event `xml:"event"`
}

type PM struct {
	SuspendToDisk *SuspendToDisk `xml:"suspend-to-disk"`
	SuspendToMem  *SuspendToMem  `xml:"suspend-to-mem"`
}

type Poll struct {
	Grow   int `xml:"grow,attr"`
	Max    int `xml:"max,attr"`
	Shrink int `xml:"shrink,attr"`
}

type PollControl struct {
	State string `xml:"state,attr"`
}

type Port struct {
	Isolated string `xml:"isolated,attr"`
}

type PortForward struct {
	Address string  `xml:"address,attr"`
	Dev     string  `xml:"dev,attr"`
	Proto   string  `xml:"proto,attr"`
	Range   []Range `xml:"range"`
}

type Product struct {
	ID string `xml:"id,attr"`
}

type Profile struct {
	Name           string `xml:"name,attr"`
	RemoveDisabled string `xml:"removeDisabled,attr"`
	Source         string `xml:"source,attr"`
}

type Property struct {
	Name  string  `xml:"name,attr"`
	Type  string  `xml:"type,attr"`
	Value *string `xml:"value,attr"`
}

type Protocol struct {
	Type string `xml:"type,attr"`
}

type PS2 struct {
	State string `xml:"state,attr"`
}

type PStore struct {
	Backend string  `xml:"backend,attr"`
	Address Address `xml:"address"`
	Path    string  `xml:"path"`
	Size    Size    `xml:"size"`
}

type PVIPI struct {
	State string `xml:"state,attr"`
}

type PVSPINLock struct {
	State string `xml:"state,attr"`
}

type Queue struct {
	ID int `xml:"id,attr"`
}

type Range struct {
	End     *int    `xml:"end,attr"`
	Exclude *string `xml:"exclude,attr"`
	Start   int     `xml:"start,attr"`
	To      *int    `xml:"to,attr"`
}

type RAS struct {
	State string `xml:"state,attr"`
}

type Rate struct {
	Bytes  int `xml:"bytes,attr,omitzero"`
	Period int `xml:"period,attr,omitzero"`
}

type Readahead struct {
	Size int `xml:"size,attr"`
}

type Reconnect struct {
	Enabled string `xml:"enabled,attr"`
	Timeout int    `xml:"timeout,attr"`
}

type RedirDev struct {
	Bus    string `xml:"bus,attr"`
	Type   string `xml:"type,attr"`
	Boot   Boot   `xml:"boot"`
	Source Source `xml:"source"`
}

type RedirFilter struct {
	Usbdev []USBDev `xml:"usbdev"`
}

type Reenlightenment struct {
	State string `xml:"state,attr"`
}

type Relaxed struct {
	State string `xml:"state,attr"`
}

type Requested struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Reservations struct {
	Managed string `xml:"managed,attr"`
	Source  Source `xml:"source"`
}

type Reset struct {
	State string `xml:"state,attr"`
}

type Resource struct {
	FiberChannel FibreChannel `xml:"fibrechannel"`
	Partition    *string      `xml:"partition"`
}

type RNG struct {
	Model   string   `xml:"model,attr"`
	Backend *Backend `xml:"backend"`
	Rate    *Rate    `xml:"rate"`
	Address *Address `xml:"address"`
}

type ROM struct {
	Bar  string  `xml:"bar,attr"`
	File *string `xml:"file,attr"`
}

type Route struct {
	Address string `xml:"address,attr"`
	Family  string `xml:"family,attr"`
	Gateway string `xml:"gateway,attr"`
	Prefix  *int   `xml:"prefix,attr"`
}

type Runtime struct {
	State string `xml:"state,attr"`
}

type Rx struct {
	Frames Frames `xml:"frames"`
}

type Sandbox struct {
	Mode string `xml:"mode,attr"`
}

type SBBC struct {
	Value string `xml:"value,attr"`
}

type Script struct {
	Path string `xml:"path,attr"`
}

type SecLabel struct {
	Model     *string `xml:"model,attr"`
	Relabel   *string `xml:"relabel,attr"`
	Type      *string `xml:"type,attr"`
	Baselabel *string `xml:"baselabel"`
	Label     *Label  `xml:"label"`
}

type Secret struct {
	Type  string `xml:"type,attr"`
	Usage string `xml:"usage,attr"`
}

type Serial struct {
	Type     *string   `xml:"type,attr"`
	CharData string    `xml:",chardata"`
	Address  *Address  `xml:"address"`
	Protocol *Protocol `xml:"protocol"`
	Source   []Source  `xml:"source"`
	Target   *Target   `xml:"target"`
}

type Server struct {
	Path string `xml:"path,attr"`
}

type Settings struct {
	Channels  int    `xml:"channels,attr"`
	Format    string `xml:"format,attr"`
	Frequency int    `xml:"frequency,attr"`
}

type SHMem struct {
	Name   string  `xml:"name,attr"`
	Role   *string `xml:"role,attr"`
	Model  Model   `xml:"model"`
	MSI    MSI     `xml:"msi"`
	Server Server  `xml:"server"`
	Size   Size    `xml:"size"`
}

type Sibling struct {
	ID    int `xml:"id,attr"`
	Value int `xml:"value,attr"`
}

type Size struct {
	Unit     string `xml:"unit,attr"`
	Value    *int   `xml:"value,attr"`
	CharData string `xml:",chardata"`
}

type Slice struct {
	Offset int    `xml:"offset,attr"`
	Size   int    `xml:"size,attr"`
	Type   string `xml:"type,attr"`
}

type Slices struct {
	Slice Slice `xml:"slice"`
}

type Smartcard struct {
	Mode         string    `xml:"mode,attr"`
	Type         string    `xml:"type,attr"`
	Address      *Address  `xml:"address"`
	Certificates []string  `xml:"certificate"`
	Database     *string   `xml:"database"`
	Protocol     *Protocol `xml:"protocol"`
	Source       *Source   `xml:"source"`
}

type SMBIOS struct {
	Mode string `xml:"mode,attr"`
}

type SMM struct {
	State string `xml:"state,attr"`
	Tseg  Tseg   `xml:"tseg"`
}

type Snapshot struct {
	Name string `xml:"name,attr"`
}

type SoftLimit struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Sound struct {
	Model        string  `xml:"model,attr"`
	Multichannel *string `xml:"multichannel,attr"`
	Streams      *int    `xml:"streams,attr"`
	Audio        *Audio  `xml:"audio"`
	Codec        *Codec  `xml:"codec"`
}

type Source struct {
	Append         *string       `xml:"append,attr"`
	Bridge         *string       `xml:"bridge,attr"`
	Channel        *string       `xml:"channel,attr"`
	ConnectionID   *int          `xml:"connectionid,attr"`
	Dev            *string       `xml:"dev,attr"`
	Dir            *string       `xml:"dir,attr"`
	EVDev          *string       `xml:"evdev,attr"`
	File           string        `xml:"file,attr,omitempty"`
	Grab           *string       `xml:"grab,attr"`
	GrabToggle     *string       `xml:"grabToggle,attr"`
	GuestReset     *string       `xml:"guestReset,attr"`
	Host           *Host         `xml:"host"`
	HostAttr       *string       `xml:"host,attr"`
	Managed        *string       `xml:"managed,attr"`
	Mode           string        `xml:"mode,attr,omitempty"`
	Name           *string       `xml:"name,attr"`
	Namespace      *int          `xml:"namespace,attr"`
	Network        string        `xml:"network,attr,omitempty"`
	Pool           *string       `xml:"pool,attr"`
	Port           *int          `xml:"port,attr"`
	PortGroup      *string       `xml:"portgroup,attr"`
	PortGroupID    *string       `xml:"portgroupid,attr"`
	PortID         *int          `xml:"portid,attr"`
	Protocol       *string       `xml:"protocol,attr"`
	Query          *string       `xml:"query,attr"`
	Repeat         *string       `xml:"repeat,attr"`
	Service        *int          `xml:"service,attr"`
	Socket         *string       `xml:"socket,attr"`
	StartupPolicy  *string       `xml:"startupPolicy,attr"`
	SwitchID       *string       `xml:"switchid,attr"`
	TLS            *string       `xml:"tls,attr"`
	Type           *string       `xml:"type,attr"`
	Volume         *string       `xml:"volume,attr"`
	WriteFiltering *string       `xml:"writeFiltering,attr"`
	WWPN           *string       `xml:"wwpn,attr"`
	Adapter        *Adapter      `xml:"adapter"`
	Address        *Address      `xml:"address"`
	AddressAttr    *string       `xml:"address,attr"`
	AlignSize      *AlignSize    `xml:"alignsize"`
	Auth           *Auth         `xml:"auth"`
	Block          *Block        `xml:"block"`
	Config         *Config       `xml:"config"`
	Cookies        *Cookies      `xml:"cookies"`
	DataStore      *DataStore    `xml:"dataStore"`
	Identity       *Identity     `xml:"identity"`
	Initiator      *Initiator    `xml:"initiator"`
	Interface      *Interface    `xml:"interface"`
	Local          *Local        `xml:"local"`
	Nodemask       *string       `xml:"nodemask"`
	Pagesize       *Pagesize     `xml:"pagesize"`
	Path           *string       `xml:"path"`
	PathAttr       *string       `xml:"path,attr"`
	PMem           *struct{}     `xml:"pmem"`
	Product        *Product      `xml:"product"`
	Readahead      *Readahead    `xml:"readahead"`
	Reconnect      *Reconnect    `xml:"reconnect"`
	Reservations   *Reservations `xml:"reservations"`
	SecLabel       *SecLabel     `xml:"seclabel"`
	Slices         *Slices       `xml:"slices"`
	Snapshot       *Snapshot     `xml:"snapshot"`
	SSL            *SSL          `xml:"ssl"`
	Timeout        *Timeout      `xml:"timeout"`
	Vendor         *Vendor       `xml:"vendor"`
}

type Spinlocks struct {
	Retries int    `xml:"retries,attr"`
	State   string `xml:"state,attr"`
}

type SSL struct {
	Verify string `xml:"verify,attr"`
}

type Stats struct {
	Period int `xml:"period,attr"`
}

type STimer struct {
	State  string `xml:"state,attr"`
	Direct Direct `xml:"direct"`
}

type Streaming struct {
	Mode string `xml:"mode,attr"`
}

type SuspendToDisk struct {
	Enabled string `xml:"enabled,attr"`
}

type SuspendToMem struct {
	Enabled string `xml:"enabled,attr"`
}

type SwapHardLimit struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Synic struct {
	State string `xml:"state,attr"`
}

type Sysinfo struct {
	Type       string      `xml:"type,attr"`
	BaseBoard  *BaseBoard  `xml:"baseBoard"`
	BIOS       *BIOS       `xml:"bios"`
	Chassis    *Chassis    `xml:"chassis"`
	Entries    []Entry     `xml:"entry"`
	OEMStrings *OEMStrings `xml:"oemStrings"`
	System     *System     `xml:"system"`
}

type System struct {
	Entries []Entry `xml:"entry"`
}

type Table struct {
	Type     string `xml:"type,attr"`
	CharData string `xml:",chardata"`
}

type Tag struct {
	ID         int    `xml:"id,attr"`
	NativeMode string `xml:"nativeMode,attr"`
}

type Target struct {
	Bus          string     `xml:"bus,attr,omitempty"`
	Dev          string     `xml:"dev,attr,omitempty"`
	Dir          *string    `xml:"dir,attr"`
	Managed      *string    `xml:"managed,attr"`
	Name         string     `xml:"name,attr,omitempty"`
	Offset       *int       `xml:"offset,attr"`
	Path         *string    `xml:"path,attr"`
	Port         string     `xml:"port,attr,omitempty"`
	RotationRate *int       `xml:"rotation_rate,attr"`
	State        *string    `xml:"state,attr"`
	Tray         *string    `xml:"tray,attr"`
	Type         string     `xml:"type,attr,omitempty"`
	Address      *Address   `xml:"address"`
	AddressAttr  *string    `xml:"address,attr"`
	Block        *Block     `xml:"block"`
	Current      *Current   `xml:"current"`
	Label        *Label     `xml:"label"`
	Model        *Model     `xml:"model"`
	Node         *Node      `xml:"node"`
	Readonly     *struct{}  `xml:"readonly"`
	Requested    *Requested `xml:"requested"`
	Size         *Size      `xml:"size"`
	Chassis      int        `xml:"chassis,attr"`
}

type TBCache struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type TCG struct {
	TBCache TBCache `xml:"tb-cache"`
}

type Teaming struct {
	Persistent *string `xml:"persistent,attr"`
	Type       string  `xml:"type,attr"`
}

type ThreadPool struct {
	Size int `xml:"size,attr"`
}

type ThrottleFilter struct {
	Group string `xml:"group,attr"`
}

type ThrottleFilters struct {
	ThrottleFilters []ThrottleFilter `xml:"throttlefilter"`
}

type ThrottleGroup struct {
	GroupName     string `xml:"group_name"`
	ReadIOPSSec   int    `xml:"read_iops_sec"`
	TotalBytesSec int    `xml:"total_bytes_sec"`
	WriteIOPSSec  int    `xml:"write_iops_sec"`
}

type ThrottleGroups struct {
	ThrottleGroup ThrottleGroup `xml:"throttlegroup"`
}

type Timeout struct {
	Seconds int `xml:"seconds,attr"`
}

type Timer struct {
	Name       string   `xml:"name,attr"`
	Present    string   `xml:"present,attr,omitempty"`
	Tickpolicy string   `xml:"tickpolicy,attr,omitempty"`
	Track      string   `xml:"track,attr,omitempty"`
	Catchup    *Catchup `xml:"catchup"`
}

type TLBFlush struct {
	State    string   `xml:"state,attr"`
	Direct   Direct   `xml:"direct"`
	Extended Extended `xml:"extended"`
}

type Topology struct {
	Clusters int `xml:"clusters,attr"`
	Cores    int `xml:"cores,attr"`
	Dies     int `xml:"dies,attr"`
	Sockets  int `xml:"sockets,attr"`
	Threads  int `xml:"threads,attr"`
}

type Tpm struct {
	Model   string  `xml:"model,attr"`
	Backend Backend `xml:"backend"`
}

type Tseg struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Tune struct {
	SNDBuf int `xml:"sndbuf"`
}

type Type struct {
	Arch     string `xml:"arch,attr"`
	Machine  string `xml:"machine,attr"`
	CharData string `xml:",chardata"`
}

type UID struct {
	Count  int `xml:"count,attr"`
	Start  int `xml:"start,attr"`
	Target int `xml:"target,attr"`
}

type USBDev struct {
	Allow   string   `xml:"allow,attr"`
	Class   *string  `xml:"class,attr"`
	Product *string  `xml:"product,attr"`
	Vendor  *string  `xml:"vendor,attr"`
	Version *float64 `xml:"version,attr"`
}

type VAPIC struct {
	State string `xml:"state,attr"`
}

type VCPU struct {
	CPUSet       *string `xml:"cpuset,attr"`
	Current      *int    `xml:"current,attr"`
	Enabled      *string `xml:"enabled,attr"`
	Hotpluggable *string `xml:"hotpluggable,attr"`
	ID           *int    `xml:"id,attr"`
	Order        *int    `xml:"order,attr"`
	Placement    string  `xml:"placement,attr,omitempty"`
	CharData     string  `xml:",chardata"`
}

type VCPIPin struct {
	CPUSet string `xml:"cpuset,attr"`
	VCPU   int    `xml:"vcpu,attr"`
}

type VCPUs struct {
	VCPUs []VCPU `xml:"vcpu"`
}

type VCPUSched struct {
	Priority  int    `xml:"priority,attr"`
	Scheduler string `xml:"scheduler,attr"`
	VCPUs     string `xml:"vcpus,attr"`
}

type Vendor struct {
	ID       *string `xml:"id,attr"`
	CharData string  `xml:",chardata"`
}

type VendorID struct {
	State string `xml:"state,attr"`
	Value string `xml:"value,attr"`
}

type Video struct {
	Alias   *Alias   `xml:"alias"`
	Driver  *Driver  `xml:"driver"`
	Model   *Model   `xml:"model"`
	Address *Address `xml:"address"`
}

type VirtualPort struct {
	Type       *string    `xml:"type,attr"`
	Parameters Parameters `xml:"parameters"`
}

type VLAN struct {
	Trunk string `xml:"trunk,attr"`
	Tag   []Tag  `xml:"tag"`
}

type VMCoreInfo struct {
	State string `xml:"state,attr"`
}

type VPIndex struct {
	State string `xml:"state,attr"`
}

type VSock struct {
	Model string `xml:"model,attr"`
	CID   CID    `xml:"cid"`
}

type Watchdog struct {
	Action string `xml:"action,attr,omitempty"`
	Model  string `xml:"model,attr,omitempty"`
}

type Xen struct {
	E820Host    E820Host    `xml:"e820_host"`
	Passthrough Passthrough `xml:"passthrough"`
}

type XMMInput struct {
	State string `xml:"state,attr"`
}
