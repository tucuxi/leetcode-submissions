-- Write your PostgreSQL query statement below
SELECT DISTINCT COALESCE(Employees.employee_id, Salaries.employee_id) AS employee_id
FROM Employees FULL OUTER JOIN Salaries ON Employees.employee_id = Salaries.employee_id
WHERE name IS NULL OR salary IS NULL 
ORDER BY employee_id