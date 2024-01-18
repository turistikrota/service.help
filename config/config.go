package config

type MongoHelp struct {
	Host       string `env:"MONGO_HELP_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_HELP_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_HELP_USERNAME" envDefault:""`
	Password   string `env:"MONGO_HELP_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_HELP_DATABASE" envDefault:"help"`
	Collection string `env:"MONGO_HELP_COLLECTION" envDefault:"helps"`
	Query      string `env:"MONGO_HELP_QUERY" envDefault:""`
}

type MongoFaq struct {
	Collection string `env:"MONGO_FAQ_COLLECTION" envDefault:"faqs"`
}

type MongoCategory struct {
	Collection string `env:"MONGO_CATEGORY_COLLECTION" envDefault:"categories"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}
type App struct {
	DB struct {
		Help     MongoHelp
		Faq      MongoFaq
		Category MongoCategory
	}
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Session     Session
	RSA         RSA
	Redis       Redis
	CacheRedis  CacheRedis
	TokenSrv    TokenSrv
}
