package logger

// Config 日志配置
type LogConfig struct {
	Devolpment    bool   `yaml:"devolpment" json:"devolpment"`       //运行模式
	Level         string `yaml:"level" json:"level"`                 //日志级别
	Encoding      string `yaml:"encoding" json:"encoding"`           //输出格式，console,json
	Directory     string `yaml:"directory" json:"directory"`         //日志目录
	FileName      string `yaml:"fileName" json:"fileName"`           //日志名
	MaxSize       uint   `yaml:"maxSize" json:"maxSize"`             //日志文件最大大小
	MaxBackups    uint   `yaml:"maxBackups" json:"maxBackups"`       //保留旧文件最大个数
	MaxAge        uint   `yaml:"maxAge" json:"maxAge"`               //保留旧文件最大天数
	Compress      bool   `yaml:"compress" json:"compress"`           //是否压缩旧文件
	EnableConsole bool   `yaml:"enableConsole" json:"enableConsole"` //是否启用控制台输出
	FileConsole   bool   `yaml:"fileConsole" json:"fileConsole"`     //是否启用文件输出
	EnableCaller  bool   `yaml:"enableCaller" json:"enableCaller"`   //是否添加调用者信息
}

func (config *LogConfig) SetDefault() {
	if config.Level == "" {
		config.Level = "info"
	}
	if config.Encoding == "" {
		config.Encoding = "json"
	}
	if config.FileName == "" {
		config.FileName = "app.log"
	}
	if config.MaxSize == 0 {
		config.MaxSize = 100 // 100MB
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 30 // 保留30个备份
	}
	if config.MaxAge == 0 {
		config.MaxAge = 90 // 保留90天
	}
}
