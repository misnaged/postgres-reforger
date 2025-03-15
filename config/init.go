package config

import (
	"github.com/spf13/viper"
)

// init initialize default config params
func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("env", "prod")

	// database
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "serega123")
	viper.SetDefault("db.name", "rbs")
	viper.SetDefault("db.dbtype", "pg") // Could be pg (Postgres) or sqlite (SQLite)

}
