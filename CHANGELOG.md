# Changelog

## 0.1.0 (2025-05-14)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/OpenExecProtocol/oxp-go/compare/v0.0.2...v0.1.0)

### Features

* **client:** add support for endpoint-specific base URLs in python ([3435e12](https://github.com/OpenExecProtocol/oxp-go/commit/3435e12d250ba62a5fb9d2fba16fe91e97b789e4))
* **client:** add support for reading base URL from environment variable ([da52f00](https://github.com/OpenExecProtocol/oxp-go/commit/da52f0027c05cc6be43fd7293589dfc7ebb14837))
* **client:** support custom http clients ([#11](https://github.com/OpenExecProtocol/oxp-go/issues/11)) ([b7b678c](https://github.com/OpenExecProtocol/oxp-go/commit/b7b678c5f4d9c146dd23871c0662b9ef5403ad5b))


### Bug Fixes

* **client:** clean up reader resources ([02e8b5f](https://github.com/OpenExecProtocol/oxp-go/commit/02e8b5fb02559a8c9add291312b9f9877cae7f53))
* **client:** correctly update body in WithJSONSet ([3190e17](https://github.com/OpenExecProtocol/oxp-go/commit/3190e17794373f3889eb2c28c8c0f90de3fd3ac4))
* **client:** return error on bad custom url instead of panic ([#10](https://github.com/OpenExecProtocol/oxp-go/issues/10)) ([8edf5ec](https://github.com/OpenExecProtocol/oxp-go/commit/8edf5eced0206863e5ae85870a85cff7e548d087))
* handle empty bodies in WithJSONSet ([43b0734](https://github.com/OpenExecProtocol/oxp-go/commit/43b073417aa7257ab842d7ce7859f29f780905a1))
* pluralize `list` response variables ([#9](https://github.com/OpenExecProtocol/oxp-go/issues/9)) ([2f32d03](https://github.com/OpenExecProtocol/oxp-go/commit/2f32d03bbadb69b873b8c57ac444b09d467d32dc))
* **test:** return early after test failure ([#7](https://github.com/OpenExecProtocol/oxp-go/issues/7)) ([e163f84](https://github.com/OpenExecProtocol/oxp-go/commit/e163f841e7943677016c18bd60f21de10a0eaa0e))


### Chores

* add request options to client tests ([#6](https://github.com/OpenExecProtocol/oxp-go/issues/6)) ([e478bcd](https://github.com/OpenExecProtocol/oxp-go/commit/e478bcd482587882013395d08abd22df061b4d12))
* **ci:** add timeout thresholds for CI jobs ([482d0da](https://github.com/OpenExecProtocol/oxp-go/commit/482d0dac6399b0e7816505e79db2336e7796f43c))
* **ci:** only use depot for staging repos ([142cc65](https://github.com/OpenExecProtocol/oxp-go/commit/142cc659c7157b718fec883d9985c89d8076d2c7))
* **docs:** document pre-request options ([20574c5](https://github.com/OpenExecProtocol/oxp-go/commit/20574c517ada77d8e9cc532842090674609b704e))
* **docs:** improve security documentation ([#4](https://github.com/OpenExecProtocol/oxp-go/issues/4)) ([9a98381](https://github.com/OpenExecProtocol/oxp-go/commit/9a98381e85ec1abb6297303695a01bf5b0bb0806))
* fix typos ([#8](https://github.com/OpenExecProtocol/oxp-go/issues/8)) ([a327a76](https://github.com/OpenExecProtocol/oxp-go/commit/a327a76a39b0c7ea966e429d769f96cda6086c5d))
* **internal:** codegen related update ([63aa3d5](https://github.com/OpenExecProtocol/oxp-go/commit/63aa3d57eb4c8ce13fa441d9e0f416ba8aee5f3e))
* **internal:** codegen related update ([d36116e](https://github.com/OpenExecProtocol/oxp-go/commit/d36116ef0fbe894a5e4e8b23e3c38a8524b4883d))
* **internal:** expand CI branch coverage ([825a88c](https://github.com/OpenExecProtocol/oxp-go/commit/825a88c183de7e804d22b0e0e34531bdddeefe03))
* **internal:** reduce CI branch coverage ([2f052f3](https://github.com/OpenExecProtocol/oxp-go/commit/2f052f3bc440f8df016ec45fd2a324748e427267))


### Documentation

* update documentation links to be more uniform ([04176d5](https://github.com/OpenExecProtocol/oxp-go/commit/04176d5fae7f921179aac873854a2ee02fd09cf0))

## 0.0.2 (2025-03-16)

Full Changelog: [v0.0.1-alpha.0...v0.0.2](https://github.com/OpenExecProtocol/oxp-go/compare/v0.0.1-alpha.0...v0.0.2)

### Chores

* configure new SDK language ([005dc1b](https://github.com/OpenExecProtocol/oxp-go/commit/005dc1b4941a9e21ca9c24995e8fc1b598732ede))
* go live ([#1](https://github.com/OpenExecProtocol/oxp-go/issues/1)) ([540e0af](https://github.com/OpenExecProtocol/oxp-go/commit/540e0af8ad66edb5cf3c852175f0a0f86e747075))
