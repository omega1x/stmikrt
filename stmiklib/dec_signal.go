package stmiklib

var (
	SIGNAL_NAME = [102]string{
		// float64 -> int32
		"titreading00heatt2", // Temperature of return heating water, [°C]",
		"titreading01netwt2", // Temperature of return network water, [°C]",
		"titreading02hotwt1", // Temperature of hot water, [°C]",
		"titreading03heatt1", // Temperature of supply heating water, [°C]",
		"titreading04netwt1", // Temperature of supply network water, [°C]",
		"titreading05utilp1", // Pressure of water utility, [kg×f/cm²]",
		"titreading06heatp2", // Pressure of return heating water, [kg×f/cm²]",
		"titreading07hotwp1", // Pressure of hot water, [kg×f/cm²]",
		"titreading08heatp1", // Pressure of supply heating water, [kg×f/cm²]",
		"titreading09netwp2", // Pressure of return network water, [kg×f/cm²]",
		"titreading10netwp1", // Pressure of supply network water, [kg×f/cm²]",

		"titlimitup00heatt2", // Upper limit setting of temperature of return heating water, [°C]",
		"titlimitup01netwt2", // Upper limit setting of temperature of return network water, [°C]",
		"titlimitup02hotwt1", // Upper limit setting of temperature of hot water, [°C]",
		"titlimitup03heatt1", // Upper limit setting of temperature of supply heating water, [°C]",
		"titlimitup04netwt1", // Upper limit setting of temperature of supply network water, [°C]",
		"titlimitup05utilp1", // Upper limit setting of pressure of water utility, [kg×f/cm²]",
		"titlimitup06heatp2", // Upper limit setting of pressure of return heating water, [kg×f/cm²]",
		"titlimitup07hotwp1", // Upper limit setting of pressure of hot water, [kg×f/cm²]",
		"titlimitup08heatp1", // Upper limit setting of pressure of supply heating water, [kg×f/cm²]",
		"titlimitup09netwp2", // Upper limit setting of pressure of return network water, [kg×f/cm²]",
		"titlimitup10netwp1", // Upper limit setting of pressure of supply network water, [kg×f/cm²]",

		"titlimitlw00heatt2", // Lower limit setting of temperature of return heating water, [°C]",
		"titlimitlw01netwt2", // Lower limit setting of temperature of return network water, [°C]",
		"titlimitlw02hotwt1", // Lower limit setting of temperature of hot water, [°C]",
		"titlimitlw03heatt1", // Lower limit setting of temperature of supply heating water, [°C]",
		"titlimitlw04netwt1", // Lower limit setting of temperature of supply network water, [°C]",
		"titlimitlw05utilp1", // Lower limit setting of pressure of water utility, [kg×f/cm²]",
		"titlimitlw06heatp2", // Lower limit setting of pressure of return heating water, [kg×f/cm²]",
		"titlimitlw07hotwp1", // Lower limit setting of pressure of hot water, [kg×f/cm²]",
		"titlimitlw08heatp1", // Lower limit setting of pressure of supply heating water, [kg×f/cm²]",
		"titlimitlw09netwp2", // Lower limit setting of pressure of return network water, [kg×f/cm²]",
		"titlimitlw10netwp1", // Lower limit setting of pressure of supply network water, [kg×f/cm²]",

		"tsregister00ts0015", // Zeroth register of TS array",
		"tsregister01ts1631", // First register of TS array",
		"tsregister02ts3247", // Second register of TS array",
		"tsregister03ts4863", // Third register of TS array",

		// bool -> uint8
		"titenabled00heatt2", // Enable of temperature of return heating water",
		"titenabled01netwt2", // Enable of temperature of return network water",
		"titenabled02hotwt1", // Enable of temperature of hot water",
		"titenabled03heatt1", // Enable of temperature of supply heating water",
		"titenabled04netwt1", // Enable of temperature of supply network water",
		"titenabled05utilp1", // Enable of pressure of water utility",
		"titenabled06heatp2", // Enable of pressure of return heating water",
		"titenabled07hotwp1", // Enable of pressure of hot water",
		"titenabled08heatp1", // Enable of pressure of supply heating water",
		"titenabled09netwp2", // Enable of pressure of return network water",
		"titenabled10netwp1", // Enable of pressure of supply network water",

		"titalarmon00heatt2", // Alarm switch of temperature of return heating water",
		"titalarmon01netwt2", // Alarm switch of temperature of return network water",
		"titalarmon02hotwt1", // Alarm switch of temperature of hot water",
		"titalarmon03heatt1", // Alarm switch of temperature of supply heating water",
		"titalarmon04netwt1", // Alarm switch of temperature of supply network water",
		"titalarmon05utilp1", // Alarm switch of pressure of water utility",
		"titalarmon06heatp2", // Alarm switch of pressure of return heating water",
		"titalarmon07hotwp1", // Alarm switch of pressure of hot water",
		"titalarmon08heatp1", // Alarm switch of pressure of supply heating water",
		"titalarmon09netwp2", // Alarm switch of pressure of return network water",
		"titalarmon10netwp1", // Alarm switch of pressure of supply network water",

		// uint8 -> uint8
		"tsbit00swchmnrelay", // Switch of relay of the drainage pit",

		"tsbit01failpmhotw1", // Failure of hot water pump-1",
		"tsbit02failpmhotw2", // Failure of hot water pump-2",
		"tsbit03failpmhotw3", // Failure of hot water pump-3",
		"tsbit04failpmheat1", // Failure of heating pump-1",
		"tsbit05failpmheat2", // Failure of heating pump-2",
		"tsbit06failpmheat3", // Failure of heating pump-3",
		"tsbit07failpmrech1", // Failure of recharge pump-1",
		"tsbit08failpmrech2", // Failure of recharge pump-2",
		"tsbit09failpmcirc1", // Failure of circulation pump-1",
		"tsbit10failpmcirc2", // Failure of circulation pump-2",

		"tsbit11swchpmdrain", // Switch of drainage pump",
		"tsbit12swchpmhotw1", // Switch of hot water pump-1",
		"tsbit13swchpmhotw2", // Switch of hot water pump-2",
		"tsbit14swchpmhotw2", // Switch of hot water pump-3",
		"tsbit15swchpmheat1", // Switch of heating pump-1",
		"tsbit16swchpmheat2", // Switch of heating pump-2",
		"tsbit17swchpmheat3", // Switch of heating pump-3",
		"tsbit18swchpmrech1", // Switch of recharge pump-1",
		"tsbit19swchpmrech2", // Switch of recharge pump-2",
		"tsbit20swchpmcirc1", // Switch of circulation pump-1",
		"tsbit21swchpmcirc2", // Switch of circulation pump-2",
		"tsbit22swchalarmon", // Switch of security alarm system",
		"tsbit23swchrchotwt", // Switch of remote control of hot water",
		"tsbit24swchrcheatq", // Switch of remote control of heating",
		"tsbit25swchrcrechg", // Switch of remote control of recharge",
		"tsbit26swchrccircl", // Switch of remote control of circulation",
		"tsbit27swchauhotwt", // Switch of hot water automation",
		"tsbit28swchauheatq", // Switch of heating automation",
		"tsbit29swchaurechg", // Switch of recharge automation",
		"tsbit30swchaucircl", // Switch of circulation automation",
		"tsbit31swchmnpower", // Switch of power supply monitoring",

		"tsbit32failrdtit00", // Fail reading of TIT01 sensor",
		"tsbit33failrdtit01", // Fail reading of TIT02 sensor",
		"tsbit34failrdtit02", // Fail reading of TIT03 sensor",
		"tsbit35failrdtit03", // Fail reading of TIT04 sensor",
		"tsbit36failrdtit04", // Fail reading of TIT05 sensor",
		"tsbit37failrdtit05", // Fail reading of TIT06 sensor",
		"tsbit38failrdtit06", // Fail reading of TIT07 sensor",
		"tsbit39failrdtit07", // Fail reading of TIT08 sensor",
		"tsbit40failrdtit08", // Fail reading of TIT09 sensor",
		"tsbit41failrdtit09", // Fail reading of TIT10 sensor",
		"tsbit42failrdtit10", // Fail reading of TIT11 sensor",
	}

	SIGNAL_DESCRIPTION = [102]string{
		"Temperature of return heating water", // titreading00heatt2
		"Temperature of return network water", // titreading01netwt2
		"Temperature of hot water",            // titreading02hotwt1
		"Temperature of supply heating water", // titreading03heatt1
		"Temperature of supply network water", // titreading04netwt1
		"Pressure of water utility",           // titreading05utilp1
		"Pressure of return heating water",    // titreading06heatp2
		"Pressure of hot water",               // titreading07hotwp1
		"Pressure of supply heating water",    // titreading08heatp1
		"Pressure of return network water",    // titreading09netwp2
		"Pressure of supply network water",    // titreading10netwp1

		"Enable of temperature of return heating water", // titenabled00heatt2
		"Enable of temperature of return network water", // titenabled01netwt2
		"Enable of temperature of hot water",            // titenabled02hotwt1
		"Enable of temperature of supply heating water", // titenabled03heatt1
		"Enable of temperature of supply network water", // titenabled04netwt1
		"Enable of pressure of water utility",           // titenabled05utilp1
		"Enable of pressure of return heating water",    // titenabled06heatp2
		"Enable of pressure of hot water",               // titenabled07hotwp1
		"Enable of pressure of supply heating water",    // titenabled08heatp1
		"Enable of pressure of return network water",    // titenabled09netwp2
		"Enable of pressure of supply network water",    // titenabled10netwp1

		"Alarm switch of temperature of return heating water", // titalarmon00heatt2
		"Alarm switch of temperature of return network water", // titalarmon01netwt2
		"Alarm switch of temperature of hot water",            // titalarmon02hotwt1
		"Alarm switch of temperature of supply heating water", // titalarmon03heatt1
		"Alarm switch of temperature of supply network water", // titalarmon04netwt1
		"Alarm switch of pressure of water utility",           // titalarmon05utilp1
		"Alarm switch of pressure of return heating water",    // titalarmon06heatp2
		"Alarm switch of pressure of hot water",               // titalarmon07hotwp1
		"Alarm switch of pressure of supply heating water",    // titalarmon08heatp1
		"Alarm switch of pressure of return network water",    // titalarmon09netwp2
		"Alarm switch of pressure of supply network water",    // titalarmon10netwp1

		"Upper limit setting of temperature of return heating water", // titlimitup00heatt2
		"Upper limit setting of temperature of return network water", // titlimitup01netwt2
		"Upper limit setting of temperature of hot water",            // titlimitup02hotwt1
		"Upper limit setting of temperature of supply heating water", // titlimitup03heatt1
		"Upper limit setting of temperature of supply network water", // titlimitup04netwt1
		"Upper limit setting of pressure of water utility",           // titlimitup05utilp1
		"Upper limit setting of pressure of return heating water",    // titlimitup06heatp2
		"Upper limit setting of pressure of hot water",               // titlimitup07hotwp1
		"Upper limit setting of pressure of supply heating water",    // titlimitup08heatp1
		"Upper limit setting of pressure of return network water",    // titlimitup09netwp2
		"Upper limit setting of pressure of supply network water",    // titlimitup10netwp1

		"Lower limit setting of temperature of return heating water", // titlimitlw00heatt2
		"Lower limit setting of temperature of return network water", // titlimitlw01netwt2
		"Lower limit setting of temperature of hot water",            // titlimitlw02hotwt1
		"Lower limit setting of temperature of supply heating water", // titlimitlw03heatt1
		"Lower limit setting of temperature of supply network water", // titlimitlw04netwt1
		"Lower limit setting of pressure of water utility",           // titlimitlw05utilp1
		"Lower limit setting of pressure of return heating water",    // titlimitlw06heatp2
		"Lower limit setting of pressure of hot water",               // titlimitlw07hotwp1
		"Lower limit setting of pressure of supply heating water",    // titlimitlw08heatp1
		"Lower limit setting of pressure of return network water",    // titlimitlw09netwp2
		"Lower limit setting of pressure of supply network water",    // titlimitlw10netwp1

		"Zeroth register of TS array", // tsregister00ts0015
		"First register of TS array",  // tsregister01ts1631
		"Second register of TS array", // tsregister02ts3247
		"Third register of TS array",  // tsregister03ts4863

		"Switch of relay of the drainage pit", // tsbit00swchmnrelay

		"Failure of hot water pump-1",   // tsbit01failpmhotw1
		"Failure of hot water pump-2",   // tsbit02failpmhotw2
		"Failure of hot water pump-3",   // tsbit03failpmhotw3
		"Failure of heating pump-1",     // tsbit04failpmheat1
		"Failure of heating pump-2",     // tsbit05failpmheat2
		"Failure of heating pump-3",     // tsbit06failpmheat3
		"Failure of recharge pump-1",    // tsbit07failpmrech1
		"Failure of recharge pump-2",    // tsbit08failpmrech2
		"Failure of circulation pump-1", // tsbit09failpmcirc1
		"Failure of circulation pump-2", // tsbit10failpmcirc2

		"Switch of drainage pump",                 // tsbit11swchpmdrain
		"Switch of hot water pump-1",              // tsbit12swchpmhotw1
		"Switch of hot water pump-2",              // tsbit13swchpmhotw2
		"Switch of hot water pump-3",              // tsbit14swchpmhotw2
		"Switch of heating pump-1",                // tsbit15swchpmheat1
		"Switch of heating pump-2",                // tsbit16swchpmheat2
		"Switch of heating pump-3",                // tsbit17swchpmheat3
		"Switch of recharge pump-1",               // tsbit18swchpmrech1
		"Switch of recharge pump-2",               // tsbit19swchpmrech2
		"Switch of circulation pump-1",            // tsbit20swchpmcirc1
		"Switch of circulation pump-2",            // tsbit21swchpmcirc2
		"Switch of security alarm system",         // tsbit22swchalarmon
		"Switch of remote control of hot water",   // tsbit23swchrchotwt
		"Switch of remote control of heating",     // tsbit24swchrcheatq
		"Switch of remote control of recharge",    // tsbit25swchrcrechg
		"Switch of remote control of circulation", // tsbit26swchrccircl
		"Switch of hot water automation",          // tsbit27swchauhotwt
		"Switch of heating automation",            // tsbit28swchauheatq
		"Switch of recharge automation",           // tsbit29swchaurechg
		"Switch of circulation automation",        // tsbit30swchaucircl
		"Switch of power supply monitoring",       // tsbit31swchmnpower

		"Fail reading of TIT01 sensor", // tsbit32failrdtit00
		"Fail reading of TIT02 sensor", // tsbit33failrdtit01
		"Fail reading of TIT03 sensor", // tsbit34failrdtit02
		"Fail reading of TIT04 sensor", // tsbit35failrdtit03
		"Fail reading of TIT05 sensor", // tsbit36failrdtit04
		"Fail reading of TIT06 sensor", // tsbit37failrdtit05
		"Fail reading of TIT07 sensor", // tsbit38failrdtit06
		"Fail reading of TIT08 sensor", // tsbit39failrdtit07
		"Fail reading of TIT09 sensor", // tsbit40failrdtit08
		"Fail reading of TIT10 sensor", // tsbit41failrdtit09
		"Fail reading of TIT11 sensor", // tsbit42failrdtit10
	}
)
