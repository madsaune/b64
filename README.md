# b64

`b64` is a small command line utility for encoding and decoding base64 strings.

## Installation

```bash
go install github.com/madsaune/b64@latest
```

## Usage

### Encoding

```bash
$ echo "my_string" | b64
bXlfc3RyaW5n

$ echo "my_string" > test.txt
$ b64 < test.txt
bXlfc3RyaW5n

$ b64 --str "my_string"
bXlfc3RyaW5n

$ b64
my_string<CR>
<CTRL+D>
bXlfc3RyaW5n
```

### Decoding

```bash
$ echo "bXlfc3RyaW5n" | b64 --decode
my_string

$ echo "bXlfc3RyaW5n" > test.txt
$ b64 --decode < test.txt
my_string

$ b64 --str "bXlfc3RyaW5n" --decode
my_string

$ b64
bXlfc3RyaW5n<CR>
<CTRL+D>
my_string
```
