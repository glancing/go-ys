package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/glancing/go-ys/tasks/yeezysupply/types"
)

func ParsePixel(body string) types.PixelConfig {
	r, err := regexp.Compile(`bazadebezolkohpepadr="(.+)defer`)
	if err != nil {
		fmt.Println(err)
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	matches := r.FindAllString(body, -1)
	if len(matches) < 1 {
		fmt.Println("error")
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	r2, _ := regexp.Compile(`bazadebezolkohpepadr="(\d+)"`)
	bazaMatches := r2.FindAllString(matches[0], -1)
	if len(bazaMatches) < 1 {
		fmt.Println("error")
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	baza := strings.Split(bazaMatches[0], "=")[1]
	//pixel get path
	pixelGetPathSplit := strings.Split(matches[0], `src="`)
	if len(strings.Split(matches[0], `src="`)[1]) < 2 {
		fmt.Println("error")
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	pixelGetPathSplit2 := strings.Split(pixelGetPathSplit[1], `" `)
	if len(pixelGetPathSplit2) < 1 {
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	//pixel post path
	r3, _ := regexp.Compile(`<noscript[^>]*>(.*?)<\/noscript>`)
	noScript := r3.FindAllString(body, -1)
	if len(noScript) < 1 {
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	noScriptSplit := strings.Split(noScript[0], `src="`)
	if len(noScriptSplit) < 2 {
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	pixelPostSplit := strings.Split(noScriptSplit[1], "?")
	if len (pixelPostSplit) < 1 {
		return types.PixelConfig{
			PixelFound: false,
		}
	}
	return types.PixelConfig{
		PixelFound: true,
		Baza: baza,
		PixelGetPath: pixelGetPathSplit2[0],
		PixelPostPath: pixelPostSplit[0],
	}
}