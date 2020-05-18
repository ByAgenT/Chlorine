package cl

import (
	"chlorine/storage"
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
	err := m.Repository.SaveMember(member)
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
	return m.Repository.SaveMember(memberToUpdate)
}

func (m ChlorineMemberService) GetMember(memberID int) (*storage.Member, error) {
	return m.Repository.GetMember(memberID)
}

func (m ChlorineMemberService) GetMemberRole(memberID int) (*storage.MemberRole, error) {
	member, err := m.Repository.GetMember(memberID)
	if err != nil {
		return nil, err
	}
	return m.Repository.GetMemberRole(member)
}
