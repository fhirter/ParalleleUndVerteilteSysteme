#!/usr/bin/env bash

sleep 10

mosquitto_pub -h broker -t 'temp' -m "test"
