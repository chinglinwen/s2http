# s2http

Convert socks proxy to a http(s) proxy

# Prerequisite (if use socks)

A socks proxy (probably a client) is listening

It may not directly connect to a socks server

## Usage

```
Usage of ./s2http:
  -port string
        port to listen (default "8080")
  -pure
        pure http/https proxy without socks
  -socks string
        socks url (default "127.0.0.1:1081")
  -v    verbose
  -version
        show version.
```

## Pure proxy

Start as a pure proxy ( without socks)

	./s2http -pure


## Credits

Inspired by [GopherFromHell](https://www.reddit.com/user/GopherFromHell),
As he provided [the example](https://play.golang.org/p/l0iLtkD1DV)

end.  
