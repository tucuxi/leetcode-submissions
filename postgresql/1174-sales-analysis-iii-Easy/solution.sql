-- Write your PostgreSQL query statement below
SELECT DISTINCT product_id, product_name 
FROM Product p JOIN Sales s USING (product_id)
WHERE EXISTS (
    SELECT product_id
    FROM Sales 
    WHERE product_id = p.product_id AND sale_date BETWEEN '2019-01-01' AND '2019-03-31' 
) AND NOT EXISTS (
    SELECT product_id
    FROM Sales 
    WHERE product_id = p.product_id AND sale_date NOT BETWEEN '2019-01-01' AND '2019-03-31' 
)