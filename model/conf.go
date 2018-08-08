package model

type Setting struct {
	SettingItem  string
	SettingValue string
	Config       *DBField
}

type DBField struct {
	FaultTask    string
	GlobalEmails string
	GlobalPhones string
	LogDir       string
	SmsPwd       string
	SmsServer    string
	SmsUser      string
	SmtpPwd      string
	SmtpServer   string
	SmtpUser     string
}
