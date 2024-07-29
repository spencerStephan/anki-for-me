package lib

type Services struct {
	DB     Database
	Config Config
}

func InitServices(config Config, db *Sqlite) *Services {
	return &Services{
		Config: config,
		DB:     db,
	}
}
