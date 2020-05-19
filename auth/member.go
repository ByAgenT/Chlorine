package auth

import (
	"chlorine/cl"
	"chlorine/storage"
	"github.com/gorilla/sessions"
	"log"
)

func GetMemberIfAuthorized(service cl.MemberService,
	session *sessions.Session) (*storage.Member, bool) {
	memberID, ok := session.Values["MemberID"].(int)
	if !ok {
		return nil, false
	}
	member, err := service.GetMember(memberID)
	if err != nil {
		log.Printf("auth: %s", err)
		return nil, false
	}
	return member, true
}

func IsMemberAdministrator(service cl.MemberService, member *storage.Member) bool {
	role, err := service.GetMemberRole(*member.ID)
	if err != nil {
		return false
	}
	return role.IsAdmin
}
