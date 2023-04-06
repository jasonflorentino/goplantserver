package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jasonflorentino/goplantserver/logger"
	"github.com/jasonflorentino/goplantserver/types"
)

func PlantsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Println("Handling request:", r.URL.Path)

	content, err := ioutil.ReadFile("data/plants.json")
	if err != nil {
		logger.Logger.Fatal("Error reading file:", err)
	}

	var plants []types.Plant
	err = json.Unmarshal(content, &plants)
	if err != nil {
		logger.Logger.Fatal("Error unmarshalling content:", err)
	}

	jsonData, err := json.Marshal(plants)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		logger.Logger.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
