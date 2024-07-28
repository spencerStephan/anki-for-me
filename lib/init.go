package lib

type InitService struct {
	Config Config
	DB     Database
}

func NewInitService(config UserConfig, db Sqlite) *InitService {
	return &InitService{
		Config: config,
		DB:     db,
	}
}
