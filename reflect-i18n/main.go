package main

import (
	"fmt"
	"reflect"
)

var (
	// i18n key map in database
	dbTable map[string]map[string]string = map[string]map[string]string{
		"zh": {
			"hello": "你好",
			"world": "世界",
		},
		"en": {
			"hello": "hello",
			"world": "world",
		},
		"jp": {
			"hello": "こんにちは",
			"world": "世界",
		},
	}
)

// sample bo data
type SampleData struct {
	HelloValue        string `i18n:"hello"`
	WorldValue        string `i18n:"world"`
	NoneSetValue      string
	WrongI18nKeyValue string `i18n:"world2"`
}

func main() {
	//sData := &SampleData{
	//	HelloValue:        "hello",
	//	WorldValue:        "world",
	//	NoneSetValue:      "noneset",
	//	WrongI18nKeyValue: "wrongvalue",
	//}
	//
	//headerLanguage := "zh"
	//fmt.Printf("=====[lang] %v=====\n", headerLanguage)
	//fmt.Printf("sData parsed before : %v\n", sData)
	//reflectForI18n(headerLanguage, sData)
	//fmt.Printf("sData parsed after : %v\n\n", sData)
	//
	//headerLanguage = "en"
	//fmt.Printf("=====[lang] %v=====\n", headerLanguage)
	//fmt.Printf("sData parsed before : %v\n", sData)
	//reflectForI18n(headerLanguage, sData)
	//fmt.Printf("sData parsed after : %v\n\n", sData)
	//
	//headerLanguage = "jp"
	//fmt.Printf("=====[lang] %v=====\n", headerLanguage)
	//fmt.Printf("sData parsed before : %v\n", sData)
	//reflectForI18n(headerLanguage, sData)
	//fmt.Printf("sData parsed after : %v\n\n", sData)

	s := struct{ A int }{0}
	field := reflect.ValueOf(s).Field(0)

	a := field.Interface()
	fmt.Println(reflect.TypeOf(a), a)
	b := reflect.Zero(field.Type())
	fmt.Println(reflect.TypeOf(b), b)

	fmt.Println(reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type())))
}

func getI18nMapByLang(lang string) map[string]string {
	m, ok := dbTable[lang]
	if !ok {
		return nil
	}

	return m
}

func reflectForI18n(lang string, data interface{}) interface{} {
	// get i18n map from header language
	getMap := getI18nMapByLang(lang)
	if getMap == nil {
		fmt.Println("no lang map found!!!")
		return data
	}

	// reflect for detecting data is struct/ptr or not
	reType := reflect.TypeOf(data)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		fmt.Println("not pointer nor struct !!!")
		return data
	}

	reValue := reflect.ValueOf(data).Elem()

	for i := 0; i < reValue.NumField(); i++ {
		structField := reValue.Type().Field(i)
		fieldTag := structField.Tag
		i18nKey := fieldTag.Get("i18n")

		// key not found
		if i18nKey == "" {
			fmt.Printf("struct field i18n key not found ... %v\n", structField.Name)
			continue
		}

		parsed, ok := getMap[i18nKey]
		if !ok {
			fmt.Printf("struct field i18n key no mapping i18n map ... %v\n", structField.Name)
			continue
		}
		reParsed := reflect.ValueOf(parsed)
		reValue.Field(i).Set(reParsed)
	}

	return data
}
