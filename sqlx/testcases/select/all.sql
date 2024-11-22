SELECT a, b, (SELECT * FROM test LIMIT 1) AS "tester" FROM test WHERE a = 1 OR (a = 2 AND b = 3) GROUP BY a ORDER BY a ASC;
