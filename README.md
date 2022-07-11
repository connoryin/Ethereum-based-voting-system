# Ethereum based voting system

### Quick start
1. Install and start `Ganache`. Click `New Workspace`. In the `Chain` tab, set `Gas Limit` to be `50000000`. Click `Save Workspace` to start the network.
2. Click the key icon on the right-hand side, and copy the private key to the `privateKey` variable on the 26th line of the file `backend/contract/contract.go`.
3. Run the file `backend/deployEvents.go` by going into the `backend` folder and running the command `go run deployEvents.go`. Copy the address printed in the console to the `addressEvents` variable of the file `backend/contract/contract.go`.
4. Run the file `backend/main.go`by going into the `backend` folder and running the command `go run main.go`.
5. Run the front end by going to the `frontend` folder and run commands `npm install` and `npm start`.
6. Go to `localhost:3000` and enjoy our voting system.

### Project structure
1. The `frontend` folder holds everything related to frontend. 
2. The `contract` folder holds the smart contract files and the compilation script. The script compiles the `sol` files into `go` files and copy them into the `backend/poll` folder. 
3. The `backend` folder holds everything related to backend and blockchain. The `model` folder defines the data structures used by the APIs. The `handler` folder defines the handlers of each API. The `contract` folder provides APIs to interact with smart contracts. 