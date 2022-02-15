package messages

const (
	INVALID_URL                 string = "[skyflow] Interface: %s - URL %s is invalid"
	INITIALIZING_SKYFLOW_CLIENT string = "[skyflow] Interface: %s - Initializing skyflow client"
	VALIDATE_INIT_CONFIG        string = "[skyflow] Interface: %s - Validating init config"
	VALIDATE_RECORDS            string = "[skyflow] Interface: %s - Validating insert records"
	VALIDATE_DETOKENIZE_INPUT   string = "[skyflow] Interface: %s - Validating detokenize input"
	VALIDATE_GET_BY_ID_INPUT    string = "[skyflow] Interface: %s - Validating getByID input"
	VALIDATE_CONNECTION_CONFIG  string = "[skyflow] Interface: %s - Validating connection config"

	INSERTING_RECORDS         string = "[skyflow] Interface: %s - Inserting records into vault with id %s"
	INSERTING_RECORDS_SUCCESS string = "[skyflow] Interface: %s - Successfully inserted records into vault with id %s"
	INSERTING_RECORDS_FAILED  string = "[skyflow] Interface: %s - Failed inserting records into vault with id %s"
	DETOKENIZING_RECORDS      string = "[skyflow] Interface: %s - Detokenizing token %s"
	DETOKENIZING_SUCCESS      string = "[skyflow] Interface: %s - Successfully detokenized the token %s"
	DETOKENIZING_FAILED       string = "[skyflow] Interface: %s - Detokenization failed for the token %s"
	GETTING_RECORDS_BY_ID     string = "[skyflow] Interface: %s - Revealing records using skyflow ids for table %s"
	GET_RECORDS_BY_ID_FAILED  string = "[skyflow] Interface: %s - Failed Revealing records using skyflow ids for table %s"
	GET_RECORDS_BY_ID_SUCCESS string = "[skyflow] Interface: %s - Successfully revealed records using skyflow ids for table %s"
	INVOKE_CONNECTION_CALLED  string = "[skyflow] Interface: %s - Invoking connection"
	INVOKE_CONNECTION_SUCCESS string = "[skyflow] Interface: %s - Invoke connection success"
	INVOKE_CONNECTION_FAILED  string = "[skyflow] Interface: %s - Failed invoking connection"

	INVALID_VAULT_ID                       string = "[skyflow] Interface: %s - vault id invalid cannot be found"
	INVALID_VAULT_URL                      string = "[skyflow] Interface: %s - vault url %s is invalid or not secure"
	EMPTY_VAULT_ID                         string = "[skyflow] Interface: %s - vault id is empty."
	EMPTY_VAULT_URL                        string = "[skyflow] Interface: %s - vault url is empty."
	INVALID_BEARER_TOKEN                   string = "[skyflow] Interface: %s - bearer token is invalid or expired"
	INVALID_TABLE_NAME                     string = "[skyflow] Interface: %s - Key 'table' doesn't have a value of type String"
	EMPTY_TABLE_NAME                       string = "[skyflow] Interface: %s - Table Name is empty"
	RECORDS_KEY_NOT_FOUND                  string = "[skyflow] Interface: %s - records object key value not found"
	EMPTY_RECORDS                          string = "[skyflow] Interface: %s - records object is empty"
	TABLE_KEY_ERROR                        string = "[skyflow] Interface: %s - key 'table' is missing or payload is incorrectly formatted"
	FIELDS_KEY_ERROR                       string = "[skyflow] Interface: %s - key 'fields' is missing or payload is incorrectly formatted"
	INVALID_COLUMN_NAME                    string = "[skyflow] Interface: %s - column with given name is not present in the vault"
	EMPTY_COLUMN_NAME                      string = "[skyflow] Interface: %s - column name is empty"
	INVALID_TOKEN_ID                       string = "[skyflow] Interface: %s - token provided is invalid "
	EMPTY_TOKEN_ID                         string = "[skyflow] Interface: %s - token id is empty"
	ID_KEY_ERROR                           string = "[skyflow] Interface: %s - key 'id' is missing in the payload provided"
	REDACTION_KEY_ERROR                    string = "[skyflow] Interface: %s - key 'redaction' is missing in the payload provided"
	INVALID_REDACTION_TYPE                 string = "[skyflow] Interface: %s - provided redaction type value doesn’t match with one of : 'plain_text' 'redacted' 'default' or 'masked'"
	INVALID_FIELD                          string = "[skyflow] Interface: %s - invalid field %s"
	MISSING_TOKEN                          string = "[skyflow] Interface: %s - missing token property"
	MISSING_KEY_IDS                        string = "[skyflow] Interface: %s - Key 'ids' is missing in the payload provided"
	EMPTY_RECORD_IDS                       string = "[skyflow] Interface: %s - record ids cannot be empty"
	INVALID_RECORD_ID_TYPE                 string = "[skyflow] Interface: %s - invalid type of records id"
	MISSING_TABLE                          string = "[skyflow] Interface: %s - missing table property"
	INVALID_RECORD_TABLE_VALUE             string = "[skyflow] Interface: %s - invalid record table value"
	INVALID_CONNECTION_URL                 string = "[skyflow] Interface: %s - invalid connection url %s"
	REQUIRED_INPUTS_NOT_PROVIDED           string = "[skyflow] Interface: %s - required inputs are not provided"
	INVALID_EVENT_TYPE                     string = "[skyflow] Interface: %s - provide a valid event type"
	INVALID_EVENT_LISTENER                 string = "[skyflow] Interface: %s - provide valid event listener"
	UNKNOWN_ERROR                          string = "[skyflow] Interface: %s - %s"
	TRANSACTION_ERROR                      string = "[skyflow] Interface: %s - an error occurred during transaction"
	CONNECTION_ERROR                       string = "[skyflow] Interface: %s - error while initializing the connection"
	MISSING_REDACTION_VALUE                string = "[skyflow] Interface: %s - missing redaction value"
	INVALID_RECORDS_TYPE                   string = "[skyflow] Interface: %s - Key 'records' is of invalid type"
	INVALID_RECORD_IDS                     string = "[skyflow] Interface: %s - ids are not valid"
	MISSING_REDACTION                      string = "[skyflow] Interface: %s - redaction is missing"
	EMPTY_KEY_IN_REQUEST_BODY              string = "[skyflow] Interface: %s - empty key present in request body"
	EMPTY_KEY_IN_QUERY_PARAMS              string = "[skyflow] Interface: %s - empty key present in query parameters"
	EMPTY_KEY_IN_PATH_PARAMS               string = "[skyflow] Interface: %s - empty key present in path parameter"
	EMPTY_KEY_IN_REQUEST_HEADER_PARAMS     string = "[skyflow] Interface: %s - empty key present in request header"
	INVALID_FIELD_IN_PATH_PARAMS           string = "[skyflow] Interface: %s - invalid data type %s present in path parameters"
	INVALID_FIELD_IN_QUERY_PARAMS          string = "[skyflow] Interface: %s - invalid data type %s present in query parameters"
	INVALID_FIELD_IN_REQUEST_HEADER_PARAMS string = "[skyflow] Interface: %s - invalid data type %s present in request header"
	INVALID_FIELD_IN_REQUEST_BODY          string = "[skyflow] Interface: %s - invalid data type %s present in request body"
	FAILED_TO_REVEAL                       string = "[skyflow] Interface: %s - Failed to reveal"
	EMPTY_CONNECTION_URL                   string = "[skyflow] Interface: %s - Empty connection url is passed"
	NOT_FOUND_IN_RESPONSE                  string = "[skyflow] Interface: %s - %s is not found in response"
	BAD_REQUEST                            string = "[skyflow] Interface: %s - bad request"
	MISSING_COLUMN                         string = "[skyflow] Interface: %s - column name is missing"
	EMPTY_FIELDS                           string = "[skyflow] Interface: %s - fields is empty"
	SERVER_ERROR                           string = "[skyflow] Interface: %s - Server error %s"
	INVALID_RECORDS                        string = "[skyflow] Interface: %s - Records are not valid"
	DEPRECATED_GENERATE_TOKEN_FUNCTION     string = "[skyflow] Interface: %s - GenerateToken method is deprecated, will be removed in future, use GenerateBearerToken()"
)
