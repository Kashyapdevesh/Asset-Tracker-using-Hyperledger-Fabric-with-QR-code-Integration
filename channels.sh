
#channel creating
docker exec cli peer channel create -o orderer.block.com:7050 -c commonchannel -f /etc/hyperledger/configtx/commonchannel.tx
#channel created

#MFG1 joining channel
docker exec cli peer channel join -b commonchannel.block
#MFG1 joined channel
#MFG2 joining channel
docker exec -e "CORE_PEER_LOCALMSPID=Mfg2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/mfg2.block.com/users/Admin@mfg2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.mfg2.block.com:7051" cli peer channel join -b commonchannel.block
#MFG2 joined channel
#SHP1 joining channel
docker exec -e "CORE_PEER_LOCALMSPID=Shp1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp1.block.com/users/Admin@shp1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp1.block.com:7051" cli peer channel join -b commonchannel.block
#SHP1 joined channel
#SHP2 joining channel
docker exec -e "CORE_PEER_LOCALMSPID=Shp2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shp2.block.com/users/Admin@shp2.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.shp2.block.com:7051" cli peer channel join -b commonchannel.block
#SHP2 joined channel
#RTL1 joining channel
docker exec -e "CORE_PEER_LOCALMSPID=Rtl1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rtl1.block.com/users/Admin@rtl1.block.com/msp" -e "CORE_PEER_ADDRESS=peer0.rtl1.block.com:7051" cli peer channel join -b commonchannel.block
#RTL1 joined channel

#all peers joined the commonchannel
