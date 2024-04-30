package g

type GlobalConfig struct {
	Debug			bool		`json:"debug"`
	LogFile			string		`json:"logfile"`
	LogMaxAge		int			`json:"log_max_age"`
	LogRotateAge	int			`json:"log_rotate_age"`
	Redis			*Rconfig	`json:"redis"`
	Http 			*HTTP		`json:"http"`
}

type HTTP struct {
	Address			string		`json:"address"`
	Port			string		`json:"port"`
}


type Rconfig struct {
	Enabled			bool			`json:"enabled"`
	Server			string			`json:"server"`
	Port			string			`json:"port"`
	MaxIdle			int				`json:"max_idle"`
	MaxActive		int				`json:"max_active"`
	IdleTimeOut		int				`json:"idle_timeout"`
	LockKey			string 			`json:"lock_key"`
	LockTime 		int				`json:"lock_time"`
	AskLockTime		int 			`json:"ask_lock_time"`
}