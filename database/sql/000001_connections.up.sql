create table if not exists connections_ (
  id_ integer primary key autoincrement,
  hypervisor_ text check(hypervisor_ in ('QEMU')) not null,
  url_ varchar(2048) not null,
  autoconnect_ boolean default false,
  ssh_ boolean default false,
  hostname_ varchar(253),
  username_ varchar(32),
  password_ varchar(255)
);
