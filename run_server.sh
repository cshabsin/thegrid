#!/bin/bash

set -e

DATA_DIR=$(dirname $(realpath "$0"))/web_zips

./server/server --data_dir=$DATA_DIR
