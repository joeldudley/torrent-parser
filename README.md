# Torrent Parser

A Go tool for parsing .torrent files (in [Bencode](https://en.wikipedia.org/wiki/Bencode) encoding) to JSON.

## Usage

Run the binary, passing in the path to a .torrent file:

```bash
go build
./torrent-parser ./data/example_torrents/big-buck-bunny.torrent
```

## Sample output

```
{
    "announce": "udp://tracker.leechers-paradise.org:6969", 
    "announce-list": [
        [
            "udp://tracker.leechers-paradise.org:6969"
        ], 
        [
            "udp://tracker.coppersurfer.tk:6969"
        ], 
        [
            "udp://tracker.opentrackr.org:1337"
        ], 
        [
            "udp://explodie.org:6969"
        ], 
        [
            "udp://tracker.empire-js.us:1337"
        ], 
        [
            "wss://tracker.btorrent.xyz"
        ], 
        [
            "wss://tracker.openwebtorrent.com"
        ], 
        [
            "wss://tracker.fastcast.nz"
        ]
    ], 
    "comment": "WebTorrent <https://webtorrent.io>", 
    "created by": "WebTorrent <https://webtorrent.io>", 
    "creation date": 1490916601, 
    "encoding": "UTF-8", 
    "info": {
        "files": [
            {
                "length": 140, 
                "path": [
                    "Big Buck Bunny.en.srt"
                ]
            }, 
            {
                "length": 276134947, 
                "path": [
                    "Big Buck Bunny.mp4"
                ]
            }, 
            {
                "length": 310380, 
                "path": [
                    "poster.jpg"
                ]
            }
        ], 
        "name": "Big Buck Bunny", 
        "piece length": 262144, 
        "pieces": "<SHA1 hashes>"
    }, 
    "url-list": [
        "https://webtorrent.io/torrents/"
    ]
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
