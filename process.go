/** @author Michael Dara **/

package main

import (
	"fmt"
	"log"
    "os"
    "bufio"
	"compress/gzip"
    "encoding/json"
    "strconv"
    "s3query/model"
)


func Process(InputRecord *model.InputRecord) {

    //Download the GZIP file from S3
    DownLoadS3File(InputRecord)

    file, err := os.Open("/tmp" + InputRecord.Key);
     
    if err != nil {
		fmt.Println("Unable to open gzfile %q, %v", file.Name(), err)
	}

    //decompress the gz file
    jsonFile, err := gzip.NewReader(file)

    defer file.Close()
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("Unable to decompress file %q, %v", file.Name(), err)
	}


    //scan through the json file filter results based on the input flags.
	scanner := bufio.NewScanner(jsonFile)

    var count int
    count = 0
    var S3Record model.S3Record 

    for scanner.Scan() {
        
      line := scanner.Text()
      json.Unmarshal([]byte(line), &S3Record)

      isWithIdMatchTrue := IsWithIdMatch(strconv.FormatInt(S3Record.ID, 10), InputRecord.WithId)
      isTimeMatchTrue := IsTimeBetweenFromAndToTime(S3Record.Time, InputRecord.FromTime, InputRecord.ToTime)
      isWithWordMatchTrue := IsWithWordMatch(S3Record.Words, InputRecord.WithWord)

       if(isWithIdMatchTrue && isTimeMatchTrue && isWithWordMatchTrue) {
            fmt.Println(line)
            count++
       }
    }

    if count == 0 {
        fmt.Printf("No matching records found")
    }
 
    //remove the gz file from the /tmp folder
    e := os.Remove(file.Name())
    if e != nil {
        log.Fatal(e)
    }

}