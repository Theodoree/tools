package mysql

/*
SELECT
    customer_id,
    COUNT(1) AS count_no_trans
FROM
    Visits AS v
WHERE
    visit_id not in (SELECT  visit_id FROM Transactions )
GROUP BY customer_id
*/
