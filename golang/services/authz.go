package services

func RespTypeCode() {
	// validate client_id and redirect_uri matches

	// User authentication

	// User Attributes from DB based on scopes

	// Consent page

	// Redirect to the redirect_uri with the authorization code (code) and the state specified by the client as query parameters (return HTTP 301 to the UserAgent)
}

//  この時、後からトークンエンドポイントで認可コードに紐づくIDトークン（ユーザの情報を含む）を発行することを考えるとバックエンドのDBにクライアントからの要求情報（client_id、redirect_uri、scope）、認証されたユーザの情報と認可コードを保存しておく必要があります。また、同時に認可コードの悪用や再利用を避けるためには認可コード自体の有効期限を極力短くしておく必要があるのと、トークンエンドポイント側で認可コードを利用したら無効化する処理などを実装する必要がありますので、それらの情報（フラグや有効期限など）もDBに保存しておく必要があるはずです。
// At this point, considering that an ID token (which includes user information) linked to
// the authorization code will be issued later at the token endpoint,
// it is necessary to save the request information from the client (client_id, redirect_uri, scope), the authenticated user's information, and the authorization code in the backend database.
// Additionally, to prevent misuse or reuse of the authorization code,
// it is essential to keep the validity period of the authorization code as short as possible.
// Also, once the authorization code is used at the token endpoint, it is necessary to implement a process to invalidate it.
// Therefore, this information (flags, expiration, etc.) must also be stored in the database.
