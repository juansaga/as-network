Diseño de una red transaccional con hyperledger fabric.

1 En primer lugar se debe crear el archivo crypto-config.yaml, en este se estable la cantidad inicial de organizaciones y es fundamental definir el orderer. En este archivo también se especifica la cantidad de usuarios y de peers de cada organización.

2 Es fundamental ejecutar el siguiente comando para habilitar las funciones de peer y cryptogen, la carpeta que se usa es el lugar donde se tiene hpl.

export PATH=$PATH:$HOME/hyperledger/fabric/fabric-samples/bin

3 El siguiente comando genera el material criptográfico necesario haciedo uso del comando cryptogen y del archivo creado al inicio. Si la consola no reconoce el comando se debe revisar que en la variable PATH se encuentre la carpeta donde tiene hpl.

cryptogen generate --config=./crypto-config.yaml

4 Cear el archivo configtx.yaml, este archivo se usa pa crear el primer bloque de la red donde está la configuración de la misma, allí están las configuraciones de las organizaciones, la versión de hpl, las politicas, el servicio de ordenamiento, el canal y los perfiles.

Se ejecuta el siguiente comando para la creación del bloque genesis.

configtxgen -profile ThreeOrgsOrdererGenesis -channelID system-channel -outputBlock ./channel-artifacts/genesis.block



configtxgen -profile ThreeOrgsChannel -channelID data -outputCreateChannelTx ./channel-artifacts/channel.tx

configtxgen -profile ThreeOrgsChannel -channelID data -outputAnchorPeersUpdate ./channel-artifacts/Hot1MSPanchors.tx -asOrg Hot1MSP

configtxgen -profile ThreeOrgsChannel -channelID data -outputAnchorPeersUpdate ./channel-artifacts/Hot2MSPanchors.tx -asOrg Hot2MSP

configtxgen -profile ThreeOrgsChannel -channelID data -outputAnchorPeersUpdate ./channel-artifacts/Hot3MSPanchors.tx -asOrg Hot3MSP


export CHANNEL_NAME=data

export VERBOSE=false

export FABRIC_CFG_PATH=$PWD

CHANNEL_NAME=$CHANNEL_NAME docker-compose -f docker-compose-cli-couchdb.yaml up -d



CAMBIAR A LA TERMINAL DEL CONTENEDOR CLI

export CHANNEL_NAME=data
export CHAINCODE_NAME=data
export CHAINCODE_VERSION=1
export CC_RUNTIME_LANGUAGE=golang
export CC_SRC_PATH="../../../chaincode/$CHAINCODE_NAME/"
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem



peer channel create -o orderer.as.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem

peer channel join -b data.block



CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot2.as.com/users/Admin@hot2.as.com/msp CORE_PEER_ADDRESS=peer0.hot2.as.com:7051 CORE_PEER_LOCALMSPID="Hot2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot2.as.com/peers/peer0.hot2.as.com/tls/ca.crt peer channel join -b data.block

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/users/Admin@hot3.as.com/msp CORE_PEER_ADDRESS=peer0.hot3.as.com:7051 CORE_PEER_LOCALMSPID="Hot3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/peers/peer0.hot3.as.com/tls/ca.crt peer channel join -b data.block





peer channel update -o orderer.as.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Hot1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem






CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot2.as.com/users/Admin@hot2.as.com/msp CORE_PEER_ADDRESS=peer0.hot2.as.com:7051 CORE_PEER_LOCALMSPID="Hot2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot2.as.com/peers/peer0.hot2.as.com/tls/ca.crt peer channel update -o orderer.as.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Hot2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem


CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/users/Admin@hot3.as.com/msp CORE_PEER_ADDRESS=peer0.hot3.as.com:7051 CORE_PEER_LOCALMSPID="Hot3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/peers/peer0.hot3.as.com/tls/ca.crt peer channel update -o orderer.as.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Hot3MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem


ver certificado:
cat /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem


export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/as.com/orderers/orderer.as.com/msp/tlscacerts/tlsca.as.com-cert.pem


empaquetado del chaincode
peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label {CHAINCODE_NAME}_${CHAINCODE_VERSION} >&log.tx



instalacion de chaincode

peer lifecycle chaincode install data.tar.gz

guardar el identificador del chaincode, para el ejemplo es:
foodcontrol_1:b472535bc08926703c5814c5e9e59051b8b64e7822a3fbeca5b0fc0aab256af5

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerHotanizations/hot2.as.com/users/Admin@hot2.as.com/msp CORE_PEER_ADDRESS=peer0.hot2.as.com:7051 CORE_PEER_LOCALMSPID="Hot2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot2.as.com/peers/peer0.hot2.as.com/tls/ca.crt peer lifecycle chaincode install foodcontrol.tar.gz

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/users/Admin@hot3.as.com/msp CORE_PEER_ADDRESS=peer0.hot3.as.com:7051 CORE_PEER_LOCALMSPID="Hot3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/peers/peer0.hot3.as.com/tls/ca.crt peer lifecycle chaincode install foodcontrol.tar.gz




peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Hot1MSP.peer','Hot3MSP.peer')" --package-id foodcontrol_1:b472535bc08926703c5814c5e9e59051b8b64e7822a3fbeca5b0fc0aab256af5


peer lifecycle chaincode commit -o orderer.as.com:7050 --tls --cafile $ORDERER_CA --peerAddresses peer0.hot1.as.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot1.as.com/peers/peer0.hot1.as.com/tls/ca.crt --peerAddresses peer0.hot3.as.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hot3.as.com/peers/peer0.hot3.as.com/tls/ca.crt --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --signature-policy "OR ('Hot1MSP.peer','Hot3MSP.peer')"

