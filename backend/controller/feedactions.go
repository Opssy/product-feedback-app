package controller

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	 	idStr := r.URL.Query().Get("id")
		if idStr == ""{
			w.WriteHeader(http.StatusBadRequest)
		    fmt.Fprintf(w, "Missing required query parameter: id")
          return
		}
		// Convert string ID to uint (replace with your actual ID type)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid ID format: %v", err)
			return
		}

	    var feedbacks []model.Feedback

	    err = model.DB.First(&feedbacks, id).Error
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

func EditFeedBack(w http.ResponseWriter, r *http.Request){
    if r.Method !=  http.MethodPatch{
	  w.WriteHeader(http.StatusMethodNotAllowed)
	  fmt.Fprintf(w, "Method not allowed: %s", r.Method)
      return
	}

		idStr := r.URL.Query().Get("id")
		if idStr == ""{
			w.WriteHeader(http.StatusBadRequest)
		    fmt.Fprintf(w, "Missing required query parameter: id")
          return
		}
		// Convert string ID to uint (replace with your actual ID type)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid ID format: %v", err)
			return
		}
		 
		var updatedFeedBack model.Feedback
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&updatedFeedBack);
		err != nil{

		}
		 // **2. Find existing feedback by ID:**
		var existingFeedback model.Feedback
		err = model.DB.First(&existingFeedback, id).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error retrieving feedback: %v", err)
			return
		}


		existingFeedback.Title = updatedFeedBack.Title
		existingFeedback.Details = updatedFeedBack.Details
		existingFeedback.Category = updatedFeedBack.Category

		err = model.DB.Save(&existingFeedback).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error saving feedback: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Feedback edited successfully")
}

func DeleteFeedBack(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed: %s", r.Method)
		return
	}
		idStr := r.URL.Query().Get("id")
		if idStr == ""{
			w.WriteHeader(http.StatusBadRequest)
		    fmt.Fprintf(w, "Missing required query parameter: id")
          return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid ID format: %v", err)
			return
		}
		result := model.DB.Delete(&model.Feedback{ID:id})

		if result.Error != nil{

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting feedback: %v", result.Error)
		return
		}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))

}