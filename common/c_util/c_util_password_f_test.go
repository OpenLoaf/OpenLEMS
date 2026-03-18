package c_util

import "testing"

func TestHashAndVerifyPassword(t *testing.T) {
	hashed, err := HashPassword("123456")
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if !IsPasswordHash(hashed) {
		t.Fatalf("expected bcrypt hash, got %q", hashed)
	}

	ok, err := VerifyPassword(hashed, "123456")
	if err != nil {
		t.Fatalf("VerifyPassword returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected password verification to succeed")
	}
}

func TestVerifyPasswordWithLegacyPlaintext(t *testing.T) {
	ok, err := VerifyPassword("111111", "111111")
	if err != nil {
		t.Fatalf("VerifyPassword returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected legacy plaintext verification to succeed")
	}
}
