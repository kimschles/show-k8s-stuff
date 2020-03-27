package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
)

func main() {
	fmt.Println("hiiii!")
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("In the else statement")
	}
}
