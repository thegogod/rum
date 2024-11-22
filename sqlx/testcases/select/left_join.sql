SELECT * FROM a LEFT JOIN b ON a.id = b.id AND b.deleted_at IS NULL;
