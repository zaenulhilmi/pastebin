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
later can be used to delete all expired contents. The implementation using a 
scheduler every 1 minute will execute ```DeleteExpiredContent()```
```
type ShortlinkService interface {
	GetContent(shortlink string) (*entities.Content, error)
	CreateContent(text string, expiryInMinutes int) (string, error)
	DeleteExpiredContent() error
}
```


## Read From and Writing To Cache

A paste can be read so many times and in parallel so it is needed to be read from
memory first instead of reading from database everytime. It will improve 
performance and reduce the load to the database. Hence, an adapter from repository
was created following the repository interface.

```
type ShortlinkRepository interface {
	FindContentByShortlink(shortlink string) (*entities.Content, error)
	CreateContent(shortlink string, text string, expiryByMinutes int) error
	DeleteExpiredContent() error
}
```
