-- Write your PostgreSQL query statement below
SELECT product_name, SUM(unit) AS unit
FROM Products JOIN Orders USING (product_id)
WHERE order_date BETWEEN '2020-02-01' AND '2020-02-29'
GROUP BY product_id, product_name
HAVING SUM(unit) >= 100