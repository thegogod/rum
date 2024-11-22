SELECT a, b, c FROM test WHERE a = b AND (SELECT * FROM tester) IS NULL;
