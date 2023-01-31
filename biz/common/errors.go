package common

type Error struct {
	error
	ErrorString string
	ErrorCode   int
	ErrorType   int
}

func (e Error) Error() string {
	return e.ErrorString
}

func (e Error) Code() int {
	return e.ErrorCode
}

func (e Error) Type() int {
	return int(e.ErrorType)
}

var (
	UNKNOWNERROR = Error{
		ErrorCode:   -1,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "未知服务器错误，请联系管理员并等待",
	}
	USERINPUTERROR = Error{
		ErrorCode:   100001,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "输入错误，请检查您的输入",
	}
	HAVENOPERMISSION = Error{
		ErrorCode:   100002,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "无权限，请联系管理员获取",
	}
	DATABASEERROR = Error{
		ErrorCode:   100003,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "服务器数据库错误，请联系管理员并等待",
	}
	USERDOESNOTEXIST = Error{
		ErrorCode:   100004,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "用户不存在，请检查",
	}
	PASSWORDISERROR = Error{
		ErrorCode:   100005,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "密码错误，请再尝试一遍",
	}
	USERNOTLOGIN = Error{
		ErrorCode:   100006,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "未登录，请重新登录",
	}
	EXCEEDTIMELIMIT = Error{
		ErrorCode:   100007,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "Token过期，请重新登录",
	}
	USEREXISTED = Error{
		ErrorCode:   100008,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "用户邮箱已存在，请尝试登录或换一个邮箱",
	}
	NONEEDNOTICE = Error{
		ErrorCode:   100009,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "该用户已经打卡，无需提醒",
	}
	EMAILSYSTEMERROR = Error{
		ErrorCode:   100010,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "邮件服务器异常，请联系管理员",
	}
	DATANOTFOUND = Error{
		ErrorCode:   100011,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "数据未找到，请检查您的输入",
	}
	FILENOTEXIST = Error{
		ErrorCode:   100012,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "该文件不存在，请检查您的输入",
	}
	NOTINDEPARTMENT = Error{
		ErrorCode:   100013,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "该用户不属于该部门，请先将该用户分配到对应的部门",
	}
	REDISERROR = Error{
		ErrorCode:   100014,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "Redis服务出现故障，请通知管理员",
	}
	MONGOERROR = Error{
		ErrorCode:   100015,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "mongo数据库异常，请通知管理员",
	}
	ACTIVECODEERROR = Error{
		ErrorCode:   100016,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "验证码错误，请检查",
	}
	IMGFORMATERROR = Error{
		ErrorCode:   100017,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "我们只支持jpg/jpeg/png图片，请上传正确格式的图片",
	}
	IMGTOOLARGEERROR = Error{
		ErrorCode:   100018,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "图片大于指定大小，请尝试减小图片大小后上传",
	}
	FILEUPLOADERROR = Error{
		ErrorCode:   100019,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "文件上传失败，请上传正确格式的文件",
	}
	CANNOTDELETEADMIN = Error{
		ErrorCode:   100020,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "不能删除管理员，请先将管理员的权限让出并降级成普通用户",
	}
)
