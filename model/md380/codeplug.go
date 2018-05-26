package md380

type Channel struct {
	AdmitCriteria           string `json:"AdmitCriteria"`
	AllowTalkaround         string `json:"AllowTalkaround"`
	Autoscan                string `json:"Autoscan"`
	Bandwidth               string `json:"Bandwidth"`
	ChannelMode             string `json:"ChannelMode"`
	ColorCode               string `json:"ColorCode"`
	CompressedUDPDataHeader string `json:"CompressedUdpDataHeader"`
	ContactName             string `json:"ContactName"`
	CtcssDecode             string `json:"CtcssDecode"`
	CtcssEncode             string `json:"CtcssEncode"`
	DataCallConfirmed       string `json:"DataCallConfirmed"`
	Decode1                 string `json:"Decode1"`
	Decode2                 string `json:"Decode2"`
	Decode3                 string `json:"Decode3"`
	Decode4                 string `json:"Decode4"`
	Decode5                 string `json:"Decode5"`
	Decode6                 string `json:"Decode6"`
	Decode7                 string `json:"Decode7"`
	Decode8                 string `json:"Decode8"`
	DisplayPTTID            string `json:"DisplayPTTID"`
	EmergencyAlarmAck       string `json:"EmergencyAlarmAck"`
	GPSSystem               string `json:"GPSSystem"`
	GroupList               string `json:"GroupList"`
	InCallCriteria          string `json:"InCallCriteria"`
	LoneWorker              string `json:"LoneWorker"`
	Name                    string `json:"Name"`
	Power                   string `json:"Power"`
	Privacy                 string `json:"Privacy"`
	PrivacyNumber           string `json:"PrivacyNumber"`
	PrivateCallConfirmed    string `json:"PrivateCallConfirmed"`
	QtReverse               string `json:"QtReverse"`
	ReceiveGPSInfo          string `json:"ReceiveGPSInfo"`
	RepeaterSlot            string `json:"RepeaterSlot"`
	ReverseBurst            string `json:"ReverseBurst"`
	RxFrequency             string `json:"RxFrequency"`
	RxOnly                  string `json:"RxOnly"`
	RxRefFrequency          string `json:"RxRefFrequency"`
	RxSignallingSystem      string `json:"RxSignallingSystem"`
	ScanList                string `json:"ScanList"`
	SendGPSInfo             string `json:"SendGPSInfo"`
	Squelch                 string `json:"Squelch"`
	Tot                     string `json:"Tot"`
	TotRekeyDelay           string `json:"TotRekeyDelay"`
	TxFrequency             string `json:"TxFrequency"`
	TxRefFrequency          string `json:"TxRefFrequency"`
	TxSignallingSystem      string `json:"TxSignallingSystem"`
	Vox                     string `json:"Vox"`
}

type Contact struct {
	CallID          string `json:"CallID"`
	CallReceiveTone string `json:"CallReceiveTone"`
	CallType        string `json:"CallType"`
	Name            string `json:"Name"`
}

type Zone struct {
	Channels []string `json:"Channel"`
	Name     string   `json:"Name"`
}

type GPSSystem struct {
	DestinationID            string `json:"DestinationID"`
	GPSDefaultReportInterval string `json:"GPSDefaultReportInterval"`
	GPSRevertChannel         string `json:"GPSRevertChannel"`
}

type GroupList struct {
	Contacts []string `json:"Contact"`
	Name     string   `json:"Name"`
}

type ScanList struct {
	Channels            []string `json:"Channel"`
	Name                string   `json:"Name"`
	PriorityChannel1    string   `json:"PriorityChannel1"`
	PriorityChannel2    string   `json:"PriorityChannel2"`
	PrioritySampleTime  string   `json:"PrioritySampleTime"`
	SignallingHoldTime  string   `json:"SignallingHoldTime"`
	TxDesignatedChannel string   `json:"TxDesignatedChannel"`
}

type BasicInformation struct {
	FrequencyRange     string `json:"FrequencyRange"`
	LastProgrammedTime string `json:"LastProgrammedTime"`
	Model              string `json:"Model"`
}

