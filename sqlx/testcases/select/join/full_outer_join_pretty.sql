SELECT
    *
FROM a
FULL OUTER JOIN b
    ON a.id = b.id
    AND b.deleted_at IS NULL;
