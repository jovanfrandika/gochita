package model

type BotConfig struct {
	Token string `json:token`
}

type DBConfig struct {
	KeyspaceName string   `json:keyspaceName`
	Clusters     []string `json:clusters`
}

type LiveChartConfig struct {
	BaseUrl string `json:baseUrl`
}

type TimeConfig struct {
	Timezone string `json:timezone`
}

type Config struct {
	Bot       BotConfig       `json:bot`
	DB        DBConfig        `json:db`
	LiveChart LiveChartConfig `json:liveChart`
	Time      TimeConfig      `json:time`
}
