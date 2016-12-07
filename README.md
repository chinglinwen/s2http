# s2http

Convert socks proxy to a http(s) proxy

# Prerequisite

A socks proxy (probably a client) is listening

It may not directly connect to a socks server

# Usage

```
Usage of ./s2http:
  -port string
        port to listen (default "8080")
  -socks string
        socks url (default "127.0.0.1:1081")
  -v    verbose
  -version
        show version.
```

# Credits

Inspired by [GopherFromHell](https://www.reddit.com/user/GopherFromHell),
As he provided [the example](https://play.golang.org/p/l0iLtkD1DV)

end.  
