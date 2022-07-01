# Cleaning Pattern GO
[![Cleaning Pattern GO](https://github.com/rizwijaya/CleaningPatternGO/actions/workflows/go.yml/badge.svg)](https://github.com/rizwijaya/CleaningPatternGO/actions/workflows/go.yml)

Golang API, HTML (Go Template), and Mysql Starter Kit

---
### Table of Contents
- [About Us Cleaning Pattern Go](#about-us)
- [Features Cleaning Pattern Go](#features)
- [Base Code Configuration](#base-code-configuration)
- [Getting Started](#getting-started)
    + [Getting Started with GO](#getting-started-with-go)
    + [Getting Started with Makefile](#getting-started-with-makefile)
- [Author](#author)

---

### About Us
Cleaning Pattern Go is a basic code that is made with the concept and pattern of the Repository-Service Pattern and the Clean Architecture.

Cleaning Pattern Go uses a Mysql database that can be customized according to your needs. In addition, it is also equipped with an HTML Parser so that you can use HTML directly such as HTML, CSS and other Front-End Frameworks such as Bootstraap, CSS Tailwind and many more.

HTML Parser is an part of Cleaning Pattern GO in one service that makes it easy to use without requiring many services to run.

---
### Features
##### Pattern
- [x] Repository-Service Pattern
- [x] Clean Architecture
##### Front-End
- [x] HTML Parser (GoTemplate)
- [x] Support All Front-End HTML Frameworks 
##### Database
- [x] Mysql/MariaDB Database
- [x] Database Can Be Customized
##### API
- [x] Support All HTTP Methods  
- [x] Custom API Response and Result
- [x] Custom API Error
- [x] API Versioning

##### Authentication
- [x] HTTP Basic Authentication
- [x] JWT Authentication
    - [x] Generate JWT Token
    - [x] Verify JWT Token
- [x] Can Be Added as Desired

##### Other
- [x] Error Handling
- [x] Support Code Versioning
- [x] Code Maintenance and Management
- [x] Support for Makefile
- [x] Support for Dockerfile
---

### Base Code Configuration
To configure the code base, it can be done by accessing the environment.
##### Setup Apps
- APP_MODE : Application Mode is debug or production 
- APP_NAME : Application Name
- APP_PORT : Application In Running Port
- APP_URL  : Application URL
- APP_SECRET : Secret Key Token Cookie and JWT
##### Setup Database
- DB_HOST : Database HostName
- DB_NAME : Database Name
- DB_USER : Database UserName
- DB_PASSWORD : Database Password
- DB_PORT : Database Port

##### Basic HTTP Authentication
- BASIC_AUTH_USER : Basic Authentication UserName
- BASIC_AUTH_PASSWORD : Basic Authentication Password

---
### Getting Started

To get started, you can clone the repository and run the following command:

```bash
git clone https://github.com/rizwijaya/CleaningPatternGO.git
```
###### Getting Started with GO
Initialize the project by running the following command:

```bash
go mod init CleaningPatternGO
go mod tidy
```

If you want to run the project directly without creating a binary file, you can use the command:

```bash
go run main.go
```
But if you want to build the project with binary files, you can run the following command to build the project:

```bash
go build main.go
```
Then, you can run the project by running the following command:

```bash
./main
```

###### Getting Started with Makefile
Initialize the project Depedencies by running the following command:

```bash
make init
```
Download and Install the project Depedencies by running the following command:

```bash
make install
```

But if you want to build the project with binary files, you can run the following command to build the project:

```bash
make run
```

Then, you can run the following command to build the project:

```bash
make build
```
Then, you can run the project by running the following command:

```bash
make start
```

Cleaning Golang Binary Project with Makefile, you can following commands:

```bash
make clean
```

Run with nodemon by running the following command:

```bash
make run-nodemon
```

---
### Author 
##### <img src="https://github.com/rizwijaya/rizwijaya/blob/main/Assets/Developer.gif" height="22px"> Rizqi Wijaya 
###### Github: https://github.com/rizwijaya
###### LinkedIn: https://www.linkedin.com/in/muhammad-rizqi-wijaya/
###### Contact Me:
<p>
    <a href="https://www.linkedin.com/in/muhammad-rizqi-wijaya" target="_blank"><img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" /></a> 
    <a href="mailto:rizwijaya241@gmail.com" target="_blank"><img alt="Gmail" src="https://img.shields.io/badge/gmail-D14836?&style=for-the-badge&logo=gmail&logoColor=white" /></a>  
    <a href="https://www.instagram.com/rizwijaya21" target="_blank"><img alt="Instagram" src="https://img.shields.io/badge/instagram-%23E4405F.svg?&style=for-the-badge&logo=instagram&logoColor=white" /></a>  
</p>
