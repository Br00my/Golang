ALTER TABLE products ADD CONSTRAINT products_price_check CHECK (price > 0);

ALTER TABLE products ADD CONSTRAINT products_category_check CHECK (category in ('Reusable Water Bottles', 'Bamboo Toothbrushes', 'Solar Charges', 'Eco-Friendly Cleaning Products', 'Recycled Paper Products'));
