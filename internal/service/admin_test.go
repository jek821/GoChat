package service

import (
	"context"
	"path/filepath"
	"testing"

	"ByteChat/internal/store/sqlite"
)

func TestCreateAdminUpdatesExistingPassword(t *testing.T) {
	ctx := context.Background()
	store, err := sqlite.New(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("sqlite.New: %v", err)
	}
	t.Cleanup(func() { store.Close() })

	auth := NewAuthService(store)
	admin := NewAdminService(store, auth, nil)

	if _, err := auth.Register(ctx, RegisterInput{Username: "admin", Password: "oldpassword1"}); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if err := admin.CreateAdmin(ctx, "admin", "newpassword1"); err != nil {
		t.Fatalf("CreateAdmin: %v", err)
	}

	if _, err := auth.Login(ctx, "admin", "oldpassword1"); err == nil {
		t.Fatal("expected old password to fail")
	}
	if _, err := admin.Login(ctx, "admin", "newpassword1"); err != nil {
		t.Fatalf("admin login with new password: %v", err)
	}
}
