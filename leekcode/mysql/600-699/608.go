package _00_699

/*
SELECT
    t.id,
    CASE
        when t.p_id IS NULL then "Root"
        when t.id IN (SELECT p_id FROM tree) then "Inner"
        else "Leaf"
		END AS type
FROM  tree AS t
*/
