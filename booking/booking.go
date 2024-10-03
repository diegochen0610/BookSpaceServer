package booking

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MeetingRoom struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"` // "available" or "reserved"
	ReservedBy string `json:"reserved_by"`
}

var meetingRooms = []MeetingRoom{
	{ID: 1, Name: "Room 101", Status: "available"},
	{ID: 2, Name: "Room 102", Status: "reserved", ReservedBy: "admin"},
}

func GetMeetingRooms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(meetingRooms)
}

func ReserveRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomIDStr := vars["id"]
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	for i, room := range meetingRooms {
		if room.ID == roomID {
			if room.Status == "available" {
				meetingRooms[i].Status = "reserved"
				meetingRooms[i].ReservedBy = "admin" // 假設是登入的用戶
				w.Write([]byte("Room reserved successfully"))
				return
			} else {
				w.Write([]byte("Room already reserved"))
				return
			}
		}
	}
	http.Error(w, "Room not found", http.StatusNotFound)
}
