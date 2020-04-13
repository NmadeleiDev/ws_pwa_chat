This project is currently in development.

To deploy it on localhost, run "make up" in the root directory. 
To make it fully functional, you also have to deploy a mongodb replica set. In order to do so, run "make down" (after you've run "make up"), uncomment commented line no. 13 in docker-compose in mongo directory and run "make mongo-up" in root directory. Then connect to the container with "docker exec -it mongodb mongo -u <your user from .env> -p <your password from .env>", run "rs.initiate()" there and exit from mongo shell. Then run "make up" in the root directory again.

On the current stage frontend part of the project is fairly simple, as my focus is on server part now.

In order to test the performance of the app, you can sign up a couple of users, and chat between them.  