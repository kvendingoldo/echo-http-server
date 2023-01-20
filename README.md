# echo-http-server

A simple HTTP server is a program that listens for incoming requests from clients on a specified port (`SERVER_PORT`), 
and then sends a response back to the client. The application is delivered as docker container.

You may try it via the following command: `docker run -it -p 8089:8089 -e SERVER_PORT=8089 kvendingoldo/echo-http-server`