package config

type Env string

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

type Config struct {
	Env        string `env:"ENV,default=developments"`
	Api        Api
	PostgresDB PostgresDB
	Telemetry  Telemetry
}

type Api struct {
	Addr string
	Host string `env:"API_HOST,default=localhost"`
	Port string `env:"API_PORT,default=8080"`
}

type PostgresDB struct {
	DSN      string
	Host     string `env:"POSTGRES_HOST,default=localhost"`
	Port     string `env:"POSTGRES_PORT,default=5432"`
	Name     string `env:"POSTGRES_NAME,default=rolldeck_db"`
	User     string `env:"POSTGRES_USER,default=admin"`
	Password string `env:"POSTGRES_PASSWORD,default=password"`
}

type Telemetry struct {
	ServiceName string `env:"TELEMETRY_SERVICE_NAME,default=rolldeck"`
	Endpoint    string `env:"TELEMETRY_ENDPOINT,default=localhost:4318"`
}
