package utils

import (
	"strings"
	"github.com/glancing/go-ys/tasks/yeezysupply/types"
)


func GetChromeUserAgentData(useragent string) types.UserAgentData {
	if (strings.Contains(useragent, "Chrome/98")) {
		if (strings.Contains(useragent, "Windows")) {
			return types.UserAgentData{
				ChUa: `" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"`,
				ChPlatform: "Windows",
				UserAgent: useragent,
			}
		} else {
			return types.UserAgentData{
				ChUa: `" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"`,
				ChPlatform: "macOS",
				UserAgent: useragent,
			}
		}
	}
	if (strings.Contains(useragent, "Chrome/99")) {
		if (strings.Contains(useragent, "Windows")) {
			return types.UserAgentData{
				ChUa: `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`,
				ChPlatform: "Windows",
				UserAgent: useragent,
			}
		} else {
			return types.UserAgentData{
				ChUa: `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`,
				ChPlatform: "macOS",
				UserAgent: useragent,
			}
		}
	}
	return types.UserAgentData{
		UserAgent: useragent,
	}
}