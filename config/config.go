package config

type Config struct {
	DB   *DBConfig
	AUTH *AuthConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

type AuthConfig struct {
	HmacSecret        []byte
	ExpirationMinutes uint
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "gwidb",
			Port:     "5432",
			Username: "gwiuser",
			Password: "secretpassword",
			Name:     "gwi",
			Charset:  "utf8",
		},
		AUTH: &AuthConfig{
			// https://mkjwk.org
			HmacSecret:        []byte("z5qkREmStJx-IFI96YH8V7jivw78FTk6ZCtDCtq-RsYik6WopxYODgZTlWxUd_jjs6-X041ZdDn4p27lWuY3fQ"),
			ExpirationMinutes: 60 * 24 * 7, // 7 days
		},
	}
}
