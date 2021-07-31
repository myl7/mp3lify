# mp3lify

Reverse proxy flac/wav to mp3 

## Why?

Sometimes, my network is not good enough to play lossless audios with streaming

## Usage

Run it, put the original URL in querystring `src`, and it serves you the MP3 file

Basic token auth is provided with header field `X-Auth-Token`

Whole file caching is provided as caching the original audio files in the specified dir

## Options & Config

Options is available in `main/main.go`, config is available in `config.go`, both with doc

## License

MIT
