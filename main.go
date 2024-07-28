package main

import (
	"context"
	"time"

	"dagger/test/internal/dagger"
)

type Test struct{}

func (m *Test) runner(opts shareOpts, cmd ...string) *dagger.Container {
	return dag.Container().
		From("alpine:latest").
		WithEnvVariable("CACHEBUSTER", time.Now().Format(time.RFC3339Nano)).
		WithMountedCache(opts.sharePath, opts.shareVol, opts.shareOpts).
		WithWorkdir(opts.sharePath).
		WithExec(cmd)
}

func (m *Test) Run(ctx context.Context) error {
	share := dag.CacheVolume(
		// Use random key ensures this volume is new.
		time.Now().Format(time.RFC3339Nano),
	)

	// Create file in the shared volume.
	opts := shareOpts{share, "/mnt/shared", dagger.ContainerWithMountedCacheOpts{Sharing: dagger.Shared}}
	_, err := m.runner(opts, "touch", "foobar").Sync(ctx)
	if err != nil {
		return err
	}

	// Stat the file we've just created.
	opts = shareOpts{share, "/mnt/shared", dagger.ContainerWithMountedCacheOpts{Sharing: dagger.Shared}}
	_, err = m.runner(opts, "stat", "foobar").Sync(ctx)
	if err != nil {
		return err
	}

	// Confusingly, the directory reads as empty when we change the owner.
	opts = shareOpts{share, "/mnt/shared", dagger.ContainerWithMountedCacheOpts{Sharing: dagger.Shared, Owner: "1000:1000"}}
	_, err = m.runner(opts, "ls", "-lha").Sync(ctx)
	if err != nil {
		return err
	}

	return nil
}

type shareOpts struct {
	shareVol  *dagger.CacheVolume
	sharePath string
	shareOpts dagger.ContainerWithMountedCacheOpts
}
