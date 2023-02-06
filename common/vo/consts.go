package vo

const (
    //正常
    ERROR_CODE_OK uint32 = 0
    //服务不可用
    ERROR_CODE_SVC_UNAVAILABLE uint32 = 100001
    //权限不足
    ERROR_CODE_PERMISSION_DENIED uint32 = 100002
    //缺少必选参数
    ERROR_CODE_REQUIRED_ARGS_MISSING uint32 = 100003
    //非法的参数
    ERROR_CODE_ARGS_ILLEGAL uint32 = 100004
    //业务处理失败
    ERROR_CODE_HANDLE_FAILED uint32 = 100005
    //token无效
    ERROR_CODE_TOKEN_INVALID uint32 = 100006
    //数据库访问错误
    ERROR_CODE_DAO_ERROR uint32 = 100007
    //用户名/密码错误
    ERROR_CODE_USER_PASSWORD_ERROR = 200001
    //用户过期
    ERROR_CODE_USER_EXPIRE = 200002
    //用户不存在
    ERROR_CODE_USER_NOT_EXIST = 200003
    //token过期
    ERROR_CODE_USER_TOKEN_EXPIRE = 200004
    //用户名已经存在
    ERROR_CODE_USER_ALREADY_EXIST = 200005
)
const (
    JWT_SECRET          = "BINGBARBYBINGGU"
    HEADER_NAME_TRACEID = "TraceId"
)
