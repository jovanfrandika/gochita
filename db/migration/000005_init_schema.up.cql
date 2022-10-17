CREATE TABLE channelShowSubscription (
  referenceId text,
  showId UUID,
  isEnabled boolean,
  PRIMARY KEY (referenceId, showId));
