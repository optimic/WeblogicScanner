package config

import (
	"github.com/EDDYCJY/fake-useragent"
)

func Fakeua() string {
	ua := browser.Random()
	//fmt.Println(ua)
	return ua
}
