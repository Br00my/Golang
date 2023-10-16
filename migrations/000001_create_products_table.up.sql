CREATE TABLE IF NOT EXISTS products (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  title text NOT NULL,
  category text NOT NULL,
  price integer NOT NULL
);
