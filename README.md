# Coreum Blockchain Backend in Golang

This repository contains the backend implementation in Golang for the Coreum blockchain. It includes REST API endpoints to retrieve balance, create and recover wallets.

## API Endpoints
### Get Balance of Account
Returns the balance of the specified account in the Coreum blockchain.

#### URL: /get-balance/{address}

#### Method: GET

#### URL Parameters:

address - The address of the account to retrieve the balance.

#### Success Response:

Code: 200 OK
Content:

```
{
    "type": "success",
    "data": "123.456",
    "message": "The Balance of the wallet is gotten successfully!"
}
```

#### Error Response:

Code: 404 Not Found

Content:

```
{
    "type": "error",
    "message": "You are missing wallet address parameter."
}
```

### Create New Wallet
Creates a new wallet in the Coreum blockchain.

#### URL: /create-new-wallet

#### Method: GET

#### Success Response:

Code: 200 OK
Content:

```
{
    "mnemonic": "example mnemonic",
    "address": "coreum1s2w0t4s4tgr8ugvpyz0ls0ufx9rvp7ns0l8grf"
}
```

### Recover Wallet from Mnemonic
Recovers a wallet in the Coreum blockchain from a mnemonic phrase.

#### URL: /recovery-wallet

#### Method: POST

#### Data Parameters:

mnemonic - The mnemonic phrase to recover the wallet.

#### Success Response:

Code: 200 OK

Content:

"coreum1s2w0t4s4tgr8ugvpyz0ls0ufx9rvp7ns0l8grf"


## Installation and Usage

Clone the repository:

```
git clone https://github.com/CoreumFoundation/coreum-backend-go.git
```

Install dependencies:

```
go
go mod tidy
```

Run the server:

```
go run main.go
```
The server will be listening at http://localhost:5432.

## Contributing
If you want to contribute to this project, please read the CONTRIBUTING.md file for more information.

## License
This project is licensed under the MIT License.
