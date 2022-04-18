# Pastebin Implementation using golang with TDD


##  Endpoints
There are two main endpoints for this pastebin, first is creating the shortlink
and second is get the content from shortlink

Create a shortlink
```
curl --location --request POST 'localhost:8080/create-paste' \
--header 'Content-Type: application/json' \
--data-raw '{
    "text": "This is all the text that I want to use",
    "expiry_in_minutes": 10
}'
```

Get a content of a shortlink
```
curl 'http://localhost:8080/paste?shortlink=abc'
```

## Deleting Expired Contents

ShortlinkService has a method to delete all expired content. The service method
later can be used to delete all expired contents.
```
type ShortlinkService interface {
	GetContent(shortlink string) (*entities.Content, error)
	CreateContent(text string, expiryInMinutes int) (string, error)
	DeleteExpiredContent() error
}
```

