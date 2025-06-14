# Setup
- ## clone git repo
    - git clone git@github.com:babbu325/Crypto_Transaction_Fetcher.git
    - cd Crypto_Transaction_Fetcher

- ## Install dependencies
    - go mod tidy

- ## create a .env file and populate it with below values
    - EXPLORER=ETHERSCAN
    - ETHERSCAN_API_KEY=YourApiKeyHere

- ## Run the script
    - go run cmd/main.go -address yourAddress -output yourOutputFile.csv
        
        ``It will create a yourOutputFile.csv at root level with desired data``

# Assumptions Made
- If any of transaction type api call fails, will return error
  assuming, Either will provide the correct set of transactions else none.
- Provided address is a valid Ether address.


# Architecture design
- Used singleton design pattern to inject config.
- Used factory design pattern to instantiate explore.

# Testing
- Test if config is being loaded properly.
- Test if any of the apis are failing, we should fail entire request.
- If any key in response mapper change.
- Test if all data are correct and valid. eg. Amount is in unit of eth and not in wei

# TODOs
- Implement logger
- Pagination for large addtess
