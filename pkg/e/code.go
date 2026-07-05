package e

const (
	EXECUTE_SUCCESS = 200 // 操作成功

	// auth
	AUTH_REGISTRY_SUCCESS             = 30000 // 用户注册成功
	AUTH_REGISTRY_ERROR_USER_EXIST    = 30001 // 用户已经存在
	AUTH_REGISTRY_ERROR_MODEL_EXECUTE = 30002 // 用户注册失败
	AUTH_REGISTRY_ERROR_GENERATE      = 30003 // 用户密码加密失败

	AUTH_LOGIN_SUCCESS                   = 30010 // 用户登录成功
	AUTH_LOGIN_ERROR_USER_NOT_EXIST      = 30011 // 用户不存在
	AUTH_LOGIN_ERROR__MODEL_EXECUTE      = 30012 // 用户登录失败
	AUTH_LOGIN_ERROR_USER_PASSWORD_ERROR = 30013 // 用户密码错误
	AUTH_LOGIN_ERROR_GENERATE            = 30014 // 用户token生成失败

	//short url
	SHORT_URL_SUCCESS              = 40000 // 生成短链接成功
	SHORT_URL_ERROR_CREATE_EXECUTE = 40001 // 创建短链接记录失败

	INVALID_PARAMS = 400 // 参数错误
	INVALID_TOKEN  = 401 // 身份认证失败
	UNKOWN_ERROR   = 600 // 未知错误
)

var codeToMsg = map[int]string{
	EXECUTE_SUCCESS: "操作成功",
	INVALID_PARAMS:  "参数不合法",
	INVALID_TOKEN:   "身份认证失败",

	// auth
	AUTH_REGISTRY_SUCCESS:             "用户注册成功",
	AUTH_REGISTRY_ERROR_USER_EXIST:    "该用户名已经存在",
	AUTH_REGISTRY_ERROR_MODEL_EXECUTE: "用户注册失败",
	AUTH_REGISTRY_ERROR_GENERATE:      "用户密码加密失败",

	AUTH_LOGIN_ERROR_USER_NOT_EXIST:      "用户不存在",
	AUTH_LOGIN_ERROR__MODEL_EXECUTE:      "用户登录失败",
	AUTH_LOGIN_SUCCESS:                   "用户登录成功",
	AUTH_LOGIN_ERROR_USER_PASSWORD_ERROR: "用户密码错误",
	AUTH_LOGIN_ERROR_GENERATE:            "用户token生成失败",

	//short url
	SHORT_URL_SUCCESS:              "生成短链接成功",
	SHORT_URL_ERROR_CREATE_EXECUTE: "创建短链接记录失败",
}

func GetResponseMsg(code int) string {
	msg, ok := codeToMsg[code]
	if ok {
		return msg
	}
	return "该code值未找到对应错误文案"
}
