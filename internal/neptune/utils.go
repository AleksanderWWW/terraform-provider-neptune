package neptune

import "terraform-provider-neptune/internal/neptune/models"

func IsMemberInProject(memberList []*models.ProjectMemberDTO, username string) bool {
	for _, member := range memberList {
		if *member.RegisteredMemberInfo.Username == username {
			return true
		}
	}
	return false
}
