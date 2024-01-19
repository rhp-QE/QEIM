package config

var NetConfig = &netConfig {
	Address: "127.0.0.1:4979",
	NetWork: "tcp",
}

type netConfig struct {
	Address  string
	NetWork  string
}