package goapiserver

type ConfigServer struct {
	BindAddr string
	LogLevel string
}

func NewConfigServer() *ConfigServer {
	return &ConfigServer{
		BindAddr: ":7070",
		LogLevel: "debug",
	}

}
