package Interface

type StorageManager interface {
	Begin() StorageManager
	Commit() error
	Rollback() error
	UserDao() StorageUserList
	TopicDao() StorageTopicList
	RoleDao() StorageRoleList
	QueueDao() StorageQueueList
	LogDao() StorageLogList
}
