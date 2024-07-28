# Dagger Cache Volumes

## Introduction

This repo shows that different containers do not see the same filesystem when a
[cache volume] is mounted with a different owner.

This seems documented but I guess I did not read carefully. In
[dagger.ContainerWithMountedCacheOpts], the `Owner` field reads:

```go
// A user:group to set for the mounted cache directory.
//
// Note that this changes the ownership of the specified mount along with the initial filesystem provided by source (if any). It does not have any effect if/when the cache has already been created.
//
// The user and group can either be an ID (1000:1000) or a name (foo:bar).
//
// If the group is omitted, it defaults to the same as the user.
Owner string
```

## Usage

Run:

    dagger call --progress=plain -m github.com/sevein/dagger-cache-volume-visibility-test run


[cache volume]: https://docs.dagger.io/manuals/developer/cache-volumes/
[dagger.ContainerWithMountedCacheOpts]: https://pkg.go.dev/dagger.io/dagger#ContainerWithMountedCacheOpts
