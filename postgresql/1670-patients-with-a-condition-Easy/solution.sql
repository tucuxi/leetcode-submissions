-- Write your PostgreSQL query statement below
SELECT *
FROM Patients
WHERE EXISTS (
    SELECT s
    FROM string_to_table(conditions, ' ') c(s)
    WHERE starts_with(s, 'DIAB1')
)