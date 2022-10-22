CREATE TABLE manga_post (
  id UUID,
  title text,
  ref text,
  published_at timestamp,
  PRIMARY KEY (id, title, published_at));
