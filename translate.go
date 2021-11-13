package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

    origin_lang := "pt-br"; // origin language
    dest_lang := "en"; // destination language
    term := "laranja"; // term to translate    

	url := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&sl="+origin_lang+"&tl="+dest_lang+"&q=" + url.QueryEscape(term)
   
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))

	defer res.Body.Close()

}