type GeneralSettings struct {
	BacklightTime          string `json:"BacklightTime"`
	CallAlertToneDuration  string `json:"CallAlertToneDuration"`
	ChFreeIndicationTone   string `json:"ChFreeIndicationTone"`
	DisableAllLeds         string `json:"DisableAllLeds"`
	DisableAllTones        string `json:"DisableAllTones"`
	GroupCallHangTime      string `json:"GroupCallHangTime"`
	IntroScreen            string `json:"IntroScreen"`
	IntroScreenLine1       string `json:"IntroScreenLine1"`
	IntroScreenLine2       string `json:"IntroScreenLine2"`
	LoneWorkerReminderTime string `json:"LoneWorkerReminderTime"`
	LoneWorkerResponseTime string `json:"LoneWorkerResponseTime"`
	Mode                   string `json:"Mode"`
	MonitorType            string `json:"MonitorType"`
	PcProgPassword         string `json:"PcProgPassword"`
	PowerOnPassword        string `json:"PowerOnPassword"`
	PrivateCallHangTime    string `json:"PrivateCallHangTime"`
	PwAndLockEnable        string `json:"PwAndLockEnable"`
	RadioID                string `json:"RadioID"`
	RadioName              string `json:"RadioName"`
	RadioProgPassword      string `json:"RadioProgPassword"`
	RxLowBatteryInterval   string `json:"RxLowBatteryInterval"`
	SaveModeReceive        string `json:"SaveModeReceive"`
	SavePreamble           string `json:"SavePreamble"`
	ScanAnalogHangTime     string `json:"ScanAnalogHangTime"`
	ScanDigitalHangTime    string `json:"ScanDigitalHangTime"`
	SetKeypadLockTime      string `json:"SetKeypadLockTime"`
	TalkPermitTone         string `json:"TalkPermitTone"`
	TxPreambleDuration     string `json:"TxPreambleDuration"`
	VoxSensitivity         string `json:"VoxSensitivity"`
}

type MenuItems struct {
	Answered        string `json:"Answered"`
	Backlight       string `json:"Backlight"`
	CallAlert       string `json:"CallAlert"`
	DisplayMode     string `json:"DisplayMode"`
	Edit            string `json:"Edit"`
	EditList        string `json:"EditList"`
	Gps             string `json:"Gps"`
	HangTime        string `json:"HangTime"`
	IntroScreen     string `json:"IntroScreen"`
	KeyboardLock    string `json:"KeyboardLock"`
	LedIndicator    string `json:"LedIndicator"`
	ManualDial      string `json:"ManualDial"`
	Missed          string `json:"Missed"`
	OutgoingRadio   string `json:"OutgoingRadio"`
	PasswordAndLock string `json:"PasswordAndLock"`
	Power           string `json:"Power"`
	ProgramKey      string `json:"ProgramKey"`
	ProgramRadio    string `json:"ProgramRadio"`
	RadioCheck      string `json:"RadioCheck"`
	RadioDisable    string `json:"RadioDisable"`
	RadioEnable     string `json:"RadioEnable"`
	RemoteMonitor   string `json:"RemoteMonitor"`
	Scan            string `json:"Scan"`
	Squelch         string `json:"Squelch"`
	Talkaround      string `json:"Talkaround"`
	TextMessage     string `json:"TextMessage"`
	ToneOrAlert     string `json:"ToneOrAlert"`
	Vox             string `json:"Vox"`
}

type Codeplug struct {
	BasicInformation BasicInformation `json:"BasicInformation"`
	GeneralSettings  GeneralSettings  `json:"GeneralSettings"`
	MenuItems        MenuItems        `json:"MenuItems"`
	Channels         []Channel        `json:"Channels"`
	Contacts         []Contact        `json:"Contacts"`
	GPSSystems       []GPSSystem      `json:"GPSSystems"`
	GroupLists       []GroupList      `json:"GroupLists"`
	ScanLists        []ScanList       `json:"ScanLists"`
	Zones            []Zone           `json:"Zones"`
}
