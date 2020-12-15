package config

func GetServerConfig() map[string]interface{} {
	serverConfig := make(map[string]interface{})

	serverConfig["host"] 	= "0.0.0.0"
	serverConfig["port"] 	= "8888"

	serverConfig["mode"]	= "debug"
	return serverConfig
}
