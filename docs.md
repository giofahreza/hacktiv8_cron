<!-- CI/CD : Build image > Create container from image (for testing purpose) > push to docker image registry > Deploy docker img -->
<!-- Local : Build image > Create container from image > Run container -->

// Docker Image
docker images //See all images
docker build -t go-echo-image . //Create docker image
docker image rm go-echo-image //Delete docker image

//Auto create container & run container
docker run -p 8080:8080 go-echo-image //Simple
docker container run --name go-echo-container --rm -itd -e PORT=8080 -e INSTANCE_ID="go echo" -p 8080:8080 go-echo-image //Complete

// Docker Container
docker container ls //See running container
docker container ls -a //See all container
docker container start go-echo-container //Start container
docker container stop go-echo-container //Stop container
docker container rm go-echo-container //Delete container
docker container create --name go-echo-container -e PORT=8080 -e INSTANCE_ID="go echo" -p 8080:8080 go-echo-image //Create container

<!-- Dockerfile -->
1. ENTRYPOINT (Can't be overridden) & CMD (Can be overridden)
2. "go mod download" only download. "go mod tidy" download and ensure that go.mod and go.sum are in a clean and accurate state
3. "apk update && apk add --no-cache git" needed if you need to fetch modules from private repositories or need to clone repositories directly. Don't do if all your module dependencies are available via the Go proxy or are public, and you’re only using go mod download to fetch them.
4. "FROM golang:1.23-alpine" Get specific go version. "FROM golang:alpine" Using latest go version
5. -itd : -i for stdin, -t for beautify to be same as your terminal config, -d detach from terminal and run in background. If run with -itd it looks not usefull because -i & -t not used. But actually you can comeback using it by type command "docker attach <container_name_or_id>" OR "docker exec -it my-container"





<!-- Cron Jobs -->
* * * * * /path/to/command
- - - - -
| | | | |
| | | | +----- Day of week (0 - 7) (Sunday is both 0 and 7)
| | | +------- Month (1 - 12)
| | +--------- Day of month (1 - 31)
| +----------- Hour (0 - 23)
+------------- Minute (0 - 59)




<!-- Upload to gcp registry image -->
docker build -t [HOSTNAME | like gcr.io]/[PROJECT-ID | go-backend-273518]/[IMAGE | example-app] .
docker build -t asia.gcr.io/trusty-mantra-434710-n9/go-echo-image .
docker push asia.gcr.io/trusty-mantra-434710-n9/go-echo-image

docker build -t \ asia-docker.pkg.dev/trusty-mantra-434710-n9/go-echo-repo/go-echo-image:v1 \ .
docker push asia-docker.pkg.dev/trusty-mantra-434710-n9/go-echo-repo/go-echo-image:v1

docker tag Nama-image gcr.io/nama-project-digcp/Nama-image
docker push gcr.io/nama-project/name-image

docker tag asdasd gcr.io/trusty-mantra-434710/asdasd
docker push gcr.io/trusty-mantra-434710-n9/asdasd

gcloud run deploy asdasd-service --image asdasd-a --platform managed --region asia-southeast2 --allow-unauthenticated --port 8080





gcloud auth login
gcloud config set project trusty-mantra-434710-n9
gcloud auth configure-docker
docker build --platform linux/amd64 -t gcr.io/trusty-mantra-434710-n9/qweqwe .
docker push gcr.io/trusty-mantra-434710-n9/qweqwe