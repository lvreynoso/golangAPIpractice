// server.go

package main

import (
"encoding/json"
"io/ioutil"
"log"
"net/http"
"github.com/labstack/echo"
)

type ChuckJoke struct {
    ID int `json:"id"`
    Joke string `json:"joke"`
    Categories []string `json:"categories"`
}

type Taco struct {
    Type string `json:"type"`
    Value ChuckJoke `json:"value"`
}

func main() {
    server := echo.New()

    server.GET("/", func(context echo.Context) error {
        response, err := http.Get("https://api.icndb.com/jokes/random")
        if err != nil {
            log.Fatalln(err)
            // panic("OMGWTFBBQ")
        }

        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalln(err)
            // panic("OMGWTFBBQ")
        }

        taco := Taco{}

        jsonErr := json.Unmarshal(body, &taco)
        if jsonErr != nil {
            log.Fatalln(jsonErr)
            // panic("OMGWTFBBQ")
        }

        return context.String(http.StatusOK, taco.Value.Joke)

        })

    server.Logger.Fatal(server.Start(":9001"))
}