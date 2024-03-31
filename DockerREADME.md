Running a Go project in a Docker container involves creating a Docker image that contains your Go application, and then running a container based on that image. Here's a step-by-step guide to accomplish this:

### Step 1: Create a Dockerfile

A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. For a simple Go application, the Dockerfile can be quite straightforward. Create a file named `Dockerfile` in the root of your Go project with the following content:

```Dockerfile
# Use the official Go image as a parent image
FROM golang:1.16 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build your program for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use a Docker multi-stage build to create a lean production image
# Start with a scratch image
FROM scratch

# Copy the binary from the builder stage to the production image
COPY --from=builder /app/myapp /myapp

# Run the binary program produced by `go install`
CMD ["/myapp"]
```

This Dockerfile does a few things:
- It starts with a base Go image (`golang:1.16`) to build your application. You can adjust the Go version as needed.
- It sets the working directory to `/app` inside the image.
- It copies your Go source code into the image.
- It compiles your application to a binary named `myapp`.
- It then starts a new stage (`FROM scratch`) to create a minimal final image, copies the binary from the first stage, and sets the command to run your application.

### Step 2: Build Your Docker Image

Navigate to the directory containing your Dockerfile and run the following command to build your Docker image. Replace `<your-image-name>` with the name you want to give to your image:

```sh
docker build -t <your-image-name> .
```

This command builds a new Docker image locally using the Dockerfile in the current directory.

### Step 3: Run Your Docker Container

Once the image is built, you can run a container based on that image. Replace `<your-image-name>` with the name you used in the previous step:

```sh
docker run --name mygoapp -d <your-image-name>
```

This command runs your Go application in a new container named `mygoapp`. The `-d` flag runs the container in detached mode, meaning it runs in the background.

### Step 4: Verify the Application is Running

To ensure your application is running as expected, you can check the logs of the container:

```sh
docker logs mygoapp
```

Replace `mygoapp` with the name of your container. This command displays the console output of your application, which is particularly useful if your application prints to stdout.

### Summary

You've created a Dockerfile for your Go application, built a Docker image from that Dockerfile, and run your application inside a Docker container. This process encapsulates your application and its environment, making it easy to deploy and run on any system that supports Docker.