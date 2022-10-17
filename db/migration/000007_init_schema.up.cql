CREATE TABLE channelShowEpisodeNotification (
  referenceId text,
  showEpisodeId UUID,
  notifiedAt timestamp,
  PRIMARY KEY (referenceId, showEpisodeId));
