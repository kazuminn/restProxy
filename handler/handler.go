package handler

import (
    "net/http"
    "io/ioutil"
    "github.com/labstack/echo"
    "fmt"
    "os"
)

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
	webpagepath := c.Param("webpagepath")
	
	resp, err := http.Get("https://" + webpagepath)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	
	byteArray, _ := ioutil.ReadAll(resp.Body)
        return c.HTML(http.StatusOK, string(byteArray))
    }
}
