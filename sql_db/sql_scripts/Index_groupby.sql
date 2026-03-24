--No index
explain analyze
select orderid, sum(price)
from order_items
group by orderid;


-- Index only on foreign key order id
CREATE INDEX idx_order_items_orderid ON order_items(orderid);

explain analyze
select orderid, sum(price)
from order_items
group by orderid;

drop index idx_order_items_orderid;


-- Index on price only
CREATE INDEX idx_order_items_price ON order_items(price);

explain analyze
select orderid, sum(price)
from order_items
group by orderid;

drop index idx_order_items_price;


-- Index on both foreign key and price, but not together
CREATE INDEX idx_order_items_orderid ON order_items(orderid);
CREATE INDEX idx_order_items_price ON order_items(price);

explain analyze
select orderid, sum(price)
from order_items
group by orderid;

drop index idx_order_items_orderid;
drop index idx_order_items_price;


-- Index on both of them together
CREATE INDEX idx_order_items_orderid_price ON order_items(orderid, price);

explain analyze
select orderid, sum(price)
from order_items
group by orderid;

drop index idx_order_items_orderid_price;

-- Index on both of them together, reverse order
CREATE INDEX idx_order_items_price_orderid ON order_items(price, orderid);

explain analyze
select orderid, sum(price)
from order_items
group by orderid;

drop index idx_order_items_price_orderid;