package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoCommit(t *testing.T) {
	context, _ := createTestContext()
	defer cleanup(context)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := GetOctopusConfig(context.repo, nil)

	// THEN doCommit should be true
	assert.True(t, octopusConfig.DoCommit)
	assert.Nil(t, err)

	// GIVEN config to false, no option
	context.repo.Git("config", "octopus.commit", "false")
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, nil)

	// THEN doCommit should be false
	assert.False(t, octopusConfig.DoCommit)
	assert.Nil(t, err)

	// Config to 0, no option. doCommit should be true
	context.repo.Git("config", "octopus.commit", "0")
	octopusConfig, err = GetOctopusConfig(context.repo, nil)

	assert.False(t, octopusConfig.DoCommit)
	assert.Nil(t, err)

	// GIVEN config to false, -c option true
	context.repo.Git("config", "octopus.commit", "false")
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, []string{"-c"})

	// THEN  doCommit should be true
	assert.True(t, octopusConfig.DoCommit)
	assert.Nil(t, err)

	// GIVEN config to true, -n option true
	context.repo.Git("config", "octopus.commit", "true")
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, []string{"-n"})

	// THEN  doCommit should be false
	assert.False(t, octopusConfig.DoCommit)
	assert.Nil(t, err)
}

func TestChunkMode(t *testing.T) {
	context, _ := createTestContext()
	defer cleanup(context)

	// GIVEN No option
	// WHEN
	octopusConfig, err := GetOctopusConfig(context.repo, nil)

	// THEN chunkSize should be 0
	assert.Equal(t, 0, octopusConfig.ChunkSize)
	assert.Nil(t, err)

	// GIVEN option -s 5
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, []string{"-s", "5"})

	// THEN chunkSize should be 5
	assert.Equal(t, 5, octopusConfig.ChunkSize)
	assert.Nil(t, err)
}

func TestExcludedPatterns(t *testing.T) {
	context, _ := createTestContext()
	defer cleanup(context)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := GetOctopusConfig(context.repo, nil)

	// THEN excludedPatterns should be empty
	assert.Empty(t, octopusConfig.ExcludedPatterns)
	assert.Nil(t, err)

	// GIVEN excludePattern config, no option
	context.repo.Git("config", "octopus.excludePattern", "excluded/*")
	context.repo.Git("config", "--add", "octopus.excludePattern", "excluded_branch")
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, nil)

	// THEN excludedPatterns should be set
	assert.Equal(t, []string{"excluded/*", "excluded_branch"}, octopusConfig.ExcludedPatterns)
	assert.Nil(t, err)

	// GIVEN excludePattern config (from previous assertion), option given
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, []string{"-e", "override_excluded"})

	// THEN option should take precedence
	assert.Equal(t, []string{"override_excluded"}, octopusConfig.ExcludedPatterns)
	assert.Nil(t, err)
}

func TestPatterns(t *testing.T) {
	context, _ := createTestContext()
	defer cleanup(context)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := GetOctopusConfig(context.repo, nil)

	// THEN excludedPatterns should be empty
	assert.Empty(t, octopusConfig.Patterns)
	assert.Nil(t, err)

	// GIVEN config, no argument.
	context.repo.Git("config", "octopus.pattern", "test")
	context.repo.Git("config", "--add", "octopus.pattern", "test2")
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, nil)

	// THEN patterns should be set
	assert.Equal(t, []string{"test", "test2"}, octopusConfig.Patterns)
	assert.Nil(t, err)

	// GIVEN config (from previous assertion), argument given
	// WHEN
	octopusConfig, err = GetOctopusConfig(context.repo, []string{"arg1", "arg2"})

	// THEN arguments should take precedence
	assert.Equal(t, []string{"arg1", "arg2"}, octopusConfig.Patterns)
	assert.Nil(t, err)
}
