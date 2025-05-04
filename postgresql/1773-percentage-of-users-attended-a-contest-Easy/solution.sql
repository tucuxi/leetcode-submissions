-- Write your PostgreSQL query statement below
SELECT contest_id, ROUND(COUNT(user_id) * 100::NUMERIC / (SELECT COUNT(*) FROM Users), 2) AS percentage
FROM Register
GROUP BY contest_id
ORDER BY 2 DESC, contest_id