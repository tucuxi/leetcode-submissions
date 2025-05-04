-- Write your PostgreSQL query statement below
SELECT name, SUM(amount) AS balance
FROM Users JOIN Transactions ON Users.account = Transactions.account
GROUP BY name
HAVING SUM(amount) > 10000