CREATE INDEX IF NOT EXISTS products_title_idx ON products USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS products_category_idx ON products USING GIN (to_tsvector('simple', category));
