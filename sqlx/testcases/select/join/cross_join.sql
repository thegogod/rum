SELECT * FROM a CROSS JOIN b ON a.id = b.id AND b.deleted_at IS NULL;
