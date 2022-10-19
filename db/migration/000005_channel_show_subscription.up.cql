CREATE TABLE channel_show_subscription (
  reference_id text,
  show_id UUID,
  is_enabled boolean,
  PRIMARY KEY (reference_id, show_id));
