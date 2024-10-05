# About
Subwrangler is a command-line utility that will rename subtitles to match video file names within the directory in which it is invoked. This allows players such as `mpv` to automatically detect subtitles, which can be useful when playing media with separate subtitles.

# Installation
```
go install github.com/samcgardner/subwrangler
```

You can also copy the provided binary from releases into your path manually, if you do not have golang installed and don't want to install it. At the moment, I only provide a Linux binary.   

# Use
```
subwrangler /path/to/directory
```

Subwrangler will operate on the current directory if one is not provided.
