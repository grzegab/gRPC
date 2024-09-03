# gRPC example app

Simple app to check out how to create and make gRPC calls.

## Proto compiler options

1. **--decode_raw Option**
    Can decode message from binary storage message:
    `cat protoc/simple.bin | protoc --decode_raw`
    Will return tags with values
2. **--decode Option**
    Can decode message but needs a message proto file:
    `cat simple.bin | protoc --decode=packageName.Simple simple.proto`
    Will return names with values
3. **--encode Option**
    Can encode to binary message:
    `cat simple.txt | protoc --encode=packageName.Simple simple.proto > simple.bin`

### Initial work
Greg, Lord of Mailgun Messages, Master of Redis Realms, Messenger of GO Secrets, Champion of Domain-Driven Development