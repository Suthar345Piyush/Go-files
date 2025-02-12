package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("playing with url")
	myURL := "https://x.com/PiyushS35"
	fmt.Printf("the  type of url is %T\n", myURL)

	//we  can parse a  url

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("the  error is  in: ", err)
		return
	}

	fmt.Println("Our parsed url is: ", parsedURL)
	fmt.Printf("the type of url is %T\n", parsedURL)
	fmt.Println("the scheme of the link is: ", parsedURL.Scheme)
	fmt.Println("the Host of the link is: ", parsedURL.Host)
	fmt.Println("the path of the link is: ", parsedURL.Path)
	fmt.Println("the rawQuery of the link is: ", parsedURL.RawQuery)

	// we can change the things of an url ,
	parsedURL.Path = "/myname"
	parsedURL.Host = "piyushS.dex"
	// scheme or port
	parsedURL.Scheme = "http"
	parsedURL.RawQuery = "mera webpage"

	newUrl := parsedURL.String()
	fmt.Printf("The  type of  new url is %T\n", newUrl)

	fmt.Println("the  new url after modifying is: ", newUrl)

}
