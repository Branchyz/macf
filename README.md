# MACF
A Windows-only Go CLI Utility for discovering the IP address associated with a specified MAC address. 
The tool performs the following steps:
1. Finds the Local Gateway IP: Utilizes the route command to identify the gateway IP address.
2. Scans the Local Network: Executes an nmap scan on the local subnet.
3. Parses Scan Results: Extracts the IP address corresponding to the given MAC address from the nmap output.

## Prerequisites
- Go: Ensure Go is installed on your system to build the application. Download it from golang.org.
- nmap: This tool relies on nmap for network scanning. Install it from nmap.org.

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/branchyz/macf.git
   ```
2. Build the application:

   ```bash
   go build -o macf.exe
   ```
3. (Optional) Add the binary to your PATH:

    To make the macf command accessible from any directory, move it to a directory that is in your system's PATH, or add the current directory to your PATH. For example:
   
   ```bash
   set PATH=%PATH%;C:\path\to\your\bin
   ```

## Usage
Run the application with the MAC address as the command-line argument:

```bash
macf <MAC_ADDRESS>
```

Replace <MAC_ADDRESS> with the MAC address you want to resolve (e.g., 00:1A:2B:3C:4D:5E).

The application will output the associated IP address or an error if it can't be found.

## Error Handling
- No MAC Address Provided: Ensure you provide a MAC address as a command-line argument.
- Invalid MAC Address: The provided MAC address format is incorrect.
- No Gateway Found: The default gateway could not be determined.
- Nmap Scan Error: The nmap scan failed.
- MAC Address Not Found: The specified MAC address was not found in the nmap scan results.

