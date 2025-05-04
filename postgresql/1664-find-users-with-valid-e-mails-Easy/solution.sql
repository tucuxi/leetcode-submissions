-- Write your PostgreSQL query statement below
SELECT user_id, name, mail
FROM Users
WHERE mail ~* '^[a-z][a-z0-9_.-]*@leetcode\.com$'