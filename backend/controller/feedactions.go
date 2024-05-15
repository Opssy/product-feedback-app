package controller

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
)


func CreateFeedBack(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
	 w.WriteHeader(http.StatusMethodNotAllowed)
	  fmt.Fprintf(w, "Method not allowed: %s", r.Method)
    return
	}

	//struct to represent the expected request body
	type createFeedBackRequest struct{
 		Title   string `json:"title"`
		Category string `json:"category"`
		Detail   string `json:"detail"`
	}

	var feedbackRequest createFeedBackRequest
    
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&feedbackRequest);
	err != nil{
	w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Error decoding request body: %v", err)
    return
	}

	// Basic validation
	if feedbackRequest.Title == "" || feedbackRequest.Category == "" || feedbackRequest.Detail == "" {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Missing required fields: title, category, or detail")
	return
    }

	//model
	feedback := model.CreateFeedBack{
		Title:  feedbackRequest.Title,
		Category: feedbackRequest.Category,
		Details: feedbackRequest.Detail,
	}

	//Save feedback to db 
	err := model.DB.Create(&feedback).Error
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving feedback: %v", err)
		return
	}
	// Feedback created successfully
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Feedback created successfully")
}
 
 func GetFeedBack(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
	 w.WriteHeader(http.StatusMethodNotAllowed)
	  fmt.Fprintf(w, "Method not allowed: %s", r.Method)
    return
	}

    var feedbacks []model.Feedback

	err := model.DB.Find(&feedbacks).Error
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error retrieving feedback: %v", err)
    return
    }

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(feedbacks); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error encoding feedback: %v", err)
    return
  }
	
}

func GetFeedBackById(w http.ResponseWriter, r *http.Request){
     if r.Method != http.MethodGet{
	  w.WriteHeader(http.StatusMethodNotAllowed)
	  fmt.Fprintf(w, "Method not allowed: %s", r.Method)
      return
	 }

}