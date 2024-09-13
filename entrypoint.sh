#!/bin/sh

if [ "${USE_DB}" = "true" ]; then
    ./myapp -d
else
    ./myapp
fi
