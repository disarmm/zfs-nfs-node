# build from stock debian image
FROM debian:latest

# install needed packages
RUN apt-get update && apt-get install -y nfs-kernel-server && rm -rf /var/lib/apt/lists/*

ADD . /tmp

# edit exports
RUN bash -c "/tmp/config.sh" && exportfs -a
RUN service nfs-kernel-server start
