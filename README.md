# Exercice 6: CI/CD

This exercice aims to understand the Continious Integration and Continous Delivery/Deployment. 

<img src="Exercice6-diagram.jpg" alt="drawing" width="350"/>

Three VMs were created, one working as the client, one for the Jenkins and another one with the Docker registry and the service.

## Jenkins VM
A Jenkins Docker was installed using the following command: 

`sudo docker run -p 8080:8080 -p 50000:50000 --name="jenkinsDocker" -v /var/run/docker.sock:/var/run/docker.sock jenkins/jenkins:lts`

**Note:** It is really important to add the volume (-v) or the Docker won't have the right permissions. 

Once it is deployed, the insecure registry list has to be updated. From both the Jenkins VM and the Docker Registry VM, go to: `/etc/docker/daemon.json` and add the IP from the Docker Registry VM: 
```
{
    "insecure-registries" : [ "192.168.176.144:5000" ]
}
```

Then run: `sudo systemctl daemon-reload` and `systemctl restart docker`. 

Now that Jenkins has been properly set up, it can be accessed through the browser from the VM accessing the localhost:8080. 

Install the recommended plugins and some additional ones: 
* Docker: All the plugins related to docker.
* Golang

Then, one has to create two pipelines for each service (post and get): one to scan for any updates in GitHub, and if there are any builds and publishes the image to the repository, and another one to pull and run the uploaded image. In this particular scenario, the following code will reference the post (registration) service, but it can be extrapolated for the get (info) service. 

### Registration Service Multibranch Pipeline
This pipeline was chosen to be multibranch in order to easily enable the source code management from GitHub. The follow configuration was followed:
* Branch sources: GitHub. Any credential were added (since the repository is public) and a "discover branches" behavior was added, with an "all branches" strategy. 
* Scan Repository Triggers: the "periodically if not otherwise run" was selected, a 1 minute interval was chosen. 
* The rest of the paremeters were untouched. 

The pipeline source code was uploaded to the repository as a Jenkinsfile, as well as the code for the service (main.go and Dockerfile). 

If everything was configured properly, from now on, everytime the user commits and pushes a modification from the code, this pipeline automatically scans, reads, builds and posts the new image. 

### Registration Publish Pipeline
This pipeline is in charge of, first of all, perform several actions as if it was the Docker Registry VM containing all the services (in reality it is using REST commands to perform this actions. This is why later in this document, in the Docker Registry VM section, docker remote API configurations are changed). This actions are pulling the image from the registry and then running it with the desired configuration. 

The code for this pipeline it is stored in Jenkins, but it was uploaded to this repository for documentation purposes with the name 

## Docker Registry VM
A docker registry was deployed using the following command: 

`docker run -p 5000:5000 -v /var/run/docker.sock:/var/run/docker.sock --name="docker_registry" registry:2`

From now on, from any local device, on could see the images uploaded in this registry by accessing its 'IP address'/v2/_catalog: 
http://192.168.176.144:5000/v2/_catalog


