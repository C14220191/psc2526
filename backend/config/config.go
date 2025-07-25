package config



var loadedconfig appConfig

type DBConfig struct {
	DBHost	 string `json:"db_host"`
	DBPort	 string `json:"db_port"`
	DBUser	 string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName	 string `json:"db_name"`
}
type appConfig struct {
	MODE string `json:"mode"`
	TZ   string `json:"timezone"`
	HTTPPort string `json:"http_port"`
	HTTPRoot string `json:"http_root"`
	DBConfig `json:"db_config"`
}


func new (envpath )