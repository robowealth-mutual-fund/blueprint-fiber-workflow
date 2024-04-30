package config

type Server struct {
	GRPCPort          string `env:"GRPC_SERVER_PORT" envDefault:"4001"`
	HTTPPort          string `env:"HTTP_SERVER_PORT" envDefault:"4002"`
	HTTPServerTimeout int64  `env:"HTTP_SERVER_TIMEOUT" envDefault:"0"`
}
