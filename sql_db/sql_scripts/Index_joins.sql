--SELECT *
--FROM orders o
--join order_items oi
--ON oi.orderid  = o.id
--join users u
--on u.id = o.userid;

explain analyze
SELECT o.id, oi.id, o.price, oi.price
FROM orders o
join order_items oi
ON oi.orderid  = o.id
where o.id < 10;


CREATE INDEX idx_order_items_orderid ON order_items(id);

explain analyze
SELECT o.id, oi.id, o.price, oi.price
FROM orders o
join order_items oi
ON oi.orderid  = o.id
where o.id < 10;

drop index idx_order_items_orderid;


CREATE INDEX idx_order_items_orderid ON order_items(orderid);
CREATE INDEX idx_orders_id ON orders(id);

explain analyze
SELECT o.id, oi.id, o.price, oi.price
FROM orders o
join order_items oi
ON oi.orderid = o.id
where o.id < 10;

drop index idx_order_items_orderid;
drop index idx_orders_id;


-- Test merge join vs hash join
SET work_mem = '1024MB';

CREATE INDEX idx_order_items_orderid ON order_items(orderid);
CREATE INDEX idx_orders_id ON orders(id);

explain analyze
SELECT o.id, oi.id, o.price, oi.price
FROM orders o
join order_items oi
ON oi.orderid = o.id;

drop index idx_order_items_orderid;
drop index idx_orders_id;

CREATE INDEX idx_order_items_orderid ON order_items(orderid);
CREATE INDEX idx_orders_id ON orders(id);
CLUSTER order_items USING idx_order_items_orderid;
SET enable_hashjoin = off;

explain analyze
SELECT o.id, oi.id, o.price, oi.price
FROM orders o
join order_items oi
ON oi.orderid = o.id;

SET enable_hashjoin = on;
drop index idx_order_items_orderid;
drop index idx_orders_id;