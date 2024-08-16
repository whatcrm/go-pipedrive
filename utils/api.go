package utils

const (
	BaseURL                           = "https://api.pipedrive.com/v1/"
	DomainBaseURL                     = "https://%s.pipedrive.com/"
	WebhooksEndPoint                  = "/webhooks?api_token=%s"
	DeleteWebhooksEndPoint            = "/webhooks:id?api_token=%s"
	TokenEndPoint                     = "https://oauth.pipedrive.com/oauth/token"
	ChannelEndPoint                   = "channels"
	DeleteChannelEndPoint             = "channels/%s"
	ReceiveMessageEndPoint            = "channels/messages/receive"
	ConversationEndPoint              = "/conversations/"
	UserEndPoint                      = "users"
	UserMeEndPoint                    = "/me"
	FilesEndPoint                     = "files"
	OneFileEndpoint                   = "files/%d"
	DownloadOneFileEndpoint           = "files/%d/download"
	RemoteFileEndpoint                = "files/remote"
	RemoteLinkFileEndpoint            = "files/remoteLink"
	DealsEndpoint                     = "deals"
	OneDealEndPoint                   = "deals/%d"
	DealFollowersEndpoint             = "deals/%d/followers"
	DealOneFollowerEndpoint           = "deals/%d/followers/%d"
	DealParticipantsEndpoint          = "deals/%d/participants"
	DealOneParticipantEndpoint        = "deals/%d/participants/%d"
	DealProductsEndpoint              = "deals/%d/products"
	DealOneProductEndpoint            = "deals/%d/products/%d"
	DealDiscountsEndpoint             = "deals/%d/discounts"
	DealOneDiscountEndpoint           = "deals/%d/discounts/%d"
	DealsMergeEndpoint                = "deals/%d/merge"
	DealsCollectionEndpoint           = "deals/collection"
	SearchDealsURL                    = "https://api.pipedrive.com/api/v2/deals/search"
	DealsSummaryEndpoint              = "deals/summary"
	DealsTimelineEndpoint             = "deals/timeline"
	DealActivitiesEndpoint            = "deals/%d/activities"
	DealFieldValuesEndpoint           = "deals/%d/changelog"
	DealFilesEndpoint                 = "deals/%d/files"
	DealUpdatesEndpoint               = "deals/%d/flow"
	DealParticipantsChangelogEndpoint = "deals/%d/participantsChangelog"
	DealMailMessagesEndpoint          = "deals/%d/mailMessages"
	DealPermittedUsersEndpoint        = "deals/%d/permittedUsers"
	DealPersonsEndpoint               = "deals/%d/persons"
	DealsProductsURL                  = "https://api.pipedrive.com/api/v2/deals/products"
	DealFieldUpdatesEndpoint          = "deals/%d/changelog"
	LeadsEndpoint                     = "leads"
	OneLeadEndpoint                   = "leads/%s"
	LeadPermittedUsersEndpoint        = "leads/%s/permittedUsers"
	LeadSearchURL                     = "https://api.pipedrive.com/api/v2/leads/search"
	LeadSourcesEndpoint               = "leadSources"
	LeadLabelsEndpoint                = "leadLabels"
	OrganizationParticipantsEndpoint  = "organizations/%d/persons"
	AddPersonEndpoint                 = "persons"
	PersonEndpoint                    = "persons/%d"
	PipelinesEndpoint                 = "pipelines"
	PipelineEndpoint                  = "pipelines/%d"
	StagesEndpoint                    = "stages"
	StagesFromPipelineEndpoint        = "stages?pipeline_id=%d"
	ItemSearchEndpoint                = "itemSearch"
)
