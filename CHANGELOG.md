# Changelog

## [0.1.8](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.7...v0.1.8) (2024-02-15)


### 🐛 Bug Fixes

* targeting key serialization fix ([#46](https://github.com/spotify/confidence-openfeature-provider-go/issues/46)) ([859dd88](https://github.com/spotify/confidence-openfeature-provider-go/commit/859dd884dbbb05277d92b9f952a8209d77585c72))

## [0.1.7](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.6...v0.1.7) (2024-02-12)


### 🐛 Bug Fixes

* add a default endpoint to the provider ([#41](https://github.com/spotify/confidence-openfeature-provider-go/issues/41)) ([c4bfe45](https://github.com/spotify/confidence-openfeature-provider-go/commit/c4bfe45d76b8fff131e795c5462d59b9a8ea2098))
* ApiConfig constructor with region defaults ([#43](https://github.com/spotify/confidence-openfeature-provider-go/issues/43)) ([064bef2](https://github.com/spotify/confidence-openfeature-provider-go/commit/064bef291ec72a6ff3facf4a9ecc7d6377ef410d))


### 📚 Documentation

* update instructions on creating ApiConfig ([#44](https://github.com/spotify/confidence-openfeature-provider-go/issues/44)) ([dd593f6](https://github.com/spotify/confidence-openfeature-provider-go/commit/dd593f676fe7309e6b0d71e01bbf8c08a05456d2))

## [0.1.6](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.5...v0.1.6) (2023-11-22)


### 🐛 Bug Fixes

* Fix nested struct bug ([#37](https://github.com/spotify/confidence-openfeature-provider-go/issues/37)) ([5fb769f](https://github.com/spotify/confidence-openfeature-provider-go/commit/5fb769f3db4c00376ea4e34b510d4391f29003c5))


### 📚 Documentation

* Small cleanups and documentation ([#38](https://github.com/spotify/confidence-openfeature-provider-go/issues/38)) ([9c56b9f](https://github.com/spotify/confidence-openfeature-provider-go/commit/9c56b9fd0cd60266bf8ff44cf9606dadc92e832d))

## [0.1.5](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.4...v0.1.5) (2023-11-15)


### ✨ New Features

* Add SDK id and version to requests ([#33](https://github.com/spotify/confidence-openfeature-provider-go/issues/33)) ([b63d8f7](https://github.com/spotify/confidence-openfeature-provider-go/commit/b63d8f7c34efe77871a6f9d1f110513bb1d2b350))


### 🧹 Chore

* Bring back demo2 ([#30](https://github.com/spotify/confidence-openfeature-provider-go/issues/30)) ([6f0451d](https://github.com/spotify/confidence-openfeature-provider-go/commit/6f0451d2b4649ca315160eddf8d6f8ee4a792c53))


### 📚 Documentation

* Fix release please setup in README ([#34](https://github.com/spotify/confidence-openfeature-provider-go/issues/34)) ([f8f08c9](https://github.com/spotify/confidence-openfeature-provider-go/commit/f8f08c948f06cf0d12af04af10b77aae2c714c65))

## [0.1.4](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.3...v0.1.4) (2023-09-08)


### 🧹 Chore

* bring back demo now that we have 0.1.3 released ([#26](https://github.com/spotify/confidence-openfeature-provider-go/issues/26)) ([ffdf177](https://github.com/spotify/confidence-openfeature-provider-go/commit/ffdf1778acd4d5f1f0c469516ef51c77f293078b))
* remove unused dep to uuid ([#28](https://github.com/spotify/confidence-openfeature-provider-go/issues/28)) ([39008d5](https://github.com/spotify/confidence-openfeature-provider-go/commit/39008d58cae05ad78fb827a5117b1f4a3814d1b6))
* update readme ([#29](https://github.com/spotify/confidence-openfeature-provider-go/issues/29)) ([b1cf5a7](https://github.com/spotify/confidence-openfeature-provider-go/commit/b1cf5a7cf3f6fabecfc2bdd896959fb09a5762c0))

## [0.1.3](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.2...v0.1.3) (2023-09-08)


### 🧹 Chore

* Make the go.mod file root level ([#24](https://github.com/spotify/confidence-openfeature-provider-go/issues/24)) ([242acf0](https://github.com/spotify/confidence-openfeature-provider-go/commit/242acf05831c8a7176e4c1ad381a0da51374366e))

## [0.1.2](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.1...v0.1.2) (2023-09-08)


### 🧹 Chore

* update module names for the go.mod files, move out demo from the pkg ([#22](https://github.com/spotify/confidence-openfeature-provider-go/issues/22)) ([895b09f](https://github.com/spotify/confidence-openfeature-provider-go/commit/895b09f87fbf286aac9570bd1cfb6eca06abc028))

## [0.1.1](https://github.com/spotify/confidence-openfeature-provider-go/compare/v0.1.0...v0.1.1) (2023-09-08)


### 🐛 Bug Fixes

* **ci:** add manifest to the release please ([#20](https://github.com/spotify/confidence-openfeature-provider-go/issues/20)) ([a3d3086](https://github.com/spotify/confidence-openfeature-provider-go/commit/a3d3086966e039fbdedbc6e696669497eaf28ab9))
* move demo app outside the provider package ([d801049](https://github.com/spotify/confidence-openfeature-provider-go/commit/d80104966f183201784a3cbc78259d2662c93109))


### 🧹 Chore

* **ci:** add pr title linter ([#7](https://github.com/spotify/confidence-openfeature-provider-go/issues/7)) ([4214439](https://github.com/spotify/confidence-openfeature-provider-go/commit/42144396f3d29aceaa0101aaccb9b09c282f71e5))
* **ci:** add release please to the provider ([5a37c10](https://github.com/spotify/confidence-openfeature-provider-go/commit/5a37c10a194814b1f272a19533e9500f263c88b5))
* **ci:** Update release-please-config.json ([fe2aaaf](https://github.com/spotify/confidence-openfeature-provider-go/commit/fe2aaaf1a27b09d1612c245445df0d89852592e0))
* **docs:** add part about how to execute the demo app ([0cb0dc2](https://github.com/spotify/confidence-openfeature-provider-go/commit/0cb0dc276140631c4861d6ca9f43c9fc0667dc3a))
