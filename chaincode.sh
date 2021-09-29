printf "\n chain code install and instanciation in mfg start\n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg1.block.com/users/Admin@mfg1.block.com/msp" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in mfg end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg1.block.com/users/Admin@mfg1.block.com/msp" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in mfg end. \n"

sleep 4

printf "\n chain code install and instanciation in mfg2 start \n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg2.block.com/users/Admin@mfg2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.mfg2.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in mfg2 end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Mfg2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg2.block.com/users/Admin@mfg2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.mfg2.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in mfg2 end \n"


sleep 4

printf "\n chain code install and instanciation in shp1 start \n"
docker exec -e "CORE_PEER_LOCALMSPID=Shp1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp1.block.com/users/Admin@shp1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp1.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in shp1 end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Shp1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp1.block.com/users/Admin@shp1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp1.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in shp1 end \n"

sleep 4

printf "\n chain code install and instanciation in shp2 start \n"
docker exec -e "CORE_PEER_LOCALMSPID=Shp2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp2.block.com/users/Admin@shp2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp2.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in shp2 end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Shp2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp2.block.com/users/Admin@shp2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp2.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in shp2 end \n"

sleep 4

printf "\n chain code install and instanciation in rtl start \n"
docker exec -e "CORE_PEER_LOCALMSPID=Rtl1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rtl1.block.com/users/Admin@rtl1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.rtl1.block.com:7051" cli peer chaincode install -n blockcc -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
printf "\n chain code install in rtl end \n"
docker exec -e "CORE_PEER_LOCALMSPID=Rtl1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rtl1.block.com/users/Admin@rtl1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.rtl1.block.com:7051" cli peer chaincode instantiate -o orderer.block.com:7050 -C commonchannel -n blockcc -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Mfg1MSP.member','Mfg2MSP.member','Shp1MSP.member','Shp2MSP.member','Rtl1MSP.member')"
printf "\n chain code install and instanciation in rtl end \n"

