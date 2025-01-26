ALTER TABLE warehouse
DROP COLUMN stock_type;

ALTER TABLE warehouse
RENAME COLUMN name TO type;