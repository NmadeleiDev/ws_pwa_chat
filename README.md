This project is currently in development.

To deploy it on localhost, run "make up" in the root directory. 
To make it fully functional, you also have to deploy a mongodb replica set. In order to do so, run "make down" (after you've run "make up"), uncomment commented line no. 13 in docker-compose in mongo directory and run "make mongo-up" in root directory. Then connect to the container with "docker exec -it mongodb mongo -u <your user from .env> -p <your password from .env>" (by default user is "admin" and password is "passwd", you can change it in mongo/.env), run "rs.initiate()" there and exit from mongo shell. Then run "make up" in the root directory again.

You also need npm to be able to build static files to nginx static dir (they are present in /backend/nginx/static folder by default, so if you don't want to rebuild them, you will be cool without npm, just don't run frontend-up in Makefile)

On the current stage frontend part of the project is fairly simple, as my focus is on server part now.

In order to test the performance of the app, you can sign up a couple of users, and chat between them.  