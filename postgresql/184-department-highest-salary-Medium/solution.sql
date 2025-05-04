-- Write your PostgreSQL query statement below
WITH HighestSalary AS (
    SELECT
        Department.id AS departmentId,
        MAX(salary) AS salary
    FROM
        Employee JOIN Department ON Employee.departmentId = Department.id
    GROUP BY Department.id
)
SELECT
    Department.name AS department,
    Employee.name AS employee,
    Employee.salary AS salary
FROM
    Employee
    JOIN Department ON Employee.departmentId = Department.id
    JOIN HighestSalary ON Employee.departmentId = HighestSalary.departmentId AND Employee.salary = HighestSalary.salary
