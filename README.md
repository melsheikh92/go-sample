## Go mux CI/CD sample (Github actions + Docker + Helm + ArgoCD)

## **Overview**

This repository contains a GoLang web server service implemented using the **mux** router. The project is containerized using Docker, and the deployment is managed using Kubernetes and Helm. Additionally, GitHub Actions are set up for continuous integration, including building the Docker image, version bumping, and publishing to Docker Hub. The application can be easily deployed and managed using Helm charts and imported into Argo CD for continuous delivery.

## **Project Structure**

*   **appversion**: File containing the application version.
*   **main.go**:  main application server code.
*   **Dockerfile**: Dockerfile for building the application image.
*   **docker-compose.yaml**: Docker Compose file for local development.
*   **charts**: Helm chart folder for deployment.
    *   **values.yaml**: Helm chart values file with an image tag and other configuration.
*   **Kubernetes**: Kubernetes deployment files.
    *   **deployment.yaml**: Deployment file for Kubernetes.
    *   **service.yaml**: Service file with NodePort configuration.
*   **argocd**: Argo CD folder.
    *   **application.yaml**: Argo CD application file for importing the application.

## **Usage**

### **Local Development**

To run the application locally using Docker Compose:

`docker-compose up`  
 
### **Building Docker Image**

`helm install/upgrade <release-name> charts/`

### **Deployment using Kubernetes**

Apply the Kubernetes deployment and service files:

`kubectl apply -f kubernetes/deployment.yaml`  
`kubectl apply -f kubernetes/service.yaml`

### **Argo CD Integration**

Import the application into Argo CD using the provided **argocd/application.yaml** file.

## **Configuration**

### **Application Version**

The application version is stored in the **appversion** file. It is automatically bumped during the GitHub Actions pipeline.

### **Helm Values**

Configure the Helm chart using the **charts/values.yaml** file. Adjust image details, replicas, and other settings as needed.

### **Kubernetes Configuration**

Modify the Kubernetes deployment and service files in the **kubernetes/** folder to customize the deployment.

### **Argo CD Application**

Update the Argo CD application file (**argocd/application.yaml**) with the appropriate details for your deployment.

## **License**

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/melsheikh92/go-sample/blob/main/LICENSE) file for details.
