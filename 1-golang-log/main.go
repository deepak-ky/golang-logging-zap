package main

import (
	"fmt"
	"log"
)

func main() {
	standardError := fmt.Errorf("connection timed out")

	/* 	log.Print("A", 1, 2, "B")
	   	log.Print("A", "B", 1, 2)
	   	log.Print(standardError)
	   	log.Print(standardError.Error())
	   	//log.Print("operation failed due to : " + standardError + " .This will be logged") // This does not work
	   	log.Print("operation failed due to : ", standardError, " .This will be logged")
	   	log.Print("operation failed due to : %s", standardError, " .This will be logged") // Incorrect output */

	/* 	log.Printf("A", 1, 2, "B") //Incorrect output
	   	log.Printf("A", "B", 1, 2) //Incorrect output
	   	log.Printf("%v %v %v %v", "A", 1, 2, "B")
	   	log.Printf("%d %d %d %d", "A", 1, 2, "B") //Incorrect ouput
	   	log.Printf("%s %s %s %s", "A", 1, 2, "B") //Incorrect ouput
	   	log.Printf("%s %d %d %s", "A", 1, 2, "B")
	   	//log.Printf(standardError) //This does not work
	   	log.Printf(standardError.Error())
	   	//log.Printf("operation failed due to : " + standardError + " .This will be logged") // This does not work
	   	log.Printf("operation failed due to : ", standardError, " .This will be logged")         // Incorrect output
	   	log.Printf("operation failed due to : ", standardError.Error(), " .This will be logged") // Incorrect output
	   	log.Printf("operation failed due to : %s .This will be logged", standardError) */

	log.Println("A", 1, 2, "B")
	log.Println("A", "B", 1, 2)
	log.Println("%v %v %v %v", "A", 1, 2, "B") //Incorrect output
	log.Println("%d %d %d %d", "A", 1, 2, "B") //Incorrect output
	log.Println("%s %s %s %s", "A", 1, 2, "B") //Incorrect output
	log.Println("%s %d %d %s", "A", 1, 2, "B") //Incorrect output
	log.Println(standardError)
	log.Println(standardError.Error())
	//log.Println("operation failed due to : " + standardError) // This does not work
	log.Println("operation failed due to :", standardError, ".This will be logged")
	log.Println("operation failed due to :", standardError.Error(), ".This will be logged")
	log.Println("operation failed due to : %s", standardError, ".This will be logged") //Incorrect ouput
}
