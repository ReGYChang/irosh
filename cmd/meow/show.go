package meow

import (
	"errors"
	"fmt"
)

func Show()error {
	if meow == nil{
		return errors.New("\nERROR: meow process have not started yet.\nPlease start the process before checking progress.\n")
	}else if meow.mCurrent != 2{
		fmt.Print(`

dBPYdaIpwl2Y0uT3tp7K6+yc7f/PLamjF2gctZ6YQIjXizMfCZzWV0MjvM/1iPrYnlz2mAwRtECEO+ykwf4h7mI3HvEeauv+PAf3rIlCSA/KujCHOfgKucf0lqLmh0EdECvFfr9bCOZjB+U4FWZYdOIv5IEYRteNKNph2y3LMgBIdrSvMUL9alaKvZ2pO8xhCIhAuyCvEGzheoIUWCkIxkC0n6jltSFSBZC9lSSX8xxQok668qcOgafXXst8hRlZqqgo3XTDYpaR4ey4iYH9nLAbeEUq3XkK+D7E9j3snNP8eBWpmr7iYVMGbLIzbhPpfVUKPwpAY687bBwySE/YVjL+G7H/7s2g07MR73Lvf0xt/ktOAZOXo3ssB1TvIstVwG49CpVK+TzEUGsR6/UkqCMv0SwK3z0xzzrGb2n+6zJ47xNYfCLD2r96AvKrZHDEg7JlgUNnVIf1IgiSyFQEfNGLc6aeAupaWKnGDyey5StHHI63ExIjzs7KQqUxtd6Wz0OMwCLRHc4znRZK70i+pg==

`)
	}else{
		fmt.Print(`
To Garra:

	喵喵喵喵喵喵喵~

	貓貓很抱歉
	在一囉生日要結束前兩個小時才在寫給一囉的信
	貓貓很多話不喜歡說在嘴上
	希望一囉都能感覺到貓貓心理想對一囉說的話
	今年幫一囉寫了一隻 shell
	它叫 irosh
	希望一囉喜歡
	也希望每年一囉生日的時候都能牽著一囉喵喵

	最後
	祝一囉生日快樂
	喵

Regy
	
`)
	}
	return nil
}