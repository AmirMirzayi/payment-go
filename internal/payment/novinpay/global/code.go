package global

var (
	Code_name = map[uint8]string{
		1: "erSucceed",
		2: "erAAS_InvalidUseridOrPass",
		3: "erAAS_InvalidSourceIp",
		4: "erAAS_InvalidData",
	}

	Code_value = map[string]uint8{
		"erSucceed":                 1,
		"erAAS_InvalidUseridOrPass": 2,
		"erAAS_InvalidSourceIp":     3,
		"erAAS_InvalidData":         4,
	}

	Code_enum = map[string]Code{
		"erSucceed":                 Success,
		"erAAS_InvalidUseridOrPass": InvalidUserOrPass,
		"erAAS_InvalidSourceIp":     InvalidSourceIp,
		"erAAS_InvalidData":         InvalidData,
	}

	ResponseMsg = map[Code]string{
		Success:           "erSucceed",
		InvalidUserOrPass: "erAAS_InvalidUseridOrPass",
		InvalidSourceIp:   "erAAS_InvalidSourceIp",
		InvalidData:       "erAAS_InvalidData",
	}

	ResponseText = map[Code]string{
		Success:           "سرویس با موفقیت اجرا شد.",
		InvalidUserOrPass: "کد کاربر یا رمزعبور صحیح نمی‌باشد.",
		InvalidSourceIp:   "آی‌پی مجار نمی‌باشد.",
		InvalidData:       "داده‌ها معتبر نمی‌باشند.",
	}
)
