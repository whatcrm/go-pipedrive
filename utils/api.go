package utils

const (
	BaseURL                            = "https://api.pipedrive.com/"
	DomainBaseURL                      = "https://%s.pipedrive.com/"
	WebhooksEndPoint                   = "v1/webhooks"
	TokenEndPoint                      = "https://oauth.pipedrive.com/oauth/token"
	ChannelEndPoint                    = "v1/channels"
	DeleteChannelEndPoint              = "v1/channels/%s"
	ReceiveMessageEndPoint             = "v1/channels/messages/receive"
	ConversationEndPoint               = "/conversations/"
	UserEndPoint                       = "v1/users"
	UserMeEndPoint                     = "/me"
	FilesEndPoint                      = "v1/files"
	OneFileEndpoint                    = "v1/files/%d"
	DownloadOneFileEndpoint            = "v1/files/%d/download"
	RemoteFileEndpoint                 = "v1/files/remote"
	RemoteLinkFileEndpoint             = "v1/files/remoteLink"
	DealsEndpoint                      = "v1/deals"
	OneDealEndpoint                    = "v1/deals/%s"
	LeadsEndpoint                      = "v1/leads"
	OneLeadEndpoint                    = "v1/leads/%s"
	LeadPermittedUsersEndpoint         = "v1/leads/%s/permittedUsers"
	LeadSearchURL                      = "api/v2/leads/search"
	LeadSourcesEndpoint                = "v1/leadSources"
	LeadLabelsEndpoint                 = "v1/leadLabels"
	OrganizationParticipantsEndpoint   = "v1/organizations/%d/persons"
	PersonsEndpoint                    = "v1/persons"
	PersonEndpoint                     = "v1/persons/%d"
	PipelinesEndpoint                  = "v1/pipelines"
	PipelineEndpoint                   = "v1/pipelines/%d"
	StagesEndpoint                     = "v1/stages"
	StagesFromPipelineEndpoint         = "v1/stages?pipeline_id=%d"
	ItemSearchV1Endpoint               = "v1/itemSearch"
	ItemSearchFieldV1Endpoint          = "v1/itemSearch/field"
	ItemSearchV2Endpoint               = "api/v2/itemSearch"
	ItemSearchFieldV2Endpoint          = "api/v2/itemSearch/field"
	CallLogsEndpoint                   = "v1/callLogs"
	CallLogRecordingEndpoint           = "v1/callLogs/%s/recordings"
	PersonFieldsEndPoint               = "v1/personFields"
	DealFieldsEndPoint                 = "v1/dealFields"
	ActivitiesEndpoint                 = "v1/activities"
	OneActivityEndpoint                = "v1/activities/%d"
	ActivitiesCollectionEndpoint       = "v1/activities/collection"
	ActivityFieldsEndpoint             = "v1/activityFields"
	ActivityTypesEndpoint              = "v1/activityTypes"
	BillingAddOnsEndpoint              = "v1/billing/subscriptions/addons"
	CurrenciesEndpoint                 = "v1/currencies"
	FiltersEndpoint                    = "v1/filters"
	FilterHelpersEndpoint              = "v1/filters/helpers"
	GoalsEndpoint                      = "v1/goals"
	FindGoalsEndpoint                  = "v1/goals/find"
	GoalResultsEndpoint                = "v1/goals/%s/results"
	MailMessageEndpoint                = "v1/mailbox/mailMessages"
	MailThreadEndpoint                 = "v1/mailbox/mailThreads"
	MailThreadMessagesEndpoint         = "v1/mailbox/mailThreads/%d/mailMessages"
	LinkUserProviderEndpoint           = "v1/meetings/userProviderLinks"
	DeleteUserProviderLinkEndpoint     = "v1/meetings/userProviderLinks/%s"
	NotesEndpoint                      = "v1/notes"
	SingleNoteEndpoint                 = "v1/notes/%d"
	CommentsEndpoint                   = "v1/notes/%d/comments"
	SingleCommentEndpoint              = "v1/notes/%d/comments/%s"
	NoteFieldsEndpoint                 = "v1/noteFields"
	OrganizationsEndpoint              = "v1/organizations"
	OrganizationSearchEndpoint         = "v1/organizations/search"
	OrganizationActivitiesEndpoint     = "v1/organizations/%d/activities"
	OrganizationDealsEndpoint          = "v1/organizations/%d/deals"
	OrganizationFollowersEndpoint      = "v1/organizations/%d/followers"
	OrganizationFilesEndpoint          = "v1/organizations/%d/files"
	OrganizationFlowEndpoint           = "v1/organizations/%d/flow"
	OrganizationMailMessagesEndpoint   = "v1/organizations/%d/mailMessages"
	OrganizationPersonsEndpoint        = "v1/organizations/%d/persons"
	OrganizationPermittedUsersEndpoint = "v1/organizations/%d/permittedUsers"
	OrganizationFieldsEndpoint         = "v1/organizationFields"
	ProductsEndpoint                   = "v1/products"
	OneProductEndpoint                 = "v1/products/%s"
	ProductDealsEndpoint               = "v1/products/%s/deals"
	ProductVariationsEndpoint          = "api/v2/products/%s/variations"
	ProductVariationEndpoint           = "api/v2/products/%s/variations/%s"
	ProductFieldsEndpoint              = "v1/productFields"
	OneProductFieldEndpoint            = "v1/productFields/%s"
)
