var refDataHelper = require('./refDataHelper');
var Fabric_Client = require('fabric-client');

var fabric_client = new Fabric_Client();

var channel = fabric_client.newChannel('commonchannel');

var mfg1peer = fabric_client.newPeer('grpc://localhost:5051');
channel.addPeer(mfg1peer);
var mfg2peer = fabric_client.newPeer('grpc://localhost:6051');
channel.addPeer(mfg2peer);
var shp1peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(shp1peer);
var shp2peer = fabric_client.newPeer('grpc://localhost:8051');
channel.addPeer(shp2peer);
var rtl1peer = fabric_client.newPeer('grpc://localhost:9051');
channel.addPeer(rtl1peer);

var order = fabric_client.newOrderer('grpc://localhost:7050')
channel.addOrderer(order);

var path = require('path');

var store_path = path.join(__dirname, 'hfc-key-store');

module.exports = {

    setUser: async (fabricclient, username) => {
        return new Promise((resolve, reject) => {
            Fabric_Client.newDefaultKeyValueStore({ path: store_path })
                .then((state_store) => {
                    fabricclient.setStateStore(state_store);
                    var crypto_suite = Fabric_Client.newCryptoSuite();
                    var crypto_store = Fabric_Client.newCryptoKeyStore({ path: store_path });
                    crypto_suite.setCryptoKeyStore(crypto_store);
                    fabricclient.setCryptoSuite(crypto_suite);
                    fabricclient.getUserContext(username, true).then((user_from_store) => {
                        if (user_from_store && user_from_store.isEnrolled()) {
                            resolve(true);
                        } else {
                            reject(false);
                        }
                    });
                });
        });
    },

    attachFabricClient: (req) => {
        req.fabricClient = fabric_client;
        req.channel = channel;
        req.Fabric_Client = Fabric_Client;
        req.org = refDataHelper.getOrgByUser(req.headers.username)
        if (req.org == "mfg1")
            req.peer = mfg1peer;
        else if (req.org == "mfg2")
            req.peer = mfg2peer;
        else if (req.org == "shp1")
            req.peer = shp1peer;
        else if (req.org == "shp2")
            req.peer = shp2peer;
        else if (req.org == "rtl1")
            req.peer = rtl1peer;


    }


}