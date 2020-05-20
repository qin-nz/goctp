package comm

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetFieldValueIgnoreEmptyPassword(field interface{}) map[string]interface{} {
	fv := make(map[string]interface{})

	t := reflect.TypeOf(field)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if strings.HasPrefix(m.Name, "Get") {
			f := m.Name[3:]

			// 日志跳过密码字段
			if strings.Contains(strings.ToLower(f), "password") {
				continue
			}

			// 调用 GetXXX 获取 XXX 的值
			fieldV := reflect.ValueOf(field)
			rtn := m.Func.Call([]reflect.Value{fieldV})
			v := rtn[0]

			// 为零时跳过
			// TODO: v != reflect.Zero(v.Type()) 返回结果不符合
			// 只能检测下 string 和 int 类型的零值
			if v.Type() == reflect.TypeOf("") && v.String() == "" {
				continue
			}
			if v.Type() == reflect.TypeOf(0) && v.Int() == 0 {
				continue
			}

			fv[f] = v
		}
	}

	return fv
}

func LogReq(reqID int, message string, reqRtn int, reqField interface{}) {
	// 获取调用者信息
	fpcs := make([]uintptr, 1)
	runtime.Callers(2, fpcs)
	callerName := runtime.FuncForPC(fpcs[0]).Name()
	callerSimpleName := callerName[strings.LastIndex(callerName, ".")+1:]

	// 获取请求字段的值
	reqFieldMap := GetFieldValueIgnoreEmptyPassword(reqField)

	// 请求信息
	reqFieldMap["request_type"] = callerSimpleName
	reqFieldMap["request_id"] = reqID
	reqFieldMap["request_rtn"] = reqRtn
	reqFieldMap["request_rtnstr"] = RequestReturnValue(reqRtn).String()

	logrus.WithFields(logrus.Fields(reqFieldMap)).Println(message)
}
