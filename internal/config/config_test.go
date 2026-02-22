package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_Valid(t *testing.T) {
	content := `
[server]
    YPPrefix = "tp"
    Port     = "9000"
    LogPath  = "test.log"
    Debug    = false

[database]
    Host     = "db.example.com"
    Port     = "3307"
    User     = "testuser"
    Password = "testpass"
    DB       = "testdb"

[peercast]
    Host         = "pc.example.com"
    Port         = "7145"
    AuthType     = "basic"
    AuthUser     = "admin"
    AuthPassword = "secret"
`
	dir := t.TempDir()
	path := filepath.Join(dir, "test.toml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Server.Port != "9000" {
		t.Errorf("Server.Port: expected '9000', got %q", cfg.Server.Port)
	}
	if cfg.Server.YPPrefix != "tp" {
		t.Errorf("Server.YPPrefix: expected 'tp', got %q", cfg.Server.YPPrefix)
	}
	if cfg.Server.Debug {
		t.Error("Server.Debug: expected false")
	}
	if cfg.Database.Host != "db.example.com" {
		t.Errorf("Database.Host: expected 'db.example.com', got %q", cfg.Database.Host)
	}
	if cfg.Database.DB != "testdb" {
		t.Errorf("Database.DB: expected 'testdb', got %q", cfg.Database.DB)
	}
	if cfg.Peercast.AuthPassword != "secret" {
		t.Errorf("Peercast.AuthPassword: expected 'secret', got %q", cfg.Peercast.AuthPassword)
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config.toml")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoad_InvalidTOML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.toml")
	if err := os.WriteFile(path, []byte("{{invalid toml}}"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(path)
	if err == nil {
		t.Error("expected error for invalid TOML")
	}
}
