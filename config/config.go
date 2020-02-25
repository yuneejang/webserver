package config

import "github.com/spf13/viper"

// DefaultConfig contains default settings for the Vote-Server.
var DefaultConfig = Config{
	Nodes:    DefaultGsymNodes,
	Database: DefaultDatabaseConfig,
	API:      DefaultAPI,
}

var DefaultDatabaseConfig = DatabaseConfig{
	Driver:   "mysql",
	Database: "vote_db",
	Host:     "localhost",
	Port:     "3306",
	Username: "root",
	Password: "1234",
}

var DefaultGsymNodes = GsymNodes{
	{Host: "localhost", WsPort: "3602", HttpPort: "3502"},
}

var DefaultAPI = API{
	Enable: true,
	Port:   "3501",
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessNode string
type GsymNodes []GsymNode
type GsymNode struct {
	Host     string `json:"host"`
	WsPort   string `json:"wsPort"`
	HttpPort string `json:"httpPort"`
}

var HttpAttach string
var WsAttach string

type BlockSwitch struct {
	StartMainBlock      uint64
	StartWarrantBlock   uint64
	StartCitizenBlock   uint64
	StartOracleBlock    uint64
	DisableMainBlock    bool
	DisableWarrantBlock bool
	DisableCitizenBlock bool
	DisableOracleBlock  bool
}

type API struct {
	Enable bool   `json:"enable"`
	Port   string `json:"port"`
}

type Config struct {
	Nodes       GsymNodes      `json:"nodes"`
	Database    DatabaseConfig `json:"database"`
	BlockSwitch BlockSwitch    `json:"blockSwitch"`
	//IsDevelop   bool           `json:"isDevelop"`
	API API `json:"api"`
}

func LoadConfigFile(filename string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}
	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
