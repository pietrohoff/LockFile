
# LockFile

LockFile is a simple command-line tool for encrypting and decrypting files using AES-256 encryption. It allows users to secure files with a password and store them safely.

## Project Structure

```plaintext
LockFile/
│
├── cmd/
│   └── main.go          # Main file to initialize the application
│
├── pkg/
│   └── encryption/
│       ├── encrypt.go   # File encryption functionality
│       ├── decrypt.go   # File decryption functionality
│       └── utils.go     # Utilities for file handling and key generation
│
├── files/
│   ├── encrypted/       # Directory for encrypted files
│   └── decrypted/       # Directory for decrypted files
│
├── go.mod               # Go module file
├── go.sum               # Dependencies summary file
└── README.md            # Project documentation
```

## Installation

1. Clone the repository:
    ```bash
    git clone <repository_url>
    cd LockFile
    ```

2. Initialize dependencies:
    ```bash
    go mod tidy
    ```

## Usage

To encrypt a file:
```bash
go run cmd/main.go <path_to_file> <password>
```

The encrypted file will be saved in the `files/encrypted/` directory.

To decrypt a file:
```bash
go run cmd/main.go <path_to_encrypted_file> <password> --decrypt
```

The decrypted file will be saved in the `files/decrypted/` directory.

## License

This project is licensed under the MIT License.
