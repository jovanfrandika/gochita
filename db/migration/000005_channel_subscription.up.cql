CREATE TABLE channel_subscription (
  id UUID,
  subscription_type int,
  reference_id text,
  context_id UUID,
  is_enabled boolean,
  PRIMARY KEY (subscription_type, reference_id, context_id));
