#!/bin/bash
if [ "$1" == "" ]; then
	tr -dc A-Za-z0-9_ < /dev/urandom | head -c 15 | xargs
else
	tr -dc A-Za-z0-9_ < /dev/urandom | head -c "$1" | xargs
fi
