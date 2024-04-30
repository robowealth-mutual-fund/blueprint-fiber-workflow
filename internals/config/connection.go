package config

type Connection struct {
	MaximumMessageSize int    `env:"MAXIMUM_MESSAGE_SIZE" envDefault:"30"`
	GRPCTimeout        int    `json:"GRPC_TIMEOUT" envDefault:"300"`
	BluePrintService   string `env:"BLUE_PRINT_SERVICE" envDefault:"0.0.0.0:3001"`
}
