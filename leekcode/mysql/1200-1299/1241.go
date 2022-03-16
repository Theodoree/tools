package mysql

/*
SELECT
	sub_id AS post_id,
	IFNULL((SELECT COUNT(DISTINCT sub_id) FROM Submissions  WHERE parent_id = s.sub_id),0) AS number_of_comments
FROM
	Submissions  AS s
WHERE
	parent_id IS NULL
GROUP BY
		sub_id
	ORDER BY sub_id


	select temp.sub_id as post_id, count(distinct s.sub_id) as number_of_comments from
(select distinct sub_id from submissions where parent_id is null) as temp left join submissions s on temp.sub_id = s.parent_id
group by post_id
*/
