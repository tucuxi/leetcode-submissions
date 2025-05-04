-- Write your PostgreSQL query statement below
SELECT SalesPerson.name
FROM SalesPerson
WHERE sales_id NOT IN (
    SELECT Orders.sales_id
    FROM Orders JOIN Company ON Orders.com_id = Company.com_id
    WHERE Company.name = 'RED'
)
