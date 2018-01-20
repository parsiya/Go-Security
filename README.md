# My Go Security Projects

<!-- MarkdownTOC -->

- [SSH Harvester](#ssh-harvester)
- [pcap Tutorial](#pcap-tutorial)
- [CLI Package Tests](#cli-package-tests)
    - [urfave/cli](#urfavecli)
    - [spf13/cobra](#spf13cobra)
- [Base64 Stream Decoder](#base64-stream-decoder)

<!-- /MarkdownTOC -->

<a name="ssh-harvester"></a>
## [SSH Harvester](ssh-harvester)
Initial version of tool written in Go that harvests SSH certificates. For explanation of code please see the blog post [Simple SSH Harvester in Go][go-sshharvester].

<a name="pcap-tutorial"></a>
## [pcap Tutorial](pcap-tutorial)
Code for the blog post [Go and pcaps][go-pcap] which explains how to use Go to extract ICMP echo payloads from a pcap file.

<a name="cli-package-tests"></a>
## [CLI Package Tests](cli-package-tests)
Sample programs and notes while learning urfave-cli and Cobra packages for creating CLI applications.

<a name="urfavecli"></a>
### [urfave/cli](cli-package-tests/urfave-cli)

<a name="spf13cobra"></a>
### [spf13/cobra](cli-package-tests/spf13-cobra)

<a name="base64-stream-decoder"></a>
## [Base64 Stream Decoder](base64-stream-decoder/b64-stream-decoder.go)
Code for blog post [Decoding Large Base64 Files with Go][go-base64-decoder] showing how to use base64 stream decoder on large files.

<!-- Links -->

[go-pcap]: https://parsiya.net/blog/2017-12-03-go-and-pcaps/
[go-sshharvester]: https://parsiya.net/blog/2017-12-28-simple-ssh-harvester-in-go/
[go-base64-decoder]: https://parsiya.net/blog/2018-01-19-decoding-large-base64-files-with-go/
