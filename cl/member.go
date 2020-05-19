package cl

import (
	"chlorine/storage"
	"errors"
	"fmt"
)

var (
	ErrorDeleteProtected = errors.New("delete protected")
)

type RawMember struct {
	Name   string
	RoomID int
	Role   int
}

type MemberService interface {
	CreateMember(rawMember RawMember) (*storage.Member, error)
	UpdateMember(memberID int, member RawMember) error
	GetMember(memberID int) (*storage.Member, error)
	GetMemberRole(memberID int) (*storage.MemberRole, error)
	Delete(memberID int, allowAdminDelete bool) error
}

type ChlorineMemberService struct {
	Repository storage.MemberRepository
}

// CreateMember creates member object and write it to the repository.
// TODO: handle the case of invalid room ID.
func (m ChlorineMemberService) CreateMember(rawMember RawMember) (*storage.Member, error) {
	member := &storage.Member{
		Name:   rawMember.Name,
		RoomID: rawMember.RoomID,
		Role:   rawMember.Role}
	err := m.Repository.Save(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (m ChlorineMemberService) UpdateMember(memberID int, member RawMember) error {
	memberStorageID := memberID
	memberToUpdate := &storage.Member{
		ID:     &memberStorageID,
		Name:   member.Name,
		RoomID: member.RoomID,
		Role:   member.Role,
	}
	return m.Repository.Save(memberToUpdate)
}

func (m ChlorineMemberService) GetMember(memberID int) (*storage.Member, error) {
	return m.Repository.Get(memberID)
}

func (m ChlorineMemberService) GetMemberRole(memberID int) (*storage.MemberRole, error) {
	member, err := m.Repository.Get(memberID)
	if err != nil {
		return nil, err
	}
	return m.Repository.GetRole(member)
}

func (m ChlorineMemberService) Delete(memberID int, allowAdminDelete bool) error {
	if !allowAdminDelete {
		role, err := m.GetMemberRole(memberID)
		if err != nil {
			return fmt.Errorf("cl: member: %w", err)
		}
		if role.IsAdmin {
			return fmt.Errorf("cl: member: %w", ErrorDeleteProtected)
		}
	}
	return m.Repository.Delete(memberID)
}
