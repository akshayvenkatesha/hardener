[swagger]: https://github.com/akshayvenkatesha/hardener/blob/master/swagger/swagger.yml

# Hardener

This Image is used for compliance scans on the target system.

## Usage
- [Create hardener container](#create-hardener-container)

- [Create target container(optional)](#create-target-container(optional))

## Create hardener container
Run below commands which will download and create a hardener container with the container name **hardener**.
```
docker run -it --privileged --name hardener akshayvenkatesha/hardener:version0 /main
```

Hardener container will expose **REST Endpoint** with [Swagger][swagger] specification for management operation.

## Get hardener container's ip
Open the terminal and run below command 
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hardener
```

### Sample Postman's Requests 
#### Create a Scan
```
POST  HTTP/1.1
Host: <hardener container's ip>:8080
Content-Type: application/json
cache-control: no-cache
Postman-Token: 58719135-f205-4cb7-93b2-bcda29cbad9c
{
    "hostname": "<Target host name or IP>",
    "id": "<any random id ex - 25>",
    "username":"<User name of the host>",
    "password":"<Password of the host>"
}------WebKitFormBoundary7MA4YWxkTrZu0gW--
```
Sample Request
```
POST  HTTP/1.1
Host: 172.17.0.3:8080
Content-Type: application/json
cache-control: no-cache
Postman-Token: 58719135-f205-4cb7-93b2-bcda29cbad9c
{
    "hostname": "172.17.0.2",
    "id": "25",
    "username":"akshay",
    "password":"akshay"
}------WebKitFormBoundary7MA4YWxkTrZu0gW--
```

#### Get Output of Scan

```
GET /<Id which will be used in above command> HTTP/1.1
Host: <hardener container's ip>:8080
Content-Type: application/json
cache-control: no-cache
Postman-Token: f96cbdd0-bfbb-4db0-8289-4c35b4f36b6f
```
Sample Request
```
GET /25 HTTP/1.1
Host: 172.17.0.3:8080
Content-Type: application/json
cache-control: no-cache
Postman-Token: f96cbdd0-bfbb-4db0-8289-4c35b4f36b6f
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