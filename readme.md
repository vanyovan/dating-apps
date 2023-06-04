# Dating App

Dating App is a web application that allows users to connect and interact with each other. Users can sign up, browse profiles, swipe left or right to indicate their interest, and upgrade to a premium package for additional features. The application is built using Go (Golang) and SQLite.

## Endpoints

The application exposes the following endpoints:

- `POST /signup`: Endpoint for user registration. Users can sign up by providing their username, password, and other required details.

- `POST /premium`: Endpoint for purchasing a premium package. Users can upgrade to a premium package to unlock unlimited swipes and receive a verified label on their profile.

## Prerequisites

Before running the application, ensure that you have the following dependencies installed:

- SQLite: Install SQLite on your machine to create and manage the application's database.

- MinGW64 (GCC): Install MinGW64 to compile and run SQLite.

## Getting Started

To get started with the Dating App, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/dating-app.git

2. Navigate to the project directory:
    ```shell
    git clone https://github.com/your-username/dating-app.git

3.  Run the SQLite command to use the database file:
     ```shell
    sqlite3 database.db

4. Build and run the application:
    ```shell
    go run main.go

5. Access the application in your web browser at http://localhost:8080


## Next Steps and Improvements
To further enhance the Dating App, consider implementing the following features:

- Swipe Transaction: Create a transaction table to track user interactions, such as when a user likes another user. This will allow you to maintain a history of swipes and facilitate matching algorithms.

- Profile Database: Extend the database schema to include a Profile table, where you can store additional user details such as name, photos, bio, and other profile-related information.