package ir

const (
	VENDOR_ID  = 0x22ea
	PRODUCT_ID = 0x003a
)

const (
	BTO_EP_IN  = 0x84
	BTO_EP_OUT = 0x04
)

const (
	IR_FREQ_MIN                   = 25000 // 赤外線周波数設定最小値 25KHz
	IR_FREQ_MAX                   = 50000 // 赤外線周波数設定最大値 50KHz
	IR_SEND_DATA_USB_SEND_MAX_LEN = 14    // USB送信１回で送信する最大ビット数
	IR_SEND_DATA_MAX_LEN          = 300   // 赤外線送信データ設定最大長[byte]
	IR_FREQ_DEFAULT               = 38000 // 赤外線周波数デフォルト値
)

const (
	MAX_BYTE_ARRAY_SIZE = 9600
)

const (
	FORMAT_AEHA_READER_CODE  = 0x007B003D // 家電協 Reader Code ON:3.2ms/26us=123(0x7B), OFF:1.6ms/26us=61(0x3D)
	FORMAT_AEHA_BIT_OFF      = 0x000F000F // 家電協 BitOFF ON:0.4ms/26us=15(0x0F), OFF:0.4ms/26us=15(0x0F)
	FORMAT_AEHA_BIT_ON       = 0x000F002E // 家電協 BitON  ON:0.4ms/26us=15(0x0F), OFF:1.2ms/26us=46(0x2E)
	FORMAT_NEC_READER_CODE   = 0x015A00AD // NEC Reader Code ON:9.0ms/26us=346(0x15A), OFF:4.5ms/26us=173(0xAD)
	FORMAT_NEC_BIT_OFF       = 0x00160016 // NEC BitOFF ON:0.56ms/26us=22(0x16), OFF:0.56ms/26us=22(0x16)
	FORMAT_NEC_BIT_ON        = 0x00160041 // NEC BitON  ON:0.56ms/26us=22(0x16), OFF:1.68ms/26us=65(0x41)
	FORMAT_SONY_READER_CODE  = 0x005C0017 // SONY Reader Code ON:2.4ms/26us=92(0x5C), OFF:0.6ms/26us=23(0x17)
	FORMAT_SONY_BIT_OFF      = 0x00170017 // SONY BitOFF ON:0.6ms/26us=23(0x17), OFF:0.6ms/26us=23(0x17)
	FORMAT_SONY_BIT_ON       = 0x002E0017 // SONY BitON  ON:1.2ms/26us=46(0x2E), OFF:0.6ms/26us=23(0x17)
	FORMAT_MITSU_READER_CODE = 0x00000000 // MITSUBISHI Reader Code なし
	FORMAT_MITSU_BIT_OFF     = 0x000F001F // MITSUBISHI BitOFF ON:0.4ms/26us=15(0x0F), OFF:0.8ms/26us=31(0x1F)
	FORMAT_MITSU_BIT_ON      = 0x000F004D // MITSUBISHI BitON  ON:0.4ms/26us=15(0x0F), OFF:2.0ms/26us=77(0x4D)
	FORMAT_STOP_CODE         = 0x00170619 // STOP CODE
)

type IR_FORMAT int

const (
	IR_FORMAT_AEHA IR_FORMAT = 1 + iota
	IR_FORMAT_NEC
	IR_FORMAT_SONY
	IR_FORMAT_MITSUBISHI
)

type PLARAIL_BAND int

const (
	PLARAIL_BAND_BAND_A PLARAIL_BAND = 1 + iota
	PLARAIL_BAND_BAND_B
)

type PLARAIL_DIRECTION int

const (
	PLARAIL_DIRECTION_FORWARD PLARAIL_DIRECTION = 1 + iota
	PLARAIL_DIRECTION_BACKWARD
)

type PLARAIL int

const (
	Plarail_StopA PLARAIL = iota
	Plarail_StopB
	Plarail_Speed_UpAF
	Plarail_Speed_UpAB
	Plarail_Speed_UpBF
	Plarail_Speed_UpBB
	Plarail_Speed_DownA
	Plarail_Speed_DownB
)

var FORMATlist []string = []string{"AEHA", "NEC", "SONY", "MITSUBISHI"}

var PLAOPTIONlist []string = []string{"Plarail_StopA", "Plarail_StopB",
	"Plarail_Speed_UpAF", "Plarail_Speed_UpAB", "Plarail_Speed_UpBF",
	"Plarail_Speed_UpBB", "Plarail_Speed_DownA", "Plarail_Speed_DownB"}
