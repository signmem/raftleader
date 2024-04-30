package g

type GlobalConfig struct {
	Debug			bool		`json:"debug"`
	LogFile			string		`json:"logfile"`
	LogMaxAge		int			`json:"logmaxage"`
	LogRotateAge	int			`json:"logrotateage"`
	Http 			*HTTP		`json:"http"`
}

type HTTP struct {
	Address			string		`json:"address"`
	Port			string		`json:"port"`
}