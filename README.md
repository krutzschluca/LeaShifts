# LeaShifts API



## Overview
LeaShifts API is a modern RESTful service designed to manage store operations including employees, stores, and shifts. It is developed in Go and hosted on Render, utilizing a PostgreSQL database hosted on Neon. This setup ensures scalable and reliable access to the application, which is ready to handle real-world operational data.

## Running Frontend
1. **Install a local Server of your choice** Start your local server from the root directory of our Project
3. **Open Website in your Borwser** Goto localhost:xxxx/frontend/ in your Browser of Choice

## Accessing the APIs
The APIs are accessible online without the need for local setup or installation. For interacting with the APIs, use the provided Postman collection which includes comprehensive documentation for each endpoint.

### Using Postman to Access the APIs
1. **Download and Install Postman**: Ensure you have Postman installed on your machine. You can download it from [Postman's official site](https://www.postman.com/downloads/).
2. **Import the Postman Collection**: The collection provided with this project contains all the endpoints configured with example requests. 
3. **Set Up Environment Variables**: Configure the necessary variables within Postman to match the deployment settings such as the base URL.

## Deployment on Render
The application is deployed on Render through a continuous deployment pipeline connected to a GitHub repository. Here’s how the deployment process works:
1. **GitHub Integration**: The Render platform is directly linked to the GitHub repository containing the Go project.
2. **Automatic Builds and Deployment**: Whenever changes are pushed to the linked branch in the GitHub repository, Render automatically triggers a build and deployment process.
3. **Immediate Updates**: Changes go live immediately after the deployment process completes, ensuring zero downtime.

## API Gateway Configuration
The API gateway is set up to route requests, handle sessions, and provide necessary security measures such as authentication and rate limiting. Here’s a brief on the setup:
- **Request Routing**: Directs incoming requests to the correct service endpoints.
- **Security Layers**: Includes configurations for secure HTTPS connections and data validation.
- **Performance Optimizations**: Caching and request rate limiting to ensure optimal performance.

## Testing the APIs
To test the APIs, follow these steps:
1. **Launch Postman**: Open the Postman application on your computer.
2. **Load the Collection**: Use the imported collection which contains configured requests.
3. **Execute Requests**: Send requests from Postman to test different functionalities of the API.

## API Documentation
The API documentation is included within the Postman collection, providing details on request formats, expected responses, and handling of different HTTP methods.

## Additional Notes
- The application does not require local execution since it is fully hosted and managed on Render, offering a scalable and maintenance-free operation.
- For detailed insights into the API functionalities, refer to the Postman documentation included in the collection.
- After previously having built the application using PHP, we decided to rewrite everything into GO. We were not able to connect to the database hosted on Neon using PHP and since it did work for a colleague of ours that was using Go, we decided to use Go as well.

## Contact Information
For more information or queries regarding the API usage, feel free to contact us, Luca Krutzsch and/or Pieter Vögele.

---
