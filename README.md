# NSD
> Netlify Short Domain
<img src="docs/fish.svg" align="right">

A short-url generator written in Go and to be used with Netlify's [Redirect & Rewrite Rules](https://www.netlify.com/docs/redirects/).

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/itsmingjie/nsd)<a href="https://codeclimate.com/github/itsmingjie/nsd/maintainability">

## User Guide
### Installation
    go get github.com/itsmingjie/nsd

### Usage
    # create a new short url at "/foo" that redirects to the argument url
    ./nsd -link=foo -url=https://example.com

    # create a new short url at a random 6-character path (e.g.: "/sdopd")
    ./nsd -url=https://example.com

    # create a new short url but rewrites the original url
    ./nsd -url=https://example.com -status=200

    # remove the short url of "/foo"
    ./nsd -remove -link=foo

    # remove the short url of "/foo" only if the url of the url of "/foo" matches the argument url
    ./nsd -remove -link=foo -url=https://example.com

    # remove the short url with target of the argument url
    # if multiple instances are found, first instance will be removed
    ./nsd -remove -url=https://example.com

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

#### Meta

<img src="https://api.codeclimate.com/v1/badges/0da119b24f797695a9a2/maintainability" /></a>