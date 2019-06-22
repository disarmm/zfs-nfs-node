#!/bin/bash
export NFSPATH
echo $NFSPATH' *(rw,sync,no_subtree_check)' >> /etc/exports
echo $HOSTNAME'	127.0.0.1' >> /etc/hosts
