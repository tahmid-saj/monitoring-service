# monitoring-service

Pub/Sub monitoring service for subscribing to sites and publishing events when they go down or come back up. Developed using Go, TypeScript, PostgreSQL, Webhooks, Encore.

<img width="1835" alt="image" src="https://github.com/user-attachments/assets/058da0cb-2c71-462e-aaac-1633f9b85a66">

<img width="845" alt="image" src="https://github.com/user-attachments/assets/4a12607c-5325-4698-b002-4b22bb82dc33">

<br/>
<br/>

## Directory structure

The directory structure is as follows:

## Directory Structure

- **monitor/**
  - Go-based service for monitoring websites.

- **site/**
  - Site-related functionalities and components for monitoring.

- **slack/**
  - Integration with Slack for notifications.

- **encore.app**
  - Encore configuration for cloud-native service deployments.

- **go.mod**: Go module dependencies.

- **go.sum**: Checksum file for Go module versions.

- **README.md**: Project documentation.

- **.gitignore**: Specifies files to be ignored by Git.

<br/>
<br/>

## Overview

### Design

The service is used mainly with a logging and notification service. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">
