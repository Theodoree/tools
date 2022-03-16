package _00_699

/*
WITH tmp_table AS (
SELECT
	*,
	row_number() OVER(PARTITION BY continent ORDER BY name) AS rk // 这里之前使用的是RANK,我猜测应该是name相同
FROM
	student
ORDER BY name
)


SELECT
    max(if(continent='America',name,null)) America,
    max(if(continent='Asia',name,null)) Asia,
    max(if(continent='Europe',name,null)) Europe
FROM
	tmp_table AS t
GROUP BY t.rk

*/
