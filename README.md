# Judah

Judah is a discord bot for notifying latest published Anime Series/Movies to your discord channel.

Judah stores livechart RSS feed that into a CassandraDB under the hood.

## Configuration

Make a configuration file at files/config.json

Example configuration file
```JSON
{
  "bot": {
    "token": ""
  },
  "db": {
    "keyspaceName": "",
    "clusters": []
  },
  "liveChart": {
    "baseUrl": "https://livechart.me"
  },
  "time": {
    "timezone": "Asia/Jakarta"
  }
}
```

## Installation

To migrate database, run this
```bash
make migrate-up db="cassandra://localhost:9042/example"
```

To use the livechart data grabber, run this
```bash
make build-livechart && make run-livechart
```

To turn on the discord bot, run this
```bash
make build-bot && make run-bot
```

## Commands

/show-list - To show channel's subscribed show

/show-subscribe-all - To subscribe new shows to a channel

/show-unsubscribe-all - To unsubscribe new shows to a channel

/show-subscribe query:<show title> - To subscribe a show to a channel

/show-unsubscribe query:<show title> - To unsubscribe a show to a channel

/headline-subscribe-all - To subscribe new headlines to a channel

/headline-unsubscribe-all - To unsubscribe new headlines to a channel

## Contributing
```
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
