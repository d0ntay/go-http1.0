package api

type Config struct {
	Addr string
}

func NewConfig(addr string) (*Config, error) {
	return &Config{
		Addr: addr,
	}, nil
}
