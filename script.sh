#!/bin/bash

# Find all files with extension .CR2
find . -name "*.CR2" -type f

# Find all files with extension .CR2 and delete
find . -name "*.CR2" -type f -delete
