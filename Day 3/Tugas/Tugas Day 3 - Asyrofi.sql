-- drop all old table
drop table if exists customers cascade;
drop table if exists products cascade;
drop table if exists orders cascade;

-- create table customers, products, orders
create table if not exists customers (
	id serial primary key,
	customer_name char(50) not null
);

create table if not exists products (
	id serial primary key,
	name char(50) not null
);

create table if not exists orders (
	order_id serial primary key,
	customer_id int not null references customers (id),
	product_id int not null references products (id),
	order_date date not null,
	total float not null
);

-- insert sample record to all table
insert into customers (customer_name)
	values ('alphandi'), ('betania'), ('charlie');

insert into products (name)
	values ('Pakan Lele'), ('Pakan Nila'), ('Pakan Udang');

insert into orders 	(customer_id, product_id, order_date, total)
	values (1, 1, now(), 45000),
		   (1, 2, now(), 52000),
		   (1, 3, now(), 82000),
		   (2, 1, now(), 45000),
		   (2, 2, now(), 52000),
		   (2, 3, now(), 82000);
			
-- update some record
update customers set customer_name = 'deltanudin' where id = 1;
update products set name = 'Pakan Bandeng' where id = 2;
update orders set total = 93000, product_id = 1 where order_id = 3;

-- delete some record
delete from customers where id = 3;
delete from orders where order_id  = 1;

-- join table customers, products, orders 
select	o.order_id, c.id as customer_id, c.customer_name, p.id as product_id, p.name as product_name, o.order_date, o.total
	from orders o
		inner join customers c on o.customer_id = c.id
		inner join products p on o.product_id = p.id;
	

