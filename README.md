VerbaChat real-time messenger project.

This is an open-source messenger project, dedicated to give people opportunity to chat with each other and be sure that their messages are not seen to anybody else.

This project is currently in development.
Technologies used in the project:
Golang for server side message transporting and users' data management (using WebSocket protocol).
PostgreSQL to store users' accounts data.
MongoDB to store and manage users' chats data and messages.
Nginx server to serve static files, and all requests to server (both standard http and websocket).
Vue.js + Bootstrap as a tool for faster frontend development.
going to use Flutter for mobile versions of VerbaChat.


HOW TO SET UP THIS PROJECT

To deploy it on localhost, run "make up" in the root directory. 
To make it fully functional, you also have to deploy a mongodb replica set. In order to do so, run "make down" (after you've run "make up"), uncomment commented line no. 13 in docker-compose in mongo directory and run "make mongo-up" in root directory. Then connect to the container with "docker exec -it mongodb mongo -u <your user from .env> -p <your password from .env>" (by default user is "admin" and password is "passwd", you can change it in mongo/.env), run "rs.initiate()" there and exit from mongo shell. Then run "make up" in the root directory again.

You also need npm to be able to build static files to nginx static dir (they are present in /backend/nginx/static folder by default, so if you don't want to rebuild them, you will be cool without npm, just don't run frontend-up in Makefile)

HOW IT WORKS?

Video demonstration: https://youtu.be/5SXtI_e_0SM

After user signs up or in, websocket connection is established. Server sees him as a client and run 2 sub-processes (goroutines) - one for reading from the connection, one for writeing to connection.
Client is also subscribed to changes in all of his active chats with mongo's collection.Watch method and change stream.
When there is a new message in some of his active chats, this message is sent to his connection and gets present on user's screen (message is sent to all the chat's participants).

On the current stage frontend part of the project is fairly simple, as my focus is on server part now.

In order to test the performance of the app, you can sign up a couple of users, and chat between them.  