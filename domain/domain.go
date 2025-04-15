package domain

type Acceleration struct {
	Accel2d string `xml:"accel2d,attr"`
	Accel3d string `xml:"accel3d,attr"`
}

type Access struct {
	Mode string `xml:"mode,attr"`
}

type Acpi struct {
	Index *int   `xml:"index,attr"`
	Table *Table `xml:"table"`
}

type ActivePcrBanks struct {
	Sha256 struct{} `xml:"sha256"`
}

type Adapter struct {
	Name string `xml:"name,attr"`
}

type Address struct {
	Base          *string `xml:"base,attr"`
	Bus           *string `xml:"bus,attr"`
	Controller    *int    `xml:"controller,attr"`
	Cssid         *string `xml:"cssid,attr"`
	Devno         *string `xml:"devno,attr"`
	Domain        *string `xml:"domain,attr"`
	Function      *string `xml:"function,attr"`
	Iobase        *string `xml:"iobase,attr"`
	Multifunction *string `xml:"multifunction,attr"`
	Port          *int    `xml:"port,attr"`
	Reg           *string `xml:"reg,attr"`
	Slot          *string `xml:"slot,attr"`
	Ssid          *string `xml:"ssid,attr"`
	Target        *int    `xml:"target,attr"`
	Type          *string `xml:"type,attr"`
	Unit          *int    `xml:"unit,attr"`
	Uuid          *string `xml:"uuid,attr"`
}

type Aia struct {
	Value string `xml:"value,attr"`
}

type Alias struct {
	Name string `xml:"name,attr"`
}

type Alignsize struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Allocation struct {
	Mode    string `xml:"mode,attr"`
	Threads int    `xml:"threads,attr"`
}

type AsyncTeardown struct {
	Enabled string `xml:"enabled,attr"`
}

type Audio struct {
	Driver      *string `xml:"driver,attr"`
	DspPolicy   *int    `xml:"dspPolicy,attr"`
	Exclusive   *string `xml:"exclusive,attr"`
	ID          int     `xml:"id,attr"`
	Path        *string `xml:"path,attr"`
	RuntimeDir  *string `xml:"runtimeDir,attr"`
	ServerName  *string `xml:"serverName,attr"`
	TimerPeriod *int    `xml:"timerPeriod,attr"`
	TryMMap     *string `xml:"tryMMap,attr"`
	Type        *string `xml:"type,attr"`
	Input       *Input  `xml:"input"`
	Output      *Output `xml:"output"`
}

type Auth struct {
	Username string `xml:"username,attr"`
	Secret   Secret `xml:"secret"`
}

type Backend struct {
	Debug          *int            `xml:"debug,attr"`
	LogFile        string          `xml:"logFile,attr"`
	Model          *string         `xml:"model,attr"`
	Queues         *int            `xml:"queues,attr"`
	Tap            *string         `xml:"tap,attr"`
	Type           *string         `xml:"type,attr"`
	Version        *float64        `xml:"version,attr"`
	Vhost          *string         `xml:"vhost,attr"`
	CharData       string          `xml:",chardata"`
	ActivePcrBanks *ActivePcrBanks `xml:"active_pcr_banks"`
	Device         *Device         `xml:"device"`
	Encryption     *Encryption     `xml:"encryption"`
	Profile        *Profile        `xml:"profile"`
	Source         []Source        `xml:"source"`
}

type Backenddomain struct {
	Name string `xml:"name,attr"`
}

type BackingStore struct {
	Type         *string       `xml:"type,attr"`
	BackingStore *BackingStore `xml:"backingStore"`
	Format       *Format       `xml:"format"`
	Source       *Source       `xml:"source"`
}

type Bandwidth struct {
	Initiator *int     `xml:"initiator,attr"`
	Target    *int     `xml:"target,attr"`
	Type      *string  `xml:"type,attr"`
	Unit      *string  `xml:"unit,attr"`
	Value     *int     `xml:"value,attr"`
	Inbound   Inbound  `xml:"inbound"`
	Outbound  Outbound `xml:"outbound"`
}

type Bar struct {
	App2     string `xml:"app2,attr"`
	CharData string `xml:",chardata"`
}

type BaseBoard struct {
	Entry []Entry `xml:"entry"`
}

type Binary struct {
	Path       string     `xml:"path,attr"`
	Xattr      string     `xml:"xattr,attr"`
	Cache      Cache      `xml:"cache"`
	Lock       Lock       `xml:"lock"`
	Sandbox    Sandbox    `xml:"sandbox"`
	ThreadPool ThreadPool `xml:"thread_pool"`
}

type Bios struct {
	RebootTimeout *int    `xml:"rebootTimeout,attr"`
	Useserial     *string `xml:"useserial,attr"`
	Entry         Entry   `xml:"entry"`
}

type Blkiotune struct {
	Device []Device `xml:"device"`
	Weight int      `xml:"weight"`
}

type Block struct {
	Unit     *string `xml:"unit,attr"`
	CharData string  `xml:",chardata"`
}

type Blockio struct {
	DiscardGranularity int `xml:"discard_granularity,attr"`
	LogicalBlockSize   int `xml:"logical_block_size,attr"`
	PhysicalBlockSize  int `xml:"physical_block_size,attr"`
}

type Boot struct {
	Dev   *string `xml:"dev,attr"`
	Order *int    `xml:"order,attr"`
}

