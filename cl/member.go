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
	UpdateMember(memberID int, member RawMember) (*storage.Member, error)
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
		RoomID: storage.Reference(rawMember.RoomID),
		Role:   storage.Reference(rawMember.Role)}
	err := m.Repository.SaveMember(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (m ChlorineMemberService) UpdateMember(memberID int, member RawMember) (*storage.Member, error) {
	panic("implement me")
}

func (m ChlorineMemberService) GetMember(memberID int) (*storage.Member, error) {
	panic("implement me")
}

func (m ChlorineMemberService) GetMemberRole(memberID int) (*storage.MemberRole, error) {
	panic("implement me")
}
