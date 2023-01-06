package main

import (
	cregex "github.com/mingrammer/commonregex"
	"github.com/rs/zerolog/log"
)

// 内置很多常用的正则表达式，开箱即用 https://github.com/mingrammer/commonregex
func main() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`

	dateList := cregex.Date(text)
	log.Print(dateList)
	// ['Jan 9th 2012']
	timeList := cregex.Time(text)
	log.Print(timeList)
	// ['5:00PM', '4:00']
	linkList := cregex.Links(text)
	log.Print(linkList)
	// ['www.linkedin.com', 'harold.smith@gmail.com']
	phoneList := cregex.PhonesWithExts(text)
	log.Print(phoneList)
	// ['(519)-236-2723x341']
	emailList := cregex.Emails(text)
	log.Print(emailList)
	// ['harold.smith@gmail.com']
}
