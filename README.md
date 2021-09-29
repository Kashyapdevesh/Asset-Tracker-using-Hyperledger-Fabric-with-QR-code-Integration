## Hyperledger Fabric - Logistics

Logistics management is that part of supply chain management that plans, implements, and controls the efficient, effective forward and reverses flow and storage of goods, services and related information between the point of origin and the point of consumption in order to meet customers' requirements.

A fully integrated system across the entire supply chain, from the moment a shipment leaves the production facility, to the final delivery on the customer's doorstep. Every shipment is federated & validated in trustless, transparent blockchain contracts. Knowing where a shipment is located and its time to arrival are key factors in keeping the B2B and B2C operations alive for companies that transport goods around the globe. The application of blockchain for logistics service providers extend to those servicing any type of supply chain.


Hyperledger Fabric based solution for consortium of logistics companies in any distribution or supply-chain industries

Here assumption is that you have installed all [prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-1.3/prereqs.html) that a standard Hyperledger Fabric requires.

Get solution
```
git clone https://github.com/dcentrum/Project-HLF-Logistics.git
cd code
```

We have three folders for three different layers

### BlockFabric
this folder contains all Fabric related cryptographic material, channel, genesis block and chaincode.

### BlockServer
this is web server based on nodejs and expressjs. facilitates following things.
* Acts as client for Fabric network by using **Fabric-Client** sdk
* Acts like a web server that wraps all the fabric interaction logic as REST API (GET, POST endpoints).
* Bridge between a user interface layer and fabric network layer.

 ### BlockClient
This is end user interface developed using Angular 7.
This consumes REST api to contact with **BlockServer** web server.

## Start Fabric
Run following commands in **BlockServer** folder to start network, install dependencies, enroll Admin and register user.

```

cd BlockServer

npm install

cd BlockFabric

./startFabric.sh

```

<!-- ## Start Web server
Run following commands in **BlockServer** folder to start the web server

```
node app.js
```

## Set up User interface
Open another terminal and run following commands in **BlockClient** folder to start the user interface application.

```
cd BlockClient

npm install -g @angular/cli

npm install

ng serve

```

After completion of above commands. open any web browser and follow this link http://localhost:4200


## Stop the network
To stop network after testing, run the given commands in **BlockFabric** folder
```
cd BlockFabric

./stop.sh
```

## Kill the network
To kill the complete network, use the **./teardown.sh** in **BlockFabric** folder

```
cd BlockFabric

./teardown.sh
``` -->
Phase 1 : [ Workflow Diagram](https://github.com/dcentrum/Project-HLF-Logistics/blob/master/docs/DCentrum_Blockchain_Block_Sequence.png)

## TEAM
| Roles                       | Names                                     |
| --------------------------- | ----------------------------------------- |
| Author , Product Owner      | Mohan Krishna G                           |
| Product Manager             | Deepak B                                  |
| Requirement Analysis        | Mohan Krishna, Tameem A                   |
| Solution Design             | Sreenivas C                               |
| UI Development              | Sesha Rao N(web), Srikanth V (Mobile)     |
| Hyperledger Orchestration   | Tameem A,Sesha Rao N, Sreenivas C         |
| Smart Contracts Dev         | Sreenivas C, Mohan Krishna G;Sesha Rao N  |
| Testing                     | Mohan Krishna G, Tameem A                 |
| UI Development              | Sesha Rao N(web), Srikanth V (Mobile)     |
| Integration                 | Sreenivas C                               |
| Deployment                  | Sreenivas C                               |



Thank you.
# Asset-Tracker-using-Hyperledger-Fabric-with-QR-code-Integration
# Kashyapdevesh-Asset-Tracker-using-Hyperledger-Fabric-with-QR-code-Integration
