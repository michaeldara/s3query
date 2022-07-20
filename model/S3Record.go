/** @author Michael Dara **/

package model

type S3Record struct {
	ID    int64    `json:"id"`
	Time  string   `json:"time"`
	Words []string `json:"words"`
}