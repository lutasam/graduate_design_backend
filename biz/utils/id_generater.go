package utils

import "github.com/bwmarrin/snowflake"

var (
	userIDGenerator       *snowflake.Node
	hospitalIDGenerator   *snowflake.Node
	departmentIDGenerator *snowflake.Node
	inquiryIDGenerator    *snowflake.Node
	doctorIDGenerator     *snowflake.Node
)

func init() {
	var err error
	userIDGenerator, err = snowflake.NewNode(100)
	if err != nil {
		panic(err)
	}

	hospitalIDGenerator, err = snowflake.NewNode(200)
	if err != nil {
		panic(err)
	}

	departmentIDGenerator, err = snowflake.NewNode(300)
	if err != nil {
		panic(err)
	}

	inquiryIDGenerator, err = snowflake.NewNode(400)
	if err != nil {
		panic(err)
	}

	doctorIDGenerator, err = snowflake.NewNode(500)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() uint64 {
	return uint64(userIDGenerator.Generate().Int64())
}

func GenerateHospitalID() uint64 {
	return uint64(hospitalIDGenerator.Generate().Int64())
}

func GenerateDepartmentID() uint64 {
	return uint64(departmentIDGenerator.Generate().Int64())
}

func GenerateInquiryID() uint64 {
	return uint64(inquiryIDGenerator.Generate().Int64())
}

func GenerateDoctorID() uint64 {
	return uint64(inquiryIDGenerator.Generate().Int64())
}
