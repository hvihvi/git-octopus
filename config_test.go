package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoCommit(t *testing.T) {
	repo := createTestRepo()
	defer cleanupTestRepo(repo)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := getOctopusConfig(repo, nil)

	// THEN doCommit should be true
	assert.True(t, octopusConfig.doCommit)
	assert.Nil(t, err)

	// GIVEN config to false, no option
	repo.git("config", "octopus.commit", "false")
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, nil)

	// THEN doCommit should be false
	assert.False(t, octopusConfig.doCommit)
	assert.Nil(t, err)

	// Config to 0, no option. doCommit should be true
	repo.git("config", "octopus.commit", "0")
	octopusConfig, err = getOctopusConfig(repo, nil)

	assert.False(t, octopusConfig.doCommit)
	assert.Nil(t, err)

	// GIVEN config to false, -c option true
	repo.git("config", "octopus.commit", "false")
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, []string{"-c"})

	// THEN  doCommit should be true
	assert.True(t, octopusConfig.doCommit)
	assert.Nil(t, err)

	// GIVEN config to true, -n option true
	repo.git("config", "octopus.commit", "true")
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, []string{"-n"})

	// THEN  doCommit should be false
	assert.False(t, octopusConfig.doCommit)
	assert.Nil(t, err)
}

func TestChunkMode(t *testing.T) {
	repo := createTestRepo()
	defer cleanupTestRepo(repo)

	// GIVEN No option
	// WHEN
	octopusConfig, err := getOctopusConfig(repo, nil)

	// THEN chunkSize should be 0
	assert.Equal(t, 0, octopusConfig.chunkSize)
	assert.Nil(t, err)

	// GIVEN option -s 5
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, []string{"-s", "5"})

	// THEN chunkSize should be 5
	assert.Equal(t, 5, octopusConfig.chunkSize)
	assert.Nil(t, err)
}

func TestExcludedPatterns(t *testing.T) {
	repo := createTestRepo()
	defer cleanupTestRepo(repo)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := getOctopusConfig(repo, nil)

	// THEN excludedPatterns should be empty
	assert.Empty(t, octopusConfig.excludedPatterns)
	assert.Nil(t, err)

	// GIVEN excludePattern config, no option
	repo.git("config", "octopus.excludePattern", "excluded/*")
	repo.git("config", "--add", "octopus.excludePattern", "excluded_branch")
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, nil)

	// THEN excludedPatterns should be set
	assert.Equal(t, []string{"excluded/*", "excluded_branch"}, octopusConfig.excludedPatterns)
	assert.Nil(t, err)

	// GIVEN excludePattern config (from previous assertion), option given
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, []string{"-e", "override_excluded"})

	// THEN option should take precedence
	assert.Equal(t, []string{"override_excluded"}, octopusConfig.excludedPatterns)
	assert.Nil(t, err)
}

func TestPatterns(t *testing.T) {
	repo := createTestRepo()
	defer cleanupTestRepo(repo)

	// GIVEN no config, no option
	// WHEN
	octopusConfig, err := getOctopusConfig(repo, nil)

	// THEN excludedPatterns should be empty
	assert.Empty(t, octopusConfig.patterns)
	assert.Nil(t, err)

	// GIVEN config, no argument.
	repo.git("config", "octopus.pattern", "test")
	repo.git("config", "--add", "octopus.pattern", "test2")
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, nil)

	// THEN patterns should be set
	assert.Equal(t, []string{"test", "test2"}, octopusConfig.patterns)
	assert.Nil(t, err)

	// GIVEN config (from previous assertion), argument given
	// WHEN
	octopusConfig, err = getOctopusConfig(repo, []string{"arg1", "arg2"})

	// THEN arguments should take precedence
	assert.Equal(t, []string{"arg1", "arg2"}, octopusConfig.patterns)
	assert.Nil(t, err)
}
