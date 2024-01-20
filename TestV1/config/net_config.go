package config

var NetConfig = &netConfig {
	Address: "127.0.0.1:1959",
	NetWork: "tcp",
}

type netConfig struct {
	Address  string
	NetWork  string
}