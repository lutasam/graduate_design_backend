package common

type Error struct {
	error
	ErrorString string
	ErrorCode   int
}

func (e Error) Error() string {
	return e.ErrorString
}

func (e Error) Code() int {
	return e.ErrorCode
}

var (
	UNKNOWNERROR = Error{
		ErrorCode:   -1,
		ErrorString: "unknown error. maybe server is error. please wait for sometime",
	}
	USERINPUTERROR = Error{
		ErrorCode:   100001,
		ErrorString: "please check your input, there is something wrong",
	}
	HAVENOPERMISSION = Error{
		ErrorCode:   100002,
		ErrorString: "you have no access to this operation",
	}
	DATABASEERROR = Error{
		ErrorCode:   100003,
		ErrorString: "server's database has some error, please try again later",
	}
	USERDOESNOTEXIST = Error{
		ErrorCode:   100004,
		ErrorString: "user does not exist. please check",
	}
	PASSWORDISERROR = Error{
		ErrorCode:   100005,
		ErrorString: "password is incorrect. please try again",
	}
	USERNOTLOGIN = Error{
		ErrorCode:   100006,
		ErrorString: "you do not login. please login",
	}
	EXCEEDTIMELIMIT = Error{
		ErrorCode:   100007,
		ErrorString: "your token has no time. please login again",
	}
	USEREXISTED = Error{
		ErrorCode:   100008,
		ErrorString: "account has existed. please rename your account",
	}
	GROUPNOTEXIST = Error{
		ErrorCode:   100009,
		ErrorString: "this group dose not exist. maybe your input is error",
	}
	GROUPEXIST = Error{
		ErrorCode:   100010,
		ErrorString: "this group has already been created. please use another name",
	}
	DATANOTFOUND = Error{
		ErrorCode:   100011,
		ErrorString: "data is not in database. please check your input",
	}
	HAVEBEENFRIEND = Error{
		ErrorCode:   100012,
		ErrorString: "you both are already friends, no need to be friend again",
	}
	NOTFRIEND = Error{
		ErrorCode:   100013,
		ErrorString: "you both are not friend, no need to delete friend",
	}
	REDISERROR = Error{
		ErrorCode:   100014,
		ErrorString: "server redis db error, please try later",
	}
)
