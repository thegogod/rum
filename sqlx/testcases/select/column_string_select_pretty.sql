SELECT
    1,
    2,
    (
        SELECT
            a,
            b,
            c
        FROM test
    ) as "results";
