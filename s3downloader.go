/** @author Michael Dara **/

package main

import (
	"fmt"
	"os"
    "s3query/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func DownLoadS3File(InputRecord *model.InputRecord) {


    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_REGION"))},
    )
  
    downloader := s3manager.NewDownloader(awsSession)
    
    file, err := os.Create("/tmp" + InputRecord.Key);
  
    if err != nil {
        ExitErrorf("Unable to fetch file %q, %v", InputRecord.Key, err)
    }
    
	defer file.Close()

    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String(InputRecord.Bucket),
            Key:    aws.String(InputRecord.Key),
        })
    if err != nil {
        fmt.Println("Unable to download file %q, %v, %s", InputRecord.Key, numBytes, err)
    }

   //fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

}