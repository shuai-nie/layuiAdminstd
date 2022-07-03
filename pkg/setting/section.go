package setting

import (
	"github.com/spf13/viper"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

type ServerSettings struct {
	RunMode       string        `json:"run_mode"`
	HttpPort      string        `json:"http_port"`
	ReadTimeout   time.Duration `json:"read_timeout"`
	WriterTimeout time.Duration `json:"writer_timeout"`
}

type AppSettings struct {
	UploadSavePath string
	UploadServerUrl string
	UploadImageMaxSize int
	UploadImageAllowExts []string
}

type DatabaseSettings struct {
	DBType string
	Username string
	Password string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host string
	Port int
	UserName string
	Password string
	IsSSL bool
	Form string
	To []string
}