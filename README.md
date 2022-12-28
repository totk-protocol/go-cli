# TOTK cli

## Build

```bash
# clone this repo
cd totkcli
go build -o totk ./cmd
```

## Usage

Generate a keypair

```
totk -g
```

List all keys

```
totk -l
```

Delete a key

```
totk -d [key id]
```

Calc totk code

```
totk -k [key id] [public key]

# Or use default key
totk [public key]
```
