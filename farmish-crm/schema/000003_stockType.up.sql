ALTER TABLE warehouse
ADD COLUMN stock_type VARCHAR NOT NULL UNIQUE;

ALTER TABLE warehouse
RENAME COLUMN type TO name;