-- Write your PostgreSQL query statement below
SELECT
    sell_date,
    count(product) AS num_sold,
    string_agg(product, ',' ORDER BY product) AS products
FROM (
    SELECT DISTINCT sell_date, product
    FROM Activities
)
GROUP BY sell_date
ORDER BY sell_date