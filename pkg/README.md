# 编写配置文档

```yaml
Server: #运行服务
  RunMode: debug #运行模式
  HttpPort: 8080 #运行端口
  ReadTimeout: 3
  WriteTimeout: 3
  DefaultContextTimeout: 5 #默认超时时间
App: #项目逻辑
  Name: test #项目名
  Version: 1.0.0 #版本号
  StartTime: 2002-03-26 #开始日期
  MachineID: 1 #机器ID(用于雪花算法生成id)
  DefaultPageSize: 5 #请求页数的大小
  MaxPageSize: 100 #最大页数
Log: #日志
  Level: debug # 日志模式(用于是否在控制台打印)
  LogSavePath: storage/logs/ # 日志保存路径
  HighLevelFile: error          # 高级别日志文件名
  LowLevelFile: info          # 低级别文件名
  LogFileExt: .log          # 日志文件后缀
  MaxSize: 200              # 最大大小(M)
  MaxAge: 30                # 最大保存时间(天)
  MaxBackups: 7             # 最大备份数
  Compress: false           # 是否压缩过期日志
Mysql: #Mysql
  DriverName: mysql #驱动名
  Username: XXX  # 填写你的数据库账号
  Password: XXX  # 填写你的数据库密码
  Host: 127.0.0.1     # ip
  Port: 3306          # 端口
  DBName: webapp      # 数据库名
  Charset: utf8       # 字符编码
  ParseTime: True     # 是否解析时间戳
  MaxIdleConns: 200
  MaxOpenConns: 50
Redis:
  Host: 
  Port: 
  DB: 0
  Password: WW876001
  PoolSize: 100 #连接池
JWT:
  Secret: raja # 密钥
  Issuer: raja # 签发者
  Expire: 7200 # 过期时间7200秒
Email:
  Host: smtp.qq.com
  Port: 465
  UserName: XXX@qq.com
  Password: XXX@qq.com
  IsSSL: true
  From: XXX@qq.com
  To:
    - XXX@qq.com

```

# 设置全局配置变量

```go
package global

import (
	"webApp/pkg/logger"
	"webApp/pkg/section"
)

var (
	ServerSetting = new(section.ServerSettingS)
	AppSetting    = new(section.AppSettingS)
	LogSetting    = new(section.LogSettingS)
	MysqlSetting  = new(section.MysqlSettingS)
	RedisSetting  = new(section.RedisSettingS)
	JWTSetting    = new(section.JWTSettingS)
	EmailSetting  = new(section.EmailSettingS)
	UploadSetting = new(section.UploadSettingS)
)
var (
	Logger = new(logger.Log)
)
```

# 从配置文件中读取配置

```go
...
var (
    port        string
    runMode     string
    configPaths string
    configName  string
    configType  string
)
//命令行参数绑定
func setupFlag() {
    flag.StringVar(&port, "port", "", "启动端口")
    flag.StringVar(&runMode, "mode", "", "启动模式")
    flag.StringVar(&configName, "name", "config", "配置文件名")
    flag.StringVar(&configType, "type", "yaml", "配置文件类型")
    flag.StringVar(&configPaths, "path", "configs/", "指定要使用的配置文件路径")
    flag.Parse()
}

//初始化配置
func setupSetting() error {
    //读取配置文件
    newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("Server", global.ServerSetting)
    if err != nil {
        return err
    }
	global.ServerSetting.ReadTimeout *= time.Second
    global.ServerSetting.WriteTimeout *= time.Second
    global.ServerSetting.DefaultContextTimeout *= time.Second
    err = newSetting.ReadSection("App", global.AppSetting)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("Log", global.LogSetting)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("Mysql", global.MysqlSetting)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("Redis", global.RedisSetting)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("JWT", global.JWTSetting)
    if err != nil {
        return err
    }
    global.JWTSetting.Expire *= time.Second
    err = newSetting.ReadSection("Email", global.EmailSetting)
    if err != nil {
        return err
    }
    err = newSetting.ReadSection("Upload", global.UploadSetting)
    if err != nil {
        return err
    }
    if runMode != "" {
        global.ServerSetting.RunMode = runMode
    }
    if port != "" {
    global.ServerSetting.HttpPort = port
    }
    return nil
}
//初始化日志
func setupLogger() {
    logger.Init(&logger.InitStruct{
        LogSavePath:   global.LogSetting.LogSavePath,
        LogFileExt:    global.LogSetting.LogFileExt,
        MaxSize:       global.LogSetting.MaxSize,
        MaxBackups:    global.LogSetting.MaxBackups,
        MaxAge:        global.LogSetting.MaxAge,
        Compress:      global.LogSetting.Compress,
        LowLevelFile:  global.LogSetting.LowLevelFile,
        HighLevelFile: global.LogSetting.HighLevelFile,
    })
    global.Logger = logger.NewLogger(global.LogSetting.Level)
}
//初始化
func init(){
    //1.初始化分页器
    app.Init(global.AppSetting.DefaultPageSize, global.AppSetting.MaxPageSize)
    //2.初始化日志
    setupLogger()
    defer global.Logger.Sync()
    var err error
	//3.初始化mysql
	if err = mysql.Init((*mysql.InitStruct)(global.MysqlSetting)); err != nil {
        global.Logger.Error("init mysql failed,err:" + err.Error())
        return
    }
    defer mysql.Close()
    //4.初始化redis
    if err = redis.Init((*redis.InitStruct)(global.RedisSetting)); err != nil {
        global.Logger.Error("init redis failed,err:" + err.Error())
        return
    }
    defer redis.Close()
    //初始化生成userID算法
	if err = snowflake.Init(global.AppSetting.StartTime, global.AppSetting.MachineID); err != nil {
        global.Logger.Error("init sonyflake failed,err:" + err.Error())
        return
    }
    //初始化JWT
    //jwt.Init(global.JWTSetting.Issuer, global.JWTSetting.Expire, global.JWTSetting.Secret)
    //初始化上传文件服务
    upload.Init(
        &upload.ServerStruct{
            SavePath:  global.UploadSetting.UploadSavePath,
            ServerUrl: global.UploadSetting.UploadServerUrl,
        },
        &upload.ImageStruct{
            ImageAllowExits:  global.UploadSetting.UploadImageAllowExts,
            SaveImageMaxSize: global.UploadSetting.UploadImageMaxSize,
        }, 
    )
}
```
