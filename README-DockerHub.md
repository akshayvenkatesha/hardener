[swagger]: https://github.com/akshayvenkatesha/hardener/blob/master/swagger/swagger.yml

# Hardener

This Image is used for compliance scans on the target system.

## Usage
- [Create hardener container](#create-hardener-container)

- [Create target container(optional)](#create-target-containeroptional)

## Create hardener container
Run below commands which will download and create a hardener container with the container name **hardener**.
```
docker run -it --privileged --name hardener akshayvenkatesha/hardener:0.2 /main
```

Hardener container will expose **REST Endpoint** with [Swagger][swagger] specification for management operation.

## Get hardener container's ip
Open the terminal and run below command 
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hardener
```

### Sample curl's Requests 
#### Create a Scan
```
curl -d '{"hostname": "<Target host name or IP>","id": "<any random id ex - 25>","username":"<Password of the host>","password":"akshay"}' -H "Content-Type: application/json" -X POST http://<hardener container's ip>:8080
```
Sample Request
```
curl -d '{"hostname": "172.17.0.2","id": "30","username":"akshay","password":"akshay"}' -H "Content-Type: application/json" -X POST http://172.17.0.2:8080
```

#### Get Output of Scan

```
curl -H "Content-Type: application/json" -X GET http://<hardener container's ip>:8080/<Id which will be used in above command>
```
Sample Request
```
curl -H "Content-Type: application/json" -X GET http://172.17.0.2:8080/30
```
Sample Response
```
{
    "hostname": "e9772dcf5057\n",
    "id": "25"
}
```

## Create target container(optional)
If it is hard to get the target machine to get in the network then you can use below commands to create an target system.

```
docker run -it --privileged --name target akshayvenkatesha/ubuntu-ssh:1.0 /bin/bash
service ssh restart
useradd --create-home --shell /bin/bash --groups sudo <User name to be added for this container>
passwd <User name to be added for this container>
```
### Above commands performs below steps
1. Creates a container.
2. Restarts **ssh service** in that container.
3. Adds user to the container.
4. Prompts for the password which will be used for that user.