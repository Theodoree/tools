package mysql

/*
SELECT
  m.member_id,
  name,
  CASE
    WHEN conversion_rate = 'No Visits' THEN 'Bronze'
    WHEN conversion_rate < 50 THEN 'Silver'
    WHEN conversion_rate >= 50 AND conversion_rate < 80 THEN 'Gold'
    ELSE 'Diamond'
  END category
FROM (
  SELECT
    m.member_id,
    IFNULL(COUNT(p.visit_id) / COUNT(v.visit_id) * 100, 'No Visits') conversion_rate
  FROM Members m
  LEFT JOIN Visits v USING(member_id)
  LEFT JOIN Purchases p USING(visit_id)
  GROUP BY 1
) t
LEFT JOIN Members m USING(member_id)
*/
