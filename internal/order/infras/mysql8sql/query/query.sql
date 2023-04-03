-- name: DelOrderById :execrows
update xf_order
set delete_is =?
where order_id = ?;

-- name: GetOrderById :one
select *
from xf_order
where order_id = ?;