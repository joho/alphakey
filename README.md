# AlphaKey

For a [Codename](http://codename.io) project I was working on I needed human readable/shareable short keys, so I ended up reimplementing the solution found on [this stack overflow answer about "Bijective Functions"](http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener/742047) in Go and figured it might be of use to anyone else who needs to put together a url shortener or the like.

## Usage

To get going, just import this

```go
import "github.com/joho/alphakey"
```

You can either call a couple of the helper functions directly

```go
key := alphakey.GetKeyForInt(198898042) // returns "ninja"
num := alphakey.GetIntForKey("dgt")     // returns 1337
```

Or if you want a bit more control you can instantiate your own converter and control the alphabet used and any offset you might want

```go
converter := &alphakey.Converter{
  alphakey.UnambiguousLowercaseAlphabet, // all lower case, no l or o, because humans
  497,                                   // an arbitrary offset so early sharers don't get a or b
}

key := converter.GetKey(1988483) // returns "ninja"
num := converter.GetInt("emq")   // returns 1337
```

You can also check the [godocs for AlphaKey](http://godoc.org/github.com/joho/alphakey)

## Build Status

[![wercker status](https://app.wercker.com/status/c6de7a7f75baca3f7e1848f674c9b8be/m "wercker status")](https://app.wercker.com/project/bykey/c6de7a7f75baca3f7e1848f674c9b8be)

---

&copy; 2013 [John Barton](https://johnbarton.co/), MIT Licenced
