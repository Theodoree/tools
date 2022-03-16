package _700_1799

/*
SELECT
	DISTINCT l.account_id
FROM
	LogInfo AS l,LogInfo AS l1
WHERE
		l.account_id = l1.account_id &&  l1.login >= l.login AND l1.login <= l.logout  && l.ip_address != l1.ip_address
*/
