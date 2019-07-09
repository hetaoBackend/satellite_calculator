create table satellite
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	g_t decimal(5,1) not null,
	s_f_d decimal(5,1) null,
	e_i_r_p decimal(5,1) not null,
	longitude decimal(5,2) not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table station
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	diameter decimal(5,1) not null,
	t_power decimal(5,1) null,
	t_g decimal(5,1) not null,
	r_g_t decimal(5,2) not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table city
(
	id int(11) unsigned auto_increment
		primary key,
	name varchar(50) not null,
	longitude decimal(5,1) not null,
	latitude decimal(5,1) null
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

