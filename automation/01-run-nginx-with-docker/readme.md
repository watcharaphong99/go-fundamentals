Run Nginx with Docker



1. Run command
$ docker-compose up -d
Creating network "1202-run-nginx-with-docker_default" with the default driver
Pulling nginx (3dsinteractive/nginx:1.12)...


2. Run command
$ docker ps
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS          PORTS                                         NAMES


3. Run command to request http from nginx
$ curl -X GET "http://localhost:8080"
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

4. Run command to stop nginx
$ docker-compose down
Stopping nginx ... done
Removing nginx ... done
Removing network 1202-run-nginx-with-docker_default

5. Run command to check nginx is stop and remove
$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
