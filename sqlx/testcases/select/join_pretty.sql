SELECT
    *
FROM a
JOIN b
    ON a.id = b.id
    AND b.deleted_at IS NULL;
