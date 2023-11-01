# Sugar for handling go errors

## Install
`go get -t github.com/dacalin/custom_error`

## Examples

### Create an error
`	
myErr := cerror.New("MyDbError", "This is my custom error")
`

### Create an error


```
err:= anerror()

if err != nil {
    myErr = myErr.AddOrigError(err).AddExtraMsg("Error when trying this.")
}

```

### Print the error
```
fmt.Println(myErr)

```

### Check error type
```
if myErr.Code() == other.Code() {
		
}
```

	


