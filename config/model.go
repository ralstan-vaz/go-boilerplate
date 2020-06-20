package config

type Config struct {
	AppVersion string `yaml:"appVersion"`
	Server     Server `yaml:"server"`
	User       User   `yaml:"user"`
}

type Server struct {
	GRPC GRPC `yaml:"grpc"`
	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	Address string `yaml:"address"`
}

type GRPC struct {
	Address string `yaml:"address"`
}

type User struct {
	RatingsUrl    string `yaml:"ratingsUrl"`
	FavouritesUrl string `yaml:"favouritesUrl"`
}
