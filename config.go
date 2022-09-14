package usrp

type Config struct {
	RXPort string // USRP UDP packets will come in here host:port.. i.e ":30000"
	TXPort string // USRP UDP packets will go out here host:port.. i.e. "127.0.0.1:30001"
	Group  string // name for what you are capturing
}

func ConfigFromTomlFile(filepath string) (*Config, error) {
	return &Config{}, nil
}
