-- Write your PostgreSQL query statement below
SELECT user_id, upper(substring(name FOR 1)) || lower(substring(name FROM 2)) AS name
FROM Users
ORDER BY user_id