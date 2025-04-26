# cb

clipboard manager command

```bash
go install github.com/notwithering/cb@latest
```

## examples

copy file to clipboard

```bash
cat file.txt | cb
```

same as

```bash
xclip -sel clip
```

---

print contents of clipboard

```bash
cb o
```

same as

```bash
xclip -o -sel clip
```
