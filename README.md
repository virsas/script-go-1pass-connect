# script-go-1pass-connect

## usage

### For login items

```bash
go run main.go -vault="test_vault" -item="login_item" -label="username"
go run main.go -vault="test_vault" -item="login_item" -label="password"
```

### For server items

create custom text label with any value and then you can retrieve your values by

```bash
go run main.go -vault="test_vault" -item="server_item" -label="custom_label"
```

### To setup your 1password connect service, please have look at

https://developer.1password.com/docs/connect
