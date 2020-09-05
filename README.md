# k8shelloworld

docker build -f Dockerfile.scratch . -t javiroman/helloworld:v1
docker build -f Dockerfile.alpine . -t javiroman/helloworld:v2

Simple Go hello world for testing kubernetes

$ docker run  -d -p 8080:8080 --name hello javiroman/helloworld:v1
$ curl localhost:8080
CPUs: 4
Hostname: 1982d2363b4c

$ docker run  -d -p 8080:8080 --cpuset-cpus=0 --name javiroman/helloworld:v1
$ curl localhost:8080
CPUs: 1
Hostname: 8240043c0558

$ docker run  -d -p 8080:8080 --name hello javiroman/helloworld:v2
$ docker exec -ti hello sh

/ # dig @8.8.8.8 www.redhat.com +short
23.214.212.161

/ # netstat -putona
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
PID/Program name     Timer
tcp6       0      0 :::8080                 :::*                    LISTEN 1/helloworld         off (0.00/0/0)







