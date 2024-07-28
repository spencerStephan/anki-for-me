package lib

type Services struct {
	Config *UserConfig
	DB     Database
}

func InitServices(config *UserConfig, db *Sqlite) *Services {
	return &Services{
		Config: config,
		DB:     db,
	}
}
