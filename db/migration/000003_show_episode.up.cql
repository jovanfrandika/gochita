CREATE TABLE show_episode (
  id UUID,
  show_id UUID,
  num int,
  published_at timestamp,
  PRIMARY KEY (show_id, num, published_at));
