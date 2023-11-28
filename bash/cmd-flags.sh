#!/bin/bash

while getopts u:a:f: flag; do
    case "${flag}" in
    u) username=${OPTARG} ;;
    a) age=${OPTARG} ;;
    f) fullname=${OPTARG} ;;
    esac
done
echo "Username: $username"
echo "Age: $age"
echo "Full Name: $fullname"

# > ./cmd-flags.sh -u user1 -a 30 -f "User One"
# Username: user1
# Age: 30
# Full Name: User One
