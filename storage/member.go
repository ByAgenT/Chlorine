package storage

import "time"

// Member is a struct representation of a Chlorine member object.
type Member struct {
	ID          *ID       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	RoomID      Reference `json:"room_id,omitempty"`
	IsAdmin     bool      `json:"is_admin,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
}

// SaveMember performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the Member object.
func (s DBStorage) SaveMember(member *Member) error {
	if member.ID == nil {
		var id ID
		member.CreatedDate = time.Now().UTC()
		err := s.QueryRow("INSERT INTO member (name, room_id, is_admin) VALUES ($1, $2, $3) RETURNING id",
			member.Name, member.RoomID, member.IsAdmin).Scan(&id)
		if err != nil {
			return err
		}
		member.ID = &id
		return nil
	}
	_, err := s.Exec("UPDATE member SET name=$2, room_id=$3, is_admin=$4 WHERE id = $1",
		member.ID, member.Name, member.RoomID, member.IsAdmin)
	if err != nil {
		return err
	}
	return nil
}

// GetMember return specific member object by it's ID.
func (s DBStorage) GetMember(memberID ID) (*Member, error) {
	member := &Member{}
	err := s.QueryRow("SELECT id, name, room_id, is_admin, created_date FROM member WHERE id = $1", memberID).Scan(
		&member.ID, &member.Name, &member.RoomID, &member.IsAdmin, &member.CreatedDate)
	if err != nil {
		return nil, err
	}
	return member, nil
}
