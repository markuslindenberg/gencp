package md380

var (
	codeplugTemplate = Codeplug{
		BasicInformation: BasicInformation{
			FrequencyRange:     "400-480 MHz",
			LastProgrammedTime: "19-Nov-2017 15:52:26",
			Model:              "MD380",
		},
		GeneralSettings: GeneralSettings{
			BacklightTime:          "15",
			CallAlertToneDuration:  "5",
			ChFreeIndicationTone:   "Off",
			DisableAllLeds:         "Off",
			DisableAllTones:        "Off",
			GroupCallHangTime:      "7000",
			IntroScreen:            "Character String",
			IntroScreenLine1:       "",
			IntroScreenLine2:       "",
			LoneWorkerReminderTime: "10",
			LoneWorkerResponseTime: "1",
			Mode:                 "Channel",
			MonitorType:          "Open Squelch",
			PcProgPassword:       "",
			PowerOnPassword:      "00000000",
			PrivateCallHangTime:  "7000",
			PwAndLockEnable:      "Off",
			RadioID:              "",
			RadioName:            "",
			RadioProgPassword:    "00000000",
			RxLowBatteryInterval: "120",
			SaveModeReceive:      "On",
			SavePreamble:         "On",
			ScanAnalogHangTime:   "7000",
			ScanDigitalHangTime:  "7000",
			SetKeypadLockTime:    "Manual",
			TalkPermitTone:       "Digital and Analog",
			TxPreambleDuration:   "300",
			VoxSensitivity:       "2",
		},
		MenuItems: MenuItems{
			Answered:        "On",
			Backlight:       "On",
			CallAlert:       "On",
			DisplayMode:     "On",
			Edit:            "On",
			EditList:        "On",
			Gps:             "Off",
			HangTime:        "Hang",
			IntroScreen:     "On",
			KeyboardLock:    "On",
			LedIndicator:    "On",
			ManualDial:      "On",
			Missed:          "On",
			OutgoingRadio:   "On",
			PasswordAndLock: "Off",
			Power:           "On",
			ProgramKey:      "On",
			ProgramRadio:    "On",
			RadioCheck:      "On",
			RadioDisable:    "Off",
			RadioEnable:     "Off",
			RemoteMonitor:   "Off",
			Scan:            "On",
			Squelch:         "On",
			Talkaround:      "On",
			TextMessage:     "On",
			ToneOrAlert:     "On",
			Vox:             "On",
		},
		GPSSystems: []GPSSystem{
			GPSSystem{
				DestinationID:            "BM-GPS",
				GPSDefaultReportInterval: "300",
				GPSRevertChannel:         "Current Channel",
			},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
			GPSSystem{GPSDefaultReportInterval: "Off"},
		},
	}

	channelTemplate = Channel{
		AdmitCriteria:           "Color code",
		AllowTalkaround:         "On",
		Autoscan:                "Off",
		Bandwidth:               "12.5",
		ChannelMode:             "Digital",
		ColorCode:               "1",
		CompressedUDPDataHeader: "On",
		ContactName:             "TG9-Lok/Refl",
		CtcssDecode:             "None",
		CtcssEncode:             "None",
		DataCallConfirmed:       "Off",
		Decode1:                 "Off",
		Decode2:                 "Off",
		Decode3:                 "Off",
		Decode4:                 "Off",
		Decode5:                 "Off",
		Decode6:                 "Off",
		Decode7:                 "Off",
		Decode8:                 "Off",
		DisplayPTTID:            "Off",
		EmergencyAlarmAck:       "Off",
		GPSSystem:               "1",
		GroupList:               "-Refl.-",
		InCallCriteria:          "Follow Admit Criteria",
		LoneWorker:              "Off",
		Name:                    "TS2-4398500",
		Power:                   "High",
		Privacy:                 "None",
		PrivacyNumber:           "0",
		PrivateCallConfirmed:    "Off",
		QtReverse:               "180",
		ReceiveGPSInfo:          "Off",
		RepeaterSlot:            "2",
		ReverseBurst:            "On",
		RxFrequency:             "439.85000",
		RxOnly:                  "Off",
		RxRefFrequency:          "Low",
		RxSignallingSystem:      "Off",
		ScanList:                "4398500",
		SendGPSInfo:             "Off",
		Squelch:                 "Normal",
		Tot:                     "90",
		TotRekeyDelay:           "0",
		TxFrequency:             "430.45000",
		TxRefFrequency:          "Low",
		TxSignallingSystem:      "Off",
		Vox:                     "Off",
	}

	scanlistTemplate = ScanList{
		Name:                "",
		PriorityChannel1:    "Selected",
		PriorityChannel2:    "Selected",
		PrioritySampleTime:  "2000",
		SignallingHoldTime:  "500",
		TxDesignatedChannel: "Last Active Channel",
	}
)
