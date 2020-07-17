package memory

import "testing"

func TestMemory(t *testing.T) {
	store := NewKVStore()
	t.Run("testing the key value store", func(t *testing.T) {
		t.Run("testing the new key value data structure", func(t *testing.T) {
			if store == nil {
				t.Errorf("calling the \"%s\" function returns \"%v\" address", "NewKVStore", nil)
			}
		})
		t.Run("getting value for a non-existent key", func(t *testing.T) {
			_, err := store.Get("key0")
			if err == nil {
				t.Errorf("expecting no error, got error \"%s\"", err.Error())
			}
		})
		t.Run("getting value for an existing key", func(t *testing.T) {
			err := store.Put("key-0", "value-0")
			if err != nil {
				t.Fatal("error in storing the record: ", err.Error())
			}
			want := "value-0"
			got, err := store.Get("key-0")
			if err != nil {
				t.Fatal("error in getting the record: ", err.Error())
			}
			if got != want {
				t.Errorf("expecting \"%s\", got \"%s\"", want, got)
			}
		})
	})
}
