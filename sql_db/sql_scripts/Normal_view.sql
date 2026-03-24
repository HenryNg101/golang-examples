-- Should be same content
select * from order_items_filter limit 100;
select * from order_items_filter_copy limit 100;
select * from order_items where id = 1;

update order_items_filter
set price = 100
where id = 1;

-- Count of order_items_filter would be one less than order_items_filter_copy
select count(*) from order_items_filter;
select count(*) from order_items_filter_copy;
select * from order_items where id = 1;		-- Data change


-- Replace the view, by enforcing rule to it
create or replace view order_items_filter as
select *
from order_items
where price > 200
with check option;


update order_items
set price = 300
where id = 1;


select count(*) from order_items_filter;

-- This would violate the rule and give error
--update order_items_filter
--set price = 100
--where id = 1;