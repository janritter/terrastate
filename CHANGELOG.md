## [1.11.2](https://github.com/janritter/terrastate/compare/1.11.1...1.11.2) (2022-08-14)


### Bug Fixes

* updated go packages and go version to 1.19 ([ab94415](https://github.com/janritter/terrastate/commit/ab944154a9bc223d076c0076f5815d201c964886))

## [1.11.1](https://github.com/janritter/terrastate/compare/1.11.0...1.11.1) (2022-06-21)


### Bug Fixes

* **deps:** update minor - go ([2acc24a](https://github.com/janritter/terrastate/commit/2acc24aedba7e70afeb3ce7f086df58e6d7317b9))

# [1.11.0](https://github.com/janritter/terrastate/compare/1.10.0...1.11.0) (2022-01-06)


### Bug Fixes

* fixed non working tf-init-upgrade flag ([ab4d382](https://github.com/janritter/terrastate/commit/ab4d3827d4d605e95e1e6c51f67ad6f425122bcf))
* fixed non working tf-init-upgrade flag ([d7546c7](https://github.com/janritter/terrastate/commit/d7546c713e75ad96dd268570fe2ecfa4160a2f03))


### Features

* added option to add the -upgrade flag to the terraform init command ([bd14aba](https://github.com/janritter/terrastate/commit/bd14aba558f26d1100d167cce65bb8f58ebd5452))
* added tf-init-upgrade flag ([86a7896](https://github.com/janritter/terrastate/commit/86a789651d0721303638396f156972b0b16be4c0))

# [1.10.0](https://github.com/janritter/terrastate/compare/1.9.0...1.10.0) (2021-09-01)


### Features

* removed unused go dep files ([84c3e75](https://github.com/janritter/terrastate/commit/84c3e75a3a00c9f6b1e965f4651e7ef31487e04e))

# [1.9.0](https://github.com/janritter/terrastate/compare/1.8.0...1.9.0) (2021-06-27)


### Features

* updated dependencies ([3632d43](https://github.com/janritter/terrastate/commit/3632d439effb28c10c97802bc8ce3cb0df609de7))

# [1.8.0](https://github.com/janritter/terrastate/compare/1.7.1...1.8.0) (2020-10-03)


### Features

* added support for refresh command - closes [#37](https://github.com/janritter/terrastate/issues/37) ([686bbc8](https://github.com/janritter/terrastate/commit/686bbc83d44b13ee829318aea12a5a2115023095))

## [1.7.1](https://github.com/janritter/terrastate/compare/1.7.0...1.7.1) (2020-09-16)


### Bug Fixes

* only append additional terraform args if set ([da7fae5](https://github.com/janritter/terrastate/commit/da7fae56cbf1693679dbc99ad8b3194be1e86fdb))

# [1.7.0](https://github.com/janritter/terrastate/compare/1.6.2...1.7.0) (2020-09-16)


### Features

* added terraform flag support to apply, plan and destroy ([7d3d520](https://github.com/janritter/terrastate/commit/7d3d52054b2550e2cebacc61e1d6a638cfd548e2))

## [1.6.2](https://github.com/janritter/terrastate/compare/1.6.1...1.6.2) (2020-09-06)


### Performance Improvements

* **rework:** changed the output function to work with the new attributes slice ([222ea3d](https://github.com/janritter/terrastate/commit/222ea3d5ebf017339c4714556bf79d3899f9c29e))
* **rework:** implemented a new creator for the state file which generates the fill fully dynamic instead of a template ([20f4bcb](https://github.com/janritter/terrastate/commit/20f4bcb18e48c91199856069cade1cbd5d3cc682))
* **rework:** implemented all attributes supported by terraform s3 backend ([e77399a](https://github.com/janritter/terrastate/commit/e77399a85874d51a4551795bb01c6c7137b894bb))
* **rework:** moved everything to interfaces to make it more testable, also made terrastate variables configurable in a central map ([2c1fe67](https://github.com/janritter/terrastate/commit/2c1fe6712c01c8b2d1104e655191729b9b466b2f))
* **rework:** new Process function which iterates through the attribute slice and collects the required information ([7b82da3](https://github.com/janritter/terrastate/commit/7b82da327700cda4cc5bc9881393ca13bed6acbe))
* **rework:** readded placeholder logic ([29da81c](https://github.com/janritter/terrastate/commit/29da81cbfffdbd68ed9f31f753108cde03b281f3))
* **rework:** reimplemented the backend specific part, implemented a struct based approach to ease adding of new backends and attributes and to lower the lines of code ([b618d3b](https://github.com/janritter/terrastate/commit/b618d3b291043bf95afca0e9f01277c6c2108718))
* **rework:** remove unnecessary return value from Generate function ([258ee8c](https://github.com/janritter/terrastate/commit/258ee8cb5d350218d90d313c330a30ec968f231d))
* **rework:** replaced errors with os.exit ([1bf1a58](https://github.com/janritter/terrastate/commit/1bf1a580596d12d1c96c5e0f32dff0176f47d3ad))
* **rework:** using new Generator command of the backend ([4d6dd85](https://github.com/janritter/terrastate/commit/4d6dd85cd70996cb51f7123fc489e2b29abcb3ba))

## [1.6.1](https://github.com/janritter/terrastate/compare/1.6.0...1.6.1) (2020-08-24)


### Bug Fixes

* changed acl to state_acl to match other variables ([a43ffa0](https://github.com/janritter/terrastate/commit/a43ffa05958d6ea31b4e2b5ba3f85426e7043c0d))

# [1.6.0](https://github.com/janritter/terrastate/compare/1.5.0...1.6.0) (2020-08-19)


### Features

* added support for s3 acl variable ([51a4a6e](https://github.com/janritter/terrastate/commit/51a4a6eb9183ea8d5a94e568a4cccc052f4220aa))

# [1.5.0](https://github.com/janritter/terrastate/compare/1.4.0...1.5.0) (2020-05-11)


### Features

* added terraform destroy command through ansible ([d4d33b0](https://github.com/janritter/terrastate/commit/d4d33b0cf0c4e0b7c185cc74068b66f899fc5ec9))
* support placeholders in all terrastate vars - solves [#23](https://github.com/janritter/terrastate/issues/23) ([36c1189](https://github.com/janritter/terrastate/commit/36c1189e84245328805a5ac08aa731c80fa66cce))

# [1.4.0](https://github.com/janritter/terrastate/compare/1.3.1...1.4.0) (2020-04-20)


### Bug Fixes

* fixed wording in parameter not found error message ([ff4e0f1](https://github.com/janritter/terrastate/commit/ff4e0f14b8279010e4e2f56bd3edf0f71ea1ccc4))


### Features

* added apply command to run terrastate, terraform init and apply in one single command ([f1df564](https://github.com/janritter/terrastate/commit/f1df5643a6bafa795d0b7ed526691f599740a198))
* added plan to call terrastate, terraform init and terraform plan with a single command ([cb1cc13](https://github.com/janritter/terrastate/commit/cb1cc13e0e6262b69c2ebf1d9f248260faff96e9))
* updated dependencies ([28ed744](https://github.com/janritter/terrastate/commit/28ed744172cd76789cb07d46cf9c02aec6a50e3f))

## [1.3.1](https://github.com/janritter/terrastate/compare/1.3.0...1.3.1) (2019-10-07)


### Bug Fixes

* Removal of old tfstate failed when file was not existent ([e1f109d](https://github.com/janritter/terrastate/commit/e1f109d))

# [1.3.0](https://github.com/janritter/terrastate/compare/1.2.3...1.3.0) (2019-10-06)


### Features

* improved removal of .terraform fodler, only remove tfstate file instaed of the whole folder ([d33b005](https://github.com/janritter/terrastate/commit/d33b005))

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
