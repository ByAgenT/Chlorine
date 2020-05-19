package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

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
	Save(member *Member) error
	Get(memberID int) (*Member, error)
	GetRole(member *Member) (*MemberRole, error)
	Delete(memberID int) error
}

type PGMemberRepository struct {
	Storage *DBStorage
}

// GetRole returns member role.
func (m PGMemberRepository) GetRole(member *Member) (*MemberRole, error) {
	role := &MemberRole{}
	err := m.Storage.QueryRow("SELECT id, role_name, is_admin FROM member_role WHERE id = $1", member.Role).Scan(
		&role.ID, &role.Name, &role.IsAdmin)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("member %w", ErrorNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("storage: member: %w", err)
	}
	return role, nil
}

// Save performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the Member object.
func (m PGMemberRepository) Save(member *Member) error {
	if member.ID == nil {
		var id int
		member.CreatedDate = time.Now().UTC()
		err := m.Storage.QueryRow("INSERT INTO member (name, room_id, role) VALUES ($1, $2, $3) RETURNING id",
			member.Name, member.RoomID, member.Role).Scan(&id)
		if err != nil {
			return fmt.Errorf("storage: member: %w", err)
		}
		member.ID = &id
		return nil
	}
	_, err := m.Storage.Exec("UPDATE member SET name=$2, room_id=$3, role=$4 WHERE id = $1",
		member.ID, member.Name, member.RoomID, member.Role)
	if err != nil {
		return fmt.Errorf("storage: member: %w", err)
	}
	return nil
}

// Get return specific member object by it's ID.
func (m PGMemberRepository) Get(memberID int) (*Member, error) {
	member := &Member{}
	err := m.Storage.QueryRow("SELECT id, name, room_id, role, created_date FROM member WHERE id = $1", memberID).Scan(
		&member.ID, &member.Name, &member.RoomID, &member.Role, &member.CreatedDate)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("member %d %w", memberID, ErrorNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("storage: member: %w", err)
	}
	return member, nil
}

func (m PGMemberRepository) Delete(memberID int) error {
	query := "DELETE FROM member WHERE id = $1"
	_, err := m.Storage.Exec(query, memberID)
	if err != nil {
		return fmt.Errorf("storage: member: %w", err)
	}
	return nil
}
