#!/bin/sh
acme-inc server &
/usr/local/bin/envoy -c /etc/envoy/envoy.yaml
