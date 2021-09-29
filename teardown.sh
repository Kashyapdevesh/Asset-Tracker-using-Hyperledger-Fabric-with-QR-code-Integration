#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -e
IMAGETAG="1.3.0"
THIRDPARTYIMAGETAG="0.4.13"
# Shut down the Docker containers for the system tests.
IMAGE_TAG=$IMAGETAG THIRDPARTYIMAGE_TAG=$THIRDPARTYIMAGETAG docker-compose -f docker-compose.yml kill 
IMAGE_TAG=$IMAGETAG THIRDPARTYIMAGE_TAG=$THIRDPARTYIMAGETAG docker-compose -f docker-compose.yml down

# remove the local state
rm -f ~/.hfc-key-store/*

# remove chaincode docker images
docker rm $(docker ps -aq)
docker rmi $(docker images dev-* -q)

# Your system is now clean
