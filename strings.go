package fasthttp

var (
	defaultServerName  = []byte("fasthttp")
	defaultUserAgent   = []byte("fasthttp")
	defaultContentType = []byte("text/plain")
)

var (
	strSlash            = []byte("/")
	strSlashSlash       = []byte("//")
	strSlashDotDot      = []byte("/..")
	strSlashDotDotSlash = []byte("/../")
	strCRLF             = []byte("\r\n")
	strHTTP             = []byte("http")
	strHTTPS            = []byte("https")
	strHTTP11           = []byte("HTTP/1.1")
	strColonSlashSlash  = []byte("://")
	strColonSpace       = []byte(": ")

	strGet  = []byte("GET")
	strHead = []byte("HEAD")
	strPost = []byte("POST")

	strConnection       = []byte("Connection")
	strContentLength    = []byte("Content-Length")
	strContentType      = []byte("Content-Type")
	strDate             = []byte("Date")
	strHost             = []byte("Host")
	strReferer          = []byte("Referer")
	strServer           = []byte("Server")
	strTransferEncoding = []byte("Transfer-Encoding")
	strUserAgent        = []byte("User-Agent")
	strCookie           = []byte("Cookie")
	strSetCookie        = []byte("Set-Cookie")

	strCookieExpires = []byte("expires")
	strCookieDomain  = []byte("domain")
	strCookiePath    = []byte("path")

	strClose               = []byte("close")
	strUpgrade             = []byte("Upgrade")
	strChunked             = []byte("chunked")
	strIdentity            = []byte("identity")
	strPostArgsContentType = []byte("application/x-www-form-urlencoded")
)
