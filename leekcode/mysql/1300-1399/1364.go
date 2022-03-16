package mysql

/*
SELECT
	invoice_id,
	customer_name,
	price,
	COUNT( c1.user_id ) contacts_cnt,
	COUNT( CASE WHEN c1.contact_name IN ( SELECT customer_name FROM Customers ) THEN 1 ELSE NULL END ) trusted_contacts_cnt
FROM
	Invoices i
	JOIN Customers c ON c.customer_id = i.user_id
	LEFT JOIN Contacts c1 ON c1.user_id = i.user_id
GROUP BY
	invoice_id
ORDER BY
    invoice_id
*/
