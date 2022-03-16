package _800_1899

/*
SELECT
    employee_id,
    IF(employee_id%2=1 AND SUBSTR(name FROM 1 FOR 1) != "M",salary,0) AS bonus
FROM
    Employees
*/
