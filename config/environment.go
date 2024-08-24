package config

type Environment struct {
	ServerAddr string    `yaml:"http_server_addr"`
	LogLevel   string    `yaml:"log_level"`
	GRPCPeers  gRPCPEers `yaml:"grpc_servers_addr"`
}
