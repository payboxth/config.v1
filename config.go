package config

// Config -- config.v1 is a shared configuration for Paybox Vending Platform
// Use by other service and application on 
// - /payboxth/ota
// - /payboxth/vending
// - /payboxth/cloud
// - /payboxth/android
type Config struct {
	TimeServers       []string          `json:"time_servers,omitempty"`
	OTP               otp               `json:"otp,omitempty"`
	Host              host              `json:"host,omitempty"`
	Client            client            `json:"client,omitempty"`
	Hardware          hardware          `json:"hardware,omitempty"`
	Cash              cash              `json:"cash,omitempty"`
	CashAcceptedValue cashAcceptedValue `json:"cash_accepted_value,omitempty"`
	ChangeRefillValue changeRefillValue `json:"change_refill_value,omitempty"`
}

type otp struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
	Key4 string `json:"key4"`
}

type host struct {
	VendingID        int    `json:"vending_id,omitempty"`
	XAccessToken     string `json:"x_access_token,omitempty"`
	MockXAccessToken string `json:"mock_x_access_token,omitempty"`
}

type client struct {
	ShopTitle      string `json:"shop_title,omitempty"`
	ShopWebsite    string `json:"shop_website,omitempty"`
	LogoImageWidth int    `json:"logo_image_width,omitempty"`
	APIServerURL   string `json:"api_server_url,omitempty"`
	CloudServerURL string `json:"cloud_server_url,omitempty"`
	DataURL        string `json:"data_url,omitempty"`
	LogoPath       string `json:"logo_path,omitempty"`
	ImagePath      string `json:"image_path,omitempty"`
}

type hardware struct {
	StatusPollingInterval int    `json:"status_polling_interval,omitempty"`
	EventPollingInterval  int    `json:"event_polling_interval,omitempty"`
	CctalkBaudrate        int    `json:"cctalk_baudrate,omitempty"`
	CctalkPort            string `json:"cctalk_port,omitempty"`
	PrinterPort           string `json:"printer_port,omitempty"`
	GpioPort              string `json:"gpio_port,omitempty"`
	NoteDispenserPort     string `json:"note_dispenser_port,omitempty"`
	QrReaderPort          string `json:"qr_reader_port,omitempty"`
}

type cash struct {
	MaxChangeCash float64 `json:"max_change_cash,omitempty"`
	MaxChangeCoin float64 `json:"max_change_coin,omitempty"`
	MinChange     float64 `json:"min_change,omitempty"`
}

// CashAcceptedValue ยอดขายขั้นต่ำที่ยอมให้รับธนบัตรแต่ละประเภท
type cashAcceptedValue struct {
	B20   float64 `json:"b20,omitempty"`
	B50   float64 `json:"b50,omitempty"`
	B100  float64 `json:"b100,omitempty"`
	B500  float64 `json:"b500,omitempty"`
	B1000 float64 `json:"b1000,omitempty"`
}

// CashRefillValue มูลค่าเงินทอนแต่ละขนาดที่ต้องเตรียมใส่ถุงไว้เติมตอนเช้าก่อนเปิดร้าน
type changeRefillValue struct {
	V1    int `json:"v1,omitempty"`
	V2    int `json:"v2,omitempty"`
	V5    int `json:"v5,omitempty"`
	V10   int `json:"v10,omitempty"`
	V20   int `json:"v20,omitempty"`
	V50   int `json:"v50,omitempty"`
	V100  int `json:"v100,omitempty"`
	V500  int `json:"v500,omitempty"`
	V1000 int `json:"v1000,omitempty"`
}