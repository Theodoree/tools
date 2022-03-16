package mysql

/*
SELECT po1.id p1,po2.id p2,
ABS(po2.x_value-po1.x_value)*ABS(po2.y_value-po1.y_value) AREA
FROM Points po1,Points po2
WHERE po2.id > po1.id AND po2.x_value!=po1.x_value AND po2.y_value!=po1.y_value
ORDER BY AREA DESC,p1,p2
*/
