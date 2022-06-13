package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var test config.Config

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	listBuckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range (*listBuckets).Buckets {
		fmt.Printf("%s %s\n", *bucket.CreationDate, *bucket.Name)
	}

	// test := context.Background()
	// var id string = "/test-ogura-iac/rds01/wordpress"
	// var params_input ssm.GetParameterInput = ssm.GetParameterInput{
	// 	Name:           &id,
	// 	WithDecryption: true,
	// }

	//region := getRegion()
	// option := ssm.Options{
	// 	Credentials: aws.NewCredentialsCache(cfg.Credentials),
	// 	Region:      "ap-northeast-1",
	// }
	// client := ssm.New(option)
	// params_output, err := client.GetParameter(test, &params_input)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(*params_output.Parameter.Value)
}

func getRegion() string {
	url := "http://169.254.169.254/latest/meta-data//placement/availability-zone"
	r, _ := http.Get(url)
	defer r.Body.Close()
	byteArray, _ := ioutil.ReadAll(r.Body)
	str := string(byteArray)
	size := len(str)
	return str[0 : size-1]
}
