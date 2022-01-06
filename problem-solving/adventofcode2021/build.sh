#!/bin/bash

# build
mkdir -p build && cd build
rm -rf *
cmake -DCMAKE_BUILD_TYPE=Debug .. && make 

# run
# ./solve