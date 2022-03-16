package _700_1799

/*
SELECT
    p.product_id,
    SUM(IF(p.store ="store1",price,null)) AS store1,
    SUM(IF(p.store ="store2",price,null)) AS store2,
    SUM(IF(p.store ="store3",price,null)) AS store3
FROM
    Products AS p
GROUP BY
    p.product_id

*/
