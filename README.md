# NSD
> Netlify Short Domain
<img src="docs/fish.svg" width="130" alt="Postman Icon" align="right">

A short-url generator written in Go and to be used with Netlify's [Redirect & Rewrite Rules](https://www.netlify.com/docs/redirects/). WIP.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/itsmingjie/nsd)

### Installation
    $ go get github.com/itsmingjie/nsd

### Usage
```
$ ./nsd -op=add -link="custom" -url="https://example.com" -status=301
```

## License

[tl;dr](https://tldrlegal.com/license/mit-license)

The MIT License (MIT)

Copyright (c) 2019 Mingjie Jiang

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.