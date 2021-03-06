# [1.4.0](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.3.0...v1.4.0) (2022-05-19)


### Features

* add velero dependencies resource, improve util functions ([#16](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/16)) ([0fbb4b4](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/0fbb4b462af3a2066f48af94aac7949eacd0c4d4))

# [1.3.0](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.2.3...v1.3.0) (2022-05-19)


### Features

* add observability dependencies component resource ([#15](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/15)) ([b59d6e4](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/b59d6e40176f8cfd509e016ff25892ce88ac498e))

## [1.2.3](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.2.2...v1.2.3) (2022-05-17)


### Bug Fixes

* refactor aws-auth configmap, restructure resources ([#14](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/14)) ([3cdb965](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/3cdb9657ee579bae06060928588f735b77d49dd2))

## [1.2.2](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.2.1...v1.2.2) (2022-05-06)


### Bug Fixes

* add condition to sdk tag creation ([#11](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/11)) ([0e2daa4](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/0e2daa44fae9175081f00042c8e1465c970a1922))
* add full stack, bootstrap examples, update mods ([#12](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/12)) ([94c0571](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/94c0571b2e3c8431e7f3042a985b420f467d8dc3))

## [1.2.1](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.2.0...v1.2.1) (2022-05-06)


### Bug Fixes

* add eks kubeconfig generation ([#10](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/10)) ([5965524](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/5965524059da243def34e1f17f5a9d7ac67c9cbd))
* **no-release:** add extra tag for sdk ([#9](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/9)) ([ec5c95e](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/ec5c95e55707b8a7942a7f3925ade34784d6cba2))

# [1.2.0](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.1.1...v1.2.0) (2022-05-05)


### Features

* initial boostrap cluster and argocd app components ([#8](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/8)) ([3985e77](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/3985e775ec6edea054b3e3069ba8b33f5e0d82bb))

## [1.1.1](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.1.0...v1.1.1) (2022-05-03)


### Bug Fixes

* k8s default version, add eks implementation example ([#7](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/7)) ([4e44a74](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/4e44a74010837de525483c7a2bec02a1ade00bc4))

# [1.1.0](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.0.1...v1.1.0) (2022-05-03)


### Bug Fixes

* add automation token to checkout task ([#5](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/5)) ([2616cb5](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/2616cb50e081b413735ac9ddf9713d216106dc77))
* arbitrary change to trigger release workflow ([#6](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/6)) ([f858f3d](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/f858f3d8db9115664a47d36159063c532ec96893))
* release PAT, remove comments ([#4](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/4)) ([1d096f3](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/1d096f3604d5babf7aaf9fe1787834041b03a129))
* typo in release workflow configuration ([6efb0f7](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/6efb0f745f9fa95e3328b9d1a3478dd61c6ec426))


### Features

* initial eks component resource, upgrade mods, boilerplate fixes ([#3](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/3)) ([334f80a](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/334f80ae8e3329246bdf784eaa065c6dfa1b4eb8))

## [1.0.1](https://github.com/catalystsquad/pulumi-catalystsquad-platform/compare/v1.0.0...v1.0.1) (2022-05-02)


### Bug Fixes

* upgrade go version in sdk to 1.17, add example vpc implementation ([#2](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/2)) ([fc24ca1](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/fc24ca1586e216836c8325a6c78c5e4507e183ff))

# 1.0.0 (2022-05-02)


### Features

* initial provider with vpc component resource ([#1](https://github.com/catalystsquad/pulumi-catalystsquad-platform/issues/1)) ([d70cd97](https://github.com/catalystsquad/pulumi-catalystsquad-platform/commit/d70cd97d8c1b3cf41a907c129e14f64b69f4e03a))
