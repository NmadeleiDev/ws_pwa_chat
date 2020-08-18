Enchat.

This is an open-source messenger project, dedicated to give people opportunity to chosenChat with each other and be sure that their messages are not seen to anybody else.

This project is currently in development.
Technologies used in the project:
Golang for server side message transporting and users' data management (using WebSocket protocol).
PostgreSQL to store users' accounts data.
MongoDB to store and manage users' chats data and messages.
Nginx server to serve static files, and all requests to server (both standard http and websocket).
Vue.js + Bootstrap as a tool for faster frontend development.
going to use Flutter for mobile versions of Enchat.


HOW TO SET UP THIS PROJECT

To deploy it on localhost, run "make up" in the root directory. 
To make it fully functional, you also have to deploy a mongodb replica set. In order to do so, run "make down" (after you've run "make up"), uncomment commented line no. 13 in docker-compose in mongo directory and run "make up" in root directory again. Then connect to the container with "docker exec -it mongodb mongo -u <your user from .env> -p <your password from .env>" (by default user is "admin" and password is "passwd", you can change it in mongo/.env), run "rs.initiate()" there and exit from mongo shell. Then run "make up" in the root directory again.

You also need npm to be able to build static files to nginx static dir (they are present in /backend/nginx/static folder by default, so if you don't want to rebuild them, you will be cool without npm, just don't run frontend-up in Makefile)
