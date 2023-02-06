package xerr

var message map[uint32]string

func init() {
    message = make(map[uint32]string)
    message[OK] = "SUCCESS"
    message[ERROR_CODE_SERVER_COMMON] = "服务器开小差啦,稍后再来试一试"
    message[ERROR_CODE_REUQEST_PARAM] = "参数错误"
    message[ERROR_CODE_TOKEN_EXPIRE] = "token失效，请重新登陆"
    message[ERROR_CODE_TOKEN_GENERATE] = "生成token失败"
    message[ERROR_CODE_DB_ERROR] = "数据库繁忙,请稍后再试"
    message[ERROR_CODE_DB_UPDATE_AFFECTED_ZERO] = "更新数据影响行数为0"

    // 用户服务
    message[ERROR_CODE_USER_NOT_EXIST] = "用户不存在"
    message[ERROR_CODE_USER_PASSWORD_ERROR] = "账号或密码不正确"

}

func MapErrMsg(errcode uint32) string {
    if msg, ok := message[errcode]; ok {
        return msg
    } else {
        return "服务器开小差啦,稍后再来试一试"
    }
}

func IsCodeErr(errcode uint32) bool {
    if _, ok := message[errcode]; ok {
        return true
    } else {
        return false
    }
}
