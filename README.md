## Event Management Service For 10 MS

> Project run instructions 

Install all the dependencies using below command 
```shell
go get
```
Then build project
```shell
make build
```
Finally, run the binary file
```shell
make run_bin
```

---

> Project build with docker

```shell
make docker_build
```

```shell
make docker_run
```

---

#### ** Check out the postman collection ** 

[Postman Collection](https://github.com/faisal-porag/event_management_service_10MS/blob/master/postman_collection/event_management_system_10MS.postman_collection.json)

---

```shell
curl --location 'localhost:8090/api/v1/get-event-list?current_page=1&item_per_page=10'
```

```shell
curl --location 'localhost:8090/api/v1/get-event-details/1'
```

```shell
curl --location 'localhost:8090/api/v1/get-workshop-list/1'
```

```shell
curl --location 'localhost:8090/api/v1/get-workshop-details/1'
```

```shell
curl --location 'localhost:8090/api/v1/workshop-reservation' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "Faisal Porag",
	"email": "faisal.porag@example.com"
}'
```

