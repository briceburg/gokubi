# gokubi

gokubi is a cli tool and Go library that sources disparate configuration and
outputs either a rendered template or the merged configuration in a unified format.

the name '_gokubi_' is an homage to the real-time language translating device
imagined by ![Kurt Vonnegut](docs/vonnegut.jpeg) in [Galapogos](https://en.wikipedia.org/wiki/Gal%C3%A1pagos_(novel)).


## supported formats

gokubi reads configuration from stdin or the filesystem. the file extension (or command line flag) is used to hint which decoder is used.

format | extension(s) | input | output
--- | --- | --- | ---
[hcl](https://github.com/hashicorp/hcl) | .hcl | yes | _planned_
[ini](https://en.wikipedia.org/wiki/INI_file) | .ini | _planned_ | _planned_
[java](https://en.wikipedia.org/wiki/.properties) | .properties | _planned_ | _planned_
[json](http://www.json.org/) | .json, .js | yes | yes
[shell](https://docs.docker.com/compose/env-file/) | .env, .sh, .vars | _planned_ |yes<sup>1</sup>
[toml](https://github.com/toml-lang/toml) | .toml | _planned_ | _planned_
[xml](https://www.w3.org/TR/REC-xml/) | .xml, .html | yes | yes
[yaml](http://yaml.org/) | .yaml, .yml | yes | yes


> **<sup>1</sup>** maps are serialized as JSON strings in shell output. see examples


## supported templates

gokubi c

* https://github.com/flosch/pongo2 ?
* https://github.com/karlseguin/liquid ?
* https://github.com/jmoiron/mandira ?
* https://github.com/karlseguin/gerb ?
* https://github.com/achun/template ?
* https://github.com/lestrrat/go-xslate ?
* https://github.com/CloudyKit/jet ?
* https://github.com/sipin/gorazor ?
* https://github.com/eknkc/amber ?
* https://github.com/aymerick/raymond ?


## example usage
