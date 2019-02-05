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
// [DONE] return an array of words in a ChuckJoke
// [WIP] call another API with similar content to original ChuckJoke

// Takes in Chuck Norris API; used in Taco struct
type ChuckJoke struct {
    ID int `json:"id"`
    Joke string `json:"joke"`
    Categories []string `json:"categories"`
}

// Takes in ChuckJoke struct; used in texasRanger
type Taco struct {
    Type string `json:"type"`
    Value ChuckJoke `json:"value"`
}

// Takes in tronalddump API; used in TrumpDump struct
type Quotes struct {
    Value string `json:"value"`
}

// Takes in Quotes struct; used in TrumpQuotes
type TrumpDump struct {
    Quotes []Quotes `json:"quotes"`
}

// Takes in TrumpDump struct; used in func newYorkBarFly
type TrumpQuotes struct {
    Embedded TrumpDump `json:"_embedded"`
}

func texasRanger() string {
    // takes in Taco struct and returns a Chuck Norris jok as a string
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
