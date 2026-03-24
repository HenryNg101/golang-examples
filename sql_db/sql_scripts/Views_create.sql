-- Create and use normal view
create or replace view order_items_filter as
select *
from order_items
where price > 200;

create materialized view order_items_filter_copy as
select *
from order_items
where price > 200;