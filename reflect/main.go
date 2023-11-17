package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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

func (s SampleData) Error() string {
	return ""
}

type StudentReq struct {
	StudentName string `kQuery:"studentName"`
	ParentPhone string `kQuery:"parentPhone"`
	Limit       int    `kQuery:"limit"`
	Offset      int    `kQuery:"offset"`
}

const (
	OPERATORS = "=,!=,>,<,>=,in,not int,like,not like,or,and"
)

func ParseReqStructToMap(input interface{}) (map[string]interface{}, error) {
	reType := reflect.TypeOf(input)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("ParseReqToQuery error : neither pointer nor struct")
	}

	reValue := reflect.ValueOf(input).Elem()

	result := make(map[string]interface{})
	for i := 0; i < reValue.NumField(); i++ {
		//structField := reValue.Type().Field(i)
		//fieldTag := structField.Tag
		//kQueryKey := fieldTag.Get("kQuery")
		fmt.Printf("reValue.Field(i).Type(): %+v\n", reValue.Field(i).Type())
		fmt.Printf("reValue.Type().Field(i): %+v\n", reValue.Type().Field(i))
		fmt.Printf("reValue.Field(i): %+v\n", reValue.Field(i))
		fmt.Printf("reValue.Field(i).Kind(): %+v\n", reValue.Field(i).Kind())

		//if kQueryKey == "" {
		//	log.Printf("ParseReqToQuery cannot get input struct field tag by reflect... %v", kQueryKey)
		//	continue
		//}

		//switch reValue.Field(i).Kind() {
		//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		//	result[kQueryKey] = reValue.Field(i).Int()
		//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		//	result[kQueryKey] = reValue.Field(i).Uint()
		//case reflect.String:
		//	result[kQueryKey] = "\"" + reValue.Field(i).String() + "\""
		//}
	}
	return result, nil

}

func ConvMapToQueryStringByAnd(m map[string]interface{}) string {
	queries := make([]string, 0, len(m))
	for k, v := range m {
		switch v.(type) {
		case int64, uint64:
			val := v.(int64)
			if val != 0 {
				queries = append(queries, fmt.Sprintf("%s %s", k, v.(int64)))
			}
		case string:
			if v == "" {
				continue
			}
			found := false
			operators := strings.Split(OPERATORS, ",")
			for _, op := range operators {
				if strings.Contains(v.(string), op) {
					queries = append(queries, fmt.Sprintf("%s %s", k, v.(string)))
					found = true
					break
				}
			}
			if !found {
				queries = append(queries, fmt.Sprintf("%s = %s", k, v.(string)))
			}
		}
	}

	return strings.Join(queries, " and ")
}

func main() {
	//m, err := ParseReqStructToMap(&StudentReq{StudentName: "李新硯", ParentPhone: "0905966970"})
	//if err == nil {
	//	fmt.Printf("%v", m)
	//fmt.Printf(ConvMapToQueryStringByAnd(m))
	//}
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
	//
	//s := struct{ A int }{0}
	//field := reflect.ValueOf(s).Field(0)
	//
	//a := field.Interface()
	//fmt.Println(reflect.TypeOf(a), a)
	//b := reflect.Zero(field.Type())
	//fmt.Println(reflect.TypeOf(b), b)
	//
	//fmt.Println(reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type())))
	var a error = &SampleData{HelloValue: "123"}
	fmt.Printf("reflect.ValueOf(a) : %#v\n", reflect.ValueOf(a))
	fmt.Printf("reflect.ValueOf(a).Kind() : %+v\n", reflect.ValueOf(a).Kind())
	fmt.Printf("reflect.ValueOf(a).Type() : %+v\n", reflect.ValueOf(a).Type())
	fmt.Printf("reflect.ValueOf(a).Type().Kind() : %+v\n", reflect.ValueOf(a).Type().Elem())
	fmt.Printf("reflect.ValueOf(a).Elem() : %#v\n", reflect.ValueOf(a).Elem())
	fmt.Printf("reflect.ValueOf(a).Type().Elem() : %+v\n", reflect.ValueOf(a).Type().Elem())
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
