## Asset Tracker using Hyperledger Fabric 
Adani-Wilmar is a multi-million dollar conglomerate that offers most of the essential kitchen commodities for Indian consumers, 
including edible oil, wheat flour, rice, pulses, and sugar. It also offers a range of staples such as wheat flour, rice, pulses, and sugar, and these products are offered under a diverse range of brands across a broad price spectrum and cater to different customer groups.

The vast array of commodities in circulation make it essential for a robust and flexible supply chain to be in place for efficient transportation and transparency within the system.
There is an immediate need for change in organizational focus from period-driven planning to event-driven execution to make the system reliable and transparent to consumers over time. 
There is a huge problem of traceability in the Adani-Wilmar supply chain as indicated by their annual traceability report and we propose to solve that problem by using **state-of-the-art permissioned blockchain technology with QR code integration and possible extension with IOT devices.**

In the initial stage, we have developed a prototype network with dummy nodes to display its potential. 
We have also displayed the QR code integration facility by constructing a dummy blockchain network and mapping the information on a OR code image. We refrained from developing a frontend component in our system in the inital stage, as we thought it would be better to discuss with mentors if we should develop an independent full-stack solution or an easily integrable API. 
We have also mentioned future scope in our projects README.md file which we plan to finish we get selected for the onsite round.

![](https://github.com/Sigsev-Dev/AdaniWilmar-Asset-Tracker-using-Hyperledger-Fabric-with-QR-code-Integration/blob/master/images/Project_Logistics_Architecture.jpg)

Here assumption is that you have installed all [prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-1.3/prereqs.html) that a standard Hyperledger Fabric requires.

Get solution
git clone https://github.com/Kashyapdevesh/Asset-Tracker-using-Hyperledger-Fabric-with-QR-code-Integration.git
cd code

### BlockServer
this is web server based on nodejs and expressjs. facilitates following things.
* Acts as client for Fabric network by using Fabric-Client sdk
* Acts like a web server that wraps all the fabric interaction logic as REST API (GET, POST endpoints).
* Bridge between a user interface layer and fabric network layer.

## Starting the Fabric Network
Run following commands in BlockServer folder to start network, install dependencies, enroll Admin and register user.


cd BlockServer

npm install

./startFabric.sh



## Stop the network
To stop network after testing, run the given command

./stop.sh

## Kill the network
To kill the complete network, use the ./teardown.sh 

./teardown.sh

 
## Future Scope :
1) We plan to develop the complete solution(API or full-stack implementation) if we get selected for the onsite round.
2) We plan to add provision for scalable IOT devices in our network architecture
3) We plan to test and display our network on the stage by simulating our mobile devices as IOT devices to display the network's functionalities in real time.

