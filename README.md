# TRINITY APP

## GitHub repository containing:

- Source code for implemented component(s):https://github.com/huynhhung30/Golang-Redis-Gin
- API Documentation: http://localhost:5001/swagger/index.html
- Setup instructions:
  - step 1: git clone https://github.com/huynhhung30/Golang-Redis-Gin.git
  - step 2: make up (build images and create containers on docker or use " make run " to run local )
  - step 3: run api http://localhost:5001/api/v1/trinity/create-table to insert table for test

## documenting

    * Technical decisions:
        + Go with Gin framework (vscode) and Mysql (Mysql workbench) for Database
        + Use Docker to create an environment for the database as well as run source code in a local environment.

    * Local setup guide:
        + make up (to set up and docker and test)
        + make run (to run local)
