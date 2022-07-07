package global

import (
	"layuiAdminstd/pkg/logger"
	"layuiAdminstd/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	EmailSetting *setting.EmailSettingS

	JWTSetting *setting.JWTSettingS

	Logger *logger.Logger


)
