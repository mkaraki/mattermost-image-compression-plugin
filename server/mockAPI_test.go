package main

import (
	"io"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
)

type MattermostMockAPI struct {
}

// AddChannelMember implements plugin.API.
func (*MattermostMockAPI) AddChannelMember(channelId string, userID string) (*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// AddReaction implements plugin.API.
func (*MattermostMockAPI) AddReaction(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	panic("unimplemented")
}

// AddUserToChannel implements plugin.API.
func (*MattermostMockAPI) AddUserToChannel(channelId string, userID string, asUserId string) (*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// CopyFileInfos implements plugin.API.
func (*MattermostMockAPI) CopyFileInfos(userID string, fileIds []string) ([]string, *model.AppError) {
	panic("unimplemented")
}

// CreateBot implements plugin.API.
func (*MattermostMockAPI) CreateBot(bot *model.Bot) (*model.Bot, *model.AppError) {
	panic("unimplemented")
}

// CreateChannel implements plugin.API.
func (*MattermostMockAPI) CreateChannel(channel *model.Channel) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// CreateChannelSidebarCategory implements plugin.API.
func (*MattermostMockAPI) CreateChannelSidebarCategory(userID string, teamID string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.AppError) {
	panic("unimplemented")
}

// CreateCommand implements plugin.API.
func (*MattermostMockAPI) CreateCommand(cmd *model.Command) (*model.Command, error) {
	panic("unimplemented")
}

// CreateOAuthApp implements plugin.API.
func (*MattermostMockAPI) CreateOAuthApp(app *model.OAuthApp) (*model.OAuthApp, *model.AppError) {
	panic("unimplemented")
}

// CreatePost implements plugin.API.
func (*MattermostMockAPI) CreatePost(post *model.Post) (*model.Post, *model.AppError) {
	panic("unimplemented")
}

// CreateSession implements plugin.API.
func (*MattermostMockAPI) CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	panic("unimplemented")
}

// CreateTeam implements plugin.API.
func (*MattermostMockAPI) CreateTeam(team *model.Team) (*model.Team, *model.AppError) {
	panic("unimplemented")
}

// CreateTeamMember implements plugin.API.
func (*MattermostMockAPI) CreateTeamMember(teamID string, userID string) (*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// CreateTeamMembers implements plugin.API.
func (*MattermostMockAPI) CreateTeamMembers(teamID string, userIds []string, requestorId string) ([]*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// CreateTeamMembersGracefully implements plugin.API.
func (*MattermostMockAPI) CreateTeamMembersGracefully(teamID string, userIds []string, requestorId string) ([]*model.TeamMemberWithError, *model.AppError) {
	panic("unimplemented")
}

// CreateUploadSession implements plugin.API.
func (*MattermostMockAPI) CreateUploadSession(us *model.UploadSession) (*model.UploadSession, error) {
	panic("unimplemented")
}

// CreateUser implements plugin.API.
func (*MattermostMockAPI) CreateUser(user *model.User) (*model.User, *model.AppError) {
	panic("unimplemented")
}

// CreateUserAccessToken implements plugin.API.
func (*MattermostMockAPI) CreateUserAccessToken(token *model.UserAccessToken) (*model.UserAccessToken, *model.AppError) {
	panic("unimplemented")
}

// DeleteChannel implements plugin.API.
func (*MattermostMockAPI) DeleteChannel(channelId string) *model.AppError {
	panic("unimplemented")
}

// DeleteChannelMember implements plugin.API.
func (*MattermostMockAPI) DeleteChannelMember(channelId string, userID string) *model.AppError {
	panic("unimplemented")
}

// DeleteCommand implements plugin.API.
func (*MattermostMockAPI) DeleteCommand(commandID string) error {
	panic("unimplemented")
}

// DeleteEphemeralPost implements plugin.API.
func (*MattermostMockAPI) DeleteEphemeralPost(userID string, postId string) {
	panic("unimplemented")
}

// DeleteOAuthApp implements plugin.API.
func (*MattermostMockAPI) DeleteOAuthApp(appID string) *model.AppError {
	panic("unimplemented")
}

// DeletePost implements plugin.API.
func (*MattermostMockAPI) DeletePost(postId string) *model.AppError {
	panic("unimplemented")
}

// DeletePreferencesForUser implements plugin.API.
func (*MattermostMockAPI) DeletePreferencesForUser(userID string, preferences []model.Preference) *model.AppError {
	panic("unimplemented")
}

// DeleteTeam implements plugin.API.
func (*MattermostMockAPI) DeleteTeam(teamID string) *model.AppError {
	panic("unimplemented")
}

// DeleteTeamMember implements plugin.API.
func (*MattermostMockAPI) DeleteTeamMember(teamID string, userID string, requestorId string) *model.AppError {
	panic("unimplemented")
}

// DeleteUser implements plugin.API.
func (*MattermostMockAPI) DeleteUser(userID string) *model.AppError {
	panic("unimplemented")
}

// DisablePlugin implements plugin.API.
func (*MattermostMockAPI) DisablePlugin(id string) *model.AppError {
	panic("unimplemented")
}

// EnablePlugin implements plugin.API.
func (*MattermostMockAPI) EnablePlugin(id string) *model.AppError {
	panic("unimplemented")
}

// EnsureBotUser implements plugin.API.
func (*MattermostMockAPI) EnsureBotUser(bot *model.Bot) (string, error) {
	panic("unimplemented")
}

// ExecuteSlashCommand implements plugin.API.
func (*MattermostMockAPI) ExecuteSlashCommand(commandArgs *model.CommandArgs) (*model.CommandResponse, error) {
	panic("unimplemented")
}

// ExtendSessionExpiry implements plugin.API.
func (*MattermostMockAPI) ExtendSessionExpiry(sessionID string, newExpiry int64) *model.AppError {
	panic("unimplemented")
}

// GetBot implements plugin.API.
func (*MattermostMockAPI) GetBot(botUserId string, includeDeleted bool) (*model.Bot, *model.AppError) {
	panic("unimplemented")
}

// GetBots implements plugin.API.
func (*MattermostMockAPI) GetBots(options *model.BotGetOptions) ([]*model.Bot, *model.AppError) {
	panic("unimplemented")
}

// GetBundlePath implements plugin.API.
func (*MattermostMockAPI) GetBundlePath() (string, error) {
	panic("unimplemented")
}

// GetChannel implements plugin.API.
func (*MattermostMockAPI) GetChannel(channelId string) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetChannelByName implements plugin.API.
func (*MattermostMockAPI) GetChannelByName(teamID string, name string, includeDeleted bool) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetChannelByNameForTeamName implements plugin.API.
func (*MattermostMockAPI) GetChannelByNameForTeamName(teamName string, channelName string, includeDeleted bool) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetChannelMember implements plugin.API.
func (*MattermostMockAPI) GetChannelMember(channelId string, userID string) (*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// GetChannelMembers implements plugin.API.
func (*MattermostMockAPI) GetChannelMembers(channelId string, page int, perPage int) (model.ChannelMembers, *model.AppError) {
	panic("unimplemented")
}

// GetChannelMembersByIds implements plugin.API.
func (*MattermostMockAPI) GetChannelMembersByIds(channelId string, userIds []string) (model.ChannelMembers, *model.AppError) {
	panic("unimplemented")
}

// GetChannelMembersForUser implements plugin.API.
func (*MattermostMockAPI) GetChannelMembersForUser(teamID string, userID string, page int, perPage int) ([]*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// GetChannelSidebarCategories implements plugin.API.
func (*MattermostMockAPI) GetChannelSidebarCategories(userID string, teamID string) (*model.OrderedSidebarCategories, *model.AppError) {
	panic("unimplemented")
}

// GetChannelStats implements plugin.API.
func (*MattermostMockAPI) GetChannelStats(channelId string) (*model.ChannelStats, *model.AppError) {
	panic("unimplemented")
}

// GetChannelsForTeamForUser implements plugin.API.
func (*MattermostMockAPI) GetChannelsForTeamForUser(teamID string, userID string, includeDeleted bool) ([]*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetCloudLimits implements plugin.API.
func (*MattermostMockAPI) GetCloudLimits() (*model.ProductLimits, error) {
	panic("unimplemented")
}

// GetCommand implements plugin.API.
func (*MattermostMockAPI) GetCommand(commandID string) (*model.Command, error) {
	panic("unimplemented")
}

// GetConfig implements plugin.API.
func (*MattermostMockAPI) GetConfig() *model.Config {
	panic("unimplemented")
}

// GetDiagnosticId implements plugin.API.
func (*MattermostMockAPI) GetDiagnosticId() string {
	panic("unimplemented")
}

// GetDirectChannel implements plugin.API.
func (*MattermostMockAPI) GetDirectChannel(userId1 string, userId2 string) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetEmoji implements plugin.API.
func (*MattermostMockAPI) GetEmoji(emojiId string) (*model.Emoji, *model.AppError) {
	panic("unimplemented")
}

// GetEmojiByName implements plugin.API.
func (*MattermostMockAPI) GetEmojiByName(name string) (*model.Emoji, *model.AppError) {
	panic("unimplemented")
}

// GetEmojiImage implements plugin.API.
func (*MattermostMockAPI) GetEmojiImage(emojiId string) ([]byte, string, *model.AppError) {
	panic("unimplemented")
}

// GetEmojiList implements plugin.API.
func (*MattermostMockAPI) GetEmojiList(sortBy string, page int, perPage int) ([]*model.Emoji, *model.AppError) {
	panic("unimplemented")
}

// GetFile implements plugin.API.
func (*MattermostMockAPI) GetFile(fileId string) ([]byte, *model.AppError) {
	panic("unimplemented")
}

// GetFileInfo implements plugin.API.
func (*MattermostMockAPI) GetFileInfo(fileId string) (*model.FileInfo, *model.AppError) {
	panic("unimplemented")
}

// GetFileInfos implements plugin.API.
func (*MattermostMockAPI) GetFileInfos(page int, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, *model.AppError) {
	panic("unimplemented")
}

// GetFileLink implements plugin.API.
func (*MattermostMockAPI) GetFileLink(fileId string) (string, *model.AppError) {
	panic("unimplemented")
}

// GetGroup implements plugin.API.
func (*MattermostMockAPI) GetGroup(groupId string) (*model.Group, *model.AppError) {
	panic("unimplemented")
}

// GetGroupByName implements plugin.API.
func (*MattermostMockAPI) GetGroupByName(name string) (*model.Group, *model.AppError) {
	panic("unimplemented")
}

// GetGroupChannel implements plugin.API.
func (*MattermostMockAPI) GetGroupChannel(userIds []string) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetGroupMemberUsers implements plugin.API.
func (*MattermostMockAPI) GetGroupMemberUsers(groupID string, page int, perPage int) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetGroupsBySource implements plugin.API.
func (*MattermostMockAPI) GetGroupsBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {
	panic("unimplemented")
}

// GetGroupsForUser implements plugin.API.
func (*MattermostMockAPI) GetGroupsForUser(userID string) ([]*model.Group, *model.AppError) {
	panic("unimplemented")
}

// GetLDAPUserAttributes implements plugin.API.
func (*MattermostMockAPI) GetLDAPUserAttributes(userID string, attributes []string) (map[string]string, *model.AppError) {
	panic("unimplemented")
}

// GetLicense implements plugin.API.
func (*MattermostMockAPI) GetLicense() *model.License {
	panic("unimplemented")
}

// GetOAuthApp implements plugin.API.
func (*MattermostMockAPI) GetOAuthApp(appID string) (*model.OAuthApp, *model.AppError) {
	panic("unimplemented")
}

// GetPluginConfig implements plugin.API.
func (*MattermostMockAPI) GetPluginConfig() map[string]any {
	panic("unimplemented")
}

// GetPluginStatus implements plugin.API.
func (*MattermostMockAPI) GetPluginStatus(id string) (*model.PluginStatus, *model.AppError) {
	panic("unimplemented")
}

// GetPlugins implements plugin.API.
func (*MattermostMockAPI) GetPlugins() ([]*model.Manifest, *model.AppError) {
	panic("unimplemented")
}

// GetPost implements plugin.API.
func (*MattermostMockAPI) GetPost(postId string) (*model.Post, *model.AppError) {
	panic("unimplemented")
}

// GetPostThread implements plugin.API.
func (*MattermostMockAPI) GetPostThread(postId string) (*model.PostList, *model.AppError) {
	panic("unimplemented")
}

// GetPostsAfter implements plugin.API.
func (*MattermostMockAPI) GetPostsAfter(channelId string, postId string, page int, perPage int) (*model.PostList, *model.AppError) {
	panic("unimplemented")
}

// GetPostsBefore implements plugin.API.
func (*MattermostMockAPI) GetPostsBefore(channelId string, postId string, page int, perPage int) (*model.PostList, *model.AppError) {
	panic("unimplemented")
}

// GetPostsForChannel implements plugin.API.
func (*MattermostMockAPI) GetPostsForChannel(channelId string, page int, perPage int) (*model.PostList, *model.AppError) {
	panic("unimplemented")
}

// GetPostsSince implements plugin.API.
func (*MattermostMockAPI) GetPostsSince(channelId string, time int64) (*model.PostList, *model.AppError) {
	panic("unimplemented")
}

// GetPreferenceForUser implements plugin.API.
func (*MattermostMockAPI) GetPreferenceForUser(userID string, category string, name string) (model.Preference, *model.AppError) {
	panic("unimplemented")
}

// GetPreferencesForUser implements plugin.API.
func (*MattermostMockAPI) GetPreferencesForUser(userID string) ([]model.Preference, *model.AppError) {
	panic("unimplemented")
}

// GetProfileImage implements plugin.API.
func (*MattermostMockAPI) GetProfileImage(userID string) ([]byte, *model.AppError) {
	panic("unimplemented")
}

// GetPublicChannelsForTeam implements plugin.API.
func (*MattermostMockAPI) GetPublicChannelsForTeam(teamID string, page int, perPage int) ([]*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// GetReactions implements plugin.API.
func (*MattermostMockAPI) GetReactions(postId string) ([]*model.Reaction, *model.AppError) {
	panic("unimplemented")
}

// GetServerVersion implements plugin.API.
func (*MattermostMockAPI) GetServerVersion() string {
	panic("unimplemented")
}

// GetSession implements plugin.API.
func (*MattermostMockAPI) GetSession(sessionID string) (*model.Session, *model.AppError) {
	panic("unimplemented")
}

// GetSystemInstallDate implements plugin.API.
func (*MattermostMockAPI) GetSystemInstallDate() (int64, *model.AppError) {
	panic("unimplemented")
}

// GetTeam implements plugin.API.
func (*MattermostMockAPI) GetTeam(teamID string) (*model.Team, *model.AppError) {
	panic("unimplemented")
}

// GetTeamByName implements plugin.API.
func (*MattermostMockAPI) GetTeamByName(name string) (*model.Team, *model.AppError) {
	panic("unimplemented")
}

// GetTeamIcon implements plugin.API.
func (*MattermostMockAPI) GetTeamIcon(teamID string) ([]byte, *model.AppError) {
	panic("unimplemented")
}

// GetTeamMember implements plugin.API.
func (*MattermostMockAPI) GetTeamMember(teamID string, userID string) (*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// GetTeamMembers implements plugin.API.
func (*MattermostMockAPI) GetTeamMembers(teamID string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// GetTeamMembersForUser implements plugin.API.
func (*MattermostMockAPI) GetTeamMembersForUser(userID string, page int, perPage int) ([]*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// GetTeamStats implements plugin.API.
func (*MattermostMockAPI) GetTeamStats(teamID string) (*model.TeamStats, *model.AppError) {
	panic("unimplemented")
}

// GetTeams implements plugin.API.
func (*MattermostMockAPI) GetTeams() ([]*model.Team, *model.AppError) {
	panic("unimplemented")
}

// GetTeamsForUser implements plugin.API.
func (*MattermostMockAPI) GetTeamsForUser(userID string) ([]*model.Team, *model.AppError) {
	panic("unimplemented")
}

// GetTeamsUnreadForUser implements plugin.API.
func (*MattermostMockAPI) GetTeamsUnreadForUser(userID string) ([]*model.TeamUnread, *model.AppError) {
	panic("unimplemented")
}

// GetTelemetryId implements plugin.API.
func (*MattermostMockAPI) GetTelemetryId() string {
	panic("unimplemented")
}

// GetUnsanitizedConfig implements plugin.API.
func (*MattermostMockAPI) GetUnsanitizedConfig() *model.Config {
	panic("unimplemented")
}

// GetUploadSession implements plugin.API.
func (*MattermostMockAPI) GetUploadSession(uploadID string) (*model.UploadSession, error) {
	panic("unimplemented")
}

// GetUser implements plugin.API.
func (*MattermostMockAPI) GetUser(userID string) (*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUserByEmail implements plugin.API.
func (*MattermostMockAPI) GetUserByEmail(email string) (*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUserByUsername implements plugin.API.
func (*MattermostMockAPI) GetUserByUsername(name string) (*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUserStatus implements plugin.API.
func (*MattermostMockAPI) GetUserStatus(userID string) (*model.Status, *model.AppError) {
	panic("unimplemented")
}

// GetUserStatusesByIds implements plugin.API.
func (*MattermostMockAPI) GetUserStatusesByIds(userIds []string) ([]*model.Status, *model.AppError) {
	panic("unimplemented")
}

// GetUsers implements plugin.API.
func (*MattermostMockAPI) GetUsers(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUsersByUsernames implements plugin.API.
func (*MattermostMockAPI) GetUsersByUsernames(usernames []string) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUsersInChannel implements plugin.API.
func (*MattermostMockAPI) GetUsersInChannel(channelID string, sortBy string, page int, perPage int) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// GetUsersInTeam implements plugin.API.
func (*MattermostMockAPI) GetUsersInTeam(teamID string, page int, perPage int) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// HasPermissionTo implements plugin.API.
func (*MattermostMockAPI) HasPermissionTo(userID string, permission *model.Permission) bool {
	panic("unimplemented")
}

// HasPermissionToChannel implements plugin.API.
func (*MattermostMockAPI) HasPermissionToChannel(userID string, channelId string, permission *model.Permission) bool {
	panic("unimplemented")
}

// HasPermissionToTeam implements plugin.API.
func (*MattermostMockAPI) HasPermissionToTeam(userID string, teamID string, permission *model.Permission) bool {
	panic("unimplemented")
}

// InstallPlugin implements plugin.API.
func (*MattermostMockAPI) InstallPlugin(file io.Reader, replace bool) (*model.Manifest, *model.AppError) {
	panic("unimplemented")
}

// InviteRemoteToChannel implements plugin.API.
func (*MattermostMockAPI) InviteRemoteToChannel(channelID string, remoteID string, userID string) error {
	panic("unimplemented")
}

// IsEnterpriseReady implements plugin.API.
func (*MattermostMockAPI) IsEnterpriseReady() bool {
	panic("unimplemented")
}

// KVCompareAndDelete implements plugin.API.
func (*MattermostMockAPI) KVCompareAndDelete(key string, oldValue []byte) (bool, *model.AppError) {
	panic("unimplemented")
}

// KVCompareAndSet implements plugin.API.
func (*MattermostMockAPI) KVCompareAndSet(key string, oldValue []byte, newValue []byte) (bool, *model.AppError) {
	panic("unimplemented")
}

// KVDelete implements plugin.API.
func (*MattermostMockAPI) KVDelete(key string) *model.AppError {
	panic("unimplemented")
}

// KVDeleteAll implements plugin.API.
func (*MattermostMockAPI) KVDeleteAll() *model.AppError {
	panic("unimplemented")
}

// KVGet implements plugin.API.
func (*MattermostMockAPI) KVGet(key string) ([]byte, *model.AppError) {
	panic("unimplemented")
}

// KVList implements plugin.API.
func (*MattermostMockAPI) KVList(page int, perPage int) ([]string, *model.AppError) {
	panic("unimplemented")
}

// KVSet implements plugin.API.
func (*MattermostMockAPI) KVSet(key string, value []byte) *model.AppError {
	panic("unimplemented")
}

// KVSetWithExpiry implements plugin.API.
func (*MattermostMockAPI) KVSetWithExpiry(key string, value []byte, expireInSeconds int64) *model.AppError {
	panic("unimplemented")
}

// KVSetWithOptions implements plugin.API.
func (*MattermostMockAPI) KVSetWithOptions(key string, value []byte, options model.PluginKVSetOptions) (bool, *model.AppError) {
	panic("unimplemented")
}

// ListBuiltInCommands implements plugin.API.
func (*MattermostMockAPI) ListBuiltInCommands() ([]*model.Command, error) {
	panic("unimplemented")
}

// ListCommands implements plugin.API.
func (*MattermostMockAPI) ListCommands(teamID string) ([]*model.Command, error) {
	panic("unimplemented")
}

// ListCustomCommands implements plugin.API.
func (*MattermostMockAPI) ListCustomCommands(teamID string) ([]*model.Command, error) {
	panic("unimplemented")
}

// ListPluginCommands implements plugin.API.
func (*MattermostMockAPI) ListPluginCommands(teamID string) ([]*model.Command, error) {
	panic("unimplemented")
}

// LoadPluginConfiguration implements plugin.API.
func (*MattermostMockAPI) LoadPluginConfiguration(dest any) error {
	panic("unimplemented")
}

// OpenInteractiveDialog implements plugin.API.
func (*MattermostMockAPI) OpenInteractiveDialog(dialog model.OpenDialogRequest) *model.AppError {
	panic("unimplemented")
}

// PatchBot implements plugin.API.
func (*MattermostMockAPI) PatchBot(botUserId string, botPatch *model.BotPatch) (*model.Bot, *model.AppError) {
	panic("unimplemented")
}

// PatchChannelMembersNotifications implements plugin.API.
func (*MattermostMockAPI) PatchChannelMembersNotifications(members []*model.ChannelMemberIdentifier, notifyProps map[string]string) *model.AppError {
	panic("unimplemented")
}

// PermanentDeleteBot implements plugin.API.
func (*MattermostMockAPI) PermanentDeleteBot(botUserId string) *model.AppError {
	panic("unimplemented")
}

// PluginHTTP implements plugin.API.
func (*MattermostMockAPI) PluginHTTP(request *http.Request) *http.Response {
	panic("unimplemented")
}

// PublishPluginClusterEvent implements plugin.API.
func (*MattermostMockAPI) PublishPluginClusterEvent(ev model.PluginClusterEvent, opts model.PluginClusterEventSendOptions) error {
	panic("unimplemented")
}

// PublishUserTyping implements plugin.API.
func (*MattermostMockAPI) PublishUserTyping(userID string, channelId string, parentId string) *model.AppError {
	panic("unimplemented")
}

// PublishWebSocketEvent implements plugin.API.
func (*MattermostMockAPI) PublishWebSocketEvent(event string, payload map[string]any, broadcast *model.WebsocketBroadcast) {
	panic("unimplemented")
}

// ReadFile implements plugin.API.
func (*MattermostMockAPI) ReadFile(path string) ([]byte, *model.AppError) {
	panic("unimplemented")
}

// RegisterCollectionAndTopic implements plugin.API.
func (*MattermostMockAPI) RegisterCollectionAndTopic(collectionType string, topicType string) error {
	panic("unimplemented")
}

// RegisterCommand implements plugin.API.
func (*MattermostMockAPI) RegisterCommand(command *model.Command) error {
	panic("unimplemented")
}

// RegisterPluginForSharedChannels implements plugin.API.
func (*MattermostMockAPI) RegisterPluginForSharedChannels(opts model.RegisterPluginOpts) (remoteID string, err error) {
	panic("unimplemented")
}

// RemovePlugin implements plugin.API.
func (*MattermostMockAPI) RemovePlugin(id string) *model.AppError {
	panic("unimplemented")
}

// RemoveReaction implements plugin.API.
func (*MattermostMockAPI) RemoveReaction(reaction *model.Reaction) *model.AppError {
	panic("unimplemented")
}

// RemoveTeamIcon implements plugin.API.
func (*MattermostMockAPI) RemoveTeamIcon(teamID string) *model.AppError {
	panic("unimplemented")
}

// RemoveUserCustomStatus implements plugin.API.
func (*MattermostMockAPI) RemoveUserCustomStatus(userID string) *model.AppError {
	panic("unimplemented")
}

// RequestTrialLicense implements plugin.API.
func (*MattermostMockAPI) RequestTrialLicense(requesterID string, users int, termsAccepted bool, receiveEmailsAccepted bool) *model.AppError {
	panic("unimplemented")
}

// RevokeSession implements plugin.API.
func (*MattermostMockAPI) RevokeSession(sessionID string) *model.AppError {
	panic("unimplemented")
}

// RevokeUserAccessToken implements plugin.API.
func (*MattermostMockAPI) RevokeUserAccessToken(tokenID string) *model.AppError {
	panic("unimplemented")
}

// RolesGrantPermission implements plugin.API.
func (*MattermostMockAPI) RolesGrantPermission(roleNames []string, permissionId string) bool {
	panic("unimplemented")
}

// SaveConfig implements plugin.API.
func (*MattermostMockAPI) SaveConfig(config *model.Config) *model.AppError {
	panic("unimplemented")
}

// SavePluginConfig implements plugin.API.
func (*MattermostMockAPI) SavePluginConfig(config map[string]any) *model.AppError {
	panic("unimplemented")
}

// SearchChannels implements plugin.API.
func (*MattermostMockAPI) SearchChannels(teamID string, term string) ([]*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// SearchPostsInTeam implements plugin.API.
func (*MattermostMockAPI) SearchPostsInTeam(teamID string, paramsList []*model.SearchParams) ([]*model.Post, *model.AppError) {
	panic("unimplemented")
}

// SearchPostsInTeamForUser implements plugin.API.
func (*MattermostMockAPI) SearchPostsInTeamForUser(teamID string, userID string, searchParams model.SearchParameter) (*model.PostSearchResults, *model.AppError) {
	panic("unimplemented")
}

// SearchTeams implements plugin.API.
func (*MattermostMockAPI) SearchTeams(term string) ([]*model.Team, *model.AppError) {
	panic("unimplemented")
}

// SearchUsers implements plugin.API.
func (*MattermostMockAPI) SearchUsers(search *model.UserSearch) ([]*model.User, *model.AppError) {
	panic("unimplemented")
}

// SendEphemeralPost implements plugin.API.
func (*MattermostMockAPI) SendEphemeralPost(userID string, post *model.Post) *model.Post {
	panic("unimplemented")
}

// SendMail implements plugin.API.
func (*MattermostMockAPI) SendMail(to string, subject string, htmlBody string) *model.AppError {
	panic("unimplemented")
}

// SendPushNotification implements plugin.API.
func (*MattermostMockAPI) SendPushNotification(notification *model.PushNotification, userID string) *model.AppError {
	panic("unimplemented")
}

// SetFileSearchableContent implements plugin.API.
func (*MattermostMockAPI) SetFileSearchableContent(fileID string, content string) *model.AppError {
	panic("unimplemented")
}

// SetProfileImage implements plugin.API.
func (*MattermostMockAPI) SetProfileImage(userID string, data []byte) *model.AppError {
	panic("unimplemented")
}

// SetTeamIcon implements plugin.API.
func (*MattermostMockAPI) SetTeamIcon(teamID string, data []byte) *model.AppError {
	panic("unimplemented")
}

// SetUserStatusTimedDND implements plugin.API.
func (*MattermostMockAPI) SetUserStatusTimedDND(userId string, endtime int64) (*model.Status, *model.AppError) {
	panic("unimplemented")
}

// ShareChannel implements plugin.API.
func (*MattermostMockAPI) ShareChannel(sc *model.SharedChannel) (*model.SharedChannel, error) {
	panic("unimplemented")
}

// SyncSharedChannel implements plugin.API.
func (*MattermostMockAPI) SyncSharedChannel(channelID string) error {
	panic("unimplemented")
}

// UninviteRemoteFromChannel implements plugin.API.
func (*MattermostMockAPI) UninviteRemoteFromChannel(channelID string, remoteID string) error {
	panic("unimplemented")
}

// UnregisterCommand implements plugin.API.
func (*MattermostMockAPI) UnregisterCommand(teamID string, trigger string) error {
	panic("unimplemented")
}

// UnregisterPluginForSharedChannels implements plugin.API.
func (*MattermostMockAPI) UnregisterPluginForSharedChannels(pluginID string) error {
	panic("unimplemented")
}

// UnshareChannel implements plugin.API.
func (*MattermostMockAPI) UnshareChannel(channelID string) (unshared bool, err error) {
	panic("unimplemented")
}

// UpdateBotActive implements plugin.API.
func (*MattermostMockAPI) UpdateBotActive(botUserId string, active bool) (*model.Bot, *model.AppError) {
	panic("unimplemented")
}

// UpdateChannel implements plugin.API.
func (*MattermostMockAPI) UpdateChannel(channel *model.Channel) (*model.Channel, *model.AppError) {
	panic("unimplemented")
}

// UpdateChannelMemberNotifications implements plugin.API.
func (*MattermostMockAPI) UpdateChannelMemberNotifications(channelId string, userID string, notifications map[string]string) (*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// UpdateChannelMemberRoles implements plugin.API.
func (*MattermostMockAPI) UpdateChannelMemberRoles(channelId string, userID string, newRoles string) (*model.ChannelMember, *model.AppError) {
	panic("unimplemented")
}

// UpdateChannelSidebarCategories implements plugin.API.
func (*MattermostMockAPI) UpdateChannelSidebarCategories(userID string, teamID string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, *model.AppError) {
	panic("unimplemented")
}

// UpdateCommand implements plugin.API.
func (*MattermostMockAPI) UpdateCommand(commandID string, updatedCmd *model.Command) (*model.Command, error) {
	panic("unimplemented")
}

// UpdateEphemeralPost implements plugin.API.
func (*MattermostMockAPI) UpdateEphemeralPost(userID string, post *model.Post) *model.Post {
	panic("unimplemented")
}

// UpdateOAuthApp implements plugin.API.
func (*MattermostMockAPI) UpdateOAuthApp(app *model.OAuthApp) (*model.OAuthApp, *model.AppError) {
	panic("unimplemented")
}

// UpdatePost implements plugin.API.
func (*MattermostMockAPI) UpdatePost(post *model.Post) (*model.Post, *model.AppError) {
	panic("unimplemented")
}

// UpdatePreferencesForUser implements plugin.API.
func (*MattermostMockAPI) UpdatePreferencesForUser(userID string, preferences []model.Preference) *model.AppError {
	panic("unimplemented")
}

// UpdateSharedChannel implements plugin.API.
func (*MattermostMockAPI) UpdateSharedChannel(sc *model.SharedChannel) (*model.SharedChannel, error) {
	panic("unimplemented")
}

// UpdateSharedChannelCursor implements plugin.API.
func (*MattermostMockAPI) UpdateSharedChannelCursor(channelID string, remoteID string, cusror model.GetPostsSinceForSyncCursor) error {
	panic("unimplemented")
}

// UpdateTeam implements plugin.API.
func (*MattermostMockAPI) UpdateTeam(team *model.Team) (*model.Team, *model.AppError) {
	panic("unimplemented")
}

// UpdateTeamMemberRoles implements plugin.API.
func (*MattermostMockAPI) UpdateTeamMemberRoles(teamID string, userID string, newRoles string) (*model.TeamMember, *model.AppError) {
	panic("unimplemented")
}

// UpdateUser implements plugin.API.
func (*MattermostMockAPI) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	panic("unimplemented")
}

// UpdateUserActive implements plugin.API.
func (*MattermostMockAPI) UpdateUserActive(userID string, active bool) *model.AppError {
	panic("unimplemented")
}

// UpdateUserAuth implements plugin.API.
func (*MattermostMockAPI) UpdateUserAuth(userID string, userAuth *model.UserAuth) (*model.UserAuth, *model.AppError) {
	panic("unimplemented")
}

// UpdateUserCustomStatus implements plugin.API.
func (*MattermostMockAPI) UpdateUserCustomStatus(userID string, customStatus *model.CustomStatus) *model.AppError {
	panic("unimplemented")
}

// UpdateUserStatus implements plugin.API.
func (*MattermostMockAPI) UpdateUserStatus(userID string, status string) (*model.Status, *model.AppError) {
	panic("unimplemented")
}

// UploadData implements plugin.API.
func (*MattermostMockAPI) UploadData(us *model.UploadSession, rd io.Reader) (*model.FileInfo, error) {
	panic("unimplemented")
}

// UploadFile implements plugin.API.
func (*MattermostMockAPI) UploadFile(data []byte, channelId string, filename string) (*model.FileInfo, *model.AppError) {
	panic("unimplemented")
}

/// LOGS

// LogDebug implements plugin.API.
func (*MattermostMockAPI) LogDebug(msg string, keyValuePairs ...any) {
	println("LogDebug: "+msg, keyValuePairs)
}

// LogError implements plugin.API.
func (*MattermostMockAPI) LogError(msg string, keyValuePairs ...any) {
	println("LogError: "+msg, keyValuePairs)
}

// LogInfo implements plugin.API.
func (*MattermostMockAPI) LogInfo(msg string, keyValuePairs ...any) {
	println("LogInfo: "+msg, keyValuePairs)
}

// LogWarn implements plugin.API.
func (*MattermostMockAPI) LogWarn(msg string, keyValuePairs ...any) {
	println("LogWarn: "+msg, keyValuePairs)
}

/// END OF LOGS
