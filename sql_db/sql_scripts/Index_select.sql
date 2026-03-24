-- Another test with multi-columns index
-- Test 1
CREATE INDEX idx_orders_items_orders_price ON order_items(orderid, price);

explain analyze
select *
from order_items
where orderid = 10;

explain analyze
select *
from order_items
where price = 200;

drop INDEX idx_orders_items_orders_price;

-- Test 2
CREATE INDEX idx_orders_items_orders_price ON order_items(price, orderid);

explain analyze
select *
from order_items
where orderid = 10;

explain analyze
select *
from order_items
where price = 200;

drop INDEX idx_orders_items_orders_price;


-- Test 3
CREATE INDEX idx_orders_items_orders_price ON order_items(orderid, price);

explain analyze
select *
from order_items
where price > 200;

drop INDEX idx_orders_items_orders_price;

-- Test 4
CREATE INDEX idx_orders_items_orders_price ON order_items(price, orderid);

explain analyze
select *
from order_items
where price > 200;

drop INDEX idx_orders_items_orders_price;

-- Test 5
CREATE INDEX idx_orders_items_price ON order_items(price);

explain analyze
select *
from order_items
where price > 200;

drop INDEX idx_orders_items_price;

-- Unique index test
create unique index idx_orders_items_price on order_items(id);
drop index idx_orders_items_price;