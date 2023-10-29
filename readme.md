# Running the program
1. Start the server by running 'go run .\Server.go' in the server folder
2. The server does not need / support any further user interaction
3. Start one or more clients, by running 'go run  .\Client.go -id [CLIENT_NAME]' in the Client folder
4. After starting a client in the terminal, messages can be sent by entering a message text followed by enter. Spaces are not supported, so any space between words will be read as the start of a new message.
5. To remove a client from the chat session, enter the message 'Disconnect'