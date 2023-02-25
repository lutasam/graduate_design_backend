package common

import (
	"time"
)

const (
	ISSUER            = "LUTASAM"                          // jwt issuer
	PASSWORDSALT      = "astaxie12798akljzmknm.ahkjkljl;k" // use only for password encryption
	OTHERSECRETSALT   = "9871267812345mn812345xyz"         // user for other encryption
	EXPIRETIME        = 86400000                           // jwt expiration time. 1 day's second
	REDISEXPIRETIME   = 24 * 3600 * time.Second            // redis normal key expire time, refers to one day
	ACTIVECODEEXPTIME = 300 * time.Second                  // active code expiration time. 5 min
	DEFAULTAVATARURL  = "http://img.duoziwang.com/2018/20/08111043560274.jpg"
	MAXIMGSPACE       = 1024 * 1024 * 1 // img upload should be less than 1 MB
)

const (
	STATUSOKCODE    = 200
	CLIENTERRORCODE = 400
	SERVERERRORCODE = 500
)

const (
	STATUSOKMSG    = "OK"
	CLIENTERRORMSG = "CLIENT ERROR"
	SERVERERRORMSG = "SERVER ERROR"
)

const (
	TALKEDUSERLISTSUFFIX = "_talked_user_list"
	ACTIVECODESUFFIX     = "_active_code"
)

type CharacterType int

const (
	USER CharacterType = iota + 1
	DOCTOR
	SYSADMIN
)

func (c CharacterType) String() string {
	switch c {
	case USER:
		return "用户"
	case DOCTOR:
		return "医生"
	case SYSADMIN:
		return "系统管理员"
	default:
		return ""
	}
}

func ParseCharacterType(i int) CharacterType {
	switch i {
	case 1:
		return USER
	case 2:
		return DOCTOR
	case 3:
		return SYSADMIN
	default:
		return 0
	}
}

func (c CharacterType) Int() int {
	switch c {
	case USER:
		return 1
	case DOCTOR:
		return 2
	case SYSADMIN:
		return 3
	default:
		return int(c)
	}
}

type ProfessionalRank int

const (
	DIRECTOR_PHYSICIAN       ProfessionalRank = iota + 1 // 主任医师
	AD_PHYSICIAN                                         // 副主任医师
	DOCTOR_IN_CHARGE_OF_CASE                             // 主治医师
	PHYSICIAN                                            // 医师
)

func (p ProfessionalRank) String() string {
	switch p {
	case DIRECTOR_PHYSICIAN:
		return "主任医师"
	case AD_PHYSICIAN:
		return "副主任医师"
	case DOCTOR_IN_CHARGE_OF_CASE:
		return "主治医师"
	case PHYSICIAN:
		return "医师"
	default:
		return ""
	}
}

func ParseProfessionalRank(i int) ProfessionalRank {
	switch i {
	case 1:
		return DIRECTOR_PHYSICIAN
	case 2:
		return AD_PHYSICIAN
	case 3:
		return DOCTOR_IN_CHARGE_OF_CASE
	case 4:
		return PHYSICIAN
	default:
		return ProfessionalRank(i)
	}
}

type HospitalRank int

const (
	FIRST_A HospitalRank = iota + 1 // 一甲
	FIRST_B
	FIRST_C
	SECOND_A
	SECOND_B
	SECOND_C
	THIRD_SPECIAL
	THIRD_A
	THIRD_B
	THIRD_C
)

func ParseHospitalRank(i int) HospitalRank {
	switch i {
	case 1:
		return FIRST_A
	case 2:
		return FIRST_B
	case 3:
		return FIRST_C
	case 4:
		return SECOND_A
	case 5:
		return SECOND_B
	case 6:
		return SECOND_C
	case 7:
		return THIRD_SPECIAL
	case 8:
		return THIRD_A
	case 9:
		return THIRD_B
	case 10:
		return THIRD_C
	default:
		return HospitalRank(i)
	}
}

func (h HospitalRank) String() string {
	switch h {
	case FIRST_A:
		return "一等甲级"
	case FIRST_B:
		return "一等乙级"
	case FIRST_C:
		return "一等丙级"
	case SECOND_A:
		return "二等甲级"
	case SECOND_B:
		return "二等乙级"
	case SECOND_C:
		return "二等丙级"
	case THIRD_SPECIAL:
		return "三等特级"
	case THIRD_A:
		return "三等甲级"
	case THIRD_B:
		return "三等乙级"
	case THIRD_C:
		return "三等丙级"
	default:
		return ""
	}
}

type ReplyStatus int

const (
	REPLIED ReplyStatus = iota + 1
	NOT_REPLIED
	ALL_STATUS
)

func ParseReplyStatus(i int) ReplyStatus {
	switch i {
	case 1:
		return REPLIED
	case 2:
		return NOT_REPLIED
	case 3:
		return ALL_STATUS
	default:
		return ReplyStatus(i)
	}
}

func (r ReplyStatus) Int() int {
	switch r {
	case REPLIED:
		return 1
	case NOT_REPLIED:
		return 2
	case ALL_STATUS:
		return 3
	default:
		return 3
	}
}

type MessageType int

const (
	NORMAL  MessageType = iota + 1 // 正常对话
	HISTORY                        // 获取历史信息
)

func (m MessageType) Int() int {
	switch m {
	case NORMAL:
		return 1
	case HISTORY:
		return 2
	default:
		return int(m)
	}
}

type SexType int

const (
	MALE SexType = iota + 1
	FEMALE
	OTHER
)

func (s SexType) Int() int {
	switch s {
	case MALE:
		return 1
	case FEMALE:
		return 2
	case OTHER:
		return 3
	default:
		return int(s)
	}
}
