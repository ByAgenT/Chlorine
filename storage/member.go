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
	ID          *int      `json:"id"`
	Name        string    `json:"name"`
	RoomID      int       `json:"room_id"`
	Role        int       `json:"role"`
	CreatedDate time.Time `json:"created_date"`
}

// MemberRole is a struct representation of a member role.
type MemberRole struct {
	ID      *int   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty"`
}

type MemberRepository interface {
	SaveMember(member *Member) error
	GetMember(memberID int) (*Member, error)
	GetMemberRole(member *Member) (*MemberRole, error)
}

type PGMemberRepository struct {
	Storage *DBStorage
}

// GetMemberRole returns member role.
func (m PGMemberRepository) GetMemberRole(member *Member) (*MemberRole, error) {
	role := &MemberRole{}
	err := m.Storage.QueryRow("SELECT id, role_name, is_admin FROM member_role WHERE id = $1", member.Role).Scan(
		&role.ID, &role.Name, &role.IsAdmin)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// SaveMember performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the Member object.
func (m PGMemberRepository) SaveMember(member *Member) error {
	if member.ID == nil {
		var id int
		member.CreatedDate = time.Now().UTC()
		err := m.Storage.QueryRow("INSERT INTO member (name, room_id, role) VALUES ($1, $2, $3) RETURNING id",
			member.Name, member.RoomID, member.Role).Scan(&id)
		if err != nil {
			return err
		}
		member.ID = &id
		return nil
	}
	_, err := m.Storage.Exec("UPDATE member SET name=$2, room_id=$3, role=$4 WHERE id = $1",
		member.ID, member.Name, member.RoomID, member.Role)
	if err != nil {
		return err
	}
	return nil
}

// GetMember return specific member object by it's ID.
func (m PGMemberRepository) GetMember(memberID int) (*Member, error) {
	member := &Member{}
	err := m.Storage.QueryRow("SELECT id, name, room_id, role, created_date FROM member WHERE id = $1", memberID).Scan(
		&member.ID, &member.Name, &member.RoomID, &member.Role, &member.CreatedDate)
	if err != nil {
		return nil, err
	}
	return member, nil
}
