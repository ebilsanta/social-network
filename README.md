# Social Network Application

## Overview
This is a social network application built to explore and learn various modern web technologies. The application allows users to:
- Create user accounts
- Post updates and content
- Follow other users
- View a curated feed of posts
- Receive notifications for new followers and posts

Some of the following parts are still in progress, such as the uploading of posts, sending of notifications and developing the frontend.

## Demo
As I am still working on developing the frontend, I haven't had time to deploy the whole project. Below is a demo video showcasing the current frontend with full backend integration:
<video src="https://github.com/user-attachments/assets/87613904-29c0-4e5a-8f52-e8050edc98fd" />

## Architecture
Below is a high-level system design diagram that outlines the key components of the application architecture:

<img width="2608" alt="Social Network System Design" src="https://github.com/user-attachments/assets/540d1fa2-90c1-489f-b008-31c0730b1174" />

## Key Technologies
This project is designed to leverage a wide range of technologies, including:
- **Backend**: Golang
- **Databases**: PostgreSQL, MongoDB, Neo4j Graph Database
- **Caching and Queuing**: Redis, Kafka
- **API Gateway**: Kong API Gateway
- **Cloud Storage and Delivery**: Amazon S3, AWS CloudFront (unimplemented)
- **Real-time Notifications**: Firebase (unimplemented)
- **Containerization**: Docker Compose
- **Inter-service Communication**: gRPC

## Key Features
- **Prefetching and Caching**: Efficient feed generation using caching mechanisms to improve user experience.
- **Event-driven Architecture**: Kafka is used to process events such as:
  - Generating user feeds when a post is created or a user is followed.
  - Updating user post/following counts asynchronously
  - Sending notifications when users follow each other or create posts.

## Getting Started
To run the project locally using Docker Compose, follow these steps:

### Prerequisites
- Docker and Docker Compose installed on your machine.
- Ensure you have the required environment variables for services like PostgreSQL, MongoDB, Neo4j, Redis.

### Setup
1. Clone the repository:
    ```bash
    git clone https://github.com/ebilsanta/social-network.git
    cd social-network
    ```
2. Set up environment variables such as service ports, database URLs, Kafka configurations and rename the `.env.sample` file to `.env`

3. Build and start the services with Docker Compose:
    ```bash
    docker-compose up --build
    ```

4. Access the backend APIs using the postman file provided.
