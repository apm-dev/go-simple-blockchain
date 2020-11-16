# Go Simple Blockchain

### Test steps:

- **Mine the genesis block**

request:
```text
curl http://localhost:8000/mine
```
result:
```json
{"message": "new block mined"}
```

- **See the Chain**

request:
```text
curl http://localhost:8000/chain
```
result:
```json
[
  {
    "Index": 1,
    "Timestamp": 1605522128,
    "Trxs": [
      {
        "Sender": "",
        "Recipient": "apm-wallet",
        "Amount": 12.5
      }
    ],
    "Nonce": 55252,
    "PreviousHash": ""
  }
]
```


- **Add new Trx**

request:
```text
curl -X POST -H 'Content-Type:application/json' -d '{"sender":"a","recipient":"b","amount":"15.4"}'
```
result:
```json
{"message":"your trx added to mempool and will get done soon"}
```

- **Mine another block and see the Chain again**
request:
```text
curl http://localhost:8000/mine
curl http://localhost:8000/chain
```
result:
```json
[
  {
    "Index": 1,
    "Timestamp": 1605522128,
    "Trxs": [
      {
        "Sender": "",
        "Recipient": "apm-wallet",
        "Amount": 12.5
      }
    ],
    "Nonce": 55252,
    "PreviousHash": ""
  },
  {
    "Index": 2,
    "Timestamp": 1605522678,
    "Trxs": [
      {
        "Sender": "a",
        "Recipient": "b",
        "Amount": 15.4
      },
      {
        "Sender": "",
        "Recipient": "apm-wallet",
        "Amount": 12.5
      }
    ],
    "Nonce": 15776,
    "PreviousHash": "51d850ca8f0ab508656e4b4c42b1b3bae4941a6388e405cea6449e1ae4980000"
  }
]
```