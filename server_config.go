package goapiserver

type ConfigServer struct {
	BindAddr string
	LogLevel string
}

func NewConfigServer() *ConfigServer {
	return &ConfigServer{
		BindAddr: ":9090",
		LogLevel: "debug",
	}

}
