CREATE TABLE headline (
  id UUID,
  title text,
  thumbnail text,
  ref text,
  published_at timestamp,
  PRIMARY KEY (id, title, published_at));
