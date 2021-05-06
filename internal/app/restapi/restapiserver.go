package restapi

import "net/http"

func Start(config *Config) error {

	serv := NewServer()
	return http.ListenAndServe(config.BindAddress, serv)
}
