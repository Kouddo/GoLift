# GoLift


### Backend Tech Stack

#### Apache
  -Used as reverse proxy, forwards requests to backend server

  -Manages rate limiting, HTTPS, and TLS for security purposes


#### CockroachDB
  -serverless database instance for scalability


#### Backend Server
  -Hosted on AWS EC2 instance

  -Written in Go, which has an HTTP implementation that automatically manages concurrency through goroutines, allowing for performance and scalability


