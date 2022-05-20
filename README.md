# b64

`b64` is a small command line utility for encoding and decoding base64 strings.

## Install

```bash
go install github.com/madsaune/b64@latest
```

## Usage

```bash
Usage of b64:
  -d    decode string
  -s string
        the string to encode/decode
```

### Encoding

```bash
$ echo "my_string" | b64
bXlfc3RyaW5n

$ echo "my_string" > test.txt
$ b64 < test.txt
bXlfc3RyaW5n

$ b64 -s "my_string"
bXlfc3RyaW5n

$ b64
my_string<CR>
<CTRL+D>
bXlfc3RyaW5n
```

### Decoding

```bash
$ echo "bXlfc3RyaW5n" | b64 -d
my_string

$ echo "bXlfc3RyaW5n" > test.txt
$ b64 -d < test.txt
my_string

$ b64 -s "bXlfc3RyaW5n" -d
my_string

$ b64 -d
bXlfc3RyaW5n<CR>
<CTRL+D>
my_string
```
