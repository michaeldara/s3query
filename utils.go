/** @author Michael Dara **/

package main

import (
	"fmt"
	"os"
	"time"
    "golang.org/x/exp/slices"
)


func utils() {

}

/**
  checks if the json string has a Time that is between from-time and to-time input flags
**/
func IsTimeBetweenFromAndToTime(timeFromJson string, fromTime string, toTime string) bool {


    var epochinMillisec int64
    var nowinMillisec int64

    timeFromJsonInMillisec := GetTimeInMillisec("jsonRecordTime", timeFromJson);

    if len(fromTime) == 0 {
        epochinMillisec = GetEpochInMillisec();
    } else {
        epochinMillisec = GetTimeInMillisec("fromTime", fromTime);
    }

    if len(toTime) == 0 {
        nowinMillisec = GetNowTimeInMillisec();
    } else {
        nowinMillisec =  GetTimeInMillisec("toTime", toTime);
    }

	//fmt.Println(fromTime, fromTime)
    return (timeFromJsonInMillisec >= epochinMillisec) && (timeFromJsonInMillisec <= nowinMillisec)
}


/**
  checks if the json string has a ID that is same as with-id flag
**/
func IsWithIdMatch(idFromJson string, withId string) bool{

    if len(withId) == 0 || idFromJson == withId {
        return true;
    } else {
        return false
    }
	return true
}


/**
  checks if the json string has a word that is same as with-word flag
**/
func IsWithWordMatch(wordListFromJson []string, withWord string) bool{

    if len(withWord) == 0 || slices.Contains(wordListFromJson, withWord) {
        return true;
    } else {
        return false
    }
	return true
}


/**
  returns time in milliseconds for a give Time string
**/
func GetTimeInMillisec(label string, jsonRecordTime string) int64 {
	t, err := time.Parse(time.RFC3339Nano, jsonRecordTime)
    if err != nil {
        ExitErrorf("Unable to parse time", label, jsonRecordTime, err)
    }
	recTimeinMillisec := t.UnixNano() / 1000000
    return recTimeinMillisec
}

/**
  returns current time in milliseconds
**/
func GetNowTimeInMillisec() int64 {
	now := time.Now()
	nowinMillisec := now.UnixNano() / 1000000
    return nowinMillisec
}

/**
  returns epoch time in milliseconds
**/
func GetEpochInMillisec() int64 {
	epochStartTime, err := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000000000-07:00")
	epochinMillisec := epochStartTime.UnixNano() / 1000000
    if err != nil {
        ExitErrorf("Unable to parse epoch to millisec", "1970-01-01T00:00:00.000000000-07:00", err)
    }
    return epochinMillisec
}


func ExitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+" ", args...)
    os.Exit(1)
}