FROM ubuntu:14.04

ADD testserver /bin/

ENTRYPOINT ["/bin/testserver"]
