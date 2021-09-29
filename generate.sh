#!/bin/sh
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/./bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME=commonchannel

# remove previous crypto material and config transactions
rm -fr config/*.
rm -fr crypto-config/*

# generate crypto material
cryptogen generate --config=./crypto-config.yaml
if [ "$?" -ne 0 ]; then
  echo "Failed to generate crypto material..."
  exit 1
fi

# generate genesis block for orderer
configtxgen -profile BlockOrdererGenesis -outputBlock ./config/commongenesis.block
if [ "$?" -ne 0 ]; then
  echo "Failed to generate orderer genesis block..."
  exit 1
fi

# generate channel configuration transaction
configtxgen -profile BlockChannel -outputCreateChannelTx ./config/commonchannel.tx -channelID $CHANNEL_NAME
if [ "$?" -ne 0 ]; then
  echo "Failed to generate channel configuration transaction..."
  exit 1
fi

# generate anchor peer transaction Mfg1MSP
configtxgen -profile BlockChannel -outputAnchorPeersUpdate ./config/Mfg1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Mfg1MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Mfg1MSP..."
  exit 1
fi
# generate anchor peer transaction Mfg1MSP
configtxgen -profile BlockChannel -outputAnchorPeersUpdate ./config/Mfg2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Mfg2MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Mfg2MSP..."
  exit 1
fi

# generate anchor peer transaction Shp1MSP
configtxgen -profile BlockChannel -outputAnchorPeersUpdate ./config/Shp1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Shp1MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Shp1MSP..."
  exit 1
fi
# generate anchor peer transaction Shp2MSP
configtxgen -profile BlockChannel -outputAnchorPeersUpdate ./config/Shp2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Shp2MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Shp2MSP..."
  exit 1
fi

# generate anchor peer transaction Rtl1MSP
configtxgen -profile BlockChannel -outputAnchorPeersUpdate ./config/Rtl1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Rtl1MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Rtl1MSP..."
  exit 1
fi