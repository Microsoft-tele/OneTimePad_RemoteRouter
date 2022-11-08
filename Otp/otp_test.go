package Otp

import "testing"

func TestOtp_AddOtp(t *testing.T) {
	otp := Otp{}
	otp.InitMysql()
	otp.AddOtp("www.baidu.com", "百度", "18697450302", "Lwj20020302", "1784929126@qq.com")
}
