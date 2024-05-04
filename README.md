## Word of Wisdom TCP Server
To protect the server from DDOS attacks, we use the Proof of Work algorithm.
In this example, we use a simple challenge-response protocol where the client sends a challenge (a random 32-byte string) to the server.

The server then performs a Proof of Work calculation on the challenge to verify its validity.
The difficulty of the Proof of Work is set to 10, which means the client must calculate a large number of hashes to find a valid solution.

How to install dependencies?:
```bash
make deps
```

How to run?:
```bash
make run
```