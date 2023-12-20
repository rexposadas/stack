
-- https://leetcode.com/problems/customers-who-never-order/description/

-- Write your MySQL query statement below
select name as 'Customers'
from Customers
where id NOT IN (select customerId from Orders)