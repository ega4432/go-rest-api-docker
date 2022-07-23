#!/bin/bash -ue

if ! type act &> /dev/null ; then
    echo "act command not found."
    exit 1
fi

echo "ok!"
act 
