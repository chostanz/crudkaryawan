-- Database: db_users

-- DROP DATABASE IF EXISTS db_users;

CREATE DATABASE db_users
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_Indonesia.1252'
    LC_CTYPE = 'English_Indonesia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
	
	
	create table users (
		id serial not null,
		name VARCHAR(100),
		phone VARCHAR(50),
		address VARCHAR(200),
		primary key (id));
		
		select * from users;
		
		insert into users(name, phone,address) values
		('konan', '0999088090', 'amega');
		
		
	ALTER TABLE users
	ADD CONSTRAINT unique_phone UNIQUE (phone);
	
	
	alter table users
	alter column name set not null;
	
	
	alter table users
	alter column phone set not null;
	
	alter table users
	alter column address set not null;