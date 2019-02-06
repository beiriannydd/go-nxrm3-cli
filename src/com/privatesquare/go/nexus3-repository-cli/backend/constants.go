package backend

const (
	ConfFileName = "nexus3-repository-cli.json"

	// API Extensions
	apiBase        = "service/rest"
	scriptAPI      = "v1/script"
	repositoryPath = "v1/repositories"

	// Script Path
	scriptBasePath = "./scripts/groovy"

	// Error Strings
	jsonMarshalError   = "JSON Marshal Error"
	jsonUnmarshalError = "JSON Unmarshal Error"

	// CLI flags and usage
	ConfCommandFlag       = "configure"
	ConfCommandUsage      = "Set nexus connection details"
	ScriptCommandFlag     = "script"
	ScriptCommandUsage    = "Nexus script operations"
	RepoCommandFlag       = "repo"
	RepoCommandUsage      = "Nexus repository operations"
	SelectorCommandFlag   = "selector"
	SelectorCommandUsage  = "Nexus content selector operations"
	PrivilegeCommandFlag  = "privilege"
	PrivilegeCommandUsage = "Nexus privilege operations"
	RoleCommandFlag       = "role"
	RoleCommandUsage      = "Nexus role operations"

	//configure
	NexusURLFlag       = "nexus-url"
	NexusURLUsage      = "Nexus 3 server URL. (Required)"
	NexusUsernameFlag  = "username"
	NexusUsernameUsage = "Nexus 3 server login user. (Required)"
	NexusPasswordFlag  = "password"
	NexusPasswordUsage = "Nexus 3 server login password. (Required)"

	connDetailsSuccessInfo = "Connection details were stored successfully in the file ./%s\n"
	connDetailsEmptyInfo   = "Server connection details are not set...First Run %q to set the connection details\n"

	//common
	TaskFlag     = "task"
	SkipTlsFlag  = "skip-tls"
	SkipTlsUsage = "Skip TLS verification for the nexus server instance"
	DebugFlag    = "debug"
	DebugUsage   = "Set Default for more information on the nexus script execution"
	VerboseFlag  = "verbose"
	VerboseUsage = "Set Verbose for detailed http request and response logs"

	TaskEmptyInfo    = "You need to select a task to be performed. Available tasks : %+q\n"
	TaskNotValidInfo = "%q is not a valid %s task. Available tasks : %+q\n"
	setVerboseInfo   = "There was an error calling the function. Set verbose flag for more information"

	//script
	ScriptTaskUsage = "Script Task (Required)  (For all tasks the script(s) should exist under the path ./scripts/groovy)\n\n" +
		"    list 	    List all the scripts available in Nexus. (Optional: name) If script name is passed the contents of the script will be printed\n" +
		"    add  	    Add a new script to nexus. (Required: name)\n" +
		"    update 	    Update a script that is available in nexus. (Required: name)\n" +
		"    add-or-update   Add or Update a script in nexus. (Required: name)\n" +
		"    delete          Delete a script from nexus. (Required: name)\n" +
		"    run 	    Run/Execute a script in nexus. (Required: name)(Optional: payload)\n"

	ScriptNameFlag     = "name"
	ScriptNameUsage    = "Name of the script to be executed in nexus"
	ScriptPayloadFlag  = "payload"
	ScriptPayloadUsage = "Arguments to be passed to a nexus script can be sent as a payload during script execution"

	scriptNameRequiredInfo = "name is a required parameter"

	//script name
	getPrivilegesScript   = "get-privileges"
	createPrivilegeScript = "create-privilege"
	updatePrivilegeScript = "update-privilege"
	deletePrivilegeScript = "delete-privilege"

	//repo
	RepoTaskUsage = "Repo Task (Required)\n\n" +
		"    list   		List all the repositories in nexus.\n" +
		"			(Optional - repo-name) If repo-name is passed the list command will get the details of the repository.\n" +
		"			(Optional - repo-format) If repo-format is passed the list command will list the repositories as per the format\n" +
		"    create-hosted	Create a hosted repository in nexus. (Required - repo-name and repo-format)\n" +
		"    create-proxy	Create a proxy repository in nexus. (Required - repo-name, repo-format and remote-url ) (Optional - proxy-user and proxy-pass)\n" +
		"    create-group	Create a group repository in nexus. (Required - repo-name,repo-format and repo-members)\n" +
		"    add-group-members	Add new members to a existing group repository. (Required - repo-name,repo-format and repo-members)\n" +
		"    delete		Delete a repository from nexus\n\n" +
		"    If you are creating a docker repository it is necessary to also provide either a docker-http-port or a docker-https-port or both.\n"

	RepoNameFlag         = "repo-name"
	RepoNameUsage        = "Nexus repository name"
	RepoFormatFlag       = "repo-format"
	RepoFormatUsage      = "Repository format. Available formats : %+q"
	RemoteURLFlag        = "remote-url"
	RemoteURLUsage       = "Remote URL to be proxied in nexus"
	RepoMembersFlag      = "repo-members"
	RepoMembersUsage     = "Comma-separated repository names that should be added to a group repo"
	ProxyUserFlag        = "proxy-user"
	ProxyUserUsage       = "Username for accessing the proxy repository"
	ProxyPassFlag        = "proxy-pass"
	ProxyPassUsage       = "Password for accessing the proxy repository"
	DockerHttpPortFlag   = "docker-http-port"
	DockerHttpPortUsage  = "Docker HTTP port"
	DockerHttpsPortFlag  = "docker-https-port"
	DockerHttpsPortUsage = "Docker HTTPs port"
	BlobStoreNameFlag    = "blob-store-name"
	BlobStoreNameUsage   = "Blob store name"
	ReleaseFlag          = "releases"
	ReleaseUsage         = "Set this flag to create a releases repository"

	RepoFormatNotValidInfo = "%q is not a valid repository format. Available repository formats are : %v\n"
	repoNameRequiredInfo   = "repo-name is a required parameter"
	repoFormatRequiredInfo = "repo-format is a required parameter"
	hostedRepoRequiredInfo = "repo-name and repo-format are required parameters to create a hosted repository"
	proxyRepoRequiredInfo  = "repo-name, repo-format and remote-url are required parameters to create a proxy repository"
	groupRequiredInfo      = "repo-name, repo-format and repo-members are required parameters"
	dockerPortsInfo        = "You need to specify either a http port or a https port or both for creating a docker repository"
	repositoryNotFoundInfo = "Repository %q was not found in nexus"

	//selector
	contentSelectorType = "csel"

	SelectorTaskUsage = "Selector Task (Required)\n\n" +
		"    list 	    List all the content selectors in nexus (Optional: name)\n" +
		"    create  	    Create a content selector in nexus (Required: name and expression) (Optional: description)\n" +
		"    update 	    Update the details of a content selector. (Required: name and expression) (Optional: description)\n" +
		"    delete          Delete a content selector (Required: name)\n"

	SelectorNameFlag        = "name"
	SelectorNameUsage       = "Content Selector name"
	SelectorDescFlag        = "description"
	SelectorDescUsage       = "Content Selector description"
	SelectorExpressionFlag  = "expression"
	SelectorExpressionUsage = "Pattern expression for the content selector"

	selectorNameRequiredInfo   = "name is a required parameter"
	createSelectorRequiredInfo = "name and expression are required parameters"
	createSelectorSuccessInfo  = "Content selector %q was created\n"
	updateSelectorSuccessInfo  = "Content selector %q was updated\n"
	deleteSelectorSuccessInfo  = "Content selector %q was deleted\n"
	selectorAlreadyExistsInfo  = "Content selector %q already exists in nexus\n"
	selectorNotFoundInfo       = "Content selector %q was not found in nexus\n"

	//privilege
	PrivilegeTaskUsage = "Privilege Task (Required)\n\n" +
		"    list 	    List all the privileges in nexus (Optional: name)\n" +
		"    create  	    Create a Privilege in nexus (Required: name, selector-name and repo-name) (Optional: description and action)\n" +
		"    update 	    Update the details of a Privilege. (Required: name and expression)\n" +
		"    delete          Delete a Privilege (Required: name)\n"

	PrivilegeNameFlag  = "name"
	PrivilegeNameUsage = "Privilege name"
	PSelectorNameFlag  = "selector-name"
	PRepoNameFlag      = "repo-name"
	ActionFlag         = "action"
	ActionUsage        = "Privilege Action. Available actions %+q"
	PrivilegeDescFlag  = "description"
	PrivilegeDescUsage = "Privilege description"

	privilegeNameRequiredInfo   = "name is a required parameter"
	privilegeNotFoundInfo       = "Privilege %q was not found in nexus\n"
	privilegeExistsInfo         = "Privilege %q already exists\n"
	createPrivilegeRequiredInfo = "name, selector-name and repo-name are required parameters"
	createPrivilegeSuccessInfo  = "Privilege %q is created"
	updatePrivilegeSuccessInfo  = "Privilege %q is updated"
	deletePrivilegeSuccessInfo  = "Privilege %q is deleted"
)
