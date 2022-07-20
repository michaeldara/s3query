/** @author Michael Dara **/

package main

import (
	"flag"
	//"fmt"
	"os"
	"net/url"
	"s3query/model"
)


func main() {

	var input string
	var withId string
    var fromTime string
    var toTime string
	var withWord string

	flag.StringVar(&input, "input", "", "s3 URI")
	flag.StringVar(&withId, "with-id", "", "json id")
    flag.StringVar(&fromTime, "from-time", "", "from time")
    flag.StringVar(&toTime, "to-time", "", "to time")
    flag.StringVar(&withWord, "with-word", "", "with word")

	flag.Parse()

    if len(input) == 0 {
		ExitErrorf("input flag cannot be blank")
	}

    //check if the s3Url is valid
	u, err := url.Parse(input)
	if err != nil {
		ExitErrorf("invalid input flag", input)
	}

    //Populate the InputRecord structure
    var InputRecord  *model.InputRecord
	InputRecord = new(model.InputRecord)

	InputRecord.Input = input
	InputRecord.WithId = withId
	InputRecord.FromTime = fromTime
	InputRecord.ToTime = toTime
	InputRecord.WithWord = withWord
	InputRecord.Bucket = u.Host;
	InputRecord.Key = u.Path;


   //fmt.Println("parameters received: %a" , InputRecord)

    //Process the query
    Process(InputRecord)

	os.Exit(1)

}