# Torrent Parser

Torrent Parser is a Go tool for parsing .torrent files (encoded in [Bencode](https://en.wikipedia.org/wiki/Bencode)) 
to JSON.

## Installation

Build the binary using Go:

```bash
go build
```

## Usage

Run the binary, passing in the path to a .torrent file:

```bash
go build

./torrent-parser ./data/example_torrents/big-buck-bunny.torrent
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
