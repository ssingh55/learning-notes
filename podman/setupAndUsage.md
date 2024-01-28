### Podman is a utility provided as part of the libpod library. It can be used to create and maintain containers. This docs will teach you how to set up Podman and perform some basic commands.

## Podman Documentation
    The documentation for Podman is located https://docs.podman.io/.

# Install podman

## 1. Update the System:
    Before you begin, ensure that your system package database is up-to-date:

    ```sudo apt update && sudo apt upgrade -y ```

## 2. Install Podman:
    Podman is available in the default Ubuntu repositories, so installing it is straightforward:

    ```sudo apt install -y podman ```

or

### Directly if  user dont want to install from ubuntu ppa then they can run this to install from kubic  repo

    ```
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL "https://download.opensuse.org/repositories/devel:kubic:libcontainers:unstable/xUbuntu_$(lsb_release -rs)/Release.key" \
    | gpg --dearmor \
    | sudo tee /etc/apt/keyrings/devel_kubic_libcontainers_unstable.gpg > /dev/null
    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/devel_kubic_libcontainers_unstable.gpg]\
        https://download.opensuse.org/repositories/devel:kubic:libcontainers:unstable/xUbuntu_$(lsb_release -rs)/ /" \
    | sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:unstable.list > /dev/null
    sudo apt-get update -qq
    sudo apt-get -qq -y install podman
    ```

## 3. Verify the Installation:
    To ensure Podman has been installed correctly:

    > podman --version 
    You should see the version of Podman printed to the console.
    ```
    shubham@shubham-prlappy:~$ podman --version
    podman version 4.6.2
    ```

# Setup podman

    Start the service
    > sudo service podman start

    if  u want to enable on boot
    > sudo service podman enable

# Set Podman registries conf (optional)
    If there is no repository to fetch and install container Images by Podman, then we need to add that manually. We can use popular repositories such as Docker.io, Redhat, and Fedora.

    ## Edit Registry configuration file 

    > sudo vim /etc/containers/registries.conf
    Scroll to the end of the file and paste the given line:

    ```
    [registries.search]
    registries=["registry.access.redhat.com", "registry.fedoraproject.org", "docker.io"]
    ```
    After that save the file by pressing Ctrl+O, hit the Enter key, and then exit using Ctrl+X.

    To get full information about Podman, you can use:
    > podman info

    U can see the information similar to this
    ```
    .....
    registries:
      search:
      - registry.fedoraproject.org
      - registry.access.redhat.com
      - docker.io
      - quay.io
    ...... 
    ```

# Basic Usage of Podman

## 1. Getting help
    
    To get some help and find out how Podman is working, you can use the help:
    > podman --help
    > podman <subcommand> --help

    For more details, you can review the manpages:
    > man podman
    > man podman-<subcommand>

## 2. Create Container
    
    Once you have the image of the application that you want, we can create a container using it. Here we have downloaded two images one is of MariaDB and the other of Ubuntu, hence we will show how to use it to create a container using Ubuntu Image.
    > podman run -dit --name ubuntucontainer ubuntu

    –name is a parameter to give the container whatever friendly name we want to assign, here we have used ubuntucontainer whereas Ubuntu is the name of the Image we have downloaded and want to use.

    Get running container command-line access:
    > podman attach container-name
    
    example:
    >> podman attach ubuntucontainer
    
## 3. Stop or Start a Container
    If you want to start or stop any container, then the command will be:

    To stop:
    > podman stop container-id or name

    example:
    >> podman stop ubuntucontainer

    To Start:
    > podman start container-id or name

    example:
    >> podman stop ubuntucontainer

