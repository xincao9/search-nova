package constant

const (
	LoggerDir                          = "logger.dir"
	LoggerLevel                        = "logger.level"
	ServerMode                         = "server.mode"
	ServerPort                         = "server.port"
	ServerCorsAccessControlAllowOrigin = "server.cors.access-control-allow-origin"
	ManagerServerPort                  = "manager.server.port"
	DataSource                         = "db.dataSource"
	PublicDir                          = "public.dir"
	ElasticsearchAddresses             = "elasticsearch.addresses"
	ElasticsearchUsername              = "elasticsearch.username"
	ElasticsearchPassword              = "elasticsearch.password"
	ElasticsearchIndex                 = "elasticsearch.index"
)

const (
	SessionUser       = "S_USER"
	SessionExpireHour = 24
	Success           = "success"
	SystemError       = "system error"
	Token             = "token"
	StatusNormal      = 0 // 默认状态
	StatusForbid      = 1 // 被封禁状态
	StatusRead        = 2 // 已读状态
)

const (
	NewStatus     = 0
	SuccessStatus = 1
	FailureStatus = 2
)
