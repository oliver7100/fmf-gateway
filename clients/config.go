package clients

type clientConfig struct {
	Url string
}

func NewConfig(url string) *clientConfig {
	return &clientConfig{
		Url: url,
	}
}
