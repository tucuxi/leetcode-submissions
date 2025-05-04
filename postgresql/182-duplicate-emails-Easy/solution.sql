-- Write your PostgreSQL query statement below
SELECT email
FROM Person
GROUP BY email
HAVING count(*) > 1