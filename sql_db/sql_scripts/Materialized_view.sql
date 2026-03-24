-- This would give errors, because you can't update through it
--update order_items_filter_copy
--set price = 100
--where id = 1;

-- Update with latest data
refresh materialized view order_items_filter_copy