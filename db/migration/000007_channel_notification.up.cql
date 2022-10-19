CREATE TABLE channel_notification (
  channel_subscription_id UUID,
  notified_at timestamp,
  PRIMARY KEY (channel_subscription_id));
