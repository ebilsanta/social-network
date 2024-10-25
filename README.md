# Social Network Application

## Overview
This is a social network application built to explore and learn various modern web technologies. The application allows users to:
- Create user accounts
- Post updates and content
- Follow other users
- View a curated feed of posts
- Receive notifications for new followers and posts

Some of the following parts are still in progress, such as the uploading of posts, user login using Firebase (or some Google integration), and developing a frontend.

## Architecture
Below is a high-level system design diagram that outlines the key components of the application architecture:

<img width="2608" alt="social-network (2)" src="https://github.com/user-attachments/assets/5be6c6ff-ed58-4ce5-947c-9d5a5ba00163">

## Key Technologies
This project leverages a wide range of technologies, including:
- **Backend**: Golang
- **Databases**: PostgreSQL, MongoDB, Neo4j Graph Database
- **Caching and Queuing**: Redis, Kafka
- **API Gateway**: Kong API Gateway
- **Cloud Storage and Delivery**: Amazon S3, AWS CloudFront
- **Real-time Notifications**: Firebase
- **Containerization**: Docker Compose
- **Inter-service Communication**: gRPC

## Key Features
- **Prefetching and Caching**: Efficient feed loading using caching mechanisms to improve user experience.
- **Event-driven Architecture**: Kafka is used to process events such as:
  - Generating user feeds when a post is created or a user is followed.
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
