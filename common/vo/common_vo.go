package vo

import (
    "context"
)

type CommonResult struct {
    TraceId string                 `json:"traceId" newtag:"traceId"`
    Code    uint32                 `json:"code" newtag:"code"`
    UserMsg string                 `json:"msg" newtag:"msg"`           //用户查看的信息，可读性更强
    LogKV   map[string]interface{} `json:"errorMsg" newtag:"errorMsg"` //打印日志的信息，携带错误详情，便于追查问题
    Data    interface{}            `json:"data" newtag:"-"`
}

func (c CommonResult) Success() bool {
    return c.Code == ERROR_CODE_OK
}

func ConstructOkResult(ctx context.Context) *CommonResult {
    result := &CommonResult{
        //TraceId: ctx.Value(HEADER_NAME_TRACEID).(string),
        Code:    ERROR_CODE_OK,
        LogKV:   make(map[string]interface{}),
    }
    return result
}

func ConstructOkResultEx(ctx context.Context, userMsg string, kvs ...KV) *CommonResult {
    result := &CommonResult{
        //TraceId: ctx.Value(HEADER_NAME_TRACEID).(string),
        Code:    ERROR_CODE_OK,
        UserMsg: userMsg,
        LogKV:   make(map[string]interface{}),
    }
    for _, kv := range kvs {
        result.LogKV[kv.K] = kv.V
    }
    return result
}

func ConstructErrorResult(ctx context.Context, code uint32, userMsg string, kvs ...KV) *CommonResult {
    result := &CommonResult{
        //TraceId: ctx.Value(HEADER_NAME_TRACEID).(string),
        Code:    code,
        UserMsg: userMsg,
        LogKV:   make(map[string]interface{}),
    }
    for _, kv := range kvs {
        result.LogKV[kv.K] = kv.V
    }
    return result
}

type CommonPageResult struct {
    PageNo   int `json:"pageNo"`
    PageSize int `json:"pageSize"`
    Total    int `json:"total"`
}

type CommonPageReq struct {
    PageNo   int
    PageSize int
}

type Count struct {
    Total int
}

type KV struct {
    K string
    V interface{}
}

func KVPair(k string, v interface{}) KV {
    return KV{
        K: k,
        V: v,
    }
}

type CommonDataResp struct {
    CommonResult `json:",inline"`
    Data         interface{} `json:"data"`
}

func ConstructCommonDataResp(ctx context.Context, code uint32, userMsg string, data interface{}) *CommonDataResp {
    return &CommonDataResp{
        CommonResult: *ConstructErrorResult(ctx, code, userMsg),
        Data:         data,
    }
}

func ConstructEmptyDataResp(ctx context.Context, code uint32, userMsg string) *CommonDataResp {
    return &CommonDataResp{
        CommonResult: *ConstructErrorResult(ctx, code, userMsg),
        Data:         "{}",
    }
}
