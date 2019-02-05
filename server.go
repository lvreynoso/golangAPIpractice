// server.go

package main

import (
"encoding/json"
"io/ioutil"
"log"
"net/http"
"github.com/labstack/echo"
"strings"
)

// STRETCH CHALLENGE OPTIONS:
// return an array of words in a ChuckJoke
// call another API with similar content to original ChuckJoke

type ChuckJoke struct {
    ID int `json:"id"`
    Joke string `json:"joke"`
    Categories []string `json:"categories"`
}

type Taco struct {
    Type string `json:"type"`
    Value ChuckJoke `json:"value"`
}

type Quotes struct {
    Value string `json:"value"`
}

type TrumpDump struct {
    Quotes []Quotes `json:"quotes"`
}

type TrumpQuotes struct {
    Embedded TrumpDump `json:"_embedded"`
}

func texasRanger() string {
    response, err := http.Get("https://api.icndb.com/jokes/random")
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }

    taco := Taco{}

    jsonErr := json.Unmarshal(body, &taco)
    if jsonErr != nil {
        log.Fatalln(jsonErr)
    }

    return taco.Value.Joke
}

func newYorkBarFly(vomit string) string {
    response, err := http.Get("https://api.tronalddump.io/search/quote?query=" + vomit)
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }

    chalupa := TrumpQuotes{}

    jsonErr := json.Unmarshal(body, &chalupa)
    if jsonErr != nil {
        log.Fatalln(jsonErr)
    }
    // returning an array of quotes
    // need to work on return a random quote from the array
    return chalupa.Embedded.Quotes
}

func main() {
    server := echo.New()

    server.GET("/", func(context echo.Context) error {
        tacoFilling := texasRanger()

        return context.HTML(http.StatusOK, "<em>" + tacoFilling + "</em>" + `<br><em>- Faith Chikwekwe</em>`)

        })

    server.GET("/tokenize", func(context echo.Context) error {
        tacoFilling := texasRanger()
        groundBeef := strings.Split(tacoFilling, " ")
        meatMap := make(map[string]int)
        for _, beef := range groundBeef {
            if _, ok := meatMap[beef]; ok {
                meatMap[beef] += 1
            } else {
                meatMap[beef] = 1
            }
        }
        return context.JSON(http.StatusOK, meatMap)
    })

    server.Logger.Fatal(server.Start(":9001"))
}
