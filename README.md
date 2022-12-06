# Garra

`Garra` is a `Go` best practices `checker` based on go vet. 

## Requirements

-`Go 1.10` and above.

## Features
- 1. not nil err but object used checker
  - example:
    - user, err := svc.User().GetUserCache(uid)
    - if nil != err{
      - printf("err occurs)
    - }
    - user.GetUid() // <-- fail! haven't check user

## Installation

- go install -a github.com/carsonfeng/garra@latest

## Usage

- go vet -vettool=$(which garra) ./...

## Contributing

The project welcomes all contributors. We appreciate your help!

## Communication:

- jiayaf@gmail.com
