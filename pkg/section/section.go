package section

import "time"

type ServerSettingS struct {
	RunMode               string
	HttpPort              string
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	DefaultContextTimeout time.Duration //默认超时时间
}

type AppSettingS struct {
	Name            string
	Version         string
	StartTime       string
	MachineID       int64
	DefaultPageSize int
	MaxPageSize     int
}

type LogSettingS struct {
	Level         string
	LogSavePath   string // 日志保存路径
	HighLevelFile string // 日志文件名
	LowLevelFile  string // 日志文件名
	LogFileExt    string // 日志文件后缀
	MaxSize       int
	MaxAge        int
	MaxBackups    int
	Compress      bool //是否压缩过期日志
}

type MysqlSettingS struct {
	DriverName   string //驱动名
	Username     string //填写你的数据库账号
	Password     string // 填写你的数据库密码
	Host         string
	Port         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Host     string
	Port     string
	DB       int
	Password string
	PoolSize int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type UploadSettingS struct {
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}
