# TODO

## Development
* Get the line number of the YAML parsing error.
* Hint to a YAML validation site.
* Implement "-f foo:bar" filter logic.
  * Adjust the help text!
* Use some kind of enum for config types.
* Show improved error message when input file is a directory! Currently the error is about a file which cannot be read.
* Adjust validation rules.
* Make sure "-v" output matches the other application's output!
* Use Go 1.8! Make sure to update docker images and stuff!
* Improve validation code.
* [kpcli](http://kpcli.sourceforge.net/)

## Links
* https://github.com/golang/go/wiki/GithubCodeLayout
* https://github.com/tcnksm-sample/travis-golang/blob/master/.travis.yml
* [awesome-go](https://github.com/avelino/awesome-go)
* [waffle.io](https://waffle.io/)

## Libraries
* https://github.com/asaskevich/govalidator
* https://github.com/go-playground/validator
* https://github.com/avelino/awesome-go
* [spew](https://github.com/davecgh/go-spew)
* [logrus](github.com/Sirupsen/logrus)

## Tools
* [gox](https://github.com/mitchellh/gox)

## Other
* https://github.com/mattn
* Plugin-System: https://awesome-go.com/#embeddable-scripting-languages
* binary stripping?
* https://onsi.github.io/ginkgo/
* http://agouti.org/
* http://onsi.github.io/gomega/
* !!! https://github.com/stretchr/testify

## Plugin
* https://appliedgo.net/plugins/
* https://github.com/natefinch/pie
* https://en.wikipedia.org/wiki/Foreign_function_interface

### External process, RPC
* https://npf.io/2015/05/pie/, Go
* https://github.com/hashicorp/go-plugin, Go
* https://github.com/grpc/grpc-go
* https://grpc.io/
* https://golang.org/pkg/net/rpc/
* Plugin does not crash app
* Any language
* Must be compiled, depending on language

### Compile-time plugins
* Must be compiled
* Must compile whole application

### Embedded scripting language
* https://github.com/yuin/gopher-lua, Lua
* https://github.com/Shopify/go-lua, Lua
* https://github.com/robertkrimen/otto, JS
* https://github.com/sbinet/go-python, Python
* In-process
* Performance overhead, negligible
