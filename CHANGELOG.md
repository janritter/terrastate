## [1.2.3](https://github.com/janritter/terrastate/compare/1.2.2...1.2.3) (2019-07-13)


### Bug Fixes

* improved s3 template to be identical to the output of the terraform code formatter ([3321145](https://github.com/janritter/terrastate/commit/3321145))

## [1.2.2](https://github.com/janritter/terrastate/compare/1.2.1...1.2.2) (2019-06-02)


### Bug Fixes

* also show the "Skipping removing of .terraform folder" message when the value is set to false and not only when the value is not set ([967f438](https://github.com/janritter/terrastate/commit/967f438))

## [1.2.1](https://github.com/janritter/terrastate/compare/1.2.0...1.2.1) (2019-06-02)


### Bug Fixes

* added missing checks and outputs for errors ([90ea548](https://github.com/janritter/terrastate/commit/90ea548))

# [1.2.0](https://github.com/janritter/terrastate/compare/1.1.0...1.2.0) (2019-03-24)


### Bug Fixes

* use golang template to generate backend configuration - fixes [#5](https://github.com/janritter/terrastate/issues/5) ([18ff8a2](https://github.com/janritter/terrastate/commit/18ff8a2))


### Features

* added option to automaticly remove the .terraform folder when creation a new backend config file - closes [#6](https://github.com/janritter/terrastate/issues/6) ([d2a834d](https://github.com/janritter/terrastate/commit/d2a834d))

# [1.1.0](https://github.com/janritter/terrastate/compare/1.0.0...1.1.0) (2019-03-02)


### Features

* added multi backend support ([9bb3cc6](https://github.com/janritter/terrastate/commit/9bb3cc6))

# 1.0.0 (2019-03-01)


### Features

* added functions read terraform var file and generate state terraform file ([0447b29](https://github.com/janritter/terrastate/commit/0447b29))
* project init ([d8c9ad0](https://github.com/janritter/terrastate/commit/d8c9ad0))
