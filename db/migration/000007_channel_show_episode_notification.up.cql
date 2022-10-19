CREATE TABLE channel_show_episode_notification (
  reference_id text,
  show_episode_id UUID,
  notified_at timestamp,
  PRIMARY KEY (reference_id, show_episode_id));
