# Matching Prefixes

## Overview

This is a simple prefix matcher built in golang, the service performs a prefix matches if the input string starts with that prefix..

## Features

- Does a preix matching using trie DS.
- Returns the longest prefix as response.
- Utilize In-memory caching for faster performance.

## Prerequisites

Before running the URL shortener service, make sure you have the following prerequisites installed:

- Go: The Go programming language.

## Getting Started

Follow these steps to set up and run the URL shortener service:

1. **Clone the Repository:**

```
   git clone https://github.com/yuvrajsingh79/matching-prefixes.git
   cd matching-prefixes
```

2. **Configuration:**

     If you plan to use any prefix file for testing, update the filepath in the server.go :

```
   filePath := filepath.Join(cwd, "../prefixes.txt")
```


3. **Build and Run the Application:**

	To start the service, run the below commands.
```
   cd cmd
   go run main.go
```

4. **Perform Testing and Verify the routes**
   
   **Prefix Matching** :
	
   Use the service to call the match prefix api using the below GET request URL:
	
```
  http://localhost:8080/prefix-match/LgSicT
```
   Response will look like :
    
```
 LgSicT
```

5. **Testing:**

Run the tests to ensure the correctness of the service:
```
go test ./...
```

## Contributing

Contributions to this project are welcome. Feel free to open issues or pull requests for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
