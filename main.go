package main

import (
	"fmt"
	"wxlogin/modules"
	"wxlogin/utils/conf"
)

func main() {
	appConfig := conf.Config.Section("58pic")
	appPage, err := modules.NewAppPage(appConfig.Key("appID").Value(), appConfig.Key("redirectURL").Value())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = appPage.VisitAppPage()
	if err != nil {
		fmt.Println(err)
		return
	}

	cookies, err := appPage.Callback()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}
