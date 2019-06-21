# build from stock debian image
FROM debian:latest

# install needed packages
RUN apt-get update && apt-get install -y nfs-kernel-server && rm -rf /var/lib/apt/lists/*

# edit exports
RUN echo $NFSPATH' *(rw,sync,no_subtree_check)' >> /etc/exports
RUN service nfs-kernel-server start
