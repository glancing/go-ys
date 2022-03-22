package types

type ApiYsConfig struct {
	CseKey                     string   `json:"cseKey"`
	PreSplashCookie            string   `json:"preSplashCookie"`
	HmacCookieName             string   `json:"hmacCookieName"`
	HmacPassedSplashCookieName string   `json:"hmacPassedSplashCookieName"`
	V3Sitekey                  string   `json:"v3Sitekey"`
	V3Action                   string   `json:"v3Action"`
	SplashDoubleTap            bool     `json:"splashDoubleTap"`
	SplashPath                 string   `json:"splashPath"`
	OverrideMetashared         bool     `json:"overrideMetashared"`
	BackupSmsEnabled           bool     `json:"backupSmsEnabled"`
	UserAgentArray             []string `json:"userAgentArray"`
	SaleStarted                bool     `json:"saleStarted"`
}

type PixelConfig struct {
	PixelFound bool
	Baza string
	PixelGetPath string
	PixelPostPath string
}

