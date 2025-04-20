# First Commando Line

### Commands to use - CLI with Cobra in Go

- hello command
```bash
    go run main.go hello <your_name>
```
- Use --uppercase flag
```bash
    go run main.go hello <your_name> --uppercase
or
    go run main.go hello --uppercase <your_name>
```

- Execute subcommand greet
```bash
    go run main.go hello greet <your_name>
```