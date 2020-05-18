# My Go Security Projects

<!-- MarkdownTOC -->

- [My Go Security Projects](#my-go-security-projects)
    - [SSH Harvester](#ssh-harvester)
    - [golnk](#golnk)
    - [pcap Tutorial](#pcap-tutorial)
    - [Go-Fuzz Adventures](#go-fuzz-adventures)
        - [iprange](#iprange)
        - [goexif2](#goexif2)
    - [CLI Package Tests](#cli-package-tests)
        - [urfave/cli](#urfavecli)
        - [spf13/cobra](#spf13cobra)
    - [Base64 Stream Decoder](#base64-stream-decoder)
    - [PNG Tests](#png-tests)
    - [Go Slices And Their Oddities](#go-slices-and-their-oddities)

<!-- /MarkdownTOC -->

## [SSH Harvester](ssh-harvester)
Initial version of tool written in Go that harvests SSH certificates. For
explanation of code please see the blog post [Simple SSH Harvester in
Go][go-sshharvester].

## [golnk](golnk)
Go package for parsing Windows shell link binary (lnk or Windows shortcut) files.

* https://github.com/parsiya/golnk

## [pcap Tutorial](pcap-tutorial)
Code for the blog post [Go and pcaps][go-pcap] which explains how to use Go to
extract ICMP echo payloads from a pcap file.

## [Go-Fuzz Adventures](go-fuzz)
Learning how to use [Go-Fuzz][go-fuzz-github].

### [iprange](go-fuzz/iprange)
Test code for fuzzing the [iprange][iprange-github] package. See blog post
[Learning Go-Fuzz 1: iprange][iprange-blog].

### [goexif2](go-fuzz/goexif2)
Code and accompanying material for blog post
[Learning Go-Fuzz 2: goexif2][goexif2-blog] discussing fuzzing the
[goexif2][goexif2-github] package.

## [CLI Package Tests](cli-package-tests)
Sample programs and notes while learning urfave-cli and Cobra packages for
creating CLI applications.

### [urfave/cli](cli-package-tests/urfave-cli)

### [spf13/cobra](cli-package-tests/spf13-cobra)

## [Base64 Stream Decoder](base64-stream-decoder/b64-stream-decoder.go)
Code for blog post [Decoding Large Base64 Files with Go][go-base64-decoder]
showing how to use base64 stream decoder on large files.

## [PNG Tests](png-tests)
Reading and writing PNG files. Includes some test scripts and code for blog post
[Extracting PNG Chunks with Go][go-png-chunk].

## [Go Slices And Their Oddities](slice-oddities)
Code for the blog post [Go Slices and Their Oddities][go-slices].
`go run slice-oddities/quiz1.go` to see the first quiz.

<!-- Links -->

[go-pcap]: https://parsiya.net/blog/2017-12-03-go-and-pcaps/
[go-sshharvester]: https://parsiya.net/blog/2017-12-28-simple-ssh-harvester-in-go/
[go-base64-decoder]: https://parsiya.net/blog/2018-01-19-decoding-large-base64-files-with-go/
[go-png-chunk]: https://parsiya.net/blog/2018-02-25-extracting-png-chunks-with-go/
[go-fuzz-github]: https://github.com/dvyukov/go-fuzz
[iprange-github]: https://github.com/malfunkt/iprange
[iprange-blog]: https://parsiya.net/blog/2018-04-29-learning-go-fuzz-1-iprange/
[goexif2-github]: https://github.com/xor-gate/goexif2
[goexif2-blog]: https://parsiya.net/blog/2018-05-05-learning-go-fuzz-2-goexif2/
[go-slices]: https://parsiya.net/blog/2020-05-17-go-slices-and-their-oddities/
