/*
 * DevNest Portal API (OpenAPI 3.0)
 *
 * The DevNest Portal API endpoints definitions based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 * Contact: contact@devnest.ro
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"database/sql"
	"encoding/json"
	"github.com/andreibr99/DevNest-Internal-Project/database"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project NewProject

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO Projects (name, clientName, description, startingDate, endingDate, currency, status) VALUES (@name, @clientName, @description, @startingDate, @endingDate, @currency, @status)",
		sql.Named("name", project.Name),
		sql.Named("clientName", project.ClientName),
		sql.Named("description", project.Description),
		sql.Named("startingDate", project.StartingDate),
		sql.Named("endingDate", project.EndingDate),
		sql.Named("currency", project.Currency),
		sql.Named("status", project.Status),
	)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, clientName, description, startingDate, endingDate, currency, status FROM Projects"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the projects
	var projects []CreatedProject

	// Iterate over the result rows and read project details
	for rows.Next() {
		var project CreatedProject
		err = rows.Scan(&project.Id, &project.Name, &project.ClientName, &project.Description, &project.StartingDate, &project.EndingDate, &project.Currency, &project.Status)
		if err != nil {
			log.Printf("Error reading project details: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		projects = append(projects, project)
	}

	// Serialize the array of projects to JSON and send it in the responsse
	jsonResponse, err := json.Marshal(projects)
	if err != nil {
		log.Printf("Error encoding projects as JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetHistoryForProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ProjectsProjectIdDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	// Parse the project ID
	projectID, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Check if the project exists in the database
	query := "SELECT COUNT(*) FROM Projects WHERE id = @id"
	var count int
	err = database.DB.QueryRow(query, sql.Named("id", projectID)).Scan(&count)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Delete the project by ID
	query = "DELETE FROM Projects WHERE id = @id"
	_, err = database.DB.Exec(query, sql.Named("id", projectID))
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

func ProjectsProjectIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	// Parse the project ID
	projectID, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Check if the project exists in the database
	query := "SELECT COUNT(*) FROM Projects WHERE id = @id"
	var count int
	err = database.DB.QueryRow(query, sql.Named("id", projectID)).Scan(&count)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Retrieve the project details by ID
	query = "SELECT id, name, clientName, description, startingDate, endingDate, currency, status FROM Projects WHERE id = @id"
	rows, err := database.DB.Query(query, sql.Named("id", projectID))

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a struct to hold the project details
	var project CreatedProject

	if rows.Next() {
		err = rows.Scan(&project.Id, &project.Name, &project.ClientName, &project.Description, &project.StartingDate, &project.EndingDate, &project.Currency, &project.Status)
		if err != nil {
			log.Printf("Error reading project details: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	}

	// Serialize the project details to JSON
	jsonResponse, err := json.Marshal(project)
	if err != nil {
		log.Printf("Error encoding project as JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func ProjectsProjectIdPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	// Parse the project ID
	projectID, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Check if the project exists in the database
	query := "SELECT COUNT(*) FROM Projects WHERE id = @id"
	var count int
	err = database.DB.QueryRow(query, sql.Named("id", projectID)).Scan(&count)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Parse the new status from the request body.
	var updatedStatus CreatedProject

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&updatedStatus); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Update the project's status in the database for the specified project ID.
	query = "UPDATE Projects SET status = @status WHERE id = @id"
	_, err = database.DB.Exec(query, sql.Named("status", updatedStatus.Status), sql.Named("id", projectID))
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ProjectsProjectIdPut(w http.ResponseWriter, r *http.Request) {
	// Extract the projectId from the URL path
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	// Parse the project ID
	projectID, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Check if the project exists in the database
	query := "SELECT COUNT(*) FROM Projects WHERE id = @id"
	var count int
	err = database.DB.QueryRow(query, sql.Named("id", projectID)).Scan(&count)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	var updatedProject CreatedProject

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&updatedProject); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE Projects SET name = @name, clientName = @clientName, description = @description, startingDate = @startingDate, endingDate = @endingDate, currency = @currency, status = @status WHERE id = @id",
		sql.Named("id", projectID),
		sql.Named("name", updatedProject.Name),
		sql.Named("clientName", updatedProject.ClientName),
		sql.Named("description", updatedProject.Description),
		sql.Named("startingDate", updatedProject.StartingDate),
		sql.Named("endingDate", updatedProject.EndingDate),
		sql.Named("currency", updatedProject.Currency),
		sql.Named("status", updatedProject.Status),
	)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