## 4. Running Containers:
    The syntax for running a container with Podman is very similar to Docker:

    > podman run -it hello-world 
    This will download the hello-world docker image (if not already present) and launch a new container. On successful run, you will see the following result:

    ```
    shubham@shubham-prlappy:~$ podman run -it hello-world 
    Resolved "hello-world" as an alias (/etc/containers/registries.conf.d/shortnames.conf)
    Trying to pull docker.io/library/hello-world:latest...
    Getting image source signatures
    Copying blob c1ec31eb5944 done  
    Copying config d2c94e258d done  
    Writing manifest to image destination

    Hello from Docker!
    This message shows that your installation appears to be working correctly.
    ```

## 5. Listing Containers:

    To list running containers:
    > podman ps 

    To list all containers (including stopped ones):
    > podman ps -a 

    ```
    shubham@shubham-prlappy:~$ podman ps -a
    CONTAINER ID  IMAGE                                 COMMAND     CREATED        STATUS                    PORTS       NAMES
    ddd3b1188fc2  docker.io/library/hello-world:latest  /hello      4 minutes ago  Exited (0) 4 minutes ago              magical_chatterjee
    ```

## 6. Removing Containers:

    To remove a container:
    > podman rm <container_id> 

## 7. Inspecting a running container
    
    You can "inspect" a running container for metadata and details about itself. podman inspect will provide lots of useful information like environment variables, network settings or allocated resources.

    Since, the container is running in rootless mode, no IP Address is assigned to the container.

    > podman inspect -l | grep IPAddress
            "IPAddress": "",

    Note: The -l is a convenience argument for latest container. You can also use the container's ID or name instead of -l or the long argument --latest.

    Viewing the container's logs
    You can view the container's logs with Podman as well:

    > podman logs -l
    ```
    127.0.0.1 - - [04/May/2020:08:33:48 +0000] "GET / HTTP/1.1" 200 45
    127.0.0.1 - - [04/May/2020:08:33:50 +0000] "GET / HTTP/1.1" 200 45
    127.0.0.1 - - [04/May/2020:08:33:51 +0000] "GET / HTTP/1.1" 200 45
    127.0.0.1 - - [04/May/2020:08:33:51 +0000] "GET / HTTP/1.1" 200 45
    127.0.0.1 - - [04/May/2020:08:33:52 +0000] "GET / HTTP/1.1" 200 45
    127.0.0.1 - - [04/May/2020:08:33:52 +0000] "GET / HTTP/1.1" 200 45
    ```

    Viewing the container's pids
    You can observe the httpd pid in the container with podman top.

    > podman top -l
    ```
    USER     PID   PPID   %CPU    ELAPSED            TTY     TIME   COMMAND
    root     1     0      0.000   22m13.33281018s    pts/0   0s     httpd -DFOREGROUND
    daemon   3     1      0.000   22m13.333132179s   pts/0   0s     httpd -DFOREGROUND
    daemon   4     1      0.000   22m13.333276305s   pts/0   0s     httpd -DFOREGROUND
    daemon   5     1      0.000   22m13.333818476s   pts/0   0s     httpd -DFOREGROUND
    ```

## 8. Handling Images:

    Podman can search for images on remote registries with some simple keywords.
    > podman search <search_term>

    You can also enhance your search with filters:
    > podman search httpd --filter=is-official

    Downloading (Pulling) an image:
    > podman pull docker.io/library/httpd
    
    List all available images on your system:
    > podman images 

    Remove an image:
    > podman rmi <image_id> 

## 9. Podman Pods:
    
    Podman has a concept called “pods”. A pod is a group of one or more containers that share the same network namespace. This is similar to Kubernetes pods. To create a new pod:
    > podman pod create --name mypod 
    
    Run a container within the pod:
    > podman run --pod mypod -d <image> 


## 10. Rootless Containers:
    One of Podman’s main features is the ability to run rootless containers. This means you can run containers as a non-root user, without any special permissions. Simply run the podman command as your regular user.

## 11. Using Volumes:
    
    You can mount volumes (directories or files from the host) into your containers. For example:
    > podman run -v /path/on/host:/path/in/container -it  /bin/bash 


### For more configurations and installation on other os u can follow this official link
    > https://podman.io/docs/installation#installing-on-linux