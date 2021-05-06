package restapi

import "net/http"

func Start(config *Config) error {

	serv, err := NewServer(config.LogLevel)
	if err != nil {
		return err
	}
	return http.ListenAndServe(config.BindAddress, serv)
}
