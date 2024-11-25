package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Info struct {
	Mail string `json:"mail"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func Load_Json(file_path string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	file, _ := os.Open(file_path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data, nil
}

func Add_Data(file_path string, data Info) error {
	json_data, _ := Load_Json(file_path)
	json_data = append(json_data, map[string]interface{}{"id": data.ID, "name": data.Name, "mail": data.Mail})
	file, _ := json.MarshalIndent(json_data, "", " ")
	_ = os.WriteFile(file_path, file, 0644)
	return nil
}

func Update_Data(file_path string, data Info, id int) error {
	found := false
	json_data, _ := Load_Json(file_path)
	for i, v := range json_data {
		if int(v["id"].(float64)) == id {
			json_data[i] = map[string]interface{}{"id": data.ID, "name": data.Name, "mail": data.Mail}
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("ID not found")
	}
	file, _ := json.MarshalIndent(json_data, "", " ")
	_ = os.WriteFile(file_path, file, 0644)
	return nil
}

func Delete_Data(file_path string, id int) error {
	found := false
	json_data, _ := Load_Json(file_path)
	for i, v := range json_data {
		if int(v["id"].(float64)) == id {
			json_data = append(json_data[:i], json_data[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("ID not found")
	}
	file, _ := json.MarshalIndent(json_data, "", " ")
	_ = os.WriteFile(file_path, file, 0644)
	return nil
}
