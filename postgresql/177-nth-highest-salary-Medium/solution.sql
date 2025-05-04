CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    SELECT Employee.salary
    FROM Employee
    GROUP BY Employee.salary
    ORDER BY Employee.salary DESC
    LIMIT CASE WHEN N > 0 THEN 1 ELSE 0 END
    OFFSET CASE WHEN N > 1 THEN N-1 ELSE 0 END
  );
END;
$$ LANGUAGE plpgsql;