type Bootmenu struct {
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

type Cachetune struct {
	Vcpus   string    `xml:"vcpus,attr"`
	Cache   []Cache   `xml:"cache"`
	Monitor []Monitor `xml:"monitor"`
}

type Catchup struct {
	Limit     int `xml:"limit,attr"`
	Slew      int `xml:"slew,attr"`
	Threshold int `xml:"threshold,attr"`
}

type CcfAssist struct {
	State string `xml:"state,attr"`
}

type Cell struct {
	Cpus      string     `xml:"cpus,attr"`
	Discard   *string    `xml:"discard,attr"`
	ID        int        `xml:"id,attr"`
	MemAccess *string    `xml:"memAccess,attr"`
	Memory    int        `xml:"memory,attr"`
	Unit      string     `xml:"unit,attr"`
	Cache     *Cache     `xml:"cache"`
	Distances *Distances `xml:"distances"`
}

type Cfpc struct {
	Value string `xml:"value,attr"`
}

type Channel struct {
	Mode   string  `xml:"mode,attr"`
	Name   string  `xml:"name,attr"`
	Type   *string `xml:"type,attr"`
	Source *Source `xml:"source"`
	Target *Target `xml:"target"`
}

type Chassis struct {
	Entry []Entry `xml:"entry"`
}

type Cid struct {
	Address int    `xml:"address,attr"`
	Auto    string `xml:"auto,attr"`
}

type Cipher struct {
	Name  string `xml:"name,attr"`
	State string `xml:"state,attr"`
}

type Clipboard struct {
	Copypaste string `xml:"copypaste,attr"`
}

type Clock struct {
	Offset *string `xml:"offset,attr"`
	Sync   string  `xml:"sync,attr"`
	Timer  []Timer `xml:"timer"`
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
	Log    Log     `xml:"log"`
	Source Source  `xml:"source"`
	Target *Target `xml:"target"`
}

type Controller struct {
	Index            *int     `xml:"index,attr"`
	MaxEventChannels *int     `xml:"maxEventChannels,attr"`
	MaxGrantFrames   *int     `xml:"maxGrantFrames,attr"`
	Model            *string  `xml:"model,attr"`
	Ports            *int     `xml:"ports,attr"`
	Type             *string  `xml:"type,attr"`
	Vectors          *int     `xml:"vectors,attr"`
	Address          *Address `xml:"address"`
	Driver           []Driver `xml:"driver"`
	Master           *Master  `xml:"master"`
}

type Cookie struct {
	Name     string `xml:"name,attr"`
	CharData string `xml:",chardata"`
}

type Cookies struct {
	Cookie Cookie `xml:"cookie"`
}

type Cpu struct {
	Match       string       `xml:"match,attr"`
	Migratable  *string      `xml:"migratable,attr"`
	Mode        *string      `xml:"mode,attr"`
	Cache       *Cache       `xml:"cache"`
	Feature     *Feature     `xml:"feature"`
	Maxphysaddr *Maxphysaddr `xml:"maxphysaddr"`
	Model       *Model       `xml:"model"`
	Numa        *Numa        `xml:"numa"`
	Topology    *Topology    `xml:"topology"`
	Vendor      Vendor       `xml:"vendor"`
}

type Cputune struct {
	Cachetune      []Cachetune   `xml:"cachetune"`
	EmulatorPeriod int           `xml:"emulator_period"`
	EmulatorQuota  int           `xml:"emulator_quota"`
	Emulatorpin    Emulatorpin   `xml:"emulatorpin"`
	GlobalPeriod   int           `xml:"global_period"`
	GlobalQuota    int           `xml:"global_quota"`
	IothreadPeriod int           `xml:"iothread_period"`
	IothreadQuota  int           `xml:"iothread_quota"`
	Iothreadpin    []Iothreadpin `xml:"iothreadpin"`
	Iothreadsched  Iothreadsched `xml:"iothreadsched"`
	Memorytune     Memorytune    `xml:"memorytune"`
	Period         int           `xml:"period"`
	Quota          int           `xml:"quota"`
	Shares         int           `xml:"shares"`
	Vcpupin        []Vcpupin     `xml:"vcpupin"`
	Vcpusched      Vcpusched     `xml:"vcpusched"`
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
	Unit     *string `xml:"unit,attr"`
	CharData string  `xml:",chardata"`
}

type DataStore struct {
	Type   string `xml:"type,attr"`
	Format Format `xml:"format"`
	Source Source `xml:"source"`
}

type Defaultiothread struct {
	ThreadPoolMax int `xml:"thread_pool_max,attr"`
	ThreadPoolMin int `xml:"thread_pool_min,attr"`
}

type Device struct {
	Alias         *string   `xml:"alias,attr"`
	Frontend      *Frontend `xml:"frontend"`
	Path          string    `xml:"path"`
	PathAttr      *string   `xml:"path,attr"`
	ReadBytesSec  int       `xml:"read_bytes_sec"`
	ReadIopsSec   int       `xml:"read_iops_sec"`
	Weight        int       `xml:"weight"`
	WriteBytesSec int       `xml:"write_bytes_sec"`
	WriteIopsSec  int       `xml:"write_iops_sec"`
}

type Devices struct {
	Audio       []Audio      `xml:"audio"`
	Channel     []Channel    `xml:"channel"`
	Console     []Console    `xml:"console"`
	Controller  []Controller `xml:"controller"`
	Crypto      *Crypto      `xml:"crypto"`
	Disk        []Disk       `xml:"disk"`
	Emulator    *string      `xml:"emulator"`
	Filesystem  []Filesystem `xml:"filesystem"`
	Graphics    []Graphics   `xml:"graphics"`
	Hostdev     []Hostdev    `xml:"hostdev"`
	Hub         *Hub         `xml:"hub"`
	Input       []Input      `xml:"input"`
	Interface   []Interface  `xml:"interface"`
	Iommu       *Iommu       `xml:"iommu"`
	Lease       *Lease       `xml:"lease"`
	Memballoon  []Memballoon `xml:"memballoon"`
	Memory      []Memory     `xml:"memory"`
	Nvram       *Nvram       `xml:"nvram"`
	Panic       []Panic      `xml:"panic"`
	Parallel    []Parallel   `xml:"parallel"`
	Pstore      *Pstore      `xml:"pstore"`
	Redirdev    []Redirdev   `xml:"redirdev"`
	Redirfilter *Redirfilter `xml:"redirfilter"`
	Rng         *Rng         `xml:"rng"`
	Serial      []Serial     `xml:"serial"`
	Shmem       []Shmem      `xml:"shmem"`
	Smartcard   []Smartcard  `xml:"smartcard"`
	Sound       []Sound      `xml:"sound"`
	Tpm         []Tpm        `xml:"tpm"`
	Video       *Video       `xml:"video"`
	Vsock       *Vsock       `xml:"vsock"`
	Watchdog    []Watchdog   `xml:"watchdog"`
}

type Direct struct {
	State string `xml:"state,attr"`
}

type DirtyRing struct {
	Size  int    `xml:"size,attr"`
	State string `xml:"state,attr"`
}

type Disk struct {
	Device          *string          `xml:"device,attr"`
	Snapshot        *string          `xml:"snapshot,attr"`
	Type            string           `xml:"type,attr"`
	Address         *Address         `xml:"address"`
	Alias           *Alias           `xml:"alias"`
	BackingStore    *BackingStore    `xml:"backingStore"`
	Blockio         *Blockio         `xml:"blockio"`
	Boot            *Boot            `xml:"boot"`
	Driver          []Driver         `xml:"driver"`
	Encryption      *Encryption      `xml:"encryption"`
	Geometry        *Geometry        `xml:"geometry"`
	Iotune          *Iotune          `xml:"iotune"`
	Readonly        *struct{}        `xml:"readonly"`
	Serial          *Serial          `xml:"serial"`
	Shareable       *struct{}        `xml:"shareable"`
	Source          *Source          `xml:"source"`
	Target          *Target          `xml:"target"`
	Throttlefilters *Throttlefilters `xml:"throttlefilters"`
	Transient       *struct{}        `xml:"transient"`
}

type Distances struct {
	Sibling []Sibling `xml:"sibling"`
}

type Domain struct {
	Baz             *string          `xml:"baz"`
	Blkiotune       *Blkiotune       `xml:"blkiotune"`
	Clock           *Clock           `xml:"clock"`
	Cpu             *Cpu             `xml:"cpu"`
	Cputune         *Cputune         `xml:"cputune"`
	CurrentMemory   *CurrentMemory   `xml:"currentMemory"`
	Defaultiothread *Defaultiothread `xml:"defaultiothread"`
	Description     *string          `xml:"description"`
	Devices         *Devices         `xml:"devices"`
	Features        *Features        `xml:"features"`
	Genid           *string          `xml:"genid"`
	Iothreadids     *Iothreadids     `xml:"iothreadids"`
	Iothreads       *Iothreads       `xml:"iothreads"`
	Keywrap         *Keywrap         `xml:"keywrap"`
	LaunchSecurity  *LaunchSecurity  `xml:"launchSecurity"`
	MaxMemory       *MaxMemory       `xml:"maxMemory"`
	Memory          *Memory          `xml:"memory"`
	MemoryBacking   *MemoryBacking   `xml:"memoryBacking"`
	Memtune         *Memtune         `xml:"memtune"`
	Metadata        *Metadata        `xml:"metadata"`
	Name            string           `xml:"name"`
	Numatune        *Numatune        `xml:"numatune"`
	OnCrash         *string          `xml:"on_crash"`
	OnLockfailure   *string          `xml:"on_lockfailure"`
	OnPoweroff      *string          `xml:"on_poweroff"`
	OnReboot        *string          `xml:"on_reboot"`
	Os              Os               `xml:"os"`
	Override        *Override        `xml:"override"`
	Perf            *Perf            `xml:"perf"`
	Pm              *Pm              `xml:"pm"`
	Resource        []Resource       `xml:"resource"`
	Seclabel        []Seclabel       `xml:"seclabel"`
	Sysinfo         []Sysinfo        `xml:"sysinfo"`
	Throttlegroups  *Throttlegroups  `xml:"throttlegroups"`
	Title           *string          `xml:"title"`
	Uuid            string           `xml:"uuid"`
	Vcpu            *Vcpu            `xml:"vcpu"`
	Vcpus           *Vcpus           `xml:"vcpus"`
}

type Downscript struct {
	Path string `xml:"path,attr"`
}

type Driver struct {
	Ats           *string        `xml:"ats,attr"`
	Cache         *string        `xml:"cache,attr"`
	EventIdx      *string        `xml:"event_idx,attr"`
	Format        *string        `xml:"format,attr"`
	Intremap      *string        `xml:"intremap,attr"`
	Io            *string        `xml:"io,attr"`
	Ioeventfd     *string        `xml:"ioeventfd,attr"`
	Iommu         *string        `xml:"iommu,attr"`
	Iothread      *int           `xml:"iothread,attr"`
	Name          *string        `xml:"name,attr"`
	Queue         *int           `xml:"queue,attr"`
	QueueSize     *int           `xml:"queue_size,attr"`
	Queues        *int           `xml:"queues,attr"`
	RxQueueSize   *int           `xml:"rx_queue_size,attr"`
	TxQueueSize   *int           `xml:"tx_queue_size,attr"`
	Txmode        *string        `xml:"txmode,attr"`
	Type          *string        `xml:"type,attr"`
	Wrpolicy      *string        `xml:"wrpolicy,attr"`
	Guest         *Guest         `xml:"guest"`
	Host          *Host          `xml:"host"`
	Iothreads     *Iothreads     `xml:"iothreads"`
	MetadataCache *MetadataCache `xml:"metadata_cache"`
}

type E820Host struct {
	State string `xml:"state,attr"`
}

type EmsrBitmap struct {
	State string `xml:"state,attr"`
}

type Emulatorpin struct {
	Cpuset string `xml:"cpuset,attr"`
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

type Evmcs struct {
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
	Acpi          *Acpi          `xml:"acpi"`
	Aia           *Aia           `xml:"aia"`
	Apic          *struct{}      `xml:"apic"`
	AsyncTeardown *AsyncTeardown `xml:"async-teardown"`
	CcfAssist     *CcfAssist     `xml:"ccf-assist"`
	Cfpc          *Cfpc          `xml:"cfpc"`
	Gic           *Gic           `xml:"gic"`
	Hap           *struct{}      `xml:"hap"`
	Hpt           *Hpt           `xml:"hpt"`
	Htm           *Htm           `xml:"htm"`
	Hyperv        *Hyperv        `xml:"hyperv"`
	Ibs           *Ibs           `xml:"ibs"`
	Ioapic        *Ioapic        `xml:"ioapic"`
	Kvm           *Kvm           `xml:"kvm"`
	Msrs          *Msrs          `xml:"msrs"`
	Pae           *struct{}      `xml:"pae"`
	Privnet       *struct{}      `xml:"privnet"`
	Ps2           *Ps2           `xml:"ps2"`
	Pvspinlock    *Pvspinlock    `xml:"pvspinlock"`
	Ras           *Ras           `xml:"ras"`
	Sbbc          *Sbbc          `xml:"sbbc"`
	Smm           *Smm           `xml:"smm"`
	Tcg           *Tcg           `xml:"tcg"`
	Vmcoreinfo    *Vmcoreinfo    `xml:"vmcoreinfo"`
	Xen           *Xen           `xml:"xen"`
}

type Fibrechannel struct {
	Appid string `xml:"appid,attr"`
}

type Filesystem struct {
	Accessmode *string   `xml:"accessmode,attr"`
	Dmode      *int      `xml:"dmode,attr"`
	Fmode      *int      `xml:"fmode,attr"`
	Multidevs  *string   `xml:"multidevs,attr"`
	Type       string    `xml:"type,attr"`
	Binary     *Binary   `xml:"binary"`
	Driver     Driver    `xml:"driver"`
	Idmap      *Idmap    `xml:"idmap"`
	Readonly   *struct{} `xml:"readonly"`
	Source     Source    `xml:"source"`
	Target     Target    `xml:"target"`
}

type Filetransfer struct {
	Enable string `xml:"enable,attr"`
}

type Filterref struct {
	Filter    string      `xml:"filter,attr"`
	Parameter []Parameter `xml:"parameter"`
}

type Foo struct {
	App1     string `xml:"app1,attr"`
	CharData string `xml:",chardata"`
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

type Gic struct {
	Version int `xml:"version,attr"`
}

type Gid struct {
	Count  int `xml:"count,attr"`
	Start  int `xml:"start,attr"`
	Target int `xml:"target,attr"`
}

type Gl struct {
	Enable     string `xml:"enable,attr"`
	Rendernode string `xml:"rendernode,attr"`
}

type Graphics struct {
	Autoport     *string       `xml:"autoport,attr"`
	Display      *string       `xml:"display,attr"`
	Fullscreen   *string       `xml:"fullscreen,attr"`
	Keymap       *string       `xml:"keymap,attr"`
	MultiUser    *string       `xml:"multiUser,attr"`
	Port         *int          `xml:"port,attr"`
	SharePolicy  *string       `xml:"sharePolicy,attr"`
	TlsPort      *int          `xml:"tlsPort,attr"`
	Type         string        `xml:"type,attr"`
	Audio        *Audio        `xml:"audio"`
	Channel      []Channel     `xml:"channel"`
	Clipboard    *Clipboard    `xml:"clipboard"`
	Filetransfer *Filetransfer `xml:"filetransfer"`
	Gl           *Gl           `xml:"gl"`
	Image        *Image        `xml:"image"`
	Listen       *Listen       `xml:"listen"`
	Mouse        *Mouse        `xml:"mouse"`
	Streaming    *Streaming    `xml:"streaming"`
}

type Guest struct {
	Csum string  `xml:"csum,attr"`
	Dev  *string `xml:"dev,attr"`
	Ecn  string  `xml:"ecn,attr"`
	Tso4 string  `xml:"tso4,attr"`
	Tso6 string  `xml:"tso6,attr"`
	Ufo  string  `xml:"ufo,attr"`
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
	Csum     *string `xml:"csum,attr"`
	Ecn      *string `xml:"ecn,attr"`
	Gso      *string `xml:"gso,attr"`
	MrgRxbuf *string `xml:"mrg_rxbuf,attr"`
	Name     string  `xml:"name,attr"`
	Port     *int    `xml:"port,attr"`
	Tso4     *string `xml:"tso4,attr"`
	Tso6     *string `xml:"tso6,attr"`
	Ufo      *string `xml:"ufo,attr"`
}

type Hostdev struct {
	Managed  *string   `xml:"managed,attr"`
	Mode     string    `xml:"mode,attr"`
	Model    *string   `xml:"model,attr"`
	Rawio    *string   `xml:"rawio,attr"`
	Type     string    `xml:"type,attr"`
	Address  *Address  `xml:"address"`
	Boot     Boot      `xml:"boot"`
	Ip       *Ip       `xml:"ip"`
	Readonly *struct{} `xml:"readonly"`
	Rom      *Rom      `xml:"rom"`
	Route    []Route   `xml:"route"`
	Source   Source    `xml:"source"`
	Teaming  *Teaming  `xml:"teaming"`
}

type Hpt struct {
	Resizing    string      `xml:"resizing,attr"`
	Maxpagesize Maxpagesize `xml:"maxpagesize"`
}

type Htm struct {
	State string `xml:"state,attr"`
}

type Hub struct {
	Type string `xml:"type,attr"`
}

type Hugepages struct {
	Page []Page `xml:"page"`
}

type Hyperv struct {
	Mode            string          `xml:"mode,attr"`
	EmsrBitmap      EmsrBitmap      `xml:"emsr_bitmap"`
	Evmcs           Evmcs           `xml:"evmcs"`
	Frequencies     Frequencies     `xml:"frequencies"`
	Ipi             Ipi             `xml:"ipi"`
	Reenlightenment Reenlightenment `xml:"reenlightenment"`
	Relaxed         Relaxed         `xml:"relaxed"`
	Reset           Reset           `xml:"reset"`
	Runtime         Runtime         `xml:"runtime"`
	Spinlocks       Spinlocks       `xml:"spinlocks"`
	Stimer          Stimer          `xml:"stimer"`
	Synic           Synic           `xml:"synic"`
	Tlbflush        Tlbflush        `xml:"tlbflush"`
	Vapic           Vapic           `xml:"vapic"`
	VendorID        VendorID        `xml:"vendor_id"`
	Vpindex         Vpindex         `xml:"vpindex"`
	XmmInput        XmmInput        `xml:"xmm_input"`
}

type Ibs struct {
	Value string `xml:"value,attr"`
}

type Identity struct {
	Group string `xml:"group,attr"`
	User  string `xml:"user,attr"`
}

type Idmap struct {
	Gid Gid `xml:"gid"`
	Uid Uid `xml:"uid"`
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

type Initenv struct {
	Name     string `xml:"name,attr"`
	CharData string `xml:",chardata"`
}

type Initiator struct {
	Iqn Iqn `xml:"iqn"`
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
	Type                *string        `xml:"type,attr"`
	CharData            string         `xml:",chardata"`
	Acpi                *Acpi          `xml:"acpi"`
	Alias               *Alias         `xml:"alias"`
	Backend             *Backend       `xml:"backend"`
	Backenddomain       *Backenddomain `xml:"backenddomain"`
	Bandwidth           *Bandwidth     `xml:"bandwidth"`
	Boot                *Boot          `xml:"boot"`
	Coalesce            *Coalesce      `xml:"coalesce"`
	Downscript          *Downscript    `xml:"downscript"`
	Driver              *Driver        `xml:"driver"`
	Filterref           *Filterref     `xml:"filterref"`
	Guest               *Guest         `xml:"guest"`
	Ip                  []Ip           `xml:"ip"`
	Link                *Link          `xml:"link"`
	Mac                 *Mac           `xml:"mac"`
	Model               *Model         `xml:"model"`
	Mtu                 *Mtu           `xml:"mtu"`
	Port                *Port          `xml:"port"`
	PortForward         []PortForward  `xml:"portForward"`
	Rom                 *Rom           `xml:"rom"`
	Route               []Route        `xml:"route"`
	Script              *Script        `xml:"script"`
	Source              []Source       `xml:"source"`
	Target              *Target        `xml:"target"`
	Teaming             *Teaming       `xml:"teaming"`
	Tune                *Tune          `xml:"tune"`
	Virtualport         *Virtualport   `xml:"virtualport"`
	Vlan                *Vlan          `xml:"vlan"`
}

type Ioapic struct {
	Driver string `xml:"driver,attr"`
}

type Iommu struct {
	Model  string `xml:"model,attr"`
	Driver Driver `xml:"driver"`
}

type Iothread struct {
	ID            int     `xml:"id,attr"`
	ThreadPoolMax int     `xml:"thread_pool_max,attr"`
	ThreadPoolMin int     `xml:"thread_pool_min,attr"`
	Poll          Poll    `xml:"poll"`
	Queue         []Queue `xml:"queue"`
}

type Iothreadids struct {
	Iothread []Iothread `xml:"iothread"`
}

type Iothreadpin struct {
	Cpuset   string `xml:"cpuset,attr"`
	Iothread int    `xml:"iothread,attr"`
}

type Iothreads struct {
	CharData string     `xml:",chardata"`
	Iothread []Iothread `xml:"iothread"`
}

type Iothreadsched struct {
	Iothreads int    `xml:"iothreads,attr"`
	Scheduler string `xml:"scheduler,attr"`
}

type Iotune struct {
	ReadIopsSec   int `xml:"read_iops_sec"`
	TotalBytesSec int `xml:"total_bytes_sec"`
	WriteIopsSec  int `xml:"write_iops_sec"`
}

type Ip struct {
	Address string  `xml:"address,attr"`
	Family  string  `xml:"family,attr"`
	Peer    *string `xml:"peer,attr"`
	Prefix  *int    `xml:"prefix,attr"`
}

type Ipi struct {
	State string `xml:"state,attr"`
}

type Iqn struct {
	Name string `xml:"name,attr"`
}

type Keywrap struct {
	Cipher Cipher `xml:"cipher"`
}

type Kvm struct {
	DirtyRing     DirtyRing     `xml:"dirty-ring"`
	Hidden        Hidden        `xml:"hidden"`
	HintDedicated HintDedicated `xml:"hint-dedicated"`
	PollControl   PollControl   `xml:"poll-control"`
	PvIpi         PvIpi         `xml:"pv-ipi"`
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
	Vcek                    string  `xml:"vcek,attr"`
	Cbitpos                 int     `xml:"cbitpos"`
	DhCert                  *string `xml:"dhCert"`
	GuestVisibleWorkarounds string  `xml:"guestVisibleWorkarounds"`
	HostData                string  `xml:"hostData"`
	IdAuth                  string  `xml:"idAuth"`
	IdBlock                 string  `xml:"idBlock"`
	Policy                  string  `xml:"policy"`
	ReducedPhysBits         int     `xml:"reducedPhysBits"`
	Session                 *string `xml:"session"`
}

type Lease struct {
	Key       string `xml:"key"`
	Lockspace string `xml:"lockspace"`
	Target    Target `xml:"target"`
}

type Libosinfo struct {
	Libosinfo string `xml:"libosinfo,attr"`
	Os        Os     `xml:"os"`
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
	Append string `xml:"append,attr"`
	File   string `xml:"file,attr"`
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

type Maxpagesize struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Maxphysaddr struct {
	Bits  int    `xml:"bits,attr"`
	Limit *int   `xml:"limit,attr"`
	Mode  string `xml:"mode,attr"`
}

type Memballoon struct {
	Model   string   `xml:"model,attr"`
	Address *Address `xml:"address"`
	Driver  *Driver  `xml:"driver"`
	Stats   *Stats   `xml:"stats"`
}

type Memnode struct {
	Cellid  int    `xml:"cellid,attr"`
	Mode    string `xml:"mode,attr"`
	Nodeset int    `xml:"nodeset,attr"`
}

type Memory struct {
	Access   *string `xml:"access,attr"`
	Discard  *string `xml:"discard,attr"`
	Mode     *string `xml:"mode,attr"`
	Model    *string `xml:"model,attr"`
	Nodeset  *string `xml:"nodeset,attr"`
	Unit     *string `xml:"unit,attr"`
	CharData string  `xml:",chardata"`
	Source   *Source `xml:"source"`
	Target   *Target `xml:"target"`
	Uuid     *string `xml:"uuid"`
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

type Memorytune struct {
	Vcpus string `xml:"vcpus,attr"`
	Node  Node   `xml:"node"`
}

type Memtune struct {
	HardLimit     HardLimit     `xml:"hard_limit"`
	MinGuarantee  MinGuarantee  `xml:"min_guarantee"`
	SoftLimit     SoftLimit     `xml:"soft_limit"`
	SwapHardLimit SwapHardLimit `xml:"swap_hard_limit"`
}

type Metadata struct {
	Bar       Bar        `xml:"bar"`
	Foo       Foo        `xml:"foo"`
	Libosinfo *Libosinfo `xml:"libosinfo"`
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
	Heads        *int          `xml:"heads,attr"`
	Name         *string       `xml:"name,attr"`
	Type         *string       `xml:"type,attr"`
	Vram         *int          `xml:"vram,attr"`
	CharData     string        `xml:",chardata"`
	Acceleration *Acceleration `xml:"acceleration"`
}

type Monitor struct {
	Level int    `xml:"level,attr"`
	Vcpus string `xml:"vcpus,attr"`
}

type Mouse struct {
	Mode string `xml:"mode,attr"`
}

type Msi struct {
	Ioeventfd string `xml:"ioeventfd,attr"`
	Vectors   int    `xml:"vectors,attr"`
}

type Msrs struct {
	Unknown string `xml:"unknown,attr"`
}

type Mtu struct {
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
	Memnode []Memnode `xml:"memnode"`
	Memory  Memory    `xml:"memory"`
}

type Nvram struct {
	Template *string  `xml:"template,attr"`
	Type     string   `xml:"type,attr"`
	CharData string   `xml:",chardata"`
	Address  *Address `xml:"address"`
	Source   Source   `xml:"source"`
}

type OemStrings struct {
	Entry []Entry `xml:"entry"`
}

type Os struct {
	Firmware       *string   `xml:"firmware,attr"`
	ID             *string   `xml:"id,attr"`
	Acpi           *Acpi     `xml:"acpi"`
	Bios           *Bios     `xml:"bios"`
	Boot           *Boot     `xml:"boot"`
	Bootloader     *string   `xml:"bootloader"`
	BootloaderArgs *string   `xml:"bootloader_args"`
	Bootmenu       *Bootmenu `xml:"bootmenu"`
	Cmdline        *string   `xml:"cmdline"`
	Dtb            *string   `xml:"dtb"`
	Idmap          *Idmap    `xml:"idmap"`
	Init           *string   `xml:"init"`
	Initarg        []string  `xml:"initarg"`
	Initdir        *string   `xml:"initdir"`
	Initenv        *Initenv  `xml:"initenv"`
	Initgroup      *int      `xml:"initgroup"`
	Initrd         *string   `xml:"initrd"`
	Inituser       *string   `xml:"inituser"`
	Kernel         *string   `xml:"kernel"`
	Loader         *Loader   `xml:"loader"`
	Nvram          *Nvram    `xml:"nvram"`
	Shim           *string   `xml:"shim"`
	Smbios         *Smbios   `xml:"smbios"`
	Type           Type      `xml:"type"`
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
	Instanceid    *string `xml:"instanceid,attr"`
	Interfaceid   *string `xml:"interfaceid,attr"`
	Managerid     *int    `xml:"managerid,attr"`
	Profileid     *string `xml:"profileid,attr"`
	Typeid        *int    `xml:"typeid,attr"`
	Typeidversion *int    `xml:"typeidversion,attr"`
}

type Passthrough struct {
	Mode  string `xml:"mode,attr"`
	State string `xml:"state,attr"`
}

type Perf struct {
	Event []Event `xml:"event"`
}

type Pm struct {
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

type Ps2 struct {
	State string `xml:"state,attr"`
}

type Pstore struct {
	Backend string  `xml:"backend,attr"`
	Address Address `xml:"address"`
	Path    string  `xml:"path"`
	Size    Size    `xml:"size"`
}

type PvIpi struct {
	State string `xml:"state,attr"`
}

type Pvspinlock struct {
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

type Ras struct {
	State string `xml:"state,attr"`
}

type Rate struct {
	Bytes  int `xml:"bytes,attr"`
	Period int `xml:"period,attr"`
}

type Readahead struct {
	Size int `xml:"size,attr"`
}

type Reconnect struct {
	Enabled string `xml:"enabled,attr"`
	Timeout int    `xml:"timeout,attr"`
}

type Redirdev struct {
	Bus    string `xml:"bus,attr"`
	Type   string `xml:"type,attr"`
	Boot   Boot   `xml:"boot"`
	Source Source `xml:"source"`
}

type Redirfilter struct {
	Usbdev []Usbdev `xml:"usbdev"`
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
	Fibrechannel Fibrechannel `xml:"fibrechannel"`
	Partition    *string      `xml:"partition"`
}

type Rng struct {
	Model   string    `xml:"model,attr"`
	Backend []Backend `xml:"backend"`
	Rate    Rate      `xml:"rate"`
}

type Rom struct {
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

type Sbbc struct {
	Value string `xml:"value,attr"`
}

type Script struct {
	Path string `xml:"path,attr"`
}

type Seclabel struct {
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

type Shmem struct {
	Name   string  `xml:"name,attr"`
	Role   *string `xml:"role,attr"`
	Model  Model   `xml:"model"`
	Msi    Msi     `xml:"msi"`
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
	Mode        string    `xml:"mode,attr"`
	Type        string    `xml:"type,attr"`
	Address     *Address  `xml:"address"`
	Certificate []string  `xml:"certificate"`
	Database    *string   `xml:"database"`
	Protocol    *Protocol `xml:"protocol"`
	Source      *Source   `xml:"source"`
}

type Smbios struct {
	Mode string `xml:"mode,attr"`
}

type Smm struct {
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
	Connectionid   *int          `xml:"connectionid,attr"`
	Dev            *string       `xml:"dev,attr"`
	Dir            *string       `xml:"dir,attr"`
	Evdev          *string       `xml:"evdev,attr"`
	File           *string       `xml:"file,attr"`
	Grab           *string       `xml:"grab,attr"`
	GrabToggle     *string       `xml:"grabToggle,attr"`
	GuestReset     *string       `xml:"guestReset,attr"`
	Host           *Host         `xml:"host"`
	HostAttr       *string       `xml:"host,attr"`
	Managed        *string       `xml:"managed,attr"`
	Mode           *string       `xml:"mode,attr"`
	Name           *string       `xml:"name,attr"`
	Namespace      *int          `xml:"namespace,attr"`
	Network        *string       `xml:"network,attr"`
	Pool           *string       `xml:"pool,attr"`
	Port           *int          `xml:"port,attr"`
	Portgroup      *string       `xml:"portgroup,attr"`
	Portgroupid    *string       `xml:"portgroupid,attr"`
	Portid         *int          `xml:"portid,attr"`
	Protocol       *string       `xml:"protocol,attr"`
	Query          *string       `xml:"query,attr"`
	Repeat         *string       `xml:"repeat,attr"`
	Service        *int          `xml:"service,attr"`
	Socket         *string       `xml:"socket,attr"`
	StartupPolicy  *string       `xml:"startupPolicy,attr"`
	Switchid       *string       `xml:"switchid,attr"`
	Tls            *string       `xml:"tls,attr"`
	Type           *string       `xml:"type,attr"`
	Volume         *string       `xml:"volume,attr"`
	WriteFiltering *string       `xml:"writeFiltering,attr"`
	Wwpn           *string       `xml:"wwpn,attr"`
	Adapter        *Adapter      `xml:"adapter"`
	Address        *Address      `xml:"address"`
	AddressAttr    *string       `xml:"address,attr"`
	Alignsize      *Alignsize    `xml:"alignsize"`
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
	Pmem           *struct{}     `xml:"pmem"`
	Product        *Product      `xml:"product"`
	Readahead      *Readahead    `xml:"readahead"`
	Reconnect      *Reconnect    `xml:"reconnect"`
	Reservations   *Reservations `xml:"reservations"`
	Seclabel       *Seclabel     `xml:"seclabel"`
	Slices         *Slices       `xml:"slices"`
	Snapshot       *Snapshot     `xml:"snapshot"`
	Ssl            *Ssl          `xml:"ssl"`
	Timeout        *Timeout      `xml:"timeout"`
	Vendor         *Vendor       `xml:"vendor"`
}

type Spinlocks struct {
	Retries int    `xml:"retries,attr"`
	State   string `xml:"state,attr"`
}

type Ssl struct {
	Verify string `xml:"verify,attr"`
}

type Stats struct {
	Period int `xml:"period,attr"`
}

type Stimer struct {
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
	Bios       *Bios       `xml:"bios"`
	Chassis    *Chassis    `xml:"chassis"`
	Entry      []Entry     `xml:"entry"`
	OemStrings *OemStrings `xml:"oemStrings"`
	System     *System     `xml:"system"`
}

type System struct {
	Entry []Entry `xml:"entry"`
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
	Bus          *string    `xml:"bus,attr"`
	Dev          *string    `xml:"dev,attr"`
	Dir          *string    `xml:"dir,attr"`
	Managed      *string    `xml:"managed,attr"`
	Name         *string    `xml:"name,attr"`
	Offset       *int       `xml:"offset,attr"`
	Path         *string    `xml:"path,attr"`
	Port         *int       `xml:"port,attr"`
	RotationRate *int       `xml:"rotation_rate,attr"`
	State        *string    `xml:"state,attr"`
	Tray         *string    `xml:"tray,attr"`
	Type         *string    `xml:"type,attr"`
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
}

type TbCache struct {
	Unit     string `xml:"unit,attr"`
	CharData string `xml:",chardata"`
}

type Tcg struct {
	TbCache TbCache `xml:"tb-cache"`
}

type Teaming struct {
	Persistent *string `xml:"persistent,attr"`
	Type       string  `xml:"type,attr"`
}

type ThreadPool struct {
	Size int `xml:"size,attr"`
}

type Throttlefilter struct {
	Group string `xml:"group,attr"`
}

type Throttlefilters struct {
	Throttlefilter []Throttlefilter `xml:"throttlefilter"`
}

type Throttlegroup struct {
	GroupName     string `xml:"group_name"`
	ReadIopsSec   int    `xml:"read_iops_sec"`
	TotalBytesSec int    `xml:"total_bytes_sec"`
	WriteIopsSec  int    `xml:"write_iops_sec"`
}

type Throttlegroups struct {
	Throttlegroup Throttlegroup `xml:"throttlegroup"`
}

type Timeout struct {
	Seconds int `xml:"seconds,attr"`
}

type Timer struct {
	Name       string   `xml:"name,attr"`
	Present    *string  `xml:"present,attr"`
	Tickpolicy *string  `xml:"tickpolicy,attr"`
	Track      *string  `xml:"track,attr"`
	Catchup    *Catchup `xml:"catchup"`
}

type Tlbflush struct {
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
	Sndbuf int `xml:"sndbuf"`
}

type Type struct {
	Arch     *string `xml:"arch,attr"`
	Machine  *string `xml:"machine,attr"`
	CharData string  `xml:",chardata"`
}

type Uid struct {
	Count  int `xml:"count,attr"`
	Start  int `xml:"start,attr"`
	Target int `xml:"target,attr"`
}

type Usbdev struct {
	Allow   string   `xml:"allow,attr"`
	Class   *string  `xml:"class,attr"`
	Product *string  `xml:"product,attr"`
	Vendor  *string  `xml:"vendor,attr"`
	Version *float64 `xml:"version,attr"`
}

type Vapic struct {
	State string `xml:"state,attr"`
}

type Vcpu struct {
	Cpuset       *string `xml:"cpuset,attr"`
	Current      *int    `xml:"current,attr"`
	Enabled      *string `xml:"enabled,attr"`
	Hotpluggable *string `xml:"hotpluggable,attr"`
	ID           *int    `xml:"id,attr"`
	Order        *int    `xml:"order,attr"`
	Placement    *string `xml:"placement,attr"`
	CharData     string  `xml:",chardata"`
}

type Vcpupin struct {
	Cpuset string `xml:"cpuset,attr"`
	Vcpu   int    `xml:"vcpu,attr"`
}

type Vcpus struct {
	Vcpu []Vcpu `xml:"vcpu"`
}

type Vcpusched struct {
	Priority  int    `xml:"priority,attr"`
	Scheduler string `xml:"scheduler,attr"`
	Vcpus     string `xml:"vcpus,attr"`
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
	Driver Driver `xml:"driver"`
	Model  Model  `xml:"model"`
}

type Virtualport struct {
	Type       *string    `xml:"type,attr"`
	Parameters Parameters `xml:"parameters"`
}

type Vlan struct {
	Trunk string `xml:"trunk,attr"`
	Tag   []Tag  `xml:"tag"`
}

type Vmcoreinfo struct {
	State string `xml:"state,attr"`
}

type Vpindex struct {
	State string `xml:"state,attr"`
}

type Vsock struct {
	Model string `xml:"model,attr"`
	Cid   Cid    `xml:"cid"`
}

type Watchdog struct {
	Action *string `xml:"action,attr"`
	Model  string  `xml:"model,attr"`
}

type Xen struct {
	E820Host    E820Host    `xml:"e820_host"`
	Passthrough Passthrough `xml:"passthrough"`
}

type XmmInput struct {
	State string `xml:"state,attr"`
}
