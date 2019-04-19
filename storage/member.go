package storage

import "time"

const (
	// RoleMember is a role that represents default app member.
	RoleMember = iota

	// RoleAdmin is a role that represents room administrator role.
	RoleAdmin
)

// Member is a struct representation of a Chlorine member object.
type Member struct {
	Model
	ID          *ID       `json:"id"`
	Name        string    `json:"name"`
	RoomID      Reference `json:"room_id"`
	Role        Reference `json:"role"`
	CreatedDate time.Time `json:"created_date"`
}

// MemberRole is a struct representation of a member role.
type MemberRole struct {
	ID      *ID    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty"`
}

// SaveMember performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the Member object.
func (s DBStorage) SaveMember(member *Member) error {
	if member.ID == nil {
		var id ID
		member.CreatedDate = time.Now().UTC()
		err := s.QueryRow("INSERT INTO member (name, room_id, role) VALUES ($1, $2, $3) RETURNING id",
			member.Name, member.RoomID, member.Role).Scan(&id)
		if err != nil {
			return err
		}
		member.ID = &id
		return nil
	}
	_, err := s.Exec("UPDATE member SET name=$2, room_id=$3, role=$4 WHERE id = $1",
		member.ID, member.Name, member.RoomID, member.Role)
	if err != nil {
		return err
	}
	return nil
}

// GetMember return specific member object by it's ID.
func (s DBStorage) GetMember(memberID ID) (*Member, error) {
	member := &Member{}
	err := s.QueryRow("SELECT id, name, room_id, role, created_date FROM member WHERE id = $1", memberID).Scan(
		&member.ID, &member.Name, &member.RoomID, &member.Role, &member.CreatedDate)
	if err != nil {
		return nil, err
	}
	return member, nil
}

// GetRole returns member role.
func (m Member) GetRole() (*MemberRole, error) {
	role := &MemberRole{}
	err := m.storage.QueryRow("SELECT id, role_name, is_admin FROM member_role WHERE id = $1", m.Role).Scan(
		&role.ID, &role.Name, &role.IsAdmin)
	if err != nil {
		return nil, err
	}
	return role, nil
}
