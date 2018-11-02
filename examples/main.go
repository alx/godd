package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CorentinB/godd"
)

var myHost = "http://localhost:8080"

func main() {
	// Get informations on the instance
	info, err := godd.GetInfo(myHost)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(info)

	// Create a service
	var service *godd.ServiceRequest
	service = godd.NewService(service)

	// Specify values for your service creation
	service.Name = "mask"
	service.Description = "And example mask detection service."
	service.Mllib = "caffe2"
	service.Parameters.Input.Connector = "image"
	service.Parameters.Input.Width = 1216
	service.Parameters.Input.Height = 800
	service.Parameters.Input.Mean = append(service.Parameters.Input.Mean, 102.9801)
	service.Parameters.Input.Mean = append(service.Parameters.Input.Mean, 115.9465)
	service.Parameters.Input.Mean = append(service.Parameters.Input.Mean, 122.7717)
	service.Parameters.Mllib.Nclasses = 81
	service.Parameters.Mllib.GPU = true
	service.Parameters.Mllib.GPUID = append(service.Parameters.Mllib.GPUID, 1)
	service.Model.Repository = "/home/corentin/test_mask/"
	service.Model.Extensions = append(service.Model.Extensions, "/home/corentin/test_mask/mask")

	// Send the service creation request
	creationResult, err := godd.CreateService(myHost, service)
	if err != nil {
		log.Fatal(err)
	}

	if creationResult.Status.Code == 200 {
		fmt.Println("Service creation: " + creationResult.Status.Msg)
	} else {
		fmt.Println("Service creation: " + creationResult.Status.Msg)
	}

	// Create predict and initialize it
	var predict *godd.PredictRequest
	predict = godd.NewPredict(predict)

	// Specify values for your prediction
	predict.Service = "mask"
	predict.Width = 1216
	predict.Height = 800
	predict.Mask = true
	predict.Data = append(predict.Data, "https://t2.ea.ltmcdn.com/fr/images/9/0/0/les_bienfaits_d_avoir_un_chien_1009_600.jpg")

	predictResult, err := godd.Predict(myHost, predict)
	if err != nil {
		log.Fatal(err)
	}

	if predictResult.Status.Code == 200 {
		fmt.Println("Prediction: " + predictResult.Status.Msg)
	} else {
		fmt.Println("Prediction: " + predictResult.Status.Msg)
	}

	// Get service informations
	serviceInfoResult, err := godd.GetServiceInfo(myHost, "mask")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service informations:")
	fmt.Println(serviceInfoResult)

	// Delete service
	serviceDeleteStatus, err := godd.DeleteService(myHost, "mask")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service deletion:")
	fmt.Println(serviceDeleteStatus)
}
