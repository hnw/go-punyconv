punyconv(1) -- Convert unicode domain names to/from punycode names
=================

## SYNOPSIS

`punyconv [filename]`

## EXAMPLE

```
$ echo "3年B組金八先生.jp" | punyconv
xn--3b-ww4c5e180e575a65lsy2b.jp
$ echo "xn--3b-ww4c5e180e575a65lsy2b.jp" | punyconv
3年b組金八先生.jp
```

## References

* [miekg/dns: DNS library in Go](https://github.com/miekg/dns)

## LICENSE

The MIT License

Copyright (c) 2016 Yoshio HANAWA

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
