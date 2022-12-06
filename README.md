# Garra

`Garra` is a `Go` best practices `checker` based on go vet.

## Requirements

-`Go 1.10` and above.

## Features
- `nilcheck` not nil err but object used checker
  - Example:
    ``` go
      user, err := svc.User().GetUserCache(uid)
      if nil != err{
        printf("err occurs)
      }
      user.GetUid() // <-- fail! haven't check user
    ```

- `sawago` go routine specification in Ziipin Sawa.
  - Example:
    - SUGGEST:
    ``` go
      func (dao *UserService)testFunc(){
        asynHandle(func(svc *Svc) {
          //XXXX
        })
      }
    ```

    - NOT SUGGEST:
    ``` go
      func (dao *UserService)testFunc(){
        go func() {
          //XXX
        }
      }
    ```

## Installation

- go install -a github.com/carsonfeng/garra@latest

## Usage

- go vet -vettool=$(which garra) ./...

## Contributing

The project welcomes all contributors. We appreciate your help!

## Communication:

- jiayaf@gmail.com
