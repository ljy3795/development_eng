# MLS-Bootcamp (by Colin)

---

## Prerequisite studies
- **Frontend Basics** (~ 1/2)
  - HTML  (ref. https://developer.mozilla.org/ko/docs/Learn/HTML)
  - Java Script  (ref. https://developer.mozilla.org/en-US/docs/Learn/JavaScript)
  - CSS (A little)
  
  
- **Backend** (~ 1/15, Reviews with Brett(1/10, 1/15)) 
  - Go Lang (ref. https://www.bogotobogo.com/GoLang/GoLang_Visual_Studio_Code.php)
  - Gin-gonic (ref. https://github.com/gin-gonic/gin)
  
- **Frontend** (~ 1/30)
  - React (ref. https://velopert.com/3613)
  
---

## Bootcamp
- **Basics**
  - [git : *learngitbraching*](https://learngitbranching.js.org/)
  - [python : *AIPE_Bootcamp*](https://docs.google.com/document/d/1p7op_bQURuW0dRkaAbcnHn5_pbxfhVEJgEigCBo2cZg/edit#heading=h.ql6ig3pd3vhi)
  - [go : *AIPE Bootcamp*](https://docs.google.com/document/d/1fh-efmn4YyOYgSutDRDWTPADTeGTuSD0NaE688jdDo0/edit#heading=h.6y0smk1l8amz)
  
- **Assignment**
  - [Take-home](https://docs.google.com/document/d/1b9CRwoiy20sWLdmb0JL3fSJJZLebRX6USyfssZZQcMY/edit#heading=h.m874wemrtk4)

- **Reviews**
  - With Bret (1/10, 1/15, 1/29)
  - With Kevin (1/29)

--- 
## Assignment
- Build a simple service of PM2.5(초미세먼지) using SEOUL Open API
- With three microservices of DB / Backend(Go / Gin) / Frotnend(React) with Docker Containers
<br>
<img src="https://user-images.githubusercontent.com/25445292/73420509-136f2e00-4366-11ea-90d9-f105e5cdaaf5.png" width="50%"/>
<br>

#### Step 1. (~1/15)
- DB (MariaDB) / Backend (Go & Gin) / Frontend (HTML on Go)
  - Tentative result without *REACT* frontend
  - http_controller.go connected to REST API renders HTML files (No frontend server)
- [Code Branch : mls-bootcamp-html](https://github.com/ljy3795/development_eng/blob/go_html_mls)

#### Step 2. (~1/30)
- DB (MariaDB) / Backend (Go & Gin) / Frontend (HTML on Go)

### Details
**Backend**
  - **structure**
  ```
      .
    ├── Dockerfile
    ├── clients
    │   └── rds.go
    ├── controller
    │   └── http_controller.go
    ├── docker_run.sh
    ├── gin.log
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
    │   └── models.go
    ├── services
        ├── api_call.go
        ├── backfill.go
        └── scheduler.go
  ```
  - **API**
    - *READ*
      ```
      [GET] /api/read/:dt/:region (ex. /api/read/20200115/강남구)
      ```
    - *DELETE*
      ```
      [POST] /api/delete/:dt/:region (ex. /api/delete/20200115/강남구)
      ```
    - *CREATE*
      ```
      [POST] /api/create (Receive a form data and insert to DB)
      ```
  - **Local Run**
    ```
      # 1) SET .env 
      $ SEOURL_API_KEY=${API_KEY}
      $ MARIADB_USER=${USER}
      $ MARIADB_PASS=${PASS}
      $ MARIADB_ADDR=${ADDR}
      $ MARIADB_PORT=${PORT}

      # 2) Run docker container(application)
      $ ./server/docker_run.sh
      
      # Could be accessed through 8080 port 
    ```
    
<br>

**Frontend**
  - **Structure**(src)  
  ```
      .
    ├── client
    │   └── Root.js
    ├── components
    │   ├── Adder.css
    │   ├── Adder.js
    │   ├── AdderTemplate.css
    │   ├── AdderTemplate.js
    │   ├── HomeViewer.js
    │   ├── Navigator.css
    │   ├── Navigator.js
    │   ├── Updater.js
    │   ├── Viewer.js
    │   ├── ViewerTemplate.css
    │   └── ViewerTemplate.js
    ├── index.css
    ├── index.js
    ├── lib
    │   └── api.js
    ├── serviceWorker.js
    ├── setupTests.js
    └── shared
        ├── App.css
        └── App.js
  ```
  - **Route**
    - *Home*
      ```
      Path : "/"
      Actions : 
        - Call Viewer Page with key values(dt / region) 
        - Add a new row
      ```
    - *Viewer*
      ```
      Path : "/view/:dt/:region"
      Actions : 
        - Rendering a view page
        - Navigating prev/next date of a page
        - Removing the selected row of the page
        - Redireting to a Update page
      ```
    - *Adder*
      ```
      Path : "/add"
      Actions : 
        - Adding a new row with HTML inputs
      ```
  - **Local run**
    ```
      # 1) Run docker container(application)
      $ ./client/docker_run.sh
      
      # Could be accessed through 3000 port 
    ```


---
#### 느낀점
