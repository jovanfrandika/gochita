# Gochita

<p align="center">
  <img src="files/images/gochita.png" alt="gochita" width="200" style="padding: 32px 0;" />
</p>

Gochita is a Discord bot for notifying you of the latest published anime series/movies, headlines, and manga chapters on your dedicated Discord channel.

Gochita stores the RSS feeds into a CassandraDB. The data will be used to figure out recent shows, headlines, and manga chapters that need to be notified to a discord channel and mark them once they are notified.

Currently, Gochita consists of two binaries. The first one is the RSS feed reader to renew fresh shows and headlines. The second one is the Discord bot client.

## Configuration

Make a configuration file at /.env

Example configuration file
```dosini
TOKEN=""

KEYSPACE_NAME= "gochita"
CLUSTER="cassandra:9042"
TIMEOUT="5"

LIVECHART_BASE_URL="https://www.livechart.me"
LIVECHART_LATEST_EPISODES_URI="/feeds/episodes"
LIVECHART_LATEST_HEADLINES_URI="/feeds/headlines"

REDDIT_BASE_URL="https://www.reddit.com"
REDDIT_LATEST_MANGA_POSTS_URI="/r/manga/new/.rss"

TIMEZONE="Asia/Jakarta"
DEFAULT_TIMEOUT="1"
NOTIFY_TIMEOUT="5"
SET_COMMANDS_TIMEOUT="10"
NOTIFY_SHOWS_INTERVAL="5"
NOTIFY_HEADLINES_INTERVAL="10"
NOTIFY_MANGAS_INTERVAL="5"
ADD_SHOWS_INTERVAL="60"
ADD_HEADLINES_INTERVAL="60"
ADD_MANGAS_INTERVAL="60"
```

## Installation

To migrate database, run this
```bash
make migrate-up DB="example"
```

## Commands

/show list - To show channel's subscribed show

/show subscribe new - To subscribe new shows to a channel

/show unsubscribe new - To unsubscribe new shows to a channel

/show subscribe one query:<show title> - To subscribe a show to a channel

/show unsubscribe one query:<show title> - To unsubscribe a show to a channel

/show unsubscribe all - To unsubscribe all shows to a channel

/headline subscribe new - To subscribe new headlines to a channel

/headline unsubscribe new - To unsubscribe new headlines to a channel

/manga subscribe new - To subscribe new manga posts to a channel

/manga unsubscribe new - To unsubscribe new manga posts to a channel

# Roadmap

- Show list pagination
- Woof voice effect
- Full text search on show titles

## Contributing
```
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
