create table if not exists connections_ (
  id_ integer primary key autoincrement,
  uri_ varchar(2048) not null,
  autoconnect_ boolean default false
);
