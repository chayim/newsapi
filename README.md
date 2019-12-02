# NEWSAPI

This library makes it about as painless as possible to retrieve articles from a variety of news sources, using [newsapi](https://www.newsapi.org).  Quick, dirty golang bindings.  Either set the *NEWSAPI* environment variable when integrating this library in your code, or set newsapi.token

## Examples

**Search**: NEWSAPI=sometoken newsapi.Search("A search term")

**Search for a specific date**: NEWSAPI=sometoken newsapi.SearchForDate("A search term", "2020-01-01")

For complete examples, see the unit tests.