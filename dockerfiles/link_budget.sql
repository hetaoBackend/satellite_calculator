create table satellite
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	gT decimal(5,1) not null,
	sFD decimal(5,1) null,
	eIRP decimal(5,1) not null,
	longitude decimal(5,2) not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table station
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	diameter decimal(5,1) not null,
	tPower decimal(5,1) null,
	tG decimal(5,1) not null,
	rGT decimal(5,2) not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table city
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	longitude decimal(5,1) not null,
	latitude decimal(5,1) null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

