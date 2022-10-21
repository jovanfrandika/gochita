# Gochita

<p align="center">
  <img src="files/images/gochita.png" alt="gochita" width="200" style="padding: 32px 0;" />
</p>

Gochita is a discord bot for notifying latest published Anime Series/Movies and headlines to your dedicated discord channel.

Gochita stores livechart RSS feed that into a CassandraDB under the hood.

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

/show list - To show channel's subscribed show

/show subscribe all - To subscribe new shows to a channel

/show unsubscribe all - To unsubscribe new shows to a channel

/show subscribe one query:<show title> - To subscribe a show to a channel

/show unsubscribe one query:<show title> - To unsubscribe a show to a channel

/headline subscribe all - To subscribe new headlines to a channel

/headline unsubscribe all - To unsubscribe new headlines to a channel

# Roadmap

- Notify new manga update

## Contributing
```
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
