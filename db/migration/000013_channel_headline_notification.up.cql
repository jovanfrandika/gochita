CREATE TABLE channel_headline_notification (
  reference_id text,
  headline_id UUID,
  notified_at timestamp,
  PRIMARY KEY (reference_id, headline_id));
