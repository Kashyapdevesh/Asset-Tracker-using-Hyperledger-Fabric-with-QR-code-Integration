#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e
IMAGETAG="1.3.0"
THIRDPARTYIMAGETAG="0.4.13"
# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
LANGUAGE=${1:-"golang"}
CC_SRC_PATH=github.com/block/go
if [ "$LANGUAGE" = "node" -o "$LANGUAGE" = "NODE" ]; then
	CC_SRC_PATH=/opt/gopath/src/github.com/block/node
fi

# clean the keystore
rm -rf ./hfc-key-store

# launch network; create channel and join peer to channel
#cd ../hlf
IMAGE_TAG=$IMAGETAG THIRDPARTYIMAGE_TAG=$THIRDPARTYIMAGETAG  ./start.sh

# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars
IMAGE_TAG=$IMAGETAG THIRDPARTYIMAGE_TAG=$THIRDPARTYIMAGETAG docker-compose -f ./docker-compose.yml up -d cli

printf "\n chain code install and instanciation in mfg1 start\n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg1.block.com/users/Admin@mfg1.block.com/msp" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in mfg1 end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg1.block.com/users/Admin@mfg1.block.com/msp" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in mfg1 end. \n"

sleep 6

# printf "\n chain code install and instanciation in mfg2 start \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Mfg2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg2.block.com/users/Admin@mfg2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.mfg2.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
# printf "\n chain code install in mfg2 end \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Mfg2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg2.block.com/users/Admin@mfg2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.mfg2.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
# printf "\n chain code install and instanciation in mfg2 end \n"


# sleep 6

# printf "\n chain code install and instanciation in shp1 start \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Shp1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp1.block.com/users/Admin@shp1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp1.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
# printf "\n chain code install in shp1 end \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Shp1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp1.block.com/users/Admin@shp1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp1.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
# printf "\n chain code install and instanciation in shp1 end \n"

# sleep 6

# printf "\n chain code install and instanciation in shp2 start \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Shp2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp2.block.com/users/Admin@shp2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp2.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
# printf "\n chain code install in shp2 end \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Shp2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp2.block.com/users/Admin@shp2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp2.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
# printf "\n chain code install and instanciation in shp2 end \n"

# sleep 6

# printf "\n chain code install and instanciation in rtl start \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Rtl1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rtl1.block.com/users/Admin@rtl1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.rtl1.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
# printf "\n chain code install in rtl end \n"
# docker exec -e "CORE_PEER_LOCALMSPID=Rtl1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rtl1.block.com/users/Admin@rtl1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.rtl1.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
# printf "\n chain code install and instanciation in rtl end \n"

printf "\n Admins enrollment start \n"

node ./../BlockServer/enrollAdmin.js mfg1
# node ./../BlockServer/enrollAdmin.js mfg2
# node ./../BlockServer/enrollAdmin.js shp1
# node ./../BlockServer/enrollAdmin.js shp2
# node ./../BlockServer/enrollAdmin.js rtl1

printf "\n Admins enrollment end \n"


printf "\n User registrations  start \n"

node ./../BlockServer/registerUser mfg1
#  node ./../BlockServer/registerUser mfg2
#  node ./../BlockServer/registerUser shp1
#  node ./../BlockServer/registerUser shp2
#  node ./../BlockServer/registerUser rtl1
#  node ./../BlockServer/registerUser shp1driver
#  node ./../BlockServer/registerUser shp2driver

printf "\n User registrations end \n"


printf 'Block Fabric network ready to connect'
printf 'One admin and one user created for each org. one driver user created for each shipper company'



#peer chaincode invoke -o orderer.block.com:7050 -C commonchannel -n blockcc -c '{"function":"addShipmentDRS","Args":["97","aaa","vvv","aaa","1","1"]}'

#peer chaincode invoke -o orderer.block.com:7050 -C commonchannel -n blockcc -c '{"function":"shipmentPickup","Args":["1122","SreeSign","2019-01-02","DeeSign","2019-02-02","17.0000","58.0000","pickedup","1"]}'


#peer chaincode invoke -o orderer.block.com:7050 -C commonchannel -n blockcc -c '{"function":"shipmentPickup","Args":["857","SreeSign","2019-01-02","DeeSign","2019-02-02","17.0000","58.0000","Delivered","1"]}'




