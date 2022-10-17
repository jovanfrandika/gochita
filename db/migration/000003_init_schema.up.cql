CREATE TABLE showEpisode (
  id UUID,
  showId UUID,
  num int,
  publishedAt timestamp,
  PRIMARY KEY (showId, num, publishedAt));
