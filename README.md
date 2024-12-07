# TRINITY APP

## GitHub repository containing:

- Source code for implemented component(s):https://github.com/huynhhung30/Trinity-App
- API Documentation: http://localhost:5001/swagger/index.html
- Setup instructions:
  - step 1: git clone https://github.com/huynhhung30/Trinity-App.git
  - step 2: make up (build images and create containers on docker or use " make run " to run local )
  - step 3: run api http://localhost:5001/api/v1/trinity/create-table to insert table for test

## documenting

    * Technical decisions:
        + Go with Gin framework (vscode) and Mysql (Mysql workbench) for Database
        + Use Docker to create an environment for the database as well as run source code in a local environment.
    * Assumptions made:
        + with basic functions first create user and login user use token to register silver member and apply promotion coupon for the first 100 people to register.
        + create a coupon table to manage possible cases like:
        + coupon expires
        + is used for the first 100 people

    * Future improvements:
        + Currently, there is no complete control over members registering for silver packages, so it is possible to create an additional table for managing members registering for other packages in the future. Add management of expired membership packages and payment methods.

    * Local setup guide:
        make up (to set up docker and test)
