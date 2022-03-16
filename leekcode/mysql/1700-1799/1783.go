package _700_1799

/*
SELECT
	p.player_id,
	p.player_name,
	SUM(IF(c.Wimbledon=p.player_id,1,0))+
	SUM(IF(c.Fr_open=p.player_id,1,0))+
	SUM(IF(c.US_open=p.player_id,1,0))+
	SUM(IF(c.Au_open=p.player_id,1,0)) AS grand_slams_count
FROM
	Players AS p , Championships AS c
group by
	p.player_id, p.player_name
HAVING grand_slams_count > 0

*/
