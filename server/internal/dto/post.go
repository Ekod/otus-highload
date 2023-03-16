package dto

type PostCreateRequest struct {
	Content string `json:"content" binding:"required"`
	UserID  int    `json:"userID" binding:"required"`
}

type PostCreateResponse struct {
	PostID string `json:"postID"`
}

type PostGetRequest struct {
	PostID int `json:"postID" binding:"required"`
}

type PostGetResponse struct {
	ID     int    `json:"id"`
	Post   string `json:"post"`
	UserID int    `json:"userID"`
}

type PostUpdateRequest struct {
	PostID  int    `json:"postID" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDeleteRequest struct {
	PostID int `json:"postID" binding:"required"`
}

type PostFeedRequest struct {
	FriendIDs []int `json:"friendIDs"`
}

type PostFeedResponse struct {
	Posts PostFeed
}

type PostFeed map[int][]string
