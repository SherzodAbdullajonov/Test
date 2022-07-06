package router

import (
	"context"
	"encoding/json"
	"strconv"

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/SherzodAbdullajonov/service1/pb"
	"github.com/SherzodAbdullajonov/service2/postpb"
)

type UpdatePostRequest struct {
	PostID int    `json:"post_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ListPostsResponse struct {
	Posts []Post `json:"posts"`
}

func toListPostsResponse(posts []*postpb.Post) ListPostsResponse {
	postsList := make([]Post, 0, len(posts))
	for _, p := range posts {
		postsList = append(postsList, Post{
			ID:     int(p.Id),
			UserID: int(p.UserId),
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return ListPostsResponse{
		Posts: postsList,
	}
}

type HTTPServer struct {
	dataServiceClient pb.DataServiceClient
	postServiceClient postpb.PostServiceClient
}

func NewHTTPServer(dataServiceClient pb.DataServiceClient, postServiceClient postpb.PostServiceClient) HTTPServer {
	return HTTPServer{
		dataServiceClient: dataServiceClient,
		postServiceClient: postServiceClient,
	}
}

func (s HTTPServer) DownloadPosts(w http.ResponseWriter, r *http.Request) {
	_, err := s.dataServiceClient.DownloadPosts(context.Background(), &pb.DownloadPostsRequest{})
	if err != nil {
		log.Fatal(err)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "download of posts started",
	}, http.StatusOK)
}
func (s HTTPServer) GetPost(w http.ResponseWriter, r *http.Request) {
	postID, err := s.extractPostIDParam(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	post, err := s.postServiceClient.GetPost(r.Context(), &postpb.GetPostRequest{
		PostId: postID,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	s.respondJSON(w, r, Post{
		ID:     int(post.Id),
		UserID: int(post.UserId),
		Title:  post.Title,
		Body:   post.Body,
	}, http.StatusOK)
}

func (s HTTPServer) ListPosts(w http.ResponseWriter, r *http.Request) {
	page, limit := s.parsePageAndLimitQueryParams(r)
	posts, err := s.postServiceClient.ListPosts(r.Context(), &postpb.ListPostsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	s.respondJSON(w, r, toListPostsResponse(posts.Posts), http.StatusOK)
}

func (s HTTPServer) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var request UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Fatal(err)
		return
	}

	_, err := s.postServiceClient.UpdatePost(r.Context(), &postpb.Post{
		Id:    int64(request.PostID),
		Title: request.Title,
		Body:  request.Body,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "ok",
	}, http.StatusOK)
}

func (s HTTPServer) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID, err := s.extractPostIDParam(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = s.postServiceClient.DeletePost(r.Context(), &postpb.DeletePostRequest{
		PostId: int64(postID),
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "ok",
	}, http.StatusOK)
}

func (s HTTPServer) respondJSON(w http.ResponseWriter, r *http.Request, body interface{}, status int) {
	bytesBody, err := json.MarshalIndent(body, " ", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(status)
	w.Write(bytesBody)
}

func (s HTTPServer) parsePageAndLimitQueryParams(r *http.Request) (int64, int64) {
	queries := r.URL.Query()
	pageVal, limitVal := queries.Get("page"), queries.Get("limit")

	page, _ := strconv.ParseInt(pageVal, 10, 64)
	limit, _ := strconv.ParseInt(limitVal, 10, 64)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return page, limit
}

func (s HTTPServer) extractPostIDParam(r *http.Request) (int64, error) {
	postIDParam := mux.Vars(r, "post-id")
	return strconv.ParseInt(postIDParam, 10, 64)
	
}
