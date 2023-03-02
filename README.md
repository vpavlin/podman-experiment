# Podman Golang Experiment

This repository contains some experimental code I produced to understand how to work with [Podman bindings](https://github.com/containers/podman/tree/main/pkg/bindings).

You can find a simple [main.go] in the root directory which executes "tasks". A task is defined as a series of actions - image pull, container start, wait for container to exit, results collection from container logs and container cleanup. The task is limited by a specified timeout.

The container image executing task must be based on [container-runtime] image and be implemented in javascript - see example [transformer].

A [Makefile] is provided for your convenience:

```
make build
make run
```

Output:
```
cd container-runtime &&\
podman build -t container-runtime .
STEP 1/4: FROM registry.access.redhat.com/ubi9/nodejs-18-minimal
STEP 2/4: ADD app/ index.js ./
--> Using cache 5e6dcd93546ba3ee410fca8365d9fdbe4b00ab419d3313ae0086848ef164d156
--> 5e6dcd93546
STEP 3/4: WORKDIR app/
--> Using cache 4e90623b7b7caa10469a579ca6e161ff6aae93b5d544db18fc942b6db9af934e
--> 4e90623b7b7
STEP 4/4: CMD node ../index.js 
--> Using cache 8c7b02efac7e17fd8f3b298fb60c8366ada8b2efd8464442f90902eb3b468ade
COMMIT container-runtime
--> 8c7b02efac7
Successfully tagged localhost/container-runtime:latest
8c7b02efac7e17fd8f3b298fb60c8366ada8b2efd8464442f90902eb3b468ade
cd transformer &&\
podman build -t transformer .
STEP 1/6: FROM container-runtime
STEP 2/6: ADD package.json package-lock.json ./
--> Using cache 34eba8351b980684860b54e82ddc12420b893df929c29d7bf510cbadf4894dac
--> 34eba8351b9
STEP 3/6: USER 0
--> Using cache be218637500365ee149fda602c65f25267e540056655b4d53a7079d091f00482
--> be218637500
STEP 4/6: RUN npm install
--> Using cache ae51088659ce143540db6e0b3358a6d1bf841fee3621df01839529c45590ea65
--> ae51088659c
STEP 5/6: USER 1001
--> Using cache 419a09978b22f0656670786ad8f566f7195985c0b6bce518f0f76e723049e6ec
--> 419a09978b2
STEP 6/6: ADD transformer.js .
--> Using cache 4467632032d21bc9d77464077dc0c39fc8dfb99e5d036b1dbafee2a63746f533
COMMIT transformer
--> 4467632032d
Successfully tagged localhost/transformer:latest
4467632032d21bc9d77464077dc0c39fc8dfb99e5d036b1dbafee2a63746f533
go run -tags "exclude_graphdriver_devicemapper exclude_graphdriver_btrfs" main.go
INFO[0000] Running task '0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85' 
INFO[0000] Pulling image: transformer                   
INFO[0000] Creating container: 0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85 
INFO[0000] Starting container: 0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85 
INFO[0000] Waiting for container 0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85 to stop 
INFO[0000] Gathering logs from container 0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85 
INFO[0000] Waiting for logs goroutine                   
INFO[0001] Extracting result from stdout:

Executing runtime
About to execute the transformer
0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85: {"Id":"changed_id","Name":"test","Number":6458.211675494412}

 
INFO[0001] Cleaning up container 0ce4d2f5-b3f7-49a0-8f07-b5f7f6c0fe85 
INFO[0001] Running task '94cb0d8d-6fa0-4faf-a15a-edf4833018ef' 
INFO[0001] Pulling image: transformer                   
INFO[0001] Creating container: 94cb0d8d-6fa0-4faf-a15a-edf4833018ef 
INFO[0001] Starting container: 94cb0d8d-6fa0-4faf-a15a-edf4833018ef 
INFO[0001] Waiting for container 94cb0d8d-6fa0-4faf-a15a-edf4833018ef to stop 
INFO[0001] Gathering logs from container 94cb0d8d-6fa0-4faf-a15a-edf4833018ef 
INFO[0001] Waiting for logs goroutine                   
INFO[0002] Extracting result from stdout:

Executing runtime
About to execute the transformer
94cb0d8d-6fa0-4faf-a15a-edf4833018ef: {"Id":"changed_id","Name":"test","Number":104.03147277338931}

 
INFO[0002] Cleaning up container 94cb0d8d-6fa0-4faf-a15a-edf4833018ef 
STDOUT: 
 {
  "Id": "changed_id",
  "Name": "test",
  "Number": 104.03147277338931
}

```