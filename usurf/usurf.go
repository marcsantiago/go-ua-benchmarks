package usurf

import (
	"strings"

	"github.com/avct/uasurfer"
)

// UserAgents ...
const (
	THBadUserAgent   = "Bot/20181113.4785 CFNetwork/975.0.3 Darwin/18.2.0"
	IOSUserAgent     = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	AndroidUserAgent = "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36"
)

// IsValid ...
func IsValid(ua string) bool {
	uaO := uasurfer.Parse(ua)
	return !(uaO.Browser.Name == uasurfer.BrowserUnknown || strings.Contains("bot", string(uaO.Browser.Name)))
}
