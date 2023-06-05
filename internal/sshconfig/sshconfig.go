package sshconfig

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
)

var (
	// sshConfigFileName is the name of the file generated by this package
	sshConfigFileName = "scaleway.config"
	// sshDefaultConfigFileName is the name of the default ssh config
	sshDefaultConfigFileName = "config"
	sshConfigFolderHomePath  = ".ssh"

	ErrFileNotFound = errors.New("file not found")
)

type Host interface {
	Config() string
}

func Generate(hosts []Host) ([]byte, error) {
	configBuffer := bytes.NewBuffer(nil)

	for _, host := range hosts {
		configBuffer.WriteString(host.Config())
		configBuffer.WriteString("\n")
	}

	return configBuffer.Bytes(), nil
}

// ConfigFilePath returns the path of the generated file
// should be ~/.ssh/scaleway.config
func ConfigFilePath(ctx context.Context) string {
	configFolder := sshConfigFolder(ctx)
	configFile := filepath.Join(configFolder, sshConfigFileName)

	return configFile
}

func Save(ctx context.Context, hosts []Host) error {
	cfg, err := Generate(hosts)
	if err != nil {
		return err
	}

	configFile := ConfigFilePath(ctx)

	return os.WriteFile(configFile, cfg, 0600)
}

func sshConfigFolder(ctx context.Context) string {
	homeDir := core.ExtractUserHomeDir(ctx)
	return filepath.Join(homeDir, sshConfigFolderHomePath)
}

func includeLine() string {
	return fmt.Sprintf("Include %s", sshConfigFileName)
}

// DefaultConfigFilePath returns the default ssh config file path
// should be ~/.ssh/config
func DefaultConfigFilePath(ctx context.Context) string {
	configFolder := sshConfigFolder(ctx)
	configFilePath := filepath.Join(configFolder, sshDefaultConfigFileName)

	return configFilePath
}

func openDefaultConfigFile(ctx context.Context) (*os.File, error) {
	configFilePath := DefaultConfigFilePath(ctx)

	configFile, err := os.Open(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, fmt.Errorf("failed to open default ssh config file: %w", err)
	}

	return configFile, nil
}

// ConfigIsIncluded checks that ssh config file is included in user's .ssh/config
// Default config file ~/.ssh/config should start with "Include scaleway.config"
func ConfigIsIncluded(ctx context.Context) (bool, error) {
	configFile, err := openDefaultConfigFile(ctx)
	if err != nil {
		return false, err
	}
	defer configFile.Close()

	expectedLine := includeLine()

	fileScanner := bufio.NewScanner(configFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), expectedLine) {
			return true, nil
		}
	}

	return false, nil
}

// IncludeConfigFile edit default ssh config to include this package generated file
// ~/.ssh/config will be prepended with "Include scaleway.config"
func IncludeConfigFile(ctx context.Context) error {
	configFileMode := os.FileMode(0600)
	fileContent := []byte(nil)

	configFile, err := openDefaultConfigFile(ctx)
	if err != nil && err != ErrFileNotFound {
		return err
	}

	if configFile != nil {
		// Keep file mode and permissions if it exists
		fi, err := configFile.Stat()
		if err != nil {
			_ = configFile.Close()
			return fmt.Errorf("failed to stat file: %w", err)
		}
		configFileMode = fi.Mode()

		fileContent, err = io.ReadAll(configFile)
		if err != nil {
			_ = configFile.Close()
			return fmt.Errorf("failed to read file: %w", err)
		}

		_ = configFile.Close()
	}

	// Prepend config file with Include line
	fileContent = append([]byte(includeLine()+"\n"), fileContent...)

	configFolder := sshConfigFolder(ctx)
	configFilePath := filepath.Join(configFolder, sshDefaultConfigFileName)

	err = os.WriteFile(configFilePath, fileContent, configFileMode)
	if err != nil {
		return fmt.Errorf("failed to write config file %s: %w", configFilePath, err)
	}

	return nil
}
