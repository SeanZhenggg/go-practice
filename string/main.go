package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

//func main() {
//fmt.Println(strings.TrimLeft("aabbccdd", "abc"))
//fmt.Println(strings.TrimRight("aabbccdd/r/n", ""))
//}

func UTF8BytesToString() {
	raw := []byte(`{"HelloWorld": "\uB155, \uC138\uC0C1(\u4E16\u4E0A). \u263a"}`)
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return
	}

	fmt.Printf("string : %s\n", str)
}

func UTF8RuneToString() {
	rs := []rune{0xb155, 0x4e16, 0x754c, 0x6742, 0x8d27, 0x94fa}
	fmt.Printf("rs : %v, size of each rune: %v\n", rs, unsafe.Sizeof(rs[0]))
	fmt.Printf("string(rs) : %v\n", string(rs))
	utf8bs := make([]byte, 0)
	for _, r := range rs {
		bs := make([]byte, 4)
		w := utf8.EncodeRune(bs, r)
		fmt.Printf("bs: %v, w: %d, size of each byte: %v\n", bs, w, unsafe.Sizeof(bs[0]))
		utf8bs = append(utf8bs, bs[:w]...)
	}
	fmt.Printf("utf8bs : %v\n", utf8bs)
	fmt.Printf("utf8 string : %v\n", string(utf8bs))
}

func UTF8DecodeFromByte() {
	for _, v := range []rune("新世界杂货铺") {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
	bs := []byte("新世界杂货铺")
	for len(bs) > 0 {
		r, w := utf8.DecodeRune(bs)
		fmt.Printf("%x ", r)
		bs = bs[w:]
	}
	fmt.Println()
	//输出:
	//65b0 4e16 754c 6742 8d27 94fa
	//65b0 4e16 754c 6742 8d27 94fa
}

func StringRegexCompile() {
	re, err := regexp.Compile(`{"default": (.*?)}`)
	if err != nil {
		return
	}

	matched := re.FindString(string([]byte(`{"default": aehfewuifhiuehvwe2131rfehriurehbr}}}}}`)))
	p := re.FindAllIndex([]byte(`{"default": aehfewuifhiuehvwe2131rfehriurehbr}}}}}`), -1)
	log.Printf("matched: %v, pos: %d\n", matched, p)
}

func StringRegexMatch() {
	str := "http: //zq.titan007.com/Image/league_match/images/20210531102310.png"
	//str := "https://image.bricblogy.com/group1/M00/15/AF/CgURt2I-96aASbgbAAAFwXWNdUY740.png"
	//str := "https://static.fastbs55.com/data/addcba48fbb4d25452dcff0577c14631.jpg"
	match, _ := regexp.MatchString(`007\.com/`, str)
	if match {
		log.Printf("matched")
	} else {
		log.Printf("not matched")
	}
}

func BackQuoteInString() {
	str := `我是 \n , 他好像是 \t, 你是 \uB155, match : 8d5`
	reg := `[1-9][a-z].`

	fmt.Printf("str : %s\n", str)
	match, err := regexp.Match(reg, []byte(str))
	if err != nil {
		return
	}

	fmt.Printf("matched : %v\n", match)

	re, _ := regexp.Compile(reg)
	matched := re.FindString(str)
	p := re.FindAllIndex([]byte(str), -1)
	fmt.Printf("matched string: %v, pos: %d\n", matched, p)
}

func main() {
	//UTF8BytesToString()

	//UTF8RuneToString()

	//UTF8DecodeFromByte()

	//StringRegexCompile()

	//StringRegexMatch()

	BackQuoteInString()
}
