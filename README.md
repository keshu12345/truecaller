

# Matching-Prefix-Service

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



### Built With

These framework/Libraries and tools are require to build Matching-Prefix Service

* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
* ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
* ![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)



<!-- GETTING STARTED -->
## Getting Started

Given below prerequisite require to install yours system.

### Application 
There ia api contract of application 

```
curl http://localhost:8080/maching-prefixes/{:prefix_name}

```

prefix_name is  string that will match with existing collections of string to get logest prefix string

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Golang
  ```sh
  Linux:
  
  1. $ sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
  2. export PATH=$PATH:/usr/local/go/bin
  
  MacOs:
  brew install go@1.21
   
  ```
* Windows
  [Golang Installation](https://go.dev/doc/install)

* Docker
 ```
  brew cask install docker
  ```

### Installation


4. Run Application
 There is three environment 
 1. Manual
    1. Local Environment
    ```
    go run main.go -config config/local
    ```
    2. Non prod Environment

    ```
    go run main.go -config config/non-prod
    ```
    3. Prod Environment

    ```
    go run main.go -config config/prod
    ```
2. Docker

   Which environment you want to run application use environment according yours requirement

   In docker-compose file change the 
   ```
    environment:
      - CONFIG_PATH=config/local or config/non-prod or config/prod
   ```

  ```
  docker-compose up 
```
## Swagger
Get logest prefix matching string
  ```sh
  http://localhost:8080/swagger/index.html
   ```

<!-- USAGE EXAMPLES -->
## Usage

Integration with third parties tools


