-- Write your PostgreSQL query statement below
SELECT Prices.product_id,
    COALESCE(ROUND(SUM(price::numeric(10, 2) * units) / SUM(units), 2), 0) AS average_price
FROM Prices LEFT JOIN UnitsSold ON UnitsSold.product_id = Prices.product_id
    AND UnitsSold.purchase_date BETWEEN Prices.start_date AND Prices.end_date
GROUP BY 1