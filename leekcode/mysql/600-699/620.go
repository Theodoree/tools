package _00_699

/*
620. 有趣的电影

SELECT
	id,
	movie,
	description,
	rating
FROM
	cinema
WHERE
	description != 'boring' AND MOD(id,2) = 1
ORDER BY rating DESC
*/
