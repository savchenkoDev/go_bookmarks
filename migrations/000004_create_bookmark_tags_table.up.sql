CREATE TABLE bookmark_tags (
  id SERIAL PRIMARY KEY,
  bookmark_id BIGINT NOT NULL,
  tag_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE bookmark_tags ADD CONSTRAINT fk_bookmark_tags_bookmark_id FOREIGN KEY (bookmark_id) REFERENCES bookmarks(id);
ALTER TABLE bookmark_tags ADD CONSTRAINT fk_bookmark_tags_tag_id FOREIGN KEY (tag_id) REFERENCES tags(id);