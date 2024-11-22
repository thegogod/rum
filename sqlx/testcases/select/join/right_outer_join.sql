SELECT * FROM a RIGHT OUTER JOIN b ON a.id = b.id AND b.deleted_at IS NULL;
