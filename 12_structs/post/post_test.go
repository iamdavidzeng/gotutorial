package post

import "testing"

func TestAdd(t *testing.T) {
	repo := New()
	repo.Add(Post{1, "test title", "test content"})
	if len(repo.Posts) != 1 {
		t.Errorf("Post was not added")
	}
}

func testGetAll(t *testing.T) {
	repo := New()
	repo.Add(Post{1, "test title", "test content"})
	results := repo.GetAll()
	if len(results) != 1 {
		t.Errorf("Post was not added")
	}
}
