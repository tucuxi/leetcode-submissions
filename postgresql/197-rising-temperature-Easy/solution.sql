-- Write your PostgreSQL query statement below
SELECT id
FROM Weather w
WHERE temperature > (SELECT temperature FROM Weather WHERE recordDate = w.recordDate - 1)