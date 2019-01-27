package main

import (
	"flag"
	"fmt"
	"hardener/gen/models"
	"hardener/gen/restapi"
	"hardener/gen/restapi/operations"
	"hardener/gen/restapi/operations/hardener"
	"log"
	"os/exec"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")
var idMap = make(map[string]string)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHardenerAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	// need to define api's in this place
	api.HardenerGetIDHandler = hardener.GetIDHandlerFunc(
		func(params hardener.GetIDParams) middleware.Responder {
			id := swag.StringValue(&params.ID)
			if id == "" {
				id = "0"
			}
			var hostname string
			value, od := idMap[id]
			if !od {
				hostname = "error - id is wrong"
			} else {
				hostname = value
			}
			job := &models.Getjob{ID: id, Hostname: hostname, Progress: 0}
			return hardener.NewGetIDOK().WithPayload(job)
		})

	api.HardenerCreateHandler = hardener.CreateHandlerFunc(
		func(params hardener.CreateParams) middleware.Responder {
			id := swag.StringValue(&params.Body.ID)
			if id == "" {
				id = "0"
			}
			hostName := swag.StringValue(params.Body.Hostname)
			if hostName == "" {
				hostName = "localhost"
			}
			userName := swag.StringValue(params.Body.Username)
			if userName == "" {
				userName = "root"
			}
			password := swag.StringValue(params.Body.Password)
			if password == "" {
				password = "password"
			}
			job := &models.Createjob{ID: id, Hostname: &hostName, Username: &userName, Password: &password}

			// cmd1 := exec.Command(fmt.Sprintf("ssh akshay@{s} \"bash -s\" < ubuntu16Harden.sh", hostName))
			cmd := fmt.Sprintf("sshpass -p %s ssh -o StrictHostKeyChecking=no %s@%s \"bash -s\" < /home/akshay/ubuntu16mini.sh", password, userName, hostName)
			fmt.Printf(cmd)
			fmt.Println()
			cmd1 := exec.Command("sh", "-c", cmd)
			output, err1 := cmd1.CombinedOutput()
			if err1 != nil {
				fmt.Println(err1)
				idMap[id] = err1.Error()
			} else {
				outputStr := string(output)
				fmt.Println(outputStr)
				idMap[id] = outputStr
			}
			return hardener.NewCreateCreated().WithPayload(job)
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
