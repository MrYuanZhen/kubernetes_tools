#!/bin/sh

haproxy -f /usr/local/etc/haproxy/haproxy.cfg -p /run/haproxy.pid -D -sf

cat /etc/kubernetes/ssl/kube-node.pem /etc/kubernetes/ssl/kube-node-key.pem > /home/ssl-crt.pem

clusterGetnodeip -rise=30

#export ENDPOINTS=$( kubectl get node -l node-role.kubernetes.io/controlplane=true --no-headers -owide | awk '{print $6}' | awk '{{printf"%s,",$0}}' | sed s'/.$//' )
#confd -log-level=debug -onetime -backend env