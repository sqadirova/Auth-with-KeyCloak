package config

type InfraConfigurations struct {
	App
	Database
	Jwt
	Swagger
}

type App struct {
	Host string
	Port int
}

type Swagger struct {
	Host   string
	Scheme string
}

type Database struct {
	Host                    string
	Port                    int
	Dialect                 string
	User                    string
	DBName                  string
	Password                string
	GormMaxIdleConn         int
	GormMaxOpenConn         int
	GormMaxConnLifetimeHour int
	SSLMode                 string
	Schema                  string
}

type Jwt struct {
	SecretKey         string
	PublicKey         string
	PrivateKey        string
	SaltKey           string
	AccessTokenExpire int64
}
