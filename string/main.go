package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//func main() {
//fmt.Println(strings.TrimLeft("aabbccdd", "abc"))
//fmt.Println(strings.TrimRight("aabbccdd/r/n", ""))
//}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	fmt.Printf("strconv.Quote(string(raw)): %v\n", strconv.Quote(string(raw)))
	fmt.Printf("strings.Replace(strconv.Quote(string(raw)), `\\\\u`, `\\u`, -1): %v\n", strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	str, err := strconv.Unquote(strconv.Quote(string(raw)))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func main() {
	//s := []byte(`{"HelloWorld": "\uB155, \uC138\uC0C1(\u4E16\u4E0A). \u263a"}`)
	//v, _ := UnescapeUnicode(s)
	//fmt.Printf("%s\n", v)

	//for _, v := range []rune("新世界杂货铺") {
	//	fmt.Printf("%x ", v)
	//}
	//fmt.Println()
	//bs := []byte("新世界杂货铺")
	//for len(bs) > 0 {
	//	r, w := utf8.DecodeRune(bs)
	//	fmt.Printf("%x ", r)
	//	bs = bs[w:]
	//}
	//fmt.Println()
	// 输出:
	// 65b0 4e16 754c 6742 8d27 94fa
	// 65b0 4e16 754c 6742 8d27 94fa

	//rs := []rune{0x65b0, 0x4e16, 0x754c, 0x6742, 0x8d27, 0x94fa}
	//fmt.Println(string(rs))
	//utf8bs := make([]byte, 0)
	//for _, r := range rs {
	//	bs := make([]byte, 4)
	//	w := utf8.EncodeRune(bs, r)
	//	utf8bs = append(utf8bs, bs[:w]...)
	//}
	//fmt.Println(string(utf8bs))

	//re, err := regexp.Compile(`{"default": (.*?)}`)
	//if err != nil {
	//	return
	//}
	//
	//matched := re.FindString(string([]byte(`{"default": aehfewuifhiuehvwe2131rfehriurehbr}}}}}`)))
	//p := re.FindAllIndex([]byte(`{"default": aehfewuifhiuehvwe2131rfehriurehbr}}}}}`), -1)
	//log.Printf("matched: %v, pos: %d\n", matched, p)

	//str := "http: //zq.titan007.com/Image/league_match/images/20210531102310.png"
	//str := "https://image.bricblogy.com/group1/M00/15/AF/CgURt2I-96aASbgbAAAFwXWNdUY740.png"
	str := "https://static.fastbs55.com/data/addcba48fbb4d25452dcff0577c14631.jpg"
	match, _ := regexp.MatchString(`007\.com/`, str)
	if match {
		log.Printf("matched")
	} else {
		log.Printf("not matched")
	}
}
