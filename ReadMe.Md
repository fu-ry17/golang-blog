# Simple Blog Service

This is a simple blog service built with Go. The project uses Docker for containerization and includes a Makefile to facilitate common development tasks.

## Requirements

- Go 1.22.4+
- Docker
- Docker Compose

## Installation

1. Clone the repository:

    ```sh
    git clone git@github.com:fu-ry17/golang-blog.git
    cd golang-blog
    ```

2. Ensure you have Docker and Docker Compose installed:
    - [Docker Installation Guide](https://docs.docker.com/get-docker/)
    - [Docker Compose Installation Guide](https://docs.docker.com/compose/install/)

## Makefile Commands

The Makefile provides several commands to help with development and deployment:

- **Build the Project**

  Compile the Go code and create an executable in the `bin` directory:

    ```sh
    make build
    ```

- **Run Tests**

  Run all tests in the project:

    ```sh
    make test
    ```

- **Run the Application**

  Build and run the application:

    ```sh
    make run
    ```

  Alternatively, you can run the application without explicitly building it first:

    ```sh
    make run
    ```

## Docker Compose Commands

- **Start the application**

  Start the application using Docker Compose. This will build the Docker image and start the containers in detached mode:

    ```sh
    make up
    ```

- **Stop the application**

  Stop the application and remove the containers:

    ```sh
    make down
    ```

- **Stop the application and remove containers, including volumes**

    ```sh
    make down-v
    ```

## Usage

### Building and Running

To build and run the application, simply use:

```sh
make run
