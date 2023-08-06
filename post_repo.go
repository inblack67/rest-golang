package main

func (s *Storage) persistPost(post *Post) error {
	res := s.db.Create(post)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Storage) getPost(id int) (*Post, error) {
	post := new(Post)
	res := s.db.Find(post, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}
