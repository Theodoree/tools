package mysql

/*
SELECT
    account_id,
    day,
    SUM(IF(type="Deposit",amount,-amount)) over (partition by account_id order by day asc) balance
FROM
    Transactions
order by account_id asc, day asc;

*